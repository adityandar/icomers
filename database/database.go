package database

import (
	"fmt"
	"icomers/models"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	var err error
	// DB, err = sql.Open("postgres", connStr)
	DB, err = gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error when opening database, %v\n", err)
	}
	if err = DB.AutoMigrate(&models.User{}, &models.Product{}); err != nil {
		log.Fatalf("Error when running auto migration, %v\n", err)
	}

	log.Println("Database connected")

}
