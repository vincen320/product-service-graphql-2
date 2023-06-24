package validator

import (
	"net/http"
	"strings"

	cError "github.com/vincen320/product-service-graphql-2/helper/error"
	productModel "github.com/vincen320/product-service-graphql-2/modules/product/model"
)

func CreateProduct(p productModel.Product) error {
	p.InitType()
	if p.Type == 0 {
		return cError.New(http.StatusInternalServerError, "please specify product type", "error product type validation")
	}
	if p.Name = strings.TrimSpace(p.Name); p.Name == "" {
		return cError.New(http.StatusInternalServerError, "product name cannot be empty", "error product name validation")
	}
	if p.Description = strings.TrimSpace(p.Description); p.Description == "" {
		return cError.New(http.StatusInternalServerError, "product description cannot be empty", "error product description validation")
	}
	if p.Price <= 0 {
		return cError.New(http.StatusInternalServerError, "product price should be bigger than zero", "error product price validation")
	}
	return nil
}
