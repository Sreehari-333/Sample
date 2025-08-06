package db

import (
	model "Sample/Model"
	"fmt"
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

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Println("ℹNo .env file loaded (this is expected in production)")
	}

	log.Println("Database connected successfully")

	// Migrating tables

	err = DB.AutoMigrate(&model.Signup{})

	if err != nil {
		log.Fatal("error with creating table", err)
	}

}
