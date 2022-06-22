package models

import "github.com/google/uuid"

type Banner struct {
	ID        uuid.UUID `json:"id"`
	ImagePath string    `json:"image_path"`
	Url       string    `json:"url"`
	CreatedAt string    `json:"-"`
	UpdatedAt string    `json:"-"`
	DeletedAt string    `json:"-"`
}
