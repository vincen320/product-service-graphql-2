package app

import (
	"os"

	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/labstack/echo/v4"

	"github.com/vincen320/product-service-graphql-2/middleware"
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
	schemeFile, err := os.ReadFile("./modules/graphql/schema.graphql")
	if err != nil {
		panic(err)
	}
	schemeString := string(schemeFile)
	rootResolver := resolver.NewRootResolver(productUseCase)
	schema := graphql.MustParseSchema(schemeString, rootResolver)
	schemaHandler := &relay.Handler{
		Schema: schema,
	}
	graphqlHandler := graphqlHandler.NewGraphqlHandler(
		schemaHandler.Schema,
	)
	e := echo.New()
	v1 := e.Group("v1")
	v1.POST("/graphql", graphqlHandler.GraphQL, middleware.JWTAuthentication())
	e.Start(":4000")
}
