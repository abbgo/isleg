package models

import "github.com/google/uuid"

type Language struct {
	ID        uuid.UUID `json:"id"`
	NameShort string    `json:"name_short"`
	Flag      string    `json:"flag"`
	CreatedAt string    `json:"-"`
	UpdatedAt string    `json:"-"`
	DeletedAt string    `json:"-"`
}
