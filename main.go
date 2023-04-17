package main

import (
	"context"
	"fmt"
	"os"

	"github.com/shurcooL/githubv4"

	"golang.org/x/oauth2"
)

func main() {
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("TOKEN")},
	)
	httpClient := oauth2.NewClient(context.Background(), src)

	client := githubv4.NewClient(httpClient)
	// Use client...

	var m struct {
		CreatePullRequest struct {
			PullRequest struct {
				ID githubv4.ID
			}
		} `graphql:"createPullRequest(input: $input)"`
	}
	input := githubv4.CreatePullRequestInput{
		BaseRefName:  "main",
		HeadRefName:  "b1",
		RepositoryID: "R_kgDOJW-w6w",
		Title:        "Jian Test",
	}

	err := client.Mutate(context.Background(), &m, input, nil)
	if err != nil {
		fmt.Println("Error")
		fmt.Println(err)
	}
	fmt.Println("ID of PR is", m.CreatePullRequest.PullRequest.ID)

}
