package web

import (
	"github.com/graphql-go/graphql"
)

func NewSchema() (graphql.Schema, error) {
	return graphql.NewSchema(graphql.SchemaConfig{
		Query:    RootQuery,
		Mutation: Mutations,
	})
}
