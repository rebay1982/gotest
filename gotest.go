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

// FetchPullRequests this is a comment
func FetchPullRequests(owner string, repo string) ([]*github.PullRequest, error) {
	client := github.NewClient(nil)
	pulls, _, err := client.PullRequests.List(context.Background(), owner, repo, nil)

	return pulls, err
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
		fmt.Printf("%v. %v, owner %v\n", i+1, repo.GetName(), repo.GetOwner().GetLogin())
		pulls, err := FetchPullRequests(repo.GetOwner().GetLogin(), repo.GetName())

		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}

		for p, pull := range pulls {
			fmt.Printf("  %v. %v\n", p+1, pull.GetTitle())
		}
	}
}
