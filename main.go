package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/TimRobillard/handicap_tracker/handlers"
	"github.com/TimRobillard/handicap_tracker/handlers/middleware"
	"github.com/TimRobillard/handicap_tracker/store"
	"github.com/TimRobillard/handicap_tracker/views/errorViews"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	pg, err := store.NewPostgresStore()

	if err != nil {
		log.Fatal(err.Error())
	}

	router := chi.NewMux()

	router.Use(middleware.AddUserToContext)
	router.Handle("/public/*", public())

	router.Get("/", handlers.Make(handlers.HandleHome, errorViews.ApiError))

	handlers.RegisterAuthRoutes(router, pg)
	handlers.RegisterCourseRoutes(router, pg)
	handlers.RegisterDashboardRoutes(router, pg)
	handlers.RegisterIndXRoutes(router)
	handlers.RegisterRoundRoutes(router, pg)

	router.Get("/*", handlers.Make(handlers.HandleNotFound, nil))

	listenAddr := os.Getenv("LISTEN_ADDR")
	slog.Info("HTTP server started", "listenAddr", listenAddr)
	if err := http.ListenAndServe(listenAddr, router); err != nil {
		slog.Info("HTTP server error", "msg", err.Error())
	}
}
