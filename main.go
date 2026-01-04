package main

import (
	"Be-Book-Padel/config"
	"Be-Book-Padel/database"
	"Be-Book-Padel/models"
	"Be-Book-Padel/routes"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	config.LoadConfig()
	database.ConnectDB()

	// migrate database
	database.DB.AutoMigrate(&models.Users{})
	database.DB.AutoMigrate(&models.RefreshToken{})

	r := gin.Default()
	routes.InitRoutes(r)

	r.Run(":" + viper.GetString("APP_PORT"))
}
