package model

import (
	"github.com/graphql-go/graphql"
)

var (
	ProductInterface = graphql.NewInterface(graphql.InterfaceConfig{
		Name:        "product",
		Description: "product is a information of the item, mandatory attribute for all product",
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
		},
		ResolveType: func(p graphql.ResolveTypeParams) *graphql.Object {
			if product, ok := p.Value.(Product); ok {
				switch product.Type {
				case ProductTypeCloth:
					return ClothType
				case ProdcutTypeVehicle:
					return VehicleType
				}
			}
			return nil
		},
	})
)
