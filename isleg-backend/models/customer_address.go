package models

import (
	"github.com/google/uuid"
)

type CustomerAddress struct {
	ID         uuid.UUID `json:"id,omitempty"`
	CustomerID uuid.UUID `json:"customer_id,omitempty"`
	Address    string    `json:"address,omitempty"`
	IsActive   bool      `json:"is_active,omitempty"`
	CreatedAt  string    `json:"-"`
	UpdatedAt  string    `json:"-"`
	DeletedAt  string    `json:"-"`
}
