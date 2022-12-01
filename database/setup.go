package database

import (
	"os"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

func InitDB() *gorm.DB {
	godotenv.Load(".env")

	DB_HOST := os.Getenv("DB_HOST")
	DB_PORT := os.Getenv("DB_PORT")
	DB_USER := os.Getenv("DB_USER")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	DB_NAME := os.Getenv("DB_NAME")

	db, err := gorm.Open("postgres", "host="+DB_HOST+" port="+DB_PORT+" user="+DB_USER+" dbname="+DB_NAME+" sslmode=disable password="+DB_PASSWORD)
	if err != nil {
		panic(err)
	}

	return db
}
