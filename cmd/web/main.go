package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/nadarashwin/learn_go_web/pkg/models/mysql"

	_ "github.com/go-sql-driver/mysql"
)

// snippets field that will allow us to make the SnippetModle object available to our handlers.
type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	snippets *mysql.SnippetModel
}

func main() {

	addr := flag.String("addr", ":4000", "default port to run on")
	dsn := flag.String("dsn", "web:web@/snippetbox?parseTime=true", "MySQL data source name")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime|log.Llongfile)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}

	// Close connection pool before main() function exits.
	defer db.Close()

	// Initailize a mysql.SnippetModel instance and add it to the application dependencies.
	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
		snippets: &mysql.SnippetModel{
			DB: db,
		},
	}

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(), // Moved all the route code to a different file (routes.go)
	}

	// log.Printf("Starting server on %s", *addr) // standard log format
	infoLog.Printf("Starting server on %s", *addr) // customised one
	// err := http.ListenAndServe(*addr, mux)  // in place of (srv := http.Server)
	err = srv.ListenAndServe()
	// log.Fatal(err)
	errorLog.Fatal(err)
}

// Wraps sql.Open function and returns a sql.DB connection pool
func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
