package handler

import (
	"log"
	"net/http"

	"github.com/graph-gophers/graphql-go"
	"github.com/labstack/echo/v4"
	cerror "github.com/vincen320/product-service-graphql-2/helper/error"
	productUseCase "github.com/vincen320/product-service-graphql-2/modules/product/usecase"
	"github.com/vincen320/product-service-graphql-2/state"
)

type graphQLHandler struct {
	graphqlSchema  *graphql.Schema
	productUseCase productUseCase.ProductUseCase
}

func NewGraphqlHandler(
	graphqlSchema *graphql.Schema,
	productUseCase productUseCase.ProductUseCase,
) *graphQLHandler {
	return &graphQLHandler{
		graphqlSchema:  graphqlSchema,
		productUseCase: productUseCase,
	}
}

func (g *graphQLHandler) GraphQL(c echo.Context) (err error) {
	var request state.GraphQLRequest
	err = c.Bind(&request)
	if err != nil {
		log.Println(err)
		return state.NewResponse(http.StatusBadRequest, "invalid payload").JSON(c.Response())
	}

	response := g.graphqlSchema.Exec(
		c.Request().Context(),
		request.Query,
		request.OperationName,
		request.Variables,
	)

	if len(response.Errors) > 0 {
		if cerr, ok := cerror.ExtractCustomError(response.Errors[0].Err); ok {
			log.Println(cerr.GetActualError())
			return state.NewResponse(cerr.GetCode(), cerr.GetErrorMessage()).JSON(c.Response())
		}
		log.Println(response.Errors)
		return state.NewResponse(http.StatusBadRequest, response.Errors).JSON(c.Response())
	}
	return state.NewResponse(http.StatusOK, response.Data).JSON(c.Response())
}
