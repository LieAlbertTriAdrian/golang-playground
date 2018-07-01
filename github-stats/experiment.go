package main

import (
	"context"
	"fmt"
	"os"

	"github.com/google/go-github/github"
    "github.com/joho/godotenv"
	"golang.org/x/oauth2"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		fmt.Println("Error loading .env file")
	}

	githubAccessToken := os.Getenv("GITHUB_ACCESS_TOKEN")

	context := context.Background()
	tokenSource := oauth2.StaticTokenSource(
		&oauth2.Token{ AccessToken: githubAccessToken },
	)
	tokenClient := oauth2.NewClient(context, tokenSource)

	githubClient := github.NewClient(tokenClient)

	repos, _, err := githubClient.Repositories.List(context, "", nil)

	fmt.Println("hello world")
	fmt.Println(repos[3])
	fmt.Println("hello")
	fmt.Println(repos[3].GetOwner().GetLogin())
	fmt.Println(repos[3].GetID())
	fmt.Println(repos[3].GetName())
	fmt.Println(repos[3].GetCloneURL())

	commits, _, err := githubClient.Repositories.ListCommits(context, repos[3].GetOwner().GetLogin(), repos[3].GetName(), nil)

	fmt.Println(commits[0].GetSHA())
	fmt.Println(commits[0].GetCommit().GetCommitter().GetDate())
	fmt.Println(commits[0].GetCommit().GetAuthor().GetName())

	fmt.Println(err)
}