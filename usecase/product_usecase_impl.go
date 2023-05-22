package usecase

import (
	"context"

	"github.com/vincen320/product-service-graphql/model"
	"github.com/vincen320/product-service-graphql/repository"
)

type productUseCase struct {
	productRepository repository.ProductRepository
}

func NewProductUseCase(productRepository repository.ProductRepository) ProductUseCase {
	return &productUseCase{
		productRepository: productRepository,
	}
}

func (p *productUseCase) FindAllProducts(ctx context.Context) (response []model.Product, err error) {
	return p.productRepository.FindAllProducts(ctx)
}
