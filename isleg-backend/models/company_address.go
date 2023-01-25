package models

import "github.com/google/uuid"

type CompanyAddress struct {
	ID        string        `json:"id,omitempty"`
	LangID    uuid.NullUUID `json:"lang_id,omitempty" binding:"required"`
	Address   string        `json:"address,omitempty" binding:"required"`
	CreatedAt string        `json:"-"`
	UpdatedAt string        `json:"-"`
	DeletedAt string        `json:"-"`
}
