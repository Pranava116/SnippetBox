package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	ts, err := template.ParseFiles("./ui/html/pages/home.tmpl.html")
	if err != nil {
		log.Fatal(err)
		http.Error(w, "Internal server error", 501)

	}
	err = ts.Execute(w, nil)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Inter nal server error", 501)
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
