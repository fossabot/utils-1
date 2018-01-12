package main

import (
	_ "expvar"
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello")
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, `{"status": "OK"}`)
}

func main() {
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/status", statusHandler)
	http.ListenAndServe(":8080", nil)
}
