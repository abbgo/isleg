package models

import "github.com/google/uuid"

type Orders struct {
	ID            string        `json:"id,omitempty"`
	CustomerID    uuid.NullUUID `json:"customer_id,omitempty"`
	CustomerMark  string        `json:"customer_mark,omitempty"`
	OrderTime     string        `json:"order_time,omitempty"`
	PaymentType   string        `json:"payment_type,omitempty"`
	TotalPrice    float64       `json:"total_price,omitempty"`
	OrderNumber   int           `json:"order_number,omitempty"`
	ShippingPrice float64       `json:"shipping_price,omitempty"`
	Excel         float64       `json:"excel,omitempty"`
	Address       string        `json:"address,omitempty"`
	CreatedAt     string        `json:"-"`
	UpdatedAt     string        `json:"-"`
	DeletedAt     string        `json:"-"`
}
