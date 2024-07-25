package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/TimRobillard/handicap_tracker/handlers/middleware"
	"github.com/TimRobillard/handicap_tracker/store"
	"github.com/TimRobillard/handicap_tracker/views/courses"
	"github.com/TimRobillard/handicap_tracker/views/dashboard"
	"github.com/go-chi/chi/v5"
)

type courseHandler struct {
	cs store.CourseStore
	us store.UserStore
}

func RegisterCourseRoutes(router *chi.Mux, pg *store.PostgresStore) {
	r := chi.NewRouter()

	c := courseHandler{
		cs: pg,
		us: pg,
	}

	r.Use(middleware.IsAuthenticated)
	r.Get("/", Make(c.handleGetCourseList, courses.CourseListError))
	r.Get("/score/{id}", Make(c.handleGetCourseForScore, nil))
	r.Get("/test", Make(c.handleTest, nil))

	router.Mount("/courses", r)
}

func (c courseHandler) handleGetCourseList(w http.ResponseWriter, r *http.Request) error {
	keyword := r.FormValue("keyword")
	if keyword == "" {
		w.Header().Add("hx-push-url", "")
		return Render(w, r, courses.CourseList(nil))
	}
	w.Header().Add("hx-push-url", fmt.Sprintf("?keyword=%s", keyword))
	if cs, err := c.cs.SearchCourses(r.Context(), keyword); err != nil {
		return err
	} else {
		return Render(w, r, courses.CourseList(cs))
	}
}

func (c courseHandler) handleGetCourseForScore(w http.ResponseWriter, r *http.Request) error {
	cId := chi.URLParam(r, "id")

	id, err := strconv.Atoi(cId)
	if err != nil {
		return err
	}

	u, err := middleware.GetUserFromRequest(r, c.us)
	if err != nil {
		return err
	}

	course, err := c.cs.GetCourseById(r.Context(), id)
	if err != nil {
		return err
	}

	return Render(w, r, dashboard.ScoreForm(u, course))
}

func (c courseHandler) handleTest(w http.ResponseWriter, r *http.Request) error {
	u, err := middleware.GetUserFromRequest(r, c.us)
	if err != nil {
		return err
	}

	course, err := c.cs.GetCourseById(r.Context(), 1)
	if err != nil {
		return err
	}

	return Render(w, r, dashboard.Test(u, course))
}
