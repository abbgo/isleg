package models

import "github.com/google/uuid"

type Orders struct {
	ID           uuid.UUID `json:"id,omitempty"`
	CustomerID   uuid.UUID `json:"customer_id,omitempty"`
	CustomerMark string    `json:"customer_mark,omitempty"`
	OrderTime    string    `json:"order_time,omitempty"`
	PaymentType  string    `json:"payment_type,omitempty"`
	TotalPrice   float64   `json:"total_price,omitempty"`
	OrderNumber  int       `json:"order_number,omitempty"`
	CreatedAt    string    `json:"-"`
	UpdatedAt    string    `json:"-"`
	DeletedAt    string    `json:"-"`
}
