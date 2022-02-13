package controllers

import (
	"encoding/json"
	"github.com/charichu/oauthgo/oauth"
	"github.com/charichu/spekulproapi/domain/species"
	"github.com/charichu/spekulproapi/services"
	"github.com/charichu/utilsgo/http_utils"
	"github.com/charichu/utilsgo/rest_errors"
	"io/ioutil"
	"net/http"
)

var (
	SpeciesController speciesControllerInterface = &speciesController{}
)

type speciesControllerInterface interface {
	Create(http.ResponseWriter, *http.Request)
	Get(http.ResponseWriter, *http.Request)
	Update(http.ResponseWriter, *http.Request)
	Delete(http.ResponseWriter, *http.Request)
}

type speciesController struct{}

func (c *speciesController) Create(w http.ResponseWriter, r *http.Request) {
	if err := oauth.AuthenticateRequest(r); err != nil {
		http_utils.RespondError(w, err)
		return
	}

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respErr := rest_errors.NewBadRequestError("invalid request body")
		http_utils.RespondError(w, respErr)
		return
	}
	defer r.Body.Close()

	var specieRequest species.Specie
	if err := json.Unmarshal(requestBody, &specieRequest); err != nil {
		respErr := rest_errors.NewBadRequestError("invalid json body")
		http_utils.RespondError(w, respErr)
		return
	}
	result, createErr := services.SpeciesService.Create(specieRequest)
	if createErr != nil {
		http_utils.RespondError(w, createErr)
		return
	}

	http_utils.RespondJson(w, http.StatusCreated, result)
	return
}

func (c *speciesController) Get(w http.ResponseWriter, r *http.Request) {

}

func (c *speciesController) Update(w http.ResponseWriter, r *http.Request) {

}

func (c *speciesController) Delete(w http.ResponseWriter, r *http.Request) {

}
