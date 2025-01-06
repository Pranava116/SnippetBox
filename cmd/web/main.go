package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {
	mux := http.NewServeMux()

	addr := flag.String("addr", ":4000", "HTTP Network Port")
	flag.Parse()
	infoLog := log.New(os.Stdout, "INFO\t", log.Ltime|log.Ldate)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ltime|log.Ldate|log.Lshortfile)
	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}
	fileserver := http.FileServer(http.Dir("./ui/static/"))
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  mux,
	}
	mux.Handle("/static/", http.StripPrefix("/static", fileserver))
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippetbox/view", app.snippetView)
	mux.HandleFunc("/snippetbox/create", app.snippetCreate)
	infoLog.Printf("Server started on port %s", *addr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}
