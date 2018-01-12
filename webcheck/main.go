package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	c "github.com/sourcegraph/checkup"
)

func main() {
	c := c.Checkup{
		Checkers: []c.Checker{
			c.HTTPChecker{
				Name:     "philoserf",
				URL:      "http://philoserf.com",
				Attempts: 5},
			c.HTTPChecker{
				Name:     "markpaulayers",
				URL:      "http://markpaulayers.com",
				Attempts: 5},
		},
	}

	result, err := c.Check()
	if err != nil {
		log.Printf("check failure:\n\t %v", err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if _, err := io.WriteString(w, fmt.Sprintf("%v", result)); err != nil {
			log.Fatalf("failed to handle request:\n\t %v", err)
		}
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("failed to listen and serve:\n\t %v", err)
	}
}
