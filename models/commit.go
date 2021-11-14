package model

type Commit struct {
	Project     string
	Message     string `json:"message"`
	DateCreated string `json:"date_created"`
	URL         string `json:"url"`
}

type Author struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Date  string `json:"date"`
	_     interface{}
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

type CommitInformation struct {
	Author       Author       `json:"author"`
	Committer    Author       `json:"committer"`
	Message      string       `json:"message"`
	Tree         Tree         `json:"tree"`
	URL          string       `json:"url"`
	CommentCount int          `json:"comment_count"`
	Verification Verification `json:"verification"`
}

type RepoCommitResponse []struct {
	Commit CommitInformation `json:"commit"`
	URL    string            `json:"url"`
}

type CommitChannelType struct {
	Data Commit
	Err  error
}
