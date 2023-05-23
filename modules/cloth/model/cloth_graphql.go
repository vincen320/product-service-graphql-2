package model

import (
	"github.com/graph-gophers/graphql-go"
)

type (
	ClothResolver struct {
		attr Cloth
	}
)

func (c *ClothResolver) ID() graphql.ID {
	var ID graphql.ID
	ID.UnmarshalGraphQL(int32(c.attr.ID))
	return ID
}

func (c *ClothResolver) Name() string {
	return c.attr.Name
}

func (c *ClothResolver) Description() string {
	return c.attr.Description
}

func (c *ClothResolver) Price() int32 {
	return int32(c.attr.Price)
}

func (c *ClothResolver) CreatedBy() int32 {
	return int32(c.attr.CreatedBy)
}

func (c *ClothResolver) CreatedAt() graphql.Time {
	var createdAt graphql.Time
	createdAt.UnmarshalGraphQL(c.attr.CreatedAt)
	return createdAt
}

func (c *ClothResolver) Type() int32 {
	return int32(c.attr.Type)
}

func (c *ClothResolver) Material() string {
	return c.attr.Material
}

func (c *ClothResolver) Size() string {
	return c.attr.Size
}
