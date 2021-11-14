package model

type Commit struct {
	Message     string `json:"message"`
	DateCreated string `json:"date_created"`
}

type Author struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Date  string `json:"date"`
}

type Tree struct {
	Sha string `json:"sha"`
	URL string `json:"url"`
}

type Verification struct {
	Verified bool   `json:"verified"`
	Reason   string `json:"reason"`
	_        interface{}
}

type RepoCommitResponse struct {
	Author       Author       `json:"author"`
	Committer    Author       `json:"committer"`
	Message      string       `json:"message"`
	Tree         Tree         `json:"tree"`
	URL          string       `json:"url"`
	CommentCount int          `json:"comment_count"`
	Verification Verification `json:"verification"`
}
