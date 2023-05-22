package app

import (
	"github.com/graphql-go/graphql"
	"github.com/labstack/echo/v4"
	"github.com/vincen320/product-service-graphql/handler"
	"github.com/vincen320/product-service-graphql/repository"
	"github.com/vincen320/product-service-graphql/usecase"
)

func Run() {
	var (
		db                = NewDB()
		productRepository = repository.NewProductRepository(db)
		productUseCase    = usecase.NewProductUseCase(productRepository)
		productGQL        = handler.NewProductGQLHandler(productUseCase)

		queryType = graphql.NewObject(graphql.ObjectConfig{
			Name: "Query",
			Fields: graphql.Fields{
				"product": productGQL.GetAllProducts(),
			},
		})

		mutationType = graphql.NewObject(graphql.ObjectConfig{
			Name:   "Mutation",
			Fields: graphql.Fields{},
		})

		schema, _ = graphql.NewSchema(graphql.SchemaConfig{
			Query:    queryType,
			Mutation: mutationType,
			// https://github.com/graphql-go/graphql/issues/486
			// Types: []graphql.Type{
			// 	model.ClothType, model.VehicleType,
			// },
		})

		graphqlPresenter = handler.NewGraphqlHandler(schema)
	)
	e := echo.New()
	v1 := e.Group("v1")
	v1.POST("/graphql", graphqlPresenter.GraphQL)
	e.Start(":4000")
}
