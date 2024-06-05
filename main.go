package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/TimRobillard/handicap_tracker/handlers"
	"github.com/TimRobillard/handicap_tracker/views/errors"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	router := chi.NewMux()

	router.Handle("/*", public())
	router.Get("/", handlers.Make(handlers.HandleHome, errors.ApiError()))

	handlers.RegisterAuthRoutes(router)

	listenAddr := os.Getenv("LISTEN_ADDR")
	slog.Info("HTTP server started", "listenAddr", listenAddr)
	http.ListenAndServe(listenAddr, router)
}
