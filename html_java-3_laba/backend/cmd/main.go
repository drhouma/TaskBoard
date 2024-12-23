package main

import (
	"github.com/sirupsen/logrus"
	"html_comments_system/internal/config"
	"html_comments_system/internal/postgres"
	"html_comments_system/internal/route"
	"net/http"
)

func main() {
	cfg, err := config.New("config.json")
	if err != nil {
		logrus.Fatal(err)
	}

	db, err := postgres.New(cfg.PostgresDSN)
	if err != nil {
		logrus.Fatal(err)
	}
	route.Init(db)

	srv := &http.Server{
		Handler: route.New(),
		Addr:    cfg.UpPort,
	}
	if err = srv.ListenAndServe(); err != nil {
		logrus.Fatal(err)
	}

	defer func() {
		if err = db.Close(); err != nil {
			logrus.Error(err)
		}
	}()
}
