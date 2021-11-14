package main

import (
	"github.com/joho/godotenv"
	"log"
	"portfolioBE/integrations"
	"portfolioBE/server"
)

func init() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	integrations.SetUp()
}

func main() {
	server.Init()
}
