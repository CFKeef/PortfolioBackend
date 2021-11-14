package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	model "portfolioBE/models"

	"github.com/contentful-labs/contentful-go"
	"github.com/gin-gonic/gin"
)

type ProjectsController struct{}

var (
	ContentfulCMA *contentful.Contentful
	Projects      map[int]model.Project
	Commits       map[int]model.Commit
	Filters       []string
)

var ProjectData ProjectsController

func SetUpProjectsController() {
	token := os.Getenv("CONTENTFUL_PERSONAL")

	ContentfulCMA = contentful.NewCMA(token)

	fetchProjects()
	fetchCommits()
}

func fetchCommits() {
	t := os.Getenv("GITHUB_TOKEN")

	ch := make(chan model.Commit)

	for _, project := range Projects {
		go fetchCommit(project.RepoName, t, ch)
	}

	fmt.Println(<-ch)
}

func fetchCommit(name string, token string, ch chan<- model.Commit) {
	req, err := http.NewRequest("GET", "https://api.github.com/repos/CFKeef/"+name+"/commits", nil)

	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		log.Fatal("Couldn't fetch")
	}

	defer resp.Body.Close()

}

func fetchProjects() {
	s := os.Getenv("CONTENTFUL_SPACEID")

	space, err := ContentfulCMA.Spaces.Get(s)

	if err != nil {
		log.Fatal(err)
	}

	collection := ContentfulCMA.Entries.List(space.Sys.ID)
	collection, err = collection.Next()

	if err != nil {
		log.Fatal(err)
	}

	spaces := collection.ToSpace()

	processProjects(spaces, space)
}

func processProjects(spaces []*contentful.Space, space *contentful.Space) {
	var projectMap = make(map[int]model.Project)
	var filterSet = make(map[string]bool)

	for _, sp := range spaces {
		entry, _ := ContentfulCMA.Entries.Get(space.Sys.ID, sp.Sys.ID)

		t := model.RawProject{}

		b, _ := json.Marshal(entry.Fields)

		_ = json.Unmarshal(b, &t)

		for _, tech := range t.Tech["en-US"] {

			_, ok := filterSet[tech]

			if !ok {
				filterSet[tech] = true
			}

		}

		projectMap[t.Order["en-US"]] = model.Project{
			Title:        t.Title["en-US"],
			Description:  t.Description["en-US"],
			Tech:         t.Tech["en-US"],
			RepoName:     t.RepoName["en-US"],
			RepoLink:     t.RepoLink["en-US"],
			DeployedLink: t.DeployedLink["en-US"],
			ImageLink:    t.ImageLink["en-US"],
			Order:        t.Order["en-US"],
		}
	}

	Projects = projectMap

	keys := make([]string, 0, len(filterSet))

	for k := range filterSet {
		keys = append(keys, k)
	}

	Filters = keys
}

func (p ProjectsController) GetAllProjects(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"count":    len(Projects),
		"projects": Projects,
		"commits":  Commits,
	})
}
