package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {

	// log.Println(r.URL.Path)
	app.infoLog.Println(r.URL.Path)

	if r.URL.Path != "/" {
		// http.NotFound(w, r)
		app.notFound(w)
		return
	}

	files := []string{
		"./ui/html/home.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		//log.Println(err.Error())
		// app.errorLog.Println(err.Error())
		// http.Error(w, "Internal Server Eroor", 500)
		app.serverError(w, err)
		return
	}

	err = ts.Execute(w, nil)

	if err != nil {
		// log.Println(err.Error())
		// app.errorLog.Println(err.Error())
		// http.Error(w, "Internal Server Eroor", 500)
		app.serverError(w, err)
		return
	}

	// w.Write([]byte("Hello from SnippetBox"))
}

func (app *application) showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil || id < 1 {
		// http.NotFound(w, r)
		app.notFound(w)
		return
	}

	fmt.Fprintf(w, "The requested snippet is %d", id)
	//w.Write([]byte("tHE REQUESTED SNIPPET IS"))
}

func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		// http.Error(w, "Http method not allowed", 405)
		app.clientError(w, 405)
		return
	}

	w.Write([]byte("Creating a Snippet....."))
}

func (c *banda) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("Welcome to the test instance %s", c.Name)))
}

func jandu(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("THE REQUESTED SNIPPET IS JANDWA"))
}
