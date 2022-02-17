package app

import (
	controllers2 "github.com/charichu/textsapi/src/controllers"
	"net/http"
)

func mapUrls() {
	router.HandleFunc("/ping", controllers2.PingController.Ping).Methods(http.MethodGet)

	router.HandleFunc("/texts", controllers2.RuleTextsController.Create).Methods(http.MethodPost)
	router.HandleFunc("/texts/{id}", controllers2.RuleTextsController.Get).Methods(http.MethodGet)
	router.HandleFunc("/texts/search", controllers2.RuleTextsController.Search).Methods(http.MethodPost)
	router.HandleFunc("/texts/match", controllers2.RuleTextsController.MatchSearch).Methods(http.MethodPost)

}
