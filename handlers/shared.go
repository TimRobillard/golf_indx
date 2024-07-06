package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"strings"

	"github.com/TimRobillard/handicap_tracker/errors"
	"github.com/a-h/templ"
)

type HTTPHandler func(w http.ResponseWriter, r *http.Request) error
type ErrorComponent func(e errors.APIError) templ.Component

func Make(h HTTPHandler, c ErrorComponent) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		accepts := r.Header.Get("Accept")
		if strings.Contains(accepts, "application/json") {
			makeJSON(h, w, r)
		} else {
			makeHTMX(h, w, r, c)
		}
	}
}

func makeHTMX(h HTTPHandler, w http.ResponseWriter, r *http.Request, c ErrorComponent) {
	if err := h(w, r); err != nil {
		// TODO get errors from APIError
		if apiErr, ok := err.(errors.APIError); ok {
			w.WriteHeader(apiErr.StatusCode)
			c(apiErr).Render(r.Context(), w)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			c(errors.NewAPIError(http.StatusInternalServerError, err)).Render(r.Context(), w)
		}
		slog.Error("HTTP handler error", "err", err, "path", r.URL.Path)
	}
}

func makeJSON(h HTTPHandler, w http.ResponseWriter, r *http.Request) {
	if err := h(w, r); err != nil {
		if apiErr, ok := err.(errors.APIError); ok {
			writeJSON(w, apiErr.StatusCode, apiErr)
		} else {
			errResp := map[string]any{
				"statusCode": http.StatusInternalServerError,
				"msg":        "internal server error",
			}
			writeJSON(w, http.StatusInternalServerError, errResp)
		}
		slog.Error("HTTP handler error", "err", err, "path", r.URL.Path)
	}
}

func writeJSON(w http.ResponseWriter, status int, v any) error {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}

func Render(w http.ResponseWriter, r *http.Request, c templ.Component) error {
	return c.Render(r.Context(), w)
}
