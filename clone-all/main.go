package main

import (
	"context"
	"fmt"
	"os"

	"github.com/shurcooL/githubql"
	"go.uber.org/zap"
	"golang.org/x/oauth2"
)

func main() {

	logger := newLogger()
	log := logger.Sugar()
	defer func() {
		if err := logger.Sync(); err != nil {
			fmt.Printf("log sync error: %v", err)
		}
	}()

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

	fmt.Println("Login:", query.Viewer.Login)
	fmt.Println("Date:", query.Viewer.CreatedAt)

	log.Info("TODO: finish the base implementation")
}

func newLogger() *zap.Logger {
	logger, err := zap.NewProduction()
	if err != nil {
		fmt.Printf("log creation error: %v", err)
	}
	return logger
}

func authn() *githubql.Client {
	return githubql.NewClient(
		oauth2.NewClient(
			context.Background(),
			oauth2.StaticTokenSource(
				&oauth2.Token{
					AccessToken: os.Getenv("GITHUB_TOKEN"),
				},
			),
		),
	)
}
