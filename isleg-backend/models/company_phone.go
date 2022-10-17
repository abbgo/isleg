package models

import "github.com/google/uuid"

type CompanyPhone struct {
	ID        uuid.UUID `json:"id,omitempty"`
	Phone     string    `json:"phone,omitempty"`
	CreatedAt string    `json:"-"`
	UpdatedAt string    `json:"-"`
	DeletedAt string    `json:"-"`
}
