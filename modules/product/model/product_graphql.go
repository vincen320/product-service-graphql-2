package model

import "github.com/graph-gophers/graphql-go"

type (
	ProductResolver struct {
		ProductInterface
	}

	ProductInterface interface {
		ID() graphql.ID
		Name() string
		Description() string
		Price() int64
		CreatedBy() int64
		CreatedAt() graphql.Time
		Type() int
	}
)
