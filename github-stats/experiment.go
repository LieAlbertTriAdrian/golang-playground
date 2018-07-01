package main

import (
	"context"
    "encoding/csv"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/google/go-github/github"
    "github.com/joho/godotenv"
	"golang.org/x/oauth2"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		fmt.Println("Error loading .env file")
	}

	bytes, err := ioutil.ReadFile("data/samples.txt")

	if err != nil {
		fmt.Println(err)
	}

	repositoryListInString := string(bytes)
	repositoryList := strings.Split(repositoryListInString, "\n")

	fmt.Println(repositoryList)

	githubAccessToken := os.Getenv("GITHUB_ACCESS_TOKEN")

	context := context.Background()
	tokenSource := oauth2.StaticTokenSource(
		&oauth2.Token{ AccessToken: githubAccessToken },
	)
	tokenClient := oauth2.NewClient(context, tokenSource)

	githubClient := github.NewClient(tokenClient)

	repos, _, err := githubClient.Repositories.List(context, "", nil)

	fmt.Println("hello world")
	// fmt.Println(repos[3])
	// fmt.Println("hello")
	// fmt.Println(repos[3].GetOwner().GetLogin())
	// fmt.Println(repos[3].GetID())
	// fmt.Println(repos[3].GetName())
	// fmt.Println(repos[3].GetCloneURL())
	repo1, _, err := githubClient.Repositories.Get(context, "LieAlbertTriAdrian", "golang-playground")

	repo2, _, err := githubClient.Repositories.Get(context, "hashicorp", "terraform")

	fmt.Println(repo1.GetCloneURL())
	fmt.Println(repo2.GetCloneURL())

	commits, _, err := githubClient.Repositories.ListCommits(context, repos[3].GetOwner().GetLogin(), repos[3].GetName(), nil)

	fmt.Println(commits[0].GetSHA())
	fmt.Println(commits[0].GetCommit().GetCommitter().GetDate())
	fmt.Println(commits[0].GetCommit().GetAuthor().GetName())

    file, err := os.Create("output/result.csv")

    if err != nil {
    	fmt.Println(err)
    }

    defer file.Close()

    writer := csv.NewWriter(file)

    defer writer.Flush()

	err1 := writer.Write(repositoryList)

	if err1 != nil {
		fmt.Println("error writing the value of repository")
	}

	// fmt.Println(err)
}