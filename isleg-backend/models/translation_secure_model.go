package models

import "github.com/google/uuid"

type TranslationSecure struct {
	ID        uuid.UUID `json:"id"`
	LangID    uuid.UUID `json:"lang_id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt string    `json:"-"`
	UpdatedAt string    `json:"-"`
	DeletedAt string    `json:"-"`
}
