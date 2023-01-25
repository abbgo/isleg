package models

import "github.com/google/uuid"

type OrderedProducts struct {
	ID                string        `json:"id,omitempty"`
	ProductID         uuid.NullUUID `json:"product_id,omitempty"`
	QuantityOfProduct uint          `json:"quantity_of_product,omitempty"`
	OrderID           uuid.NullUUID `json:"order_id,omitempty"`
	CreatedAt         string        `json:"-"`
	UpdatedAt         string        `json:"-"`
	DeletedAt         string        `json:"-"`
}
