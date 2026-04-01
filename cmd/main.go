package main

import (
	"log"
	"os"
	"strconv"

	"github.com/erwindrsno/Quotation-Builder/internal/app"
	"github.com/joho/godotenv"
)

func main() {
	//First of all, remember to clear tmp files before running the app. If not, configuring line 17 to the path is necessary
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	server := app.New()
	port, err := strconv.Atoi(os.Getenv("APP_PORT"))
	if err != nil {
		log.Fatal("Failed to convert string to int for SERVER_PORT")
	}
	server.Run(port)
}
