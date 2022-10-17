package models

import "github.com/google/uuid"

type Banner struct {
	ID        uuid.UUID `json:"id,omitempty"`
	Image     string    `json:"image,omitempty"`
	Url       string    `json:"url,omitempty"`
	CreatedAt string    `json:"-"`
	UpdatedAt string    `json:"-"`
	DeletedAt string    `json:"-"`
}
