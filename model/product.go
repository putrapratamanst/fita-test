package model

type Add struct {
	SKU      string `json:"sku"`
	Name     string `json:"name"`
	Price    string `json:"price"`
	Quantity int    `json:"quantity"`
}

type Cart struct {
	SKU      string `json:"sku" validate:"required"`
	Quantity int    `json:"quantity" validate:"required"`
}
