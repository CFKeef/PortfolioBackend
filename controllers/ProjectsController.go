package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"portfolioBE/integrations"
	model "portfolioBE/models"
	"strconv"
	"strings"
)

type ProjectsController struct{}

func (p ProjectsController) All(c *gin.Context) {
	if val, ok := c.GetQuery("tech"); ok {
		var filtered []model.Project
		var commits = make(map[string]model.Commit)

		for _,v := range integrations.Projects {
			var shouldAdd = false

			for _,tech := range v.Tech {
				if val == strings.ToLower(tech) {
					shouldAdd = true
					break
				}
			}

			if shouldAdd {
				filtered = append(filtered, v)
			}
		}

		for _,proj := range filtered {
			commits[proj.RepoName]=integrations.Commits[proj.RepoName]
		}

		c.JSON(http.StatusOK, gin.H{
			"count":    len(filtered),
			"projects": filtered,
			"commits":  commits,
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"count":    len(integrations.Projects),
		"projects": integrations.Projects,
		"commits":  integrations.Commits,
	})
}

func (p ProjectsController) One(c *gin.Context) {
	id := c.Param("id")
	if id != "" {
		intId, _ := strconv.Atoi(id)

		target := integrations.Projects[intId]

		c.JSON(http.StatusOK, target)
		return
	}

	c.JSON(http.StatusBadRequest, gin.H{
		"message": "Couldn't find project by that ID",
	})
}

func (p ProjectsController) Filters(c *gin.Context) {
	c.JSON(http.StatusOK, integrations.Filters)
}
