package main

import (
	"context"
	"fmt"

	"github.com/google/go-github/github"
)

// FetchRepositories This is a comment.
//
func FetchRepositories(username string) ([]*github.Repository, error) {
	client := github.NewClient(nil)
	repos, _, err := client.Repositories.List(context.Background(), username, nil)

	return repos, err
}

func main() {
	var username string
	fmt.Print("Enter GitHub username: ")
	fmt.Scanf("%s", &username)

	repos, err := FetchRepositories(username)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	for i, repo := range repos {
		fmt.Printf("%v. %v\n", i+1, repo.GetName())
	}
}
