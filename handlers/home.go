package handlers

import (
	"github.com/TimRobillard/handicap_tracker/views/home"
	"net/http"
)

func HandleHome(w http.ResponseWriter, r *http.Request) error {
	return Render(w, r, home.Index())
}