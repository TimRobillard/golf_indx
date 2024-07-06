package handlers

import (
	"fmt"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/TimRobillard/handicap_tracker/handlers/middleware"
	"github.com/TimRobillard/handicap_tracker/store"
	"github.com/TimRobillard/handicap_tracker/views/dashboard"
	"github.com/TimRobillard/handicap_tracker/views/errorViews"
	"github.com/go-chi/chi/v5"
)

type roundHandler struct {
	us store.UserStore
	rs store.RoundStore
}

func RegisterRoundRoutes(router *chi.Mux, pg *store.PostgresStore) {
	r := chi.NewRouter()

	h := &roundHandler{
		us: pg,
		rs: pg,
	}

	r.Post("/", Make(h.handlePostRound, errorViews.ApiError))

	router.Mount("/rounds", r)
}

func (h roundHandler) handlePostRound(w http.ResponseWriter, r *http.Request) error {
	u, err := middleware.GetUserFromRequest(r, h.us)

	if err != nil {
		return err
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
		Course: *c,
		Front:  [9]int{},
		Back:   [9]int{},
	}

	var errs []string
	for i := 0; i < 18; i++ {
		key := fmt.Sprintf("hole-%d", i+1)
		s := r.FormValue(key)
		slog.Info(key + " " + s)
		if v, err := strconv.Atoi(s); err != nil {
			errs = []string{"Invalid input"}
			break
		} else {
			if i < 9 {
				round.Front[i] = v
			} else {
				round.Back[i-9] = v
			}
		}

	}

	if len(errs) > 0 {
		dashboard.Error(errs).Render(r.Context(), w)
		return nil
	}

	w.Header().Set("Hx-Redirect", "/dashboard")
	return Render(w, r, nil)
}
