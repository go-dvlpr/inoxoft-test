package server

import (
	"github.com/go-chi/chi/v5"
	"inoxoft-test/server/handlers"
)

func BindRoutes(h *handlers.Handlers) *chi.Mux {
	r := chi.NewRouter()

	r.Post("/jobs", h.CreateJob)
	r.Get("/jobs/logs", h.StreamAllLogs)
	r.Get("/jobs/{id}/logs", h.StreamLogs)

	return r
}
