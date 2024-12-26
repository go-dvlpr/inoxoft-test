package server

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
	"inoxoft-test/config"
	"net/http"
	"time"
)

func Run(cfg *config.Config, router *chi.Mux) error {
	server := &http.Server{
		Addr:              fmt.Sprintf(":%s", cfg.Port),
		ReadTimeout:       60 * time.Minute,
		WriteTimeout:      60 * time.Minute,
		IdleTimeout:       60 * time.Minute,
		ReadHeaderTimeout: 60 * time.Second,
		Handler:           router,
	}

	logrus.Infof("server started on %s port", cfg.Port)

	return server.ListenAndServe()
}
