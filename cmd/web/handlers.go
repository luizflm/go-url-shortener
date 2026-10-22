package main

import (
	"errors"
	"net/http"

	"github.com/luizflm/go-url-shortener/internal/base62"
	"github.com/luizflm/go-url-shortener/internal/models"
	"github.com/luizflm/go-url-shortener/internal/validator"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	err := app.renderHome(w, http.StatusOK, templateData{})
	if err != nil {
		app.serverError(w, r, err)
	}
}

func (app *application) shortenUrl(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	longUrl := r.PostForm.Get("url")
	if !validator.NotBlank(longUrl) {
		if err := app.renderHome(w, http.StatusBadRequest, templateData{
			LongURL: longUrl,
			Error:   "O campo não pode ser vazio.",
		}); err != nil {
			app.serverError(w, r, err)
		}
		return
	}

	if !validator.ValidURL(longUrl) {
		if err := app.renderHome(w, http.StatusBadRequest, templateData{
			LongURL: longUrl,
			Error:   "Por favor, informe uma URL válida (http ou https).",
		}); err != nil {
			app.serverError(w, r, err)
		}
		return
	}

	dbId, err := app.urls.GetIdByUrl(r.Context(), longUrl)
	if err != nil {
		switch {
		case errors.Is(err, models.ErrNotFound):
			pId, err := app.urls.Insert(r.Context(), longUrl)
			if err != nil {
				app.serverError(w, r, err)
				return
			}
			dbId = pId
		default:
			app.serverError(w, r, err)
			return
		}
	}

	code := base62.ConvertToBase62(dbId)
	shortUrl := app.baseURL + "/" + code
	if err := app.renderHome(w, http.StatusOK, templateData{ShortURL: shortUrl, LongURL: longUrl}); err != nil {
		app.serverError(w, r, err)
	}
}

func (app *application) handleUrl(w http.ResponseWriter, r *http.Request) {
	code := r.PathValue("code")

	if !validator.Matches(code, validator.Base62RX) {
		http.NotFound(w, r)
		return
	}

	id, err := base62.ConvertToDecimal(code)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	url, err := app.urls.GetById(r.Context(), id)
	if err != nil {
		switch {
		case errors.Is(err, models.ErrNotFound):
			http.NotFound(w, r)
		default:
			app.serverError(w, r, err)
		}
		return
	}

	http.Redirect(w, r, url, http.StatusMovedPermanently)
}
