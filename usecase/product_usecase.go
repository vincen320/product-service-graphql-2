package usecase

import (
	"context"

	"github.com/vincen320/product-service-graphql/model"
)

type (
	ProductUseCase interface {
		FindAllProducts(ctx context.Context) (response []model.Product, err error)
	}
)
