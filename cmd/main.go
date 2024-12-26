package main

import (
	"github.com/sirupsen/logrus"
	"inoxoft-test/config"
	"inoxoft-test/jobs"
	"inoxoft-test/server"
	"inoxoft-test/server/handlers"
)

func main() {
	cfg, err := config.Parse()
	if err != nil {
		logrus.Fatal("failed to load .env file", logrus.WithError(err))
	}

	jobProcessor := jobs.NewJobProcessor()
	h := handlers.New(jobProcessor)

	err = server.Run(cfg, server.BindRoutes(h))
	if err != nil {
		logrus.Fatal("failed to run server", logrus.WithError(err))
	}
}
