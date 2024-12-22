package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

func main() {
	mux := http.NewServeMux()
	addr := flag.String("addr", ":4000", "HTTP Network Port")
	flag.Parse()
	infoLogger := log.New(os.Stdout, "INFO\t", log.Ltime|log.Ldate)
	erroLogger := log.New(os.Stderr, "ERROR\t", log.Ltime|log.Ldate|log.Lshortfile)
	fileserver := http.FileServer(http.Dir("./ui/static/"))
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: erroLogger,
		Handler:  mux,
	}
	mux.Handle("/static/", http.StripPrefix("/static", fileserver))
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippetbox/view", snippetView)
	mux.HandleFunc("/snippetbox/create", snippetCreate)
	infoLogger.Printf("Server started on port %s", *addr)
	err := srv.ListenAndServe()
	erroLogger.Fatal(err)
}
