package model

type RawProject struct {
	Title        map[string]string
	Description  map[string]string
	Tech         map[string][]string
	RepoName     map[string]string
	RepoLink     map[string]string
	DeployedLink map[string]string
	ImageLink    map[string]string
	Order        map[string]int
}

type Project struct {
	Title        string   `json:"title"`
	Description  string   `json:"description"`
	Tech         []string `json:"tech"`
	RepoName     string   `json:"repoName"`
	RepoLink     string   `json:"repoLink"`
	DeployedLink string   `json:"deployedLink"`
	ImageLink    string   `json:"imageLink"`
	Order        int      `json:"id"`
}
