package web

import (
	"fmt"

	"github.com/graphql-go/graphql"
	"github.com/squishedfox/webservice-prototype/db/mongodb"
	"github.com/squishedfox/webservice-prototype/models"
)

var (
	Mutations *graphql.Object = graphql.NewObject(graphql.ObjectConfig{
		Name: "mutations",
		Fields: graphql.Fields{
			"createPerson": &graphql.Field{
				Type:        IDObject,
				Description: "Create new Person",
				Args: graphql.FieldConfigArgument{
					"firstName": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"lastName": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					model := models.Person{
						FirstName: params.Args["firstName"].(string),
						LastName:  params.Args["lastName"].(string),
					}

					mgr := mongodb.FromContext(params.Context)
					id, err := mgr.CreatePerson(model)
					if err != nil {
						return nil, err
					}
					return id, err
				},
			},
			"deletePerson": &graphql.Field{
				Type:        graphql.Boolean,
				Description: "Delete an existing Person resource",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					mgr := mongodb.FromContext(params.Context)
					err := mgr.DeletePerson(params.Args["id"].(string))
					return true, err
				},
			},
			"updatePerson": &graphql.Field{
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
					if err := mgr.UpdatePerson(id, *p); err != nil {
						return nil, err
					}
					return true, nil
				}, // end Resolve
			}, // ends updatePerson Field type definition
		},
	})
)
