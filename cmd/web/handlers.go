package main

import (
	"fmt"
	// "html/template"
	"net/http"
	"strconv"

	"github.com/nadarashwin/learn_go_web/pkg/models"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {

	// log.Println(r.URL.Path)
	app.infoLog.Println(r.URL.Path)

	if r.URL.Path != "/" {
		// http.NotFound(w, r)
		app.notFound(w)
		return
	}

	s, err := app.snippets.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}

	for _, snippet := range s {
		fmt.Fprintf(w, "%v\n", snippet)
	}

	// files := []string{
	// 	"./ui/html/home.page.tmpl",
	// 	"./ui/html/base.layout.tmpl",
	// 	"./ui/html/footer.partial.tmpl",
	// }

	// ts, err := template.ParseFiles(files...)
	// if err != nil {
	// 	//log.Println(err.Error())
	// 	// app.errorLog.Println(err.Error())
	// 	// http.Error(w, "Internal Server Eroor", 500)
	// 	app.serverError(w, err)
	// 	return
	// }

	// err = ts.Execute(w, nil)

	// if err != nil {
	// 	// log.Println(err.Error())
	// 	// app.errorLog.Println(err.Error())
	// 	// http.Error(w, "Internal Server Eroor", 500)
	// 	app.serverError(w, err)
	// 	return
	// }

	// w.Write([]byte("Hello from SnippetBox"))
}

func (app *application) showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil || id < 1 {
		// http.NotFound(w, r)
		app.notFound(w)
		return
	}

	s, err := app.snippets.Get(id)
	if err == models.ErrNoRecord {
		app.notFound(w)
		return
	} else if err != nil {
		app.serverError(w, err)
		return
	}

	fmt.Fprintf(w, "%v", s)
	// fmt.Fprintf(w, "The requested snippet is %d\n", id)
	//w.Write([]byte("tHE REQUESTED SNIPPET IS"))
}

func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		// http.Error(w, "Http method not allowed", 405)
		app.clientError(w, 405)
		return
	}

	// Dummy data
	title := "3test phase thress"
	content := "the start of the test phase three\n follwoed by some\n\n contents"
	expires := "7"

	// pass the data to the SnippetModel.Insert() method, a ID of the new block will be returned.
	id, err := app.snippets.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, err)
	}

	// Redirect the user to the relevant page for the snippet
	http.Redirect(w, r, fmt.Sprintf("/snippet?id=%d", id), http.StatusSeeOther)

	w.Write([]byte("Creating a Snippet....."))
}

func (c *banda) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("Welcome to the test instance %s", c.Name)))
}

func jandu(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("THE REQUESTED SNIPPET IS JANDWA"))
}
