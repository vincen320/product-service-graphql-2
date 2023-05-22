package repository

import (
	"context"

	"github.com/vincen320/product-service-graphql/model"
)

type (
	ProductRepository interface {
		FindAllProducts(ctx context.Context) (response []model.Product, err error)
	}
)
