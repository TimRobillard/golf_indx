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
	cs store.CourseStore
	rs store.RoundStore
	us store.UserStore
}

func RegisterRoundRoutes(router *chi.Mux, pg *store.PostgresStore) {
	r := chi.NewRouter()

	h := &roundHandler{
		cs: pg,
		rs: pg,
		us: pg,
	}

	r.Use(middleware.IsAuthenticated)
	r.Post("/", Make(h.handlePostRound, errorViews.ApiError))

	router.Mount("/rounds", r)
}

func (h roundHandler) handlePostRound(w http.ResponseWriter, r *http.Request) error {
	u, err := middleware.GetUserFromRequest(r, h.us)

	if err != nil {
		return err
	}

	cId := r.FormValue("course_id")
	courseId, err := strconv.Atoi(cId)
	if err != nil {
		return err
	}

	c, err := h.cs.GetCourseById(r.Context(), courseId)
	if err != nil {
		slog.Info(fmt.Sprintf("1.3 course id %s", err.Error()))
		return err
	}

	if !r.PostForm.Has("date") {
		return fmt.Errorf("missing date")
	}
	if !r.PostForm.Has("time") {
		return fmt.Errorf("missing time")
	}
	d := r.FormValue("date")
	time := r.FormValue("time")
	d = d + " " + time + ":00"

	f := [9]int{}
	b := [9]int{}

	var errs []string
	for i := 0; i < 18; i++ {
		key := fmt.Sprintf("hole-%d", i+1)
		s := r.FormValue(key)
		if v, err := strconv.Atoi(s); err != nil {
			errs = []string{"Invalid input"}
			break
		} else {
			if i < 9 {
				f[i] = v
			} else {
				b[i-9] = v
			}
		}

	}

	if len(errs) > 0 {
		return Render(w, r, dashboard.Error(errs))
	}

	_, err = h.rs.CreateRound(r.Context(), u, *c, f, b, d)
	if err != nil {
		return err
	}

	w.Header().Set("Hx-Redirect", "/dashboard")
	return nil
}
