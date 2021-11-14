package main

import (
	"github.com/joho/godotenv"
	"portfolioBE/integrations"
	"portfolioBE/server"
)

func init() {
	godotenv.Load(".env")

	integrations.SetUp()
}

func main() {
	server.Init()
}
