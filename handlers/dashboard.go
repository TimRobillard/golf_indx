package handlers

import (
	"net/http"

	"github.com/TimRobillard/handicap_tracker/handlers/middleware"
	"github.com/TimRobillard/handicap_tracker/store"
	"github.com/TimRobillard/handicap_tracker/views/dashboard"
	"github.com/TimRobillard/handicap_tracker/views/errorViews"
	"github.com/go-chi/chi/v5"
)

type dashboardHandler struct {
	cs store.CourseStore
	hs store.HandicapStore
	rs store.RoundStore
	us store.UserStore
}

func RegisterDashboardRoutes(r *chi.Mux, pg *store.PostgresStore) {
	d := chi.NewRouter()
	h := &dashboardHandler{
		cs: pg,
		hs: pg,
		rs: pg,
		us: pg,
	}

	d.Use(middleware.IsAuthenticated)
	d.Get("/", Make(h.handleDashboard, errorViews.ApiError))
	d.Get("/score", Make(h.handleScore, errorViews.ApiError))
	d.Get("/chart/me", Make(h.handleChartMe, nil))

	r.Mount("/dashboard", d)
}

func (h dashboardHandler) handleDashboard(w http.ResponseWriter, r *http.Request) error {
	u, err := middleware.GetUserFromRequest(r, h.us)
	if err != nil {
		return err
	}

	rounds, err := h.rs.GetCalcRoundsByUserId(r.Context(), u.Id)
	if err != nil {
		return err
	}
	return Render(w, r, dashboard.Me(u, rounds))
}

func (h dashboardHandler) handleScore(w http.ResponseWriter, r *http.Request) error {
	keyword := r.URL.Query().Get("keyword")

	if keyword != "" {
		recents, err := h.cs.SearchCourses(r.Context(), keyword)
		if err != nil {
			return err
		}

		return Render(w, r, dashboard.ScorePage(recents, keyword))

	}
	return Render(w, r, dashboard.ScorePage(nil, ""))
}

func (h dashboardHandler) handleChartMe(w http.ResponseWriter, r *http.Request) error {
	u, err := middleware.GetUserFromRequest(r, h.us)
	if err != nil {
		return err
	}
	data, err := h.hs.GetChartDataForUser(r.Context(), u.Id, 10)
	if err != nil {
		return err
	}

	return writeJSON(w, http.StatusOK, data)
}
