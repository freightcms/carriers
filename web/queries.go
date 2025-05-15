package web

import (
	"github.com/graphql-go/graphql"
)

var (
	CarrierQuery *graphql.Object = graphql.NewObject(graphql.ObjectConfig{
		Name: "CarrierQuery",
		Fields: graphql.Fields{
			"carriers": &CarriersField,
		}, // end Fields
	}) // ends object
) // end var
