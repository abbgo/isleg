package models

import "github.com/google/uuid"

type TranslationFooter struct {
	ID        uuid.UUID `json:"id"`
	LangID    uuid.UUID `json:"lang_id"`
	About     string    `json:"about"`
	Payment   string    `json:"payment"`
	Contact   string    `json:"contact"`
	Secure    string    `json:"secure"`
	Word      string    `json:"word"`
	CreatedAt string    `json:"-"`
	UpdatedAt string    `json:"-"`
	DeletedAt string    `json:"-"`
}
