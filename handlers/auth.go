package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/TimRobillard/handicap_tracker/views/auth"
	"github.com/go-chi/chi/v5"
)

func RegisterAuthRoutes(router *chi.Mux) {
	router.Get("/register", Make(handleRegister))
	router.Post("/register", Make(handlePostRegister))
}

type RegisterRequest struct {
	Username string `json:"username"`
}

func handlePostRegister(w http.ResponseWriter, r *http.Request) error {
	var req RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return InvalidJSON()
	}
	return nil
}

func handleRegister(w http.ResponseWriter, r *http.Request) error {
	return Render(w, r, auth.Register())
}
