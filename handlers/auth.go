package handlers

import (
	"net/http"

	"github.com/TimRobillard/handicap_tracker/errors"
	"github.com/TimRobillard/handicap_tracker/views/auth"
	"github.com/TimRobillard/handicap_tracker/views/errorViews"
	"github.com/go-chi/chi/v5"
)

func RegisterAuthRoutes(router *chi.Mux) {
	router.Get("/login", Make(handleLogin, errorViews.ApiError))
	router.Get("/register", Make(handleRegister, errorViews.ApiError))

	router.Post("/auth/login", Make(handlePostLogin, errorViews.ApiError))
	router.Post("/auth/register", Make(handlePostRegister, auth.BadLogin))
}

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func handlePostLogin(w http.ResponseWriter, r *http.Request) error {
	return errors.NotImplementedError()
}

func handlePostRegister(w http.ResponseWriter, r *http.Request) error {
	username := r.FormValue("username")
	password := r.FormValue("password")

	errorMap := make(map[string]string)

	if len(username) == 0 {
		errorMap["username"] = "username required"
	}
	if len(password) == 0 {
		errorMap["password"] = "password required"
	}

	if len(errorMap) > 0 {
		return errors.BadRequestError("username and password required", errorMap)
	}

	return nil
}

func handleLogin(w http.ResponseWriter, r *http.Request) error {
	return Render(w, r, auth.Login())
}

func handleRegister(w http.ResponseWriter, r *http.Request) error {
	return Render(w, r, auth.Register())
}
