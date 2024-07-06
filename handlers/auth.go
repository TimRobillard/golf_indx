package handlers

import (
	"net/http"
	"strings"
	"time"

	"github.com/TimRobillard/handicap_tracker/errors"
	"github.com/TimRobillard/handicap_tracker/handlers/middleware"
	"github.com/TimRobillard/handicap_tracker/store"
	"github.com/TimRobillard/handicap_tracker/views/auth"
	"github.com/TimRobillard/handicap_tracker/views/errorViews"
	"github.com/go-chi/chi/v5"
)

type authHandler struct {
	us store.UserStore
}

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func RegisterAuthRoutes(router *chi.Mux, us store.UserStore) {
	ah := &authHandler{
		us: us,
	}

	router.Get("/login", Make(ah.handleLogin, errorViews.ApiError))
	router.Get("/register", Make(ah.handleRegister, errorViews.ApiError))

	router.Post("/auth/login", Make(ah.handlePostLogin, auth.BadLogin))
	router.Post("/auth/register", Make(ah.handlePostRegister, auth.BadLogin))
}

func (a authHandler) handlePostLogin(w http.ResponseWriter, r *http.Request) error {
	username := r.FormValue("username")
	password := r.FormValue("password")

	user, err := a.us.GetUserByUsername(username)

	if err != nil {
		return errors.UnauthorizedError("unauthorized", nil)
	}

	if ok := user.ValidatePassword(password); !ok {
		return errors.UnauthorizedError("unauthorized", nil)
	}

	token, err := middleware.GenerateToken(user.Id)
	if err != nil {
		return errors.InternalServerError(err)
	}

	cookie := &http.Cookie{
		Value:   token,
		Name:    "_q",
		Path:    "/",
		Expires: time.Now().Add(24 * time.Hour),
	}

	http.SetCookie(w, cookie)
	w.Header().Add("Hx-Redirect", "/dashboard")

	return nil
}

func (a authHandler) handlePostRegister(w http.ResponseWriter, r *http.Request) error {
	username := r.FormValue("username")
	password := r.FormValue("password")

	if len(username) == 0 || len(password) == 0 {
		return errors.BadRequestError("username and password required", nil)
	}

	user, err := a.us.CreateUser(username, password)

	if err != nil {
		if strings.Contains(err.Error(), "duplicate key") {
			return errors.BadRequestError("username taken", nil)
		}
		return errors.InternalServerError(err)
	}

	token, err := middleware.GenerateToken(user.Id)
	if err != nil {
		return errors.InternalServerError(err)
	}

	cookie := &http.Cookie{
		Value:   token,
		Name:    "_q",
		Path:    "/",
		Expires: time.Now().Add(24 * time.Hour),
	}

	http.SetCookie(w, cookie)
	w.Header().Add("Hx-Redirect", "/dashboard")

	return nil
}

func (a authHandler) handleLogin(w http.ResponseWriter, r *http.Request) error {
	return Render(w, r, auth.Login())
}

func (a authHandler) handleRegister(w http.ResponseWriter, r *http.Request) error {
	return Render(w, r, auth.Register())
}
