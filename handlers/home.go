package handlers

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/TimRobillard/handicap_tracker/handlers/middleware"
	"github.com/TimRobillard/handicap_tracker/views/errorViews"
	"github.com/TimRobillard/handicap_tracker/views/home"
)

func HandleHome(w http.ResponseWriter, r *http.Request) error {
	userId := r.Context().Value(middleware.ContextId)
	slog.Info(fmt.Sprintf("%x", userId))
	if userId != nil {
		http.Redirect(w, r, "/dashboard", http.StatusTemporaryRedirect)
		return nil
	}
	return Render(w, r, home.Index())
}

func HandleNotFound(w http.ResponseWriter, r *http.Request) error {
	return Render(w, r, errorViews.NotFound())
}
