package usecase

import (
	"context"

	productModel "github.com/vincen320/product-service-graphql-2/modules/product/model"
	prductRepository "github.com/vincen320/product-service-graphql-2/modules/product/repository"
	"github.com/vincen320/product-service-graphql-2/state"
	"github.com/vincen320/product-service-graphql-2/validator"
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
	err = validator.CreateProduct(&request)
	if err != nil {
		return
	}
	userID := ctx.Value(state.UserID{}).(int64)
	request.CreatedBy = userID
	return p.productRepository.CreateProduct(ctx, request)
}
