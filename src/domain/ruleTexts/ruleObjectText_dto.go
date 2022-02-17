package ruleTexts

type RuleObjectText struct {
	Id               string `json:"id"`
	RuleObjectId     string `json:"rule_object_id"`
	RuleObjectType   string `json:"rule_object_type"`
	Title            string `json:"title"`
	Description      string `json:"description"`
	ShortDescription string `json:"short_description"`
	Note             string `json:"note"`
	Effect           string `json:"effect"`
	Language         string `json:"language"`
}
