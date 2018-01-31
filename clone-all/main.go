package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/shurcooL/githubql"
	"golang.org/x/oauth2"
)

func init() {
	log.SetPrefix("clone-all: ")
}

func main() {
	ctx := context.Background()
	client := authn()

	var query struct {
		Viewer struct {
			Login     githubql.String
			CreatedAt githubql.DateTime
		}
	}

	if err := client.Query(ctx, &query, nil); err != nil {
		log.Fatal(err)
	}

	fmt.Println("    Login:", query.Viewer.Login)
	fmt.Println("CreatedAt:", query.Viewer.CreatedAt)

	log.Fatal("TODO: finish the base implementation")
}

func authn() *githubql.Client {
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
	)

	httpClient := oauth2.NewClient(context.Background(), src)

	return githubql.NewClient(httpClient)
}
