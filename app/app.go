package app

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/labstack/echo/v4"
	"golang.org/x/sync/errgroup"

	"github.com/vincen320/product-service-graphql-2/middleware"
	graphqlHandler "github.com/vincen320/product-service-graphql-2/modules/graphql/handler"
	"github.com/vincen320/product-service-graphql-2/modules/graphql/resolver"
	productRepository "github.com/vincen320/product-service-graphql-2/modules/product/repository"
	productUseCase "github.com/vincen320/product-service-graphql-2/modules/product/usecase"
)

func Run() {
	ctx, cancel := context.WithCancel(context.Background())
	db := NewDB()
	productRepository := productRepository.NewProductRepository(db)
	productUseCase := productUseCase.NewProductUseCase(productRepository)
	schemeFile, err := os.ReadFile("./modules/graphql/schema.graphql")
	if err != nil {
		panic(err)
	}
	go func() {
		stop := make(chan os.Signal, 1)
		signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
		<-stop
		cancel()
	}()
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
	eg, egCtx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		err = e.Start(":4000")
		log.Println("successfully gracefully shutdown server")
		return err
	})
	eg.Go(func() error {
		<-egCtx.Done()
		err = db.Close()
		log.Println("DB Successfully closed")
		err = e.Shutdown(context.Background())
		log.Println("HTTP server finish graceful shutdown")
		return err
	})
	if err := eg.Wait(); err != nil {
		log.Println("Exit Reason", err)
	}
	fmt.Println("Clean up")
}
