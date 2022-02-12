package services

import (
	"github.com/charichu/spekulproapi/domain/species"
	"github.com/charichu/utilsgo/rest_errors"
	"net/http"
)

var (
	SpeciesService speciesServiceInterface = &speciesService{}
)

type speciesServiceInterface interface {
	Create(specie species.Specie) (*species.Specie, rest_errors.RestErr)
	Get(string) (*species.Specie, rest_errors.RestErr)
	Update(specie species.Specie) (*species.Specie, rest_errors.RestErr)
	Delete(string) (*species.Specie, rest_errors.RestErr)
}

type speciesService struct{}

func (s *speciesService) Create(specie species.Specie) (*species.Specie, rest_errors.RestErr) {
	return nil, rest_errors.NewRestError("implemente me!", http.StatusNotImplemented, "not_implemented", nil)
}

func (s *speciesService) Get(id string) (*species.Specie, rest_errors.RestErr) {
	return nil, rest_errors.NewRestError("implemente me!", http.StatusNotImplemented, "not_implemented", nil)
}

func (s *speciesService) Update(specie species.Specie) (*species.Specie, rest_errors.RestErr) {
	return nil, rest_errors.NewRestError("implemente me!", http.StatusNotImplemented, "not_implemented", nil)
}

func (s *speciesService) Delete(id string) (*species.Specie, rest_errors.RestErr) {
	return nil, rest_errors.NewRestError("implemente me!", http.StatusNotImplemented, "not_implemented", nil)
}
