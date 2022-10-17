package models

import "github.com/google/uuid"

type Category struct {
	ID               uuid.UUID `json:"id,omitempty"`
	ParentCategoryID uuid.UUID `json:"parent_category_id,omitempty"`
	Image            string    `json:"image,omitempty"`
	IsHomeCategory   bool      `json:"is_home_category,omitempty"`
	CreatedAt        string    `json:"-"`
	UpdatedAt        string    `json:"-"`
	DeletedAt        string    `json:"-"`
}

type TranslationCategory struct {
	ID         uuid.UUID `json:"id,omitempty"`
	LangID     uuid.UUID `json:"lang_id,omitempty"`
	CategoryID uuid.UUID `json:"category_id,omitempty"`
	Name       string    `json:"name,omitempty"`
	CreatedAt  string    `json:"-"`
	UpdatedAt  string    `json:"-"`
	DeletedAt  string    `json:"-"`
}
