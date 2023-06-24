package handler

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/graph-gophers/graphql-go"
	"github.com/labstack/echo/v4"
	cerror "github.com/vincen320/product-service-graphql-2/helper/error"
	"github.com/vincen320/product-service-graphql-2/state"
)

type graphQLHandler struct {
	graphqlSchema *graphql.Schema
}

func NewGraphqlHandler(
	graphqlSchema *graphql.Schema,
) *graphQLHandler {
	return &graphQLHandler{
		graphqlSchema: graphqlSchema,
	}
}

func (g *graphQLHandler) GraphQL(c echo.Context) (err error) {
	var request state.GraphQLRequest
	err = c.Bind(&request)
	if err != nil {
		log.Println(err)
		return state.NewResponse(http.StatusBadRequest, "invalid payload").JSON(c.Response())
	}
	userID := c.Get("user-id").(string)
	userIDInt, _ := strconv.ParseInt(userID, 10, 64)
	ctx := context.WithValue(c.Request().Context(), state.UserID{}, userIDInt)
	response := g.graphqlSchema.Exec(
		ctx,
		request.Query,
		request.OperationName,
		request.Variables,
	)

	if len(response.Errors) > 0 {
		if cerr, ok := cerror.ExtractCustomError(response.Errors[0].Unwrap()); ok {
			log.Println(cerr.GetActualError())
			return state.NewResponse(cerr.GetCode(), cerr.GetErrorMessage()).JSON(c.Response())
		}
		log.Println(response.Errors)
		return state.NewResponse(http.StatusBadRequest, response.Errors).JSON(c.Response())
	}
	return state.NewResponse(http.StatusOK, response.Data).JSON(c.Response())
}
