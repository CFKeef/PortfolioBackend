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
	Title        string   `json:"title,omitempty"`
	Description  string   `json:"description,omitempty"`
	Tech         []string `json:"tech,omitempty"`
	RepoName     string   `json:"repoName,omitempty"`
	RepoLink     string   `json:"repoLink,omitempty"`
	DeployedLink string   `json:"deployedLink,omitempty"`
	ImageLink    string   `json:"imageLink,omitempty"`
	Order        int      `json:"id,omitempty"`
}
