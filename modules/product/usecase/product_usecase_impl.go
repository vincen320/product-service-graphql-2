package usecase

import (
	"context"

	productModel "github.com/vincen320/product-service-graphql-2/modules/product/model"
	prductRepository "github.com/vincen320/product-service-graphql-2/modules/product/repository"
)

type productUseCase struct {
	productRepository prductRepository.ProductRepository
}

func NewProductUseCase(productRepository prductRepository.ProductRepository) ProductUseCase {
	return &productUseCase{
		productRepository: productRepository,
	}
}

func (p *productUseCase) FindAllProducts(ctx context.Context) (response []productModel.Product, err error) {
	return p.productRepository.FindAllProducts(ctx)
}

func (p *productUseCase) CreateProduct(ctx context.Context, request productModel.Product) (response productModel.Product, err error) {
	// Add Validator
	request.InitType()
	return p.productRepository.CreateProduct(ctx, request)
}
