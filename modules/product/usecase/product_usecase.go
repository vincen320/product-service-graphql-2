package usecase

import (
	"context"

	productModel "github.com/vincen320/product-service-graphql-2/modules/product/model"
)

type (
	ProductUseCase interface {
		FindAllProducts(ctx context.Context) (response []productModel.Product, err error)
		CreateProduct(ctx context.Context, request productModel.Product) (response productModel.Product, err error)
	}
)
