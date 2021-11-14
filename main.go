package main

import (
	"log"
	"portfolioBE/controllers"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	controllers.SetUpProjectsController()

	v1 := router.Group("/v1")
	{
		// Start PROJECTS
		projects := new(controllers.ProjectsController)
		v1.GET("/projects/", projects.GetAllProjects)
	}

	router.Run("localhost:8080")
}
