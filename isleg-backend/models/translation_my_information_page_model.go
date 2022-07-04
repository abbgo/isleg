package models

import (
	"github.com/google/uuid"
)

type TranslationMyInformationPage struct {
	ID        uuid.UUID `json:"id"`
	LangID    uuid.UUID `json:"lang_id"`
	Birthday  string    `json:"birthday"`
	Address   string    `json:"address"`
	CreatedAt string    `json:"-"`
	UpdatedAt string    `json:"-"`
	DeletedAt string    `json:"-"`
}
