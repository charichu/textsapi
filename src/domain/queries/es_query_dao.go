package queries

import (
	"github.com/olivere/elastic"
)

func (q EsQuery) Build() elastic.Query {
	// equals
	query := elastic.NewBoolQuery()

	equalsQueries := make([]elastic.Query, 0)
	for _, eq := range q.Equals {
		equalsQueries = append(equalsQueries, elastic.NewMatchQuery(eq.Field, eq.Value))
	}
	query.Must(equalsQueries...)

	// not equals
	notEqualsQueries := make([]elastic.Query, 0)
	for _, neq := range q.NotEquals {
		notEqualsQueries = append(notEqualsQueries, elastic.NewMatchQuery(neq.Field, neq.Value))
	}
	query.MustNot(notEqualsQueries...)

	// filters
	filterQueries := make([]elastic.Query, 0)
	for _, filter := range q.Filters {
		filterQueries = append(filterQueries, elastic.NewMatchQuery(filter.Field, filter.Value))
	}
	query.Filter(filterQueries...)

	return query
}

func (q EsQuery) BuildMatchQuery() elastic.Query {
	return elastic.NewMultiMatchQuery(q.Match.Term, q.Match.Fields...).Type("phrase_prefix")
}
