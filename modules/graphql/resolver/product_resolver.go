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
		switch product.Type {
		case productModel.ProductTypeCloth:
			cloth, err := helper.ConvertPayload[productModel.Cloth](product)
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
			vehicle, err := helper.ConvertPayload[productModel.Vehicle](product)
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
