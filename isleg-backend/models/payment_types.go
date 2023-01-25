package models

import "github.com/google/uuid"

type PaymentTypes struct {
	ID        string        `json:"id,omitempty"`
	LangID    uuid.NullUUID `json:"lang_id,omitempty" binding:"required"`
	Type      string        `json:"type,omitempty" binding:"required"`
	CreatedAt string        `json:"-"`
	UpdatedAt string        `json:"-"`
	DeletedAt string        `json:"-"`
}
