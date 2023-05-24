package model

type Vehicle struct {
	Product
	Engine string `json:"engine"`
	Wheel  int    `json:"wheel"`
}
