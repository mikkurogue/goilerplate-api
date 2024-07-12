package main

import (
	"goilerplate-api/api"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	// env load as first thing.
	err := godotenv.Load()
	if err != nil {
		log.Fatal("no .env file exists...")
	}

	// bootstrap our api here
	api.Start()
}
