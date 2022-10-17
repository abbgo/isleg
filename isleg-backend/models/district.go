package models

import "github.com/google/uuid"

type District struct {
	ID        uuid.UUID `json:"id,omitempty"`
	Price     float64   `json:"price,omitempty"`
	CreatedAt string    `json:"-"`
	UpdatedAt string    `json:"-"`
	DeletedAt string    `json:"-"`
}

type TranslationDistrict struct {
	ID         uuid.UUID `json:"id,omitempty"`
	LangID     uuid.UUID `json:"lang_id,omitempty"`
	DistrictID uuid.UUID `json:"district_id,omitempty"`
	CreatedAt  string    `json:"-"`
	UpdatedAt  string    `json:"-"`
	DeletedAt  string    `json:"-"`
}
