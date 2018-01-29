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
<<<<<<< HEAD
	client := authn()
=======
	client := auth()
>>>>>>> Various clean up changes

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

<<<<<<< HEAD
func authn() *githubql.Client {
=======
func auth() *githubql.Client {
>>>>>>> Various clean up changes
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
	)

	httpClient := oauth2.NewClient(context.Background(), src)

	return githubql.NewClient(httpClient)
}
