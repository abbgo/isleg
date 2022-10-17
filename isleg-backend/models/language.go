package models

import "github.com/google/uuid"

type Language struct {
	ID        uuid.UUID `json:"id,omitempty"`
	NameShort string    `json:"name_short,omitempty"`
	Flag      string    `json:"flag,omitempty"`
	CreatedAt string    `json:"-"`
	UpdatedAt string    `json:"-"`
	DeletedAt string    `json:"-"`
}
