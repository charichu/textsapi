package controllers

import (
	"fmt"
	"github.com/charichu/oauthgo/oauth"
	"github.com/charichu/spekulproapi/domain/species"
	"github.com/charichu/spekulproapi/services"
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
		// TODO return error
		return
	}
	var specie species.Specie
	result, err := services.SpeciesService.Create(specie)
	if err != nil {
		// TODO return error
		return
	}
	fmt.Println(result)
	// TODO return created item with http status 201
}

func (c *speciesController) Get(w http.ResponseWriter, r *http.Request) {

}

func (c *speciesController) Update(w http.ResponseWriter, r *http.Request) {

}

func (c *speciesController) Delete(w http.ResponseWriter, r *http.Request) {

}
