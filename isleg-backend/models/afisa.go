package models

import (
	"github.com/google/uuid"
)

type Afisa struct {
	ID               string             `json:"id,omitempty"`
	Image            string             `json:"image,omitempty"`
	CreatedAt        string             `json:"-"`
	UpdatedAt        string             `json:"-"`
	DeletedAt        string             `json:"-"`
	TranslationAfisa []TranslationAfisa `json:"translation_afisa,omitempty"` // one to many
}

type TranslationAfisa struct {
	ID          string        `json:"id,omitempty"`
	AfisaID     uuid.NullUUID `json:"afisa_id,omitempty"`
	LangID      uuid.NullUUID `json:"lang_id,omitempty" binding:"required"`
	Title       string        `json:"title,omitempty" binding:"required"`
	Description string        `json:"description,omitempty" binding:"required"`
	CreatedAt   string        `json:"-"`
	UpdatedAt   string        `json:"-"`
	DeletedAt   string        `json:"-"`
}
