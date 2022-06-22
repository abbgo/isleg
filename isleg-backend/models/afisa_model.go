package models

import "github.com/google/uuid"

type Afisa struct {
	ID        uuid.UUID `json:"id"`
	ImagePath string    `json:"image_path"`
	CreatedAt string    `json:"-"`
	UpdatedAt string    `json:"-"`
	DeletedAt string    `json:"-"`
}

type TranslationAfisa struct {
	ID          uuid.UUID `json:"id"`
	AfisaID     uuid.UUID `json:"afisa_id"`
	LangID      uuid.UUID `json:"lang_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   string    `json:"-"`
	UpdatedAt   string    `json:"-"`
	DeletedAt   string    `json:"-"`
}
