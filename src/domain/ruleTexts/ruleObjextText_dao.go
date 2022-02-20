package ruleTexts

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/charichu/textsapi/clients/elastic_search"
	"github.com/charichu/textsapi/domain/queries"
	"github.com/charichu/utilsgo/rest_errors"
	"strings"
)

const (
	indexRuleTexts  = "texts"
	docTypeRuleText = "ruleText"
)

func (r *RuleObjectText) Save() rest_errors.RestErr {
	result, err := elastic_search.Client.Index(indexRuleTexts, docTypeRuleText, r)
	if err != nil {
		return rest_errors.NewInternalServerError("error when trying to save ruleText", errors.New("db error"))
	}
	r.Id = result.Id
	return nil
}

func (r *RuleObjectText) Get() rest_errors.RestErr {
	textId := r.Id
	result, err := elastic_search.Client.Get(indexRuleTexts, docTypeRuleText, r.Id)
	if err != nil {
		if strings.Contains(err.Error(), "404") {
			return rest_errors.NewNotFoundError(fmt.Sprintf("no item found with id %s", r.Id))
		}
		return rest_errors.NewInternalServerError(fmt.Sprintf("error when trying to get id %s", r.Id), errors.New("db error"))
	}

	bytes, err := result.Source.MarshalJSON()
	if err != nil {
		return rest_errors.NewInternalServerError("error when trying to parse db response", errors.New("db error"))
	}
	if err := json.Unmarshal(bytes, &r); err != nil {
		return rest_errors.NewInternalServerError("error when trying to parse db response", errors.New("db error"))
	}
	r.Id = textId
	return nil
}

func (r *RuleObjectText) Search(query queries.EsQuery) ([]RuleObjectText, rest_errors.RestErr) {
	result, err := elastic_search.Client.Search(indexRuleTexts, query.Build())
	fmt.Println(query.Build())
	if err != nil {
		return nil, rest_errors.NewInternalServerError("error when trying to search documents", errors.New("db error"))
	}
	results := make([]RuleObjectText, result.TotalHits())
	for index, hit := range result.Hits.Hits {
		bytes, _ := hit.Source.MarshalJSON()
		var result RuleObjectText
		if err := json.Unmarshal(bytes, &result); err != nil {
			return nil, rest_errors.NewInternalServerError("error when trying to parse response", errors.New("db error"))
		}
		result.Id = hit.Id
		results[index] = result
	}
	if len(results) == 0 {
		return nil, rest_errors.NewNotFoundError("no result found matching given criteria")
	}

	return results, nil
}

func (r *RuleObjectText) MatchSearch(query queries.EsQuery) ([]RuleObjectText, rest_errors.RestErr) {
	result, err := elastic_search.Client.Search(indexRuleTexts, query.BuildMatchQuery())
	if err != nil {
		return nil, rest_errors.NewInternalServerError("error when trying to search documentsby match", errors.New("db error"))
	}
	results := make([]RuleObjectText, result.TotalHits())
	for index, hit := range result.Hits.Hits {
		bytes, _ := hit.Source.MarshalJSON()
		var result RuleObjectText
		if err := json.Unmarshal(bytes, &result); err != nil {
			return nil, rest_errors.NewInternalServerError("error when trying to parse response of matching documents", errors.New("db error"))
		}
		result.Id = hit.Id
		results[index] = result
	}
	if len(results) == 0 {
		return nil, rest_errors.NewNotFoundError("no result found matching given criteria")
	}

	return results, nil
}
