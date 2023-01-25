package models

import "github.com/google/uuid"

type CustomerAddress struct {
	ID         string        `json:"id,omitempty"`
	CustomerID uuid.NullUUID `json:"customer_id,omitempty"`
	Address    string        `json:"address,omitempty"`
	IsActive   bool          `json:"is_active,omitempty"`
	CreatedAt  string        `json:"-"`
	UpdatedAt  string        `json:"-"`
	DeletedAt  string        `json:"-"`
}
