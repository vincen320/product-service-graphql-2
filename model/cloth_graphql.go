package model

import (
	"github.com/graphql-go/graphql"
)

var (
	ClothType *graphql.Object
)

func init() {
	ClothType = graphql.NewObject(graphql.ObjectConfig{
		Name:        "cloth",
		Description: "object of cloth, mandatory attributes for cloth",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"description": &graphql.Field{
				Type: graphql.String,
			},
			"price": &graphql.Field{
				Type: graphql.Int,
			},
			"created_by": &graphql.Field{
				Type: graphql.Int,
			},
			"created_at": &graphql.Field{
				Type: graphql.DateTime,
			},
			"material": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if _, ok := p.Source.(Product); ok {
						return "query", nil
					}
					return "", nil
				},
			},
			"size": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if _, ok := p.Source.(Product); ok {
						return "query", nil
					}
					return "", nil
				},
			},
		},
		Interfaces: []*graphql.Interface{
			ProductInterface,
		},
	})
}
