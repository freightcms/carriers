package web

import (
	"fmt"

	"github.com/freightcms/carriers/db"
	"github.com/freightcms/carriers/db/mongodb"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/language/ast"
)

func getSelectedFields(params graphql.ResolveParams) (map[string]interface{}, error) {
	fieldASTs := params.Info.FieldASTs
	if len(fieldASTs) == 0 {
		return nil, fmt.Errorf("getSelectedFields: ResolveParams has no fields")
	}
	return selectedFieldsFromSelections(params, fieldASTs[0].SelectionSet.Selections)
}

func selectedFieldsFromSelections(params graphql.ResolveParams, selections []ast.Selection) (selected map[string]interface{}, err error) {
	selected = map[string]interface{}{}

	for _, s := range selections {
		switch s := s.(type) {
		case *ast.Field:
			if s.SelectionSet == nil {
				selected[s.Name.Value] = true
			} else {
				selected[s.Name.Value], err = selectedFieldsFromSelections(params, s.SelectionSet.Selections)
				if err != nil {
					return
				}
			}
		case *ast.FragmentSpread:
			n := s.Name.Value
			frag, ok := params.Info.Fragments[n]
			if !ok {
				err = fmt.Errorf("getSelectedFields: no fragment found with name %v", n)

				return
			}
			selected[s.Name.Value], err = selectedFieldsFromSelections(params, frag.GetSelectionSet().Selections)
			if err != nil {
				return
			}
		default:
			err = fmt.Errorf("getSelectedFields: found unexpected selection type %v", s)

			return
		}
	}

	return
}

var (
	CarrierQuery = graphql.NewObject(graphql.ObjectConfig{
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
					fields, err := getSelectedFields(p)
					if err != nil {
						return nil, err
					}

					fmt.Println(fields)
					people, err := mgr.Get(q)
					return people, err
				},
			}, // end people field
		}, // end Fields
	}) // ends object
) // end var
