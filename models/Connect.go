package models

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB

func ConnectDatabase() {
	dataSourceName := "host=" + os.Getenv("POSTGRES_HOST") + " user=" + os.Getenv("POSTGRES_USER") + " password=" + os.Getenv("POSTGRES_PASSWORD") + " dbname=" + os.Getenv("POSTGRES_DB") + " port=" + os.Getenv("POSTGRES_PORT")
	database, err := gorm.Open(postgres.Open(dataSourceName), &gorm.Config{})

	if err != nil {
		fmt.Print(err)
		panic("Failed to connect to database!")
	}

	DB = database
}
