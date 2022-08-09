package main

import (
	"github.com/go-resty/resty/v2"
)

type GithubRepo struct {
	ID         int         `json:"id"`
	Name       string      `json:"name"`
	Private    bool        `json:"private"`
	FullName   string      `json:"full_name"`
	Fork       bool        `json:"fork"`
	URL        string      `json:"url"`
	Visibility string      `json:"visibility"`
	Owner      GithubOwner `json:"owner"`
}

type GithubOwner struct {
	Login string `json:"login"`
	ID    int    `json:"id"`
}
type GitHubContext struct {
	Token  string
	Client *resty.Client
}

type NewOwner struct {
	NewOwner string `json:"new_owner"`
}

type GithubResponse struct {
	Message string                `json:"message"`
	Errors  []GithubErrorResponse `json:"errors"`
}

type GithubErrorResponse struct {
	Message string `json:"message"`
}
