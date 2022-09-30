package models

import "github.com/google/uuid"

type PaymentTypes struct {
	ID        uuid.UUID `json:"id"`
	LangID    uuid.UUID `json:"lang_id"`
	Type      string    `json:"type"`
	CreatedAt string    `json:"-"`
	UpdatedAt string    `json:"-"`
	DeletedAt string    `json:"-"`
}
