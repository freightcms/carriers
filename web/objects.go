package web

import (
	orgWeb "github.com/freightcms/organizations/web"
	"github.com/graphql-go/graphql"
)

var InsuranceObject *graphql.Object = &graphql.NewObject(graphql.ObjectConfig{
    Name: "Insurance",
    Fields: graphql.Fields{
        "policyProvider": &graphql.Field{
            Description: "Whom is the financial backing provider or broker for the insurance",
            Type: &graphql.NewNonNull(graphql.String),
        },
        "policyNumber": &graphql.Field{
            Type: &graphql.NewNonNull(graphql.String),
        },
        "effectiveDate", &graphql.Field{
            Type: &graphql.NewNonNull(graphql.DateTime)
        },
        "expirationDate": &graphql.Field{
            Type: &graphql.NewNonNull(graphql.DateTime)
        },
        "insuranceType": &graphql.Field{
            Type: &graphql.NewNonNull(graphql.String),
        },
        "insuredAmount": &graphql.Field{
            Type: &graphql.NewNonNull(graphql.Float),
        }
    },
})
var IDObject *graphql.Object = graphql.NewObject(graphql.ObjectConfig{
		Name: "ID",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
		},
	})
var CarrierObject *graphql.Object = graphql.NewObject(graphql.ObjectConfig{
		Name: "Carrier",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
			},
            "isActive": &graphql.Field{
                Description: "Flag for indicating if the carrier is active in the market",
                Type: graphql.NewNonNull(graphql.Boolean),
            },
            "insurance" &graphql.NewList(InsuranceObject)
        },
	})
)
