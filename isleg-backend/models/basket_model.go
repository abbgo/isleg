package models

import "github.com/google/uuid"

type Basket struct {
	ID                uuid.UUID `json:"id"`
	ProductID         uuid.UUID `json:"product_id"`
	CustomerID        uuid.UUID `json:"customer_id"`
	QuantityOfProduct uint      `json:"quantity_of_product"`
	CreatedAt         string    `json:"-"`
	UpdatedAt         string    `json:"-"`
	DeletedAt         string    `json:"-"`
}
