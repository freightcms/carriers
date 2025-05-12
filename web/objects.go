package web

import (
	"github.com/graphql-go/graphql"
)

var (
	InsuranceObject = graphql.NewObject(graphql.ObjectConfig{
		Name: "Insurance",
		Fields: graphql.Fields{
			"policyProvider": &graphql.Field{
				Description: "Whom is the financial backing provider or broker for the insurance",
				Type:        graphql.NewNonNull(graphql.String),
			},
			"policyNumber": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
			},
			"effectiveDate": &graphql.Field{
				Type: graphql.NewNonNull(graphql.DateTime),
			},
			"expirationDate": &graphql.Field{
				Type: graphql.NewNonNull(graphql.DateTime),
			},
			"insuranceType": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
			},
			"insuredAmount": &graphql.Field{
				Type: graphql.NewNonNull(graphql.Float),
			},
		},
	})
	IDObject = graphql.NewObject(graphql.ObjectConfig{
		Name: "ID",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
		},
	})
	CarrierIdentifierObject = graphql.NewObject(graphql.ObjectConfig{
		Name: "CarrierIdentifiers",
		Fields: graphql.Fields{
			"type": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
			},
			"value": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
	})
	CarrierObject = graphql.NewObject(graphql.ObjectConfig{
		Name: "Carrier",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
			},
			"dba": &graphql.Field{
				Description: "Doing Business As can be the alias for a company",
				Type:        graphql.String,
			},
			"isActive": &graphql.Field{
				Description: "Flag for indicating if the carrier is active in the market",
				Type:        graphql.NewNonNull(graphql.Boolean),
			},
			"insurance": &graphql.Field{
				Type: graphql.NewList(InsuranceObject),
			},
			"identifeirs": &graphql.Field{
				Type: graphql.NewList(CarrierIdentifierObject),
			},
		},
	})
)
