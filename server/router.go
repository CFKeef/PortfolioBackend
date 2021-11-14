package server

import (
	"portfolioBE/controllers"

	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "https://www.patryck.dev/")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Max")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func NewRouter() *gin.Engine {
	router := gin.Default()

	router.Use(CORSMiddleware())

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
