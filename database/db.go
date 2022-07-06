package database

import (
	"Simple_Web_API_Login/models"
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func StartDB() {
	config := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_NAME"), os.Getenv("DB_PASSWORD"))

	DB, err = gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		fmt.Println("Error in connecting to database: ", err)
	}

	DB.AutoMigrate(&models.User{})
}

func GetDB() *gorm.DB {
	return DB
}
