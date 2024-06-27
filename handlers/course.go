package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/TimRobillard/handicap_tracker/store"
	"github.com/TimRobillard/handicap_tracker/views/courses"
	"github.com/TimRobillard/handicap_tracker/views/dashboard"
	"github.com/go-chi/chi/v5"
)

type courseHandler struct {
	cs store.CourseStore
}

func RegisterCourseRoutes(router *chi.Mux, cs store.CourseStore) {
	r := chi.NewRouter()

	c := courseHandler{}
	c.cs = cs

	r.Get("/", Make(c.handleGetCourseList, courses.CourseListError))
	r.Get("/score/{id}", Make(c.handleGetCourseForScore, nil))

	router.Mount("/courses", r)
}

func (c courseHandler) handleGetCourseList(w http.ResponseWriter, r *http.Request) error {
	keyword := r.FormValue("keyword")
	w.Header().Add("hx-push-url", fmt.Sprintf("?keyword=%s", keyword))
	if cs, err := c.cs.SearchCourses(keyword); err != nil {
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
	u := store.User{
		Id:       1,
		Username: "albatrooss",
	}
	course := store.Course{
		Id:     id,
		Name:   "Manderley Central North",
		Front:  [9]int{5, 3, 5, 3, 4, 3, 5, 4, 4},
		Back:   [9]int{5, 4, 4, 3, 5, 4, 3, 4, 4},
		Slope:  110,
		Rating: 67.1,
	}

	return Render(w, r, dashboard.ScoreForm(&u, &course))
}
