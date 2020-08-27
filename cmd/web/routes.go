package main

import "net/http"

type banda struct {
	Name string
}

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet", app.showSnippet)
	mux.HandleFunc("/snippet/create", app.createSnippet)

	fileserver := http.FileServer(http.Dir("./ui/static"))

	mux.Handle("/static/", http.StripPrefix("/static", fileserver))

	s := &banda{
		Name: "Roger",
	}

	mux.Handle("/test", s)

	mux.Handle("/pest", http.HandlerFunc(jandu))

	return mux

}
