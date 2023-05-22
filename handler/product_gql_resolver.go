package handler

import (
	"github.com/graphql-go/graphql"
	"github.com/vincen320/product-service-graphql/model"
	"github.com/vincen320/product-service-graphql/usecase"
)

type productGQLHandler struct {
	productUseCase usecase.ProductUseCase
}

func NewProductGQLHandler(productUseCase usecase.ProductUseCase) *productGQLHandler {
	return &productGQLHandler{
		productUseCase: productUseCase,
	}
}

func (pr *productGQLHandler) GetAllProducts() *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewList(model.ProductInterface),
		Resolve: func(p graphql.ResolveParams) (response interface{}, err error) {
			return pr.productUseCase.FindAllProducts(p.Context)
		},
	}
}
