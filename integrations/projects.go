package integrations

import (
	"encoding/json"
	"github.com/contentful-labs/contentful-go"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	model "portfolioBE/models"
	"sync"
)

var (
	ContentfulCMA *contentful.Contentful
	Projects      map[int]model.Project
	Commits       map[string]model.Commit
	Filters       []string
)

func SetUp() {
	token := os.Getenv("CONTENTFUL_PERSONAL")

	ContentfulCMA = contentful.NewCMA(token)

	fetchProjects()
	fetchCommits()
}

func fetchCommits() {
	t := os.Getenv("GITHUB_TOKEN")

	ch := make(chan model.CommitChannelType)
	var wg sync.WaitGroup

	for _, project := range Projects {
		wg.Add(1)
		go fetchCommit(project.RepoName, t, ch, &wg)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	commitMap := make(map[string]model.Commit)

	for result := range ch {
		commitMap[result.Data.Project] = result.Data
	}

	Commits = commitMap
}

func fetchCommit(name string, token string, ch chan<- model.CommitChannelType, wg *sync.WaitGroup) {
	defer (*wg).Done()
	req, err := http.NewRequest("GET", "https://api.github.com/repos/CFKeef/"+name+"/commits", nil)
	res := new(model.CommitChannelType)

	if err != nil {
		res.Err = err
		ch <- *res
	}

	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		res.Err = err
		ch <- *res
	}

	defer resp.Body.Close()

	responseData, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		res.Err = err
		ch <- *res
	}

	var arr model.RepoCommitResponse

	err = json.Unmarshal(responseData, &arr)

	temp := arr[0]

	if err != nil {
		res.Err = err
		ch <- *res
	}


	res.Data = model.Commit{
		Project: name,
		Message:     temp.Commit.Message,
		DateCreated: temp.Commit.Committer.Date,
		URL:         temp.URL,
	}
	res.Err = nil
	ch <- *res
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
		entry, err := ContentfulCMA.Entries.Get(space.Sys.ID, sp.Sys.ID)

		if err != nil {
			log.Println(err)
		}

		t := model.RawProject{}

		b, err := json.Marshal(entry.Fields)

		if err != nil {
			log.Fatal(err)
		}

		err = json.Unmarshal(b, &t)

		if err != nil {
			log.Fatal(err)
		}

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

