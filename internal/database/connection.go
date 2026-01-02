package database

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	cred := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s search_path=tiket_bioskop",
		viper.GetString("DB_HOST"),
		viper.GetString("DB_USER"),
		viper.GetString("DB_PASSWORD"),
		viper.GetString("DB_NAME"),
		viper.GetString("DB_PORT"),
		viper.GetString("DB_SSLMODE"),
	)

	db, err := gorm.Open(postgres.Open(cred), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect database: ", err)
	}

	DB = db
	log.Println("Database connected")
}