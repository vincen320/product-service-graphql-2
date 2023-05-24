package model

import "github.com/graph-gophers/graphql-go"

type (
	VehicleResolver struct {
		attr Vehicle
	}
)

func (v *VehicleResolver) Name() string {
	return v.attr.Name
}

func (v *VehicleResolver) Description() string {
	return v.attr.Description
}

func (v *VehicleResolver) Price() int32 {
	return int32(v.attr.Price)
}

func (v *VehicleResolver) CreatedBy() int32 {
	return int32(v.attr.CreatedBy)
}

func (v *VehicleResolver) CreatedAt() graphql.Time {
	var createdAt graphql.Time
	createdAt.UnmarshalGraphQL(v.attr.CreatedAt)
	return createdAt
}

func (v *VehicleResolver) Type() int32 {
	return int32(v.attr.Type)
}

func (v *VehicleResolver) Engine() string {
	return v.attr.Engine
}

func (v *VehicleResolver) Wheel() int32 {
	return int32(v.attr.Wheel)
}
