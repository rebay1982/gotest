package main

import (
	"context"
	"fmt"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

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
