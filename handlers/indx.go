package handlers

import (
	"github.com/go-chi/chi/v5"
)

func RegisterIndXRoutes(r *chi.Mux) {
	d := chi.NewRouter()

	r.Mount("/indx", d)
}
