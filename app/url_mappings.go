package app

import (
	"github.com/charichu/spekulproapi/controllers"
	"net/http"
)

func mapUrls() {
	router.HandleFunc("/ping", controllers.PingController.Ping).Methods(http.MethodGet)
	router.HandleFunc("/species", controllers.SpeciesController.Create).Methods(http.MethodPost)
}
