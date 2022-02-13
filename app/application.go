package app

import (
	"github.com/charichu/spekulproapi/clients/elastic_search"
	"github.com/charichu/utilsgo/logger"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

var (
	router = mux.NewRouter()
)

func StartApplication() {
	mapUrls()

	logger.Info("starting the application...")
	elastic_search.Init()

	srv := &http.Server{
		Handler:      router,
		Addr:         "localhost:8082",
		WriteTimeout: 500 * time.Millisecond,
		ReadTimeout:  2 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}
