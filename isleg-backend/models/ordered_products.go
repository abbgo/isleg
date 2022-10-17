package models

import "github.com/google/uuid"

type OrderedProducts struct {
	ID                uuid.UUID `json:"id,omitempty"`
	ProductID         uuid.UUID `json:"product_id,omitempty"`
	QuantityOfProduct uint      `json:"quantity_of_product,omitempty"`
	OrderID           uuid.UUID `json:"order_id,omitempty"`
	CreatedAt         string    `json:"-"`
	UpdatedAt         string    `json:"-"`
	DeletedAt         string    `json:"-"`
}
