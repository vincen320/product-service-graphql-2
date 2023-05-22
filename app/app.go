package app

import (
	"os"

	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/labstack/echo/v4"

	graphqlHandler "github.com/vincen320/product-service-graphql-2/modules/graphql/handler"
	"github.com/vincen320/product-service-graphql-2/modules/graphql/resolver"
	productRepository "github.com/vincen320/product-service-graphql-2/modules/product/repository"
	productUseCase "github.com/vincen320/product-service-graphql-2/modules/product/usecase"
)

func Run() {
	var (
		db                = NewDB()
		productRepository = productRepository.NewProductRepository(db)
		productUseCase    = productUseCase.NewProductUseCase(productRepository)
	)
	schemeFile, err := os.ReadFile("schema.graphql")
	if err != nil {
		panic(err)
	}
	schemeString := string(schemeFile)
	schema := graphql.MustParseSchema(schemeString, &resolver.RootResolver{})
	schemaHandler := &relay.Handler{
		Schema: schema,
	}
	graphqlHandler := graphqlHandler.NewGraphqlHandler(
		schemaHandler.Schema,
		productUseCase,
	)
	e := echo.New()
	v1 := e.Group("v1")
	v1.POST("/graphql", graphqlHandler.GraphQL)
	e.Start(":4000")
}
