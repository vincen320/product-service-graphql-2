package model

import productModel "github.com/vincen320/product-service-graphql-2/modules/product/model"

type Cloth struct {
	productModel.Product
	Material string `json:"material"`
	Size     string `json:"size"`
}
