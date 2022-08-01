package models

import "github.com/google/uuid"

type Banner struct {
	ID        uuid.UUID `json:"id"`
	Image     string    `json:"image"`
	Url       string    `json:"url"`
	CreatedAt string    `json:"-"`
	UpdatedAt string    `json:"-"`
	DeletedAt string    `json:"-"`
}
