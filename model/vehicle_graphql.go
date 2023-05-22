package model

import (
	"github.com/graphql-go/graphql"
)

var (
	VehicleType *graphql.Object
)

func init() {
	VehicleType = graphql.NewObject(graphql.ObjectConfig{
		Name:        "vehicle",
		Description: "object of vehicle, mandatory attributes for vehicle",
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
			"engine": &graphql.Field{
				Type: graphql.String,
			},
			"wheel": &graphql.Field{
				Type: graphql.Int,
			},
		},
		Interfaces: []*graphql.Interface{
			ProductInterface,
		},
	})
}
