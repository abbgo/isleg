package models

import "github.com/google/uuid"

type PaymentTypes struct {
	ID        uuid.UUID `json:"id,omitempty"`
	LangID    uuid.UUID `json:"lang_id,omitempty"`
	Type      string    `json:"type,omitempty"`
	CreatedAt string    `json:"-"`
	UpdatedAt string    `json:"-"`
	DeletedAt string    `json:"-"`
}
