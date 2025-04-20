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

var (
	Mutations *graphql.Object = graphql.NewObject(graphql.ObjectConfig{
		Name: "CarrierMutations",
		Fields: graphql.Fields{
			"createCarrier": &graphql.Field{
				Type:        IDObject,
				Description: "Create new Carrier",
				Args:        graphql.FieldConfigArgument{},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
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
			},
			"deleteCarrier": &graphql.Field{
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
			},
			"updateCarrier": &graphql.Field{
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
			}, // ends updateCarrier Field type definition
		},
	})
)
