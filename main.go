package main

import (
	"fmt"
	"log"

	resty "github.com/go-resty/resty/v2"
	"github.com/stargazer39/github-utility/utils"
)

func main() {
	// Read token
	token, err := utils.ReadToken()
	check(err)

	// Initialize HTTP client
	client := resty.New()

	github := &GitHubContext{
		Client: client,
		Token:  token,
	}
	// Get Organization from user
	from := utils.GetString("Transfer from: ")

	fmt.Printf("\nGetting Repositories of %s...\n", from)

	// List all repositories
	repos, _, err := list(from, github)
	check(err)

	// Display all repos
	displayRepos(*repos)

	// Get repos to transfer
	choices, err2 := utils.GetIntChoice("\nEnter one or more Repo numbers (Ex: 1 5 10): ")
	check(err2)

	// Fix the index
	for i := range choices {
		choices[i]--
	}

	// Get where to transfer
	to := utils.GetString("Transfer to: ")

	// Show what will happen before
	for _, r := range choices {
		fmt.Printf("Repo %s will be transferred to %s\n", (*repos)[r].FullName, to)
	}

	// If confirmed run the tasks
	if utils.GetConfirm("Continue ? [N]: ", true) {
		for _, r := range choices {
			if err, resp := transfer(to, (*repos)[r], github); err != nil {
				log.Printf("\n%s", resp)
				continue
			}

			fmt.Printf("Repo %s transferred to %s\n", (*repos)[r].FullName, to)
		}
	} else {
		log.Println("\nYou cancelled the transfer.")
	}
}

func check(err error) {
	if err != nil {
		log.Panic(err)
	}
}
