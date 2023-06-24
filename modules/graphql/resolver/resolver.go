package resolver

import productUseCase "github.com/vincen320/product-service-graphql-2/modules/product/usecase"

type (
	rootResolver struct {
		productUseCase productUseCase.ProductUseCase
	}
	queryResolver struct {
		productUseCase productUseCase.ProductUseCase
	}
	mutationResolver struct {
		productUseCase productUseCase.ProductUseCase
	}
)

func NewRootResolver(productUseCase productUseCase.ProductUseCase) *rootResolver {
	return &rootResolver{
		productUseCase: productUseCase,
	}
}
func (r *rootResolver) Query() *queryResolver {
	return &queryResolver{
		productUseCase: r.productUseCase,
	}
}

func (r *rootResolver) Mutation() *mutationResolver {
	return &mutationResolver{
		productUseCase: r.productUseCase}
}
