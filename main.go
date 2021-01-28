package main

import (
	"./models"
	"./routes"
	"github.com/joho/godotenv"
)


func main() {

	err := godotenv.Load()
	if err != nil {
		panic("Error getting env")
	}
	models.ConnectDatabase()
	models.InitDatabase()
	routes.ConfigRouter()
}
