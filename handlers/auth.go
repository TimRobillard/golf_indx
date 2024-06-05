package handlers

import (
	"net/http"

	"github.com/TimRobillard/handicap_tracker/views/auth"
	"github.com/TimRobillard/handicap_tracker/views/errors"
	"github.com/go-chi/chi/v5"
)

func RegisterAuthRoutes(router *chi.Mux) {
	router.Get("/register", Make(handleRegister, errors.ApiError()))
	router.Post("/register", Make(handlePostRegister, errors.ApiError()))
}

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func handlePostRegister(w http.ResponseWriter, r *http.Request) error {
	return NotImplementedError()
	// username := r.FormValue("username")
	// password := r.FormValue("password")

	// if len(username) ==
	// return nil
}

func handleRegister(w http.ResponseWriter, r *http.Request) error {
	return Render(w, r, auth.Register())
}
