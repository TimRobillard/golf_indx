package handlers

import (
	"fmt"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/TimRobillard/handicap_tracker/store"
	"github.com/TimRobillard/handicap_tracker/views/dashboard"
	"github.com/go-chi/chi/v5"
)

func RegisterRoundRoutes(router *chi.Mux) {
	r := chi.NewRouter()
	r.Post("/", handlePostRound)

	router.Mount("/rounds", r)
}

func handlePostRound(w http.ResponseWriter, r *http.Request) {
	u := &store.User{
		Id:       1,
		Username: "albatrooss",
	}
	c := &store.Course{
		Id:     1,
		Name:   "Manderley Central North",
		Front:  [9]int{5, 3, 5, 3, 4, 3, 5, 4, 4},
		Back:   [9]int{5, 4, 4, 3, 5, 4, 3, 4, 4},
		Slope:  110,
		Rating: 67.1,
	}
	round := &store.Round{
		Id:     1,
		User:   u,
		Course: c,
		Front:  [9]int{},
		Back:   [9]int{},
	}

	var errors []string
	for i := 0; i < 18; i++ {
		key := fmt.Sprintf("hole-%d", i+1)
		s := r.FormValue(key)
		slog.Info(key + " " + s)
		if v, err := strconv.Atoi(s); err != nil {
			errors = []string{"Invalid input"}
			break
		} else {
			if i < 9 {
				round.Front[i] = v
			} else {
				round.Back[i-9] = v
			}
		}

	}

	if len(errors) > 0 {
		dashboard.Error(errors).Render(r.Context(), w)
		return
	}

	w.Header().Set("Hx-Redirect", "/dashboard")
	w.Write(nil)
}
