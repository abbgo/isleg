package models

import "github.com/google/uuid"

type Brend struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	ImagePath string    `json:"image_path"`
	CreatedAt string    `json:"-"`
	UpdatedAt string    `json:"-"`
	DeletedAt string    `json:"-"`
}
