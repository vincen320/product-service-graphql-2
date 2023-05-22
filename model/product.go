package model

import "time"

const (
	ProductTypeCloth   = 1
	ProdcutTypeVehicle = 2
)

type (
	Product struct {
		ID          int64     `json:"id"`
		Name        string    `json:"name"`
		Description string    `json:"description"`
		Price       int64     `json:"price"`
		CreatedBy   int64     `json:"created_by"`
		CreatedAt   time.Time `json:"created_at"`
		Type        int       `json:"type"`
	}
)
