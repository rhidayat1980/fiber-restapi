package database

import (
	"log"
	"os"

	"github.com/rhidayat1980/fiber-restapi/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB gorm connector
var DB *gorm.DB

// function for conecting to DB
func ConnectDB() {
	var err error

	env := os.Getenv("DATABASE_URL")
	DB, err := gorm.Open(postgres.Open(env), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	err = DB.AutoMigrate(&models.User{}, &models.Session{}, &models.Product{})
	if err != nil {
		log.Fatal(err)
	}
}
