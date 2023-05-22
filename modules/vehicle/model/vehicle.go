package model

import productModel "github.com/vincen320/product-service-graphql-2/modules/product/model"

type Vehicle struct {
	productModel.Product
	Engine string `json:"engine"`
	Wheel  int    `json:"wheel"`
}
