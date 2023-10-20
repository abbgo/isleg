package models

import "github.com/google/uuid"

type PaymentTypes struct {
	ID        string        `json:"id,omitempty"`
	LangID    uuid.NullUUID `json:"lang_id,omitempty" binding:"required"`
	Name      string        `json:"name,omitempty" binding:"required"`
	Value     uint8         `json:"value,omitempty" binding:"required"`
	CreatedAt string        `json:"-"`
	UpdatedAt string        `json:"-"`
	DeletedAt string        `json:"-"`
}
