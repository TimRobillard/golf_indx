package handlers

import (
	"net/http"

	"github.com/TimRobillard/handicap_tracker/views/dashboard"
	"github.com/TimRobillard/handicap_tracker/views/errorViews"
	"github.com/go-chi/chi/v5"
)

func RegisterDashboardRoutes(r *chi.Mux) {
	d := chi.NewRouter()
	d.Get("/", Make(handleDashboard, errorViews.ApiError))

	r.Mount("/dashboard", d)
}

func handleDashboard(w http.ResponseWriter, r *http.Request) error {
	return Render(w, r, dashboard.Me())
}
