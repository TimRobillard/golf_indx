package handlers

import (
	"net/http"

	"github.com/TimRobillard/handicap_tracker/store"
	"github.com/TimRobillard/handicap_tracker/views/dashboard"
	"github.com/TimRobillard/handicap_tracker/views/errorViews"
	"github.com/go-chi/chi/v5"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func RegisterDashboardRoutes(r *chi.Mux) {
	d := chi.NewRouter()
	d.Get("/", Make(handleDashboard, errorViews.ApiError))

	r.Mount("/dashboard", d)
}

func handleDashboard(w http.ResponseWriter, r *http.Request) error {
	caser := cases.Title(language.English)
	manderley := caser.String("manderley on the green")
	dragonfly := caser.String("drangonfly")
	amberwood := caser.String("amberwood")
	cedarhill := caser.String("cedarhill")
	profile_pic := "https://t4.ftcdn.net/jpg/03/64/21/11/360_F_364211147_1qgLVxv1Tcq0Ohz3FawUfrtONzz8nq3e.jpg"
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
	return Render(w, r, dashboard.Me("20.3", &profile_pic, rounds))
}
