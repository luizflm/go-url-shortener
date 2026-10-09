package main

import (
	"html/template"
	"net/http"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	file := "./assets/html/pages/home.tmpl"

	ts, err := template.ParseFiles(file)
	if err != nil {
		app.serverError(w, r, err)
	}

	err = ts.Execute(w, nil)
	if err != nil {
		app.serverError(w, r, err)
	}
}
