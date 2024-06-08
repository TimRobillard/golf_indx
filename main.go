package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/TimRobillard/handicap_tracker/handlers"
	"github.com/TimRobillard/handicap_tracker/views/errorViews"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	router := chi.NewMux()

	router.Handle("/public/*", public())
	router.Get("/", handlers.Make(handlers.HandleHome, errorViews.ApiError))

	handlers.RegisterAuthRoutes(router)

	router.Get("/*", handlers.Make(handlers.HandleNotFound, nil))

	listenAddr := os.Getenv("LISTEN_ADDR")
	slog.Info("HTTP server started", "listenAddr", listenAddr)
	if err := http.ListenAndServe(listenAddr, router); err != nil {
		slog.Info("HTTP server error", "msg", err.Error())
	}
}
