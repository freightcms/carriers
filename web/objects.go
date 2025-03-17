package web

import (
	"github.com/graphql-go/graphql"
)

var (
	IDObject *graphql.Object = graphql.NewObject(graphql.ObjectConfig{
		Name: "ID",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
		},
	})
	CarrierObject *graphql.Object = graphql.NewObject(graphql.ObjectConfig{
		Name: "Carrier",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
			"firstName": &graphql.Field{
				Type: graphql.String,
			},
			"lastName": &graphql.Field{
				Type: graphql.String,
			},
		},
	})
)
