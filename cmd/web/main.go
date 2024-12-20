package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippetbox/view", snippetView)
	mux.HandleFunc("/snippetbox/create", snippetCreate)
	log.Print("Server started on port 4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
