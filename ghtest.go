package main

import (
	"context"
	"fmt"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
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
	ghToken := ""

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: ghToken},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)
	prs, _, err := client.PullRequests.List(ctx, "", "", &github.PullRequestListOptions{State: "open"})

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	for p, pr := range prs {
		fmt.Printf("  %v. %v\n", p+1, pr.GetTitle())
	}

}
