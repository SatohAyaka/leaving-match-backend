package main

import (
	"SatohAyaka/leaving-match-backend/lib"
	"SatohAyaka/leaving-match-backend/router"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	lib.InitDB()
	router.Router()

	if err != nil {
		log.Fatal("Error loading .env file")
	}
	log.Println("MYSQL_USER:", os.Getenv("MYSQL_USER"))
	log.Println("MYSQL_HOST:", os.Getenv("MYSQL_HOST"))
}
