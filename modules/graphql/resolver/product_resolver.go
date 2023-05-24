package resolver

import (
	"context"
	"net/http"

	"github.com/vincen320/product-service-graphql-2/helper"
	cErr "github.com/vincen320/product-service-graphql-2/helper/error"
	productModel "github.com/vincen320/product-service-graphql-2/modules/product/model"
)

func (q *queryResolver) Product(ctx context.Context) (response []*productModel.ProductResolver, err error) {
	products, err := q.productUseCase.FindAllProducts(ctx)
	response = []*productModel.ProductResolver{}
	for _, product := range products {
		productResolver, err := helper.ConvertPayload[productModel.ProductResolver](product)
		if err != nil {
			err = cErr.New(http.StatusInternalServerError, "internal server error", err.Error())
			return response, err
		}
		response = append(response, &productResolver)
	}
	return
}
