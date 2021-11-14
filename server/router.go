package server

import (
	"portfolioBE/controllers"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		projectsGroup := v1.Group("/projects")
		{
			projects := new(controllers.ProjectsController)
			projectsGroup.GET("", projects.All)
			projectsGroup.GET("/:id", projects.One)
			projectsGroup.GET("/filters", projects.Filters)
		}
	}

	return router
}
