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
	profile_pic := "https://t4.ftcdn.net/jpg/03/64/21/11/360_F_364211147_1qgLVxv1Tcq0Ohz3FawUfrtONzz8nq3e.jpg"
	return Render(w, r, dashboard.Me("20.3", &profile_pic))
}
