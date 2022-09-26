package models

import "github.com/google/uuid"

type OrderTimes struct {
	ID          uuid.UUID `json:"id"`
	OrderDateID uuid.UUID `json:"order_date_id"`
	Time        string    `json:"time"`
	CreatedAt   string    `json:"-"`
	UpdatedAt   string    `json:"-"`
	DeletedAt   string    `json:"-"`
}
