package main

import (
	"log"
	"os"

	"github.com/gabbottron/catpix-api/pkg/api"
	"github.com/gabbottron/catpix-api/pkg/datastore"
	"github.com/joho/godotenv"
)

// check for env file
func GetAppEnv() error {
	env := os.Getenv("APP_ENV")

	if len(env) == 0 {
		err := godotenv.Load()
		if err != nil {
			log.Panic("Error loading .env file")
			return err
		}
	}

	return nil
}

func main() {
	log.Println("Starting the catpix API...")

	// Make sure the application can load env
	err := GetAppEnv()
	if err != nil {
		log.Panic("Error loading APP ENV!")
	}

	// Connect to the database
	err = datastore.InitDB()
	if err != nil {
		log.Panic("Error connecting to the DB")
	}

	// Initialize the router and run it
	router := api.InitRouter()
	router.Run()
}
