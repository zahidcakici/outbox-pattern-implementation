package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	ProductID uint    `json:"productId"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`
	Status    string  `json:"status"`
}

type Outbox struct {
	gorm.Model
	EventType string
	Payload   string
	Processed bool
}
