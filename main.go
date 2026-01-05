package main

import (
	"Be-Book-Padel/config"
	"Be-Book-Padel/database"
	"Be-Book-Padel/routes"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	config.LoadConfig()
	database.ConnectDB()

	r := gin.Default()
	routes.InitRoutes(r)

	r.Run(":" + viper.GetString("APP_PORT"))
}
