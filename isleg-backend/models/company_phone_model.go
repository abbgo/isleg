package models

import "github.com/google/uuid"

type CompanyPhone struct {
	ID        uuid.UUID `json:"id"`
	Phone     string    `json:"phone"`
	CreatedAt string    `json:"-"`
	UpdatedAt string    `json:"-"`
	DeletedAt string    `json:"-"`
}
