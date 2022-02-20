package services

import (
	"github.com/charichu/textsapi/domain/queries"
	"github.com/charichu/textsapi/domain/ruleTexts"
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
	Search(queries.EsQuery) ([]ruleTexts.RuleObjectText, rest_errors.RestErr)
	MatchSearch(queries.EsQuery) ([]ruleTexts.RuleObjectText, rest_errors.RestErr)
}

type ruleTextsService struct{}

func (s *ruleTextsService) Create(ruleTextRequest ruleTexts.RuleObjectText) (*ruleTexts.RuleObjectText, rest_errors.RestErr) {

	if err := ruleTextRequest.Save(); err != nil {
		return nil, err
	}
	return &ruleTextRequest, nil
}

func (s *ruleTextsService) Get(id string) (*ruleTexts.RuleObjectText, rest_errors.RestErr) {
	ruleText := ruleTexts.RuleObjectText{Id: id}

	if err := ruleText.Get(); err != nil {
		return nil, err
	}
	return &ruleText, nil
}

func (s *ruleTextsService) Search(query queries.EsQuery) ([]ruleTexts.RuleObjectText, rest_errors.RestErr) {
	dao := ruleTexts.RuleObjectText{}
	return dao.Search(query)
}

func (s *ruleTextsService) MatchSearch(query queries.EsQuery) ([]ruleTexts.RuleObjectText, rest_errors.RestErr) {
	dao := ruleTexts.RuleObjectText{}
	return dao.MatchSearch(query)
}

func (s *ruleTextsService) Update(ruleText ruleTexts.RuleObjectText) (*ruleTexts.RuleObjectText, rest_errors.RestErr) {
	return nil, rest_errors.NewRestError("implemente me!", http.StatusNotImplemented, "not_implemented", nil)
}

func (s *ruleTextsService) Delete(id string) (*ruleTexts.RuleObjectText, rest_errors.RestErr) {
	return nil, rest_errors.NewRestError("implemente me!", http.StatusNotImplemented, "not_implemented", nil)
}
