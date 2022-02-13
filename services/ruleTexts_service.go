package services

import (
	"github.com/charichu/spekulproapi/domain/ruleTexts"
	"github.com/charichu/utilsgo/rest_errors"
	"net/http"
)

var (
	RuleTextsService ruleTextsServiceInterface = &ruleTextsService{}
)

type ruleTextsServiceInterface interface {
	Create(specie ruleTexts.RuleObjectText) (*ruleTexts.RuleObjectText, rest_errors.RestErr)
	Get(string) (*ruleTexts.RuleObjectText, rest_errors.RestErr)
	Update(specie ruleTexts.RuleObjectText) (*ruleTexts.RuleObjectText, rest_errors.RestErr)
	Delete(string) (*ruleTexts.RuleObjectText, rest_errors.RestErr)
}

type ruleTextsService struct{}

func (s *ruleTextsService) Create(ruleText ruleTexts.RuleObjectText) (*ruleTexts.RuleObjectText, rest_errors.RestErr) {
	return nil, rest_errors.NewRestError("implemente me!", http.StatusNotImplemented, "not_implemented", nil)
}

func (s *ruleTextsService) Get(id string) (*ruleTexts.RuleObjectText, rest_errors.RestErr) {
	return nil, rest_errors.NewRestError("implemente me!", http.StatusNotImplemented, "not_implemented", nil)
}

func (s *ruleTextsService) Update(ruleText ruleTexts.RuleObjectText) (*ruleTexts.RuleObjectText, rest_errors.RestErr) {
	return nil, rest_errors.NewRestError("implemente me!", http.StatusNotImplemented, "not_implemented", nil)
}

func (s *ruleTextsService) Delete(id string) (*ruleTexts.RuleObjectText, rest_errors.RestErr) {
	return nil, rest_errors.NewRestError("implemente me!", http.StatusNotImplemented, "not_implemented", nil)
}
