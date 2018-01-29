package main

import (
	_ "expvar"
	"fmt"
	"log"
	"net/http"
)

func init() {
	log.SetPrefix("expvars-example: ")
}

func main() {
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/status", statusHandler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if _, err := fmt.Fprintln(w, "Hello"); err != nil {
		log.Print(err)
	}
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
	if _, err := fmt.Fprintln(w, `{"status": "OK"}`); err != nil {
		log.Print(err)
	}
}
