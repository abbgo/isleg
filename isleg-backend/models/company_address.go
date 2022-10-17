package models

import "github.com/google/uuid"

type CompanyAddress struct {
	ID        uuid.UUID `json:"id,omitempty"`
	LangID    uuid.UUID `json:"lang_id,omitempty"`
	Address   string    `json:"address,omitempty"`
	CreatedAt string    `json:"-"`
	UpdatedAt string    `json:"-"`
	DeletedAt string    `json:"-"`
}
