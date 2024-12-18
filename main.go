package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte("Hello World"))
	log.Print(r)
}
func snippetView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display the snippet here..."))
}
func snippetCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Method invalid", http.StatusMethodNotAllowed)
	}
	w.Write([]byte("Create a new snippet here..."))
}
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Print(("Staarting on server 4000"))
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
