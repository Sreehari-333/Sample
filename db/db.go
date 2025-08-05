package db

import (
	model "Sample/Model"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Connecting Database

var DB *gorm.DB

func ConnectDb() {

	var err error

	err = godotenv.Load()

	if err != nil {
		log.Fatal("failed to load env")
	}

	dsn := os.Getenv("DB_DSN")

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Database connection failed", err)
	}

	log.Println("Database connected successfully")

	// Migrating tables

	err = DB.AutoMigrate(&model.Signup{})

	if err != nil {
		log.Fatal("error with creating table", err)
	}

}
