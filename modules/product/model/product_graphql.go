package model

import (
	"github.com/graph-gophers/graphql-go"
)

type (
	ProductResolver struct {
		ProductInterface
	}

	ProductInterface interface {
		ID() graphql.ID
		Name() string
		Description() string
		Price() int32
		CreatedBy() int32
		CreatedAt() graphql.Time
		Type() int32
	}
)

// *model.ProductResolver does not resolve "Product": missing method "ToCloth" to convert to "Cloth"
func (p *ProductResolver) ToCloth() (cloth *ClothResolver, ok bool) {
	cloth, ok = p.ProductInterface.(*ClothResolver)
	return
}

// *model.ProductResolver does not resolve "Product": missing method "ToVehicle" to convert to "Vehicle"
func (p *ProductResolver) ToVehicle() (vehicle *VehicleResolver, ok bool) {
	vehicle, ok = p.ProductInterface.(*VehicleResolver)
	return
}
