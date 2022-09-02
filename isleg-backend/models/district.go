package models

import "github.com/google/uuid"

type District struct {
	ID        uuid.UUID `json:"id"`
	Price     float64   `json:"price"`
	CreatedAt string    `json:"-"`
	UpdatedAt string    `json:"-"`
	DeletedAt string    `json:"-"`
}

type TranslationDistrict struct {
	ID         uuid.UUID `json:"id"`
	LangID     uuid.UUID `json:"lang_id"`
	DistrictID uuid.UUID `json:"district_id"`
	CreatedAt  string    `json:"-"`
	UpdatedAt  string    `json:"-"`
	DeletedAt  string    `json:"-"`
}
