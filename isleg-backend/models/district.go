package models

import "github.com/google/uuid"

type District struct {
	ID                  string                `json:"id,omitempty"`
	Price               float64               `json:"price,omitempty"`
	CreatedAt           string                `json:"-"`
	UpdatedAt           string                `json:"-"`
	DeletedAt           string                `json:"-"`
	TranslationDistrict []TranslationDistrict `json:"translation_district,omitempty"` // one to many
}

type TranslationDistrict struct {
	ID         string        `json:"id,omitempty"`
	LangID     uuid.NullUUID `json:"lang_id,omitempty"`
	DistrictID uuid.NullUUID `json:"district_id,omitempty"`
	CreatedAt  string        `json:"-"`
	UpdatedAt  string        `json:"-"`
	DeletedAt  string        `json:"-"`
}
