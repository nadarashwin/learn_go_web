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

	addr := flag.String("addr", ":4000", "default port to run on")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime|log.Llongfile)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	srv := http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(), // Moved all the route code to a different file (routes.go)
	}

	// log.Printf("Starting server on %s", *addr) // standard log format
	infoLog.Printf("Starting server on %s", *addr) // customised one
	// err := http.ListenAndServe(*addr, mux)  // in place of (srv := http.Server)
	err := srv.ListenAndServe()
	// log.Fatal(err)
	errorLog.Fatal(err)
}
