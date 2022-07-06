package main

import (
	"Simple_Web_API_Login/database"
	"Simple_Web_API_Login/routes"
	"log"

	"github.com/joho/godotenv"
)

func init() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	database.StartDB()
	r := routes.StartApp()
	r.Run(":9000")
}
