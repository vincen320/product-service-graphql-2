package model

type Cloth struct {
	Product
	Material string `json:"material"`
	Size     string `json:"size"`
}
