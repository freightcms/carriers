package web

import (
	"fmt"

	"github.com/freightcms/carriers/db/mongodb"
	"github.com/freightcms/carriers/models"
	commonModels "github.com/freightcms/common/models"
	orgWeb "github.com/freightcms/organizations/web"
	"github.com/graphql-go/graphql"
)

func mapCreateCarrierParams(params graphql.ResolveParams) models.Carrier {
	model := models.Carrier{
		IsActive:     params.Args["isActive"].(bool),
		Insurance:    []*models.InsuranceInfo{},
		Modes:        []commonModels.TransportationMode{},
		Organization: orgWeb.OrganizationFromParams(params),
	}

	return model
}

func mapCarrierIdentifierParams(params graphql.ResolveParams) models.CarrierIdentifier {
	model := models.CarrierIdentifier{
		Type:  params.Args["type"].(string),
		Value: params.Args["value"].(string),
	}

	return model
}

var (
	createCarrierMutation = &graphql.Field{
		Type:        IDObject,
		Description: "Create new Carrier",
		Args:        graphql.FieldConfigArgument{},
		Resolve: func(params graphql.ResolveParams) (any, error) {
			model := mapCreateCarrierParams(params)

			mgr := mongodb.FromContext(params.Context)
			id, err := mgr.CreateCarrier(model)
			if err != nil {
				return nil, err
			}
			resp := struct {
				ID string `json:"id" bson:"id"`
			}{
				ID: id.(string),
			}
			return &resp, err
		},
	}
	deleteCarrierMutation = &graphql.Field{
		Type:        graphql.Boolean,
		Description: "Delete an existing Carrier resource",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			mgr := mongodb.FromContext(params.Context)
			err := mgr.DeleteCarrier(params.Args["id"].(string))
			return err != nil, err
		},
	}
	updateCarrierMutation = &graphql.Field{
		Type:        graphql.Boolean,
		Description: "Update an existing person object",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		}, // ends aarguments
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			mgr := mongodb.FromContext(params.Context)
			id := params.Args["id"]
			p, err := mgr.GetById(id)
			if err != nil {
				return nil, err
			}
			if p == nil {
				return nil, fmt.Errorf("could not find person with ID %s", id)
			}
			carrier := mapCreateCarrierParams(params)
			if err := mgr.UpdateCarrier(id, carrier); err != nil {
				return nil, err
			}
			return true, nil
		}, // end Resolve
	} // ends updateCarrier Field type definition
	createCarrierIdentifierMutation = &graphql.Field{
		Type:        graphql.NewNonNull(graphql.Boolean),
		Description: "Adds an identifier code to the carrier specified",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "Identifier of the carrier to add an identifier to",
			},
		},
		Resolve: func(params graphql.ResolveParams) (any, error) {
			model := mapCarrierIdentifierParams(params)

			if params.Args["id"] == nil {
				return false, nil
			}
			mgr := mongodb.FromContext(params.Context)
			err := mgr.AddIdentifier(params.Args["id"], model)

			return err != nil, err
		},
	}
	deleteCarrierIdentifier = &graphql.Field{
		Type:        graphql.NewNonNull(graphql.Boolean),
		Description: "Deletes an identifier and returns a boolean. True if was deleted, else false if not found",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "ID of the carrier the Identifier belongs to",
			},
			"type": &graphql.ArgumentConfig{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "Type of identifier to delete",
			},
		},
	}
	Mutations *graphql.Object = graphql.NewObject(graphql.ObjectConfig{
		Name: "CarrierMutations",
		Fields: graphql.Fields{
			"createCarrier":    createCarrierMutation,
			"deleteCarrier":    deleteCarrierMutation,
			"updateCarrier":    updateCarrierMutation,
			"createIdentifier": createCarrierIdentifierMutation,
			"deleteIdentifier": deleteCarrierIdentifier,
		},
	})
)
