package web

import (
	"fmt"

	"github.com/freightcms/carriers/db/mongodb"
	"github.com/freightcms/carriers/models"
	"github.com/graphql-go/graphql"
)

var (
	Mutations *graphql.Object = graphql.NewObject(graphql.ObjectConfig{
		Name: "mutations",
		Fields: graphql.Fields{
			"createCarrier": &graphql.Field{
				Type:        IDObject,
				Description: "Create new Carrier",
				Args: graphql.FieldConfigArgument{
					"firstName": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"lastName": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					model := models.Carrier{
						FirstName: params.Args["firstName"].(string),
						LastName:  params.Args["lastName"].(string),
					}

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
					return true, err
				},
			},
			"updateCarrier": &graphql.Field{
				Type:        graphql.Boolean,
				Description: "Update an existing person object",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"firstName": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"lastName": &graphql.ArgumentConfig{
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
					if params.Args["firstName"] != nil {
						p.FirstName = params.Args["firstName"].(string)
					}
					if params.Args["lastName"] != nil {
						p.LastName = params.Args["lastName"].(string)
					}
					if err := mgr.UpdateCarrier(id, *p); err != nil {
						return nil, err
					}
					return true, nil
				}, // end Resolve
			}, // ends updateCarrier Field type definition
		},
	})
)
