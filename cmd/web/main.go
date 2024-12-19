package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Displaying all the snippets")
}
func snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Displaying the snippet with the id %d", id)
}
func snippetCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allowed", http.MethodPost)
		http.Error(w, "Mthos not allowed", http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprintf(w, "Displaying the create page")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippetbox/view", snippetView)
	mux.HandleFunc("/snippetbox/create", snippetCreate)
	log.Print("Server started on port 4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
