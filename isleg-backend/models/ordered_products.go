package models

import "github.com/google/uuid"

type OrderedProducts struct {
	ID                uuid.UUID `json:"id"`
	ProductID         uuid.UUID `json:"product_id"`
	QuantityOfProduct uint      `json:"quantity_of_product"`
	OrderID           uuid.UUID `json:"order_id"`
	UpdatedAt         string    `json:"-"`
	DeletedAt         string    `json:"-"`
}
