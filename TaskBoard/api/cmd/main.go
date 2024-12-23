package main

import (
	"net/http"
	"task_board/internal/config"
	"task_board/internal/postgres"
	"task_board/internal/route"

	"github.com/sirupsen/logrus"
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
