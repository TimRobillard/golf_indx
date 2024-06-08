package handlers

import (
	"net/http"

	"github.com/TimRobillard/handicap_tracker/views/errorViews"
	"github.com/TimRobillard/handicap_tracker/views/home"
)

func HandleHome(w http.ResponseWriter, r *http.Request) error {
	return Render(w, r, home.Index())
}

func HandleNotFound(w http.ResponseWriter, r *http.Request) error {
	return Render(w, r, errorViews.NotFound())
}