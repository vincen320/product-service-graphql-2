package resolver

import (
	"context"
	"net/http"

	cErr "github.com/vincen320/product-service-graphql-2/helper/error"
	productModel "github.com/vincen320/product-service-graphql-2/modules/product/model"
)

func (q *queryResolver) Product(ctx context.Context) (response []*productModel.ProductResolver, err error) {
	products, err := q.productUseCase.FindAllProducts(ctx)
	if err != nil {
		return
	}
	response = []*productModel.ProductResolver{}
	for _, product := range products {
		switch product.Type {
		case productModel.ProductTypeCloth:
			cloth := product.ToCloth()
			var clothResolver productModel.ClothResolver
			clothResolver.SetAttr(cloth)
			if err != nil {
				err = cErr.New(http.StatusInternalServerError, "internal server error", err.Error())
				return response, err
			}
			response = append(response, &productModel.ProductResolver{
				ProductInterface: &clothResolver,
			})
		case productModel.ProdcutTypeVehicle:
			vehicle := product.ToVehicle()
			var vehicleResolver productModel.VehicleResolver
			vehicleResolver.SetAttr(vehicle)
			if err != nil {
				err = cErr.New(http.StatusInternalServerError, "internal server error", err.Error())
				return response, err
			}
			response = append(response, &productModel.ProductResolver{
				ProductInterface: &vehicleResolver,
			})
		}
	}
	return
}

func (m *mutationResolver) CreateProduct(ctx context.Context, args struct{ Product productModel.ProductInput }) (response *productModel.ProductResolver, err error) {
	product, err := m.productUseCase.CreateProduct(ctx, args.Product.ToProduct())
	if err != nil {
		return
	}
	switch product.Type {
	case productModel.ProductTypeCloth:
		cloth := product.ToCloth()
		var clothResolver productModel.ClothResolver
		clothResolver.SetAttr(cloth)
		if err != nil {
			err = cErr.New(http.StatusInternalServerError, "internal server error", err.Error())
			return response, err
		}
		response = &productModel.ProductResolver{
			ProductInterface: &clothResolver,
		}
	case productModel.ProdcutTypeVehicle:
		vehicle := product.ToVehicle()
		var vehicleResolver productModel.VehicleResolver
		vehicleResolver.SetAttr(vehicle)
		if err != nil {
			err = cErr.New(http.StatusInternalServerError, "internal server error", err.Error())
			return response, err
		}
		response = &productModel.ProductResolver{
			ProductInterface: &vehicleResolver,
		}
	}
	return
}
