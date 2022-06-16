package models

import "github.com/google/uuid"

type Category struct {
	ID               uuid.UUID `json:"id"`
	ParentCategoryID uuid.UUID `json:"parent_category_id"`
	ImagePath        string    `json:"image_path"`
	IsHomeCategory   bool      `json:"is_home_category"`
	CreatedAt        string    `json:"-"`
	UpdatedAt        string    `json:"-"`
	DeletedAt        string    `json:"-"`
}

type TranslationCategory struct {
	ID         uuid.UUID `json:"id"`
	LangID     uuid.UUID `json:"lang_id"`
	CategoryID uuid.UUID `json:"category_id"`
	Name       string    `json:"name"`
	CreatedAt  string    `json:"-"`
	UpdatedAt  string    `json:"-"`
	DeletedAt  string    `json:"-"`
}
