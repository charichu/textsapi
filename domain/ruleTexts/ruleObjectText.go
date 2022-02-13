package ruleTexts

type RuleObjectText struct {
	Id             string `json:"id"`
	RuleObjectId   string `json:"rule_object_id"`
	RuleObjectName string `json:"rule_object_name"`
	RuleObjectType string `json:"rule_object_type"`
	TextType       string `json:"text_type"`
	Text           string `json:"text"`
	Language       string `json:"notes"`
}
