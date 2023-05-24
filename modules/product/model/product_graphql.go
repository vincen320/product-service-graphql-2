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
