package main

import (
	"github.com/joho/godotenv"
	"log"
	"portfolioBE/integrations"
	"portfolioBE/server"
)

func init() {
	err := godotenv.Load(".env")

	integrations.SetUp()
}

func main() {
	server.Init()
}
