package models

import "github.com/google/uuid"

type TranslationAbout struct {
	ID        string        `json:"id,omitempty"`
	LangID    uuid.NullUUID `json:"lang_id,omitempty" binding:"required"`
	Title     string        `json:"title,omitempty" binding:"required"`
	Content   string        `json:"content,omitempty" binding:"required"`
	CreatedAt string        `json:"-"`
	UpdatedAt string        `json:"-"`
	DeletedAt string        `json:"-"`
}
