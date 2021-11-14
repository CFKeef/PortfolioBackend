package server

import (
	"portfolioBE/controllers"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	controllers.SetUpProjectsController()

	projectsGroup := router.Group("/projects")
	{
		projects := new(controllers.ProjectsController)
		projectsGroup.GET("/", projects.GetAllProjects)
	}

	return router
}
