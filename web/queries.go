package web

import (
	"github.com/freightcms/carriers/db"
	"github.com/freightcms/carriers/db/mongodb"
	"github.com/graphql-go/graphql"
)

var (
	CarrierQuery *graphql.Object = graphql.NewObject(graphql.ObjectConfig{
		Name: "CarrierQuery",
		Fields: graphql.Fields{
			"carriers": &graphql.Field{
				Type: graphql.NewList(CarrierObject),
				Args: graphql.FieldConfigArgument{
					"page": &graphql.ArgumentConfig{
						Type:         graphql.Int,
						DefaultValue: 0,
					},
					"pageSize": &graphql.ArgumentConfig{
						Type:         graphql.Int,
						DefaultValue: 10,
					},
					"sortBy": &graphql.ArgumentConfig{
						Type: graphql.NewEnum(graphql.EnumConfig{
							Name:        "SortByFields",
							Description: "Valid field to sort data by in the return",
							Values: graphql.EnumValueConfigMap{
								"id": &graphql.EnumValueConfig{
									Value: "_id",
								},
							},
						}),
						DefaultValue: "_id",
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					mgr := mongodb.FromContext(p.Context)
					q := db.NewQuery()
					if p.Args["page"] != nil {
						q.SetPage(p.Args["page"].(int))
					}
					if p.Args["pageSize"] != nil {
						q.SetPageSize(p.Args["pageSize"].(int))
					}
					if p.Args["sortBy"] != nil {
						q.SetSortBy(p.Args["sortBy"].(string))
					}
					people, err := mgr.Get(q)
					return people, err
				},
			}, // end people field
		}, // end Fields
	}) // ends object
) // end var
