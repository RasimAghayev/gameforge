package rest

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Router struct {
	*chi.Mux
	logger *slog.Logger
}

func NewRouter(logger *slog.Logger) *Router {
	r := &Router{
		Mux:   chi.NewRouter(),
		logger: logger,
	}
	r.get("/healthz", r.healthz)
	r.get("/readyz", r.readyz)
	return r
}

func (r *Router) get(path string, h http.HandlerFunc) {
	r.Mux.Get(path, h)
}

func (r *Router) healthz(w http.ResponseWriter, req *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

func (r *Router) readyz(w http.ResponseWriter, req *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{"status": "ready"})
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}
