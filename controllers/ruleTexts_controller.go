package controllers

import (
	"encoding/json"
	"github.com/charichu/oauthgo/oauth"
	"github.com/charichu/textsapi/domain/queries"
	"github.com/charichu/textsapi/domain/ruleTexts"
	"github.com/charichu/textsapi/services"
	"github.com/charichu/utilsgo/http_utils"
	"github.com/charichu/utilsgo/rest_errors"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strings"
)

var (
	RuleTextsController ruleTextsControllerInterface = &ruleTextsController{}
)

type ruleTextsControllerInterface interface {
	Create(http.ResponseWriter, *http.Request)
	Get(http.ResponseWriter, *http.Request)
	Update(http.ResponseWriter, *http.Request)
	Delete(http.ResponseWriter, *http.Request)
	Search(http.ResponseWriter, *http.Request)
}

type ruleTextsController struct{}

func (c *ruleTextsController) Create(w http.ResponseWriter, r *http.Request) {
	if err := oauth.AuthenticateRequest(r); err != nil {
		http_utils.RespondError(w, err)
		return
	}
	if oauth.GetCallerId(r) == 0 {
		respErr := rest_errors.NewUnauthorizedError("no machting caller id")
		http_utils.RespondError(w, respErr)
		return
	}

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respErr := rest_errors.NewBadRequestError("invalid request body")
		http_utils.RespondError(w, respErr)
		return
	}
	defer r.Body.Close()

	var ruleTextRequest ruleTexts.RuleObjectText
	if err := json.Unmarshal(requestBody, &ruleTextRequest); err != nil {
		respErr := rest_errors.NewBadRequestError("invalid json body")
		http_utils.RespondError(w, respErr)
		return
	}
	result, createErr := services.RuleTextsService.Create(ruleTextRequest)
	if createErr != nil {
		http_utils.RespondError(w, createErr)
		return
	}

	http_utils.RespondJson(w, http.StatusCreated, result)
	return
}

func (c *ruleTextsController) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	textId := strings.TrimSpace(vars["id"])

	ruleText, err := services.RuleTextsService.Get(textId)
	if err != nil {
		http_utils.RespondError(w, err)
		return
	}
	http_utils.RespondJson(w, http.StatusOK, ruleText)
	return
}

func (c *ruleTextsController) Search(w http.ResponseWriter, r *http.Request) {
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		apiErr := rest_errors.NewBadRequestError("invalid json body")
		http_utils.RespondError(w, apiErr)
		return
	}
	defer r.Body.Close()

	var query queries.EsQuery
	if err := json.Unmarshal(bytes, &query); err != nil {
		apiErr := rest_errors.NewBadRequestError("invalid json body")
		http_utils.RespondError(w, apiErr)
	}

	results, searchErr := services.RuleTextsService.Search(query)
	if searchErr != nil {
		http_utils.RespondError(w, searchErr)
		return
	}
	http_utils.RespondJson(w, http.StatusOK, results)
}

func (c *ruleTextsController) Update(w http.ResponseWriter, r *http.Request) {

}

func (c *ruleTextsController) Delete(w http.ResponseWriter, r *http.Request) {

}
