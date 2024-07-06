package handlers

import (
	"net/http"

	"github.com/TimRobillard/handicap_tracker/handlers/middleware"
	"github.com/TimRobillard/handicap_tracker/store"
	"github.com/TimRobillard/handicap_tracker/views/dashboard"
	"github.com/TimRobillard/handicap_tracker/views/errorViews"
	"github.com/go-chi/chi/v5"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type dashboardHandler struct {
	us store.UserStore
	cs store.CourseStore
	rs store.RoundStore
}

func RegisterDashboardRoutes(r *chi.Mux, pg *store.PostgresStore) {
	d := chi.NewRouter()
	h := &dashboardHandler{
		us: pg,
		cs: pg,
		rs: pg,
	}

	d.Use(middleware.JwtAuth)
	d.Get("/", Make(h.handleDashboard, errorViews.ApiError))
	d.Get("/score", Make(handleScore, errorViews.ApiError))
	d.Get("/chart/me", Make(handleChartMe, nil))

	r.Mount("/dashboard", d)
}

func (h dashboardHandler) handleDashboard(w http.ResponseWriter, r *http.Request) error {
	u, err := middleware.GetUserFromRequest(r, h.us)
	if err != nil {
		return err
	}
	caser := cases.Title(language.English)
	manderley := caser.String("manderley on the green")
	dragonfly := caser.String("drangonfly")
	amberwood := caser.String("amberwood")
	cedarhill := caser.String("cedarhill")

	rounds := [20]store.CalcRound{
		{Score: "102", Course: manderley, TimeAgo: "2 days ago"},
		{Score: "92", Course: dragonfly, TimeAgo: "2 days ago"},
		{Score: "96", Course: manderley, TimeAgo: "2 days ago"},
		{Score: "96", Course: cedarhill, TimeAgo: "2 days ago"},
		{Score: "34", Course: amberwood, TimeAgo: "2 days ago"},
		{Score: "101", Course: "Some fancy Golf and Country Club", TimeAgo: "2 days ago"},
		{Score: "100", Course: cedarhill, TimeAgo: "2 days ago"},
		{Score: "98", Course: cedarhill, TimeAgo: "2 days ago"},
		{Score: "94", Course: manderley, TimeAgo: "2 days ago"},
		{Score: "93", Course: manderley, TimeAgo: "2 days ago"},
		{Score: "92", Course: manderley, TimeAgo: "2 days ago"},
		{Score: "102", Course: manderley, TimeAgo: "2 days ago"},
		{Score: "92", Course: dragonfly, TimeAgo: "2 days ago"},
		{Score: "96", Course: manderley, TimeAgo: "2 days ago"},
		{Score: "96", Course: cedarhill, TimeAgo: "2 days ago"},
		{Score: "34", Course: amberwood, TimeAgo: "2 days ago"},
		{Score: "101", Course: manderley, TimeAgo: "2 days ago"},
		{Score: "100", Course: cedarhill, TimeAgo: "2 days ago"},
		{Score: "98", Course: cedarhill, TimeAgo: "2 days ago"},
		{Score: "94", Course: manderley, TimeAgo: "2 days ago"},
	}
	return Render(w, r, dashboard.Me(u, rounds))
}

func handleScore(w http.ResponseWriter, r *http.Request) error {
	recents :=
		[]*store.UICourse{
			{Id: 1, Name: "Manderley On The Green North South", Thumbnail: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRYsa5s9fz-agjYOZtBTJDSaDV_78gxOiRTQw&usqp=CAU", Par: "72"},
			{Id: 2, Name: "Amberwood Golf Club", Thumbnail: "https://stittsvillecentral.ca/wp-content/uploads/amberwood-village-golf-green.jpg", Par: "32"},
		}

	return Render(w, r, dashboard.ScorePage(recents))
}

type Data struct {
	Labels []string  `json:"labels"`
	Data   []float32 `json:"data"`
	Min    int       `json:"min"`
	Max    int       `json:"max"`
}

func handleChartMe(w http.ResponseWriter, _r *http.Request) error {
	data := Data{
		Labels: []string{"", "May", "June", "", "July", ""},
		Data:   []float32{20.1, 20.3, 20.0, 19.2, 18.9, 21},
		Min:    18,
		Max:    22,
	}
	return writeJSON(w, http.StatusOK, data)
}
