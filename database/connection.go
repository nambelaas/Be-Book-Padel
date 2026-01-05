package database

import (
	"Be-Book-Padel/models"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDB() {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Silent,
			IgnoreRecordNotFoundError: true,
			ParameterizedQueries:      true,
			Colorful:                  true,
		},
	)
	cred := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s search_path=%s",
		viper.GetString("DB_HOST"),
		viper.GetString("DB_USER"),
		viper.GetString("DB_PASSWORD"),
		viper.GetString("DB_NAME"),
		viper.GetString("DB_PORT"),
		viper.GetString("DB_SSLMODE"),
		viper.GetString("DB_SCHEMA"),
	)

	db, err := gorm.Open(postgres.Open(cred), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		log.Fatal("Failed to connect database: ", err)
	}

	DB = db
	log.Println("Database connected")

	InitEnums(db)
	MigrateDB()
	InsertAdminUser()
}

func MigrateDB() {
	if err := DB.AutoMigrate(
		&models.Users{},
		&models.RefreshToken{},
		&models.Field{},
		&models.FieldPricing{},
	); err != nil {
		log.Fatal("Failed to migrate database: ", err)
	}

	log.Println("Database migrated")
}

func InitEnums(db *gorm.DB) {
	if err := db.Exec(`
		DO $$
		BEGIN
			IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'user_roles') THEN
				CREATE TYPE be_book_padel_db.user_roles AS ENUM ('admin','user','staff');
			END IF;
		END$$;
	`).Error; err != nil {
		log.Fatal("Failed to create user role enum", err)
	}

	if err := db.Exec(`
		DO $$
		BEGIN
			IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'day_type') THEN
				CREATE TYPE be_book_padel_db.day_type AS ENUM ('weekday','weekend');
			END IF;
		END$$;
	`).Error; err != nil {
		log.Fatal("Failed to create day type enum: ", err)
	}
}

func InsertAdminUser() {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(viper.GetString("ADMIN_PASSWORD")), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal("Failed to hash admin password: ", err)
		return
	}

	var count int64
	DB.Model(&models.Users{}).Where("role = ?", "admin").Count(&count)
	if count == 0 {
		admin := models.Users{
			FirstName: "Admin",
			LastName:  "User",
			Email:     "admin@example.com",
			Password:  string(hashedPassword),
			Role:      models.Admin,
		}
		DB.Create(&admin)
	}
}
