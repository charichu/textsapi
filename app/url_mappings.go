package app

import (
	"github.com/charichu/textsapi/controllers"
	"net/http"
)

func mapUrls() {
	router.HandleFunc("/ping", controllers.PingController.Ping).Methods(http.MethodGet)

	router.HandleFunc("/texts", controllers.RuleTextsController.Create).Methods(http.MethodPost)
	router.HandleFunc("/texts/{id}", controllers.RuleTextsController.Get).Methods(http.MethodGet)
	router.HandleFunc("/texts/search", controllers.RuleTextsController.Search).Methods(http.MethodPost)

}
