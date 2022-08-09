package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/stargazer39/github-utility/utils"
)

func list(username string, ctx *GitHubContext) (*[]GithubRepo, string, error) {
	var repos []GithubRepo

	req, err := ctx.Client.
		R().
		SetHeader("Authorization", fmt.Sprintf("token %s", ctx.Token)).
		SetResult(&repos).
		Get(fmt.Sprintf("https://api.github.com/users/%s/repos", username))

	if err != nil {
		return nil, "", err
	}

	if req.StatusCode() != http.StatusOK {
		return nil, utils.PrettyPrintJSONBytes(req.Body()), &GithubError{
			message: "Error with github",
			code:    req.StatusCode(),
		}
	}

	return &repos, utils.PrettyPrintJSONBytes(req.Body()), nil
}

func transfer(to string, repo GithubRepo, ctx *GitHubContext) (error, string) {
	bytes, _ := json.Marshal(&NewOwner{NewOwner: to})

	req, err := ctx.Client.
		R().
		SetHeader("Authorization", fmt.Sprintf("token %s", ctx.Token)).
		SetBody(bytes).
		Post(fmt.Sprintf("https://api.github.com/repos/%s/%s/transfer", repo.Owner.Login, repo.Name))

	if err != nil {
		return err, ""
	}

	if req.StatusCode() != http.StatusAccepted {
		return &GithubError{
			message: "Error with github " + req.Status(),
			code:    req.StatusCode(),
		}, utils.PrettyPrintJSONBytes(req.Body())
	}

	return nil, utils.PrettyPrintJSONBytes(req.Body())
}

func displayRepos(repos []GithubRepo) {
	for i, r := range repos {
		fmt.Printf("%d. %s\t%s\t%s\n", i+1, r.FullName, r.Visibility, r.Owner.Login)
	}
}
