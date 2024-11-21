package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	UserId string `json:"user_id"`
	Items  []Item `json:"items"`
	Status string `json:"status"`
}

type Item struct {
	gorm.Model
	OrderID   int     `json:"order_id"`
	ProductID string  `json:"product_id"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`
}