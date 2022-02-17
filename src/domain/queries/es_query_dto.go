package queries

type EsQuery struct {
	Equals    []FieldValue
	NotEquals []FieldValue
	Filters   []FieldValue
	Match     MatchMessage
}

type FieldValue struct {
	Field string      `json:"field"`
	Value interface{} `json:"value"`
}

type MatchMessage struct {
	Fields []string    `json:"fields"`
	Term   interface{} `json:"term"`
}
