package main

import (
	"context"
	"fmt"

	"github.com/google/go-github/github"
)

// FetchOrganizations Fetch all the public organizations' membership of a user.
//
func FetchOrganizations(username string) ([]*github.Organization, error) {
	client := github.NewClient(nil)
	orgs, _, err := client.Organizations.List(context.Background(), username, nil)
	return orgs, err
}

func main() {
	fmt.Println("GoTest")
}
