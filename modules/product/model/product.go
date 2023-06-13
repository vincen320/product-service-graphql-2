package model

import (
	"strconv"
	"time"
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
	}
	if materialOk || sizeOk {
		p.Type = ProductTypeCloth
	}
}
