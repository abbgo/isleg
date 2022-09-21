package models

import "github.com/google/uuid"

type Orders struct {
	ID           uuid.UUID `json:"id"`
	CustomerID   uuid.UUID `json:"customer_id"`
	CustomerMark string    `json:"customer_mark"`
	OrderTime    string    `json:"order_time"`
	PaymentType  string    `json:"payment_type"`
	TotalPrice   float64   `json:"total_price"`
	UpdatedAt    string    `json:"-"`
	DeletedAt    string    `json:"-"`
}
