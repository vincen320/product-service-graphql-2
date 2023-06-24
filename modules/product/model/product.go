package model

import (
	"strconv"
	"time"

	"github.com/vincen320/product-service-graphql-2/helper"
)

const (
	ProductTypeCloth   = 1
	ProdcutTypeVehicle = 2
)

type (
	Product struct {
		ID             int64                  `json:"id"`
		Name           string                 `json:"name"`
		Description    string                 `json:"description"`
		Price          int64                  `json:"price"`
		CreatedBy      int64                  `json:"created_by"`
		CreatedAt      time.Time              `json:"created_at"`
		Type           int                    `json:"type"`
		AdditionalAttr map[string]interface{} `json:"-"`
	}
	ProductInput struct {
		Name        string  `json:"name"`
		Description string  `json:"description"`
		Price       float64 `json:"price"`
		Material    *string `json:"material"`
		Size        *string `json:"size"`
		Engine      *string `json:"engine"`
		Wheel       *int    `json:"wheel"`
	}
)

func (p Product) ToCloth() Cloth {
	return Cloth{
		Product:  p,
		Material: p.AdditionalAttr["material"].(string),
		Size:     p.AdditionalAttr["size"].(string),
	}
}

func (p Product) ToVehicle() Vehicle {
	wheelStr, _ := p.AdditionalAttr["wheel"].(string)
	wheel, _ := strconv.Atoi(wheelStr)
	return Vehicle{
		Product: p,
		Engine:  p.AdditionalAttr["engine"].(string),
		Wheel:   wheel,
	}
}

func (p *Product) InitType() {
	_, wheelOk := p.AdditionalAttr["wheel"]
	_, engineOk := p.AdditionalAttr["engine"]
	_, materialOk := p.AdditionalAttr["material"]
	_, sizeOk := p.AdditionalAttr["size"]

	if wheelOk || engineOk {
		p.Type = ProdcutTypeVehicle
	} else if materialOk || sizeOk {
		p.Type = ProductTypeCloth
	}
}

func (p ProductInput) ToProduct() Product {
	product, _ := helper.ConvertPayload[Product](p)
	product.AdditionalAttr = map[string]interface{}{}
	if p.Wheel != nil {
		product.AdditionalAttr["wheel"] = p.Wheel
	}
	if p.Engine != nil {
		product.AdditionalAttr["engine"] = p.Engine
		return product
	}
	if p.Material != nil {
		product.AdditionalAttr["material"] = p.Material
	}
	if p.Size != nil {
		product.AdditionalAttr["size"] = p.Size
	}
	return product
}
