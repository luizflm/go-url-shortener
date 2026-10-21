package main

import (
	"bytes"
	"html/template"
	"net/http"
)

func (app *application) serverError(w http.ResponseWriter, r *http.Request, err error) {
	var (
		method = r.Method
		uri    = r.URL.RequestURI()
	)

	app.logger.Error(err.Error(), "method", method, "uri", uri)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func (app *application) renderHome(w http.ResponseWriter, status int, data templateData) error {
	ts, err := template.ParseFiles("./assets/html/pages/home.tmpl")
	if err != nil {
		return err
	}

	var body bytes.Buffer
	if err := ts.Execute(&body, data); err != nil {
		return err
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(status)
	_, err = body.WriteTo(w)
	return err
}
