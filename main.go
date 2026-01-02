package main

import (
	"Be-Book-Padel/internal/config"
	"Be-Book-Padel/internal/database"
)

func main(){
	config.LoadConfig()
	database.ConnectDB()
}