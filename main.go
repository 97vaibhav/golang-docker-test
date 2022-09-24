package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "homepage %q", html.EscapeString(r.URL.Path))
	})
	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello about page ")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
