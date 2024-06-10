package handlers

import (
	"net/http"
	"time"

	"github.com/TimRobillard/handicap_tracker/views/dashboard"
	"github.com/TimRobillard/handicap_tracker/views/errorViews"
	"github.com/go-chi/chi/v5"
)

func RegisterIndXRoutes(r *chi.Mux) {
	d := chi.NewRouter()
	d.Get("/current", Make(handleCurrent, errorViews.ApiError))
	d.Get("/lowest", Make(handleLowest, errorViews.ApiError))
	d.Get("/highest", Make(handleHighest, errorViews.ApiError))

	r.Mount("/indx", d)
}

func handleCurrent(w http.ResponseWriter, r *http.Request) error {
	idx := make(chan string)
	go func() {

		time.Sleep(time.Second * 3)
		idx <- "20.1"
	}()
	for i := range idx {
		return Render(w, r, dashboard.Card("current", i, "", "green-800", true))
	}
	return nil
}

func handleLowest(w http.ResponseWriter, r *http.Request) error {
	idx := make(chan string)
	go func() {

		time.Sleep(time.Second + time.Millisecond*500)
		idx <- "18.0"
	}()
	for i := range idx {
		return Render(w, r, dashboard.Card("lowest", i, "", "green-400", true))
	}
	return nil
}

func handleHighest(w http.ResponseWriter, r *http.Request) error {
	idx := make(chan string)
	go func() {

		time.Sleep(time.Second * 2)
		idx <- "27.2"
	}()
	for i := range idx {
		return Render(w, r, dashboard.Card("highest", i, "", "red-400", true))
	}
	return nil
}
