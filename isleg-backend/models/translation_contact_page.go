package models

import "github.com/google/uuid"

type TranslationContact struct {
	ID           string        `json:"id,omitempty"`
	LangID       uuid.NullUUID `json:"lang_id,omitempty" binding:"required"`
	FullName     string        `json:"full_name,omitempty" binding:"required"`
	Email        string        `json:"email,omitempty" binding:"required"`
	Phone        string        `json:"phone,omitempty" binding:"required"`
	Letter       string        `json:"letter,omitempty" binding:"required"`
	CompanyPhone string        `json:"company_phone,omitempty" binding:"required"`
	Imo          string        `json:"imo,omitempty" binding:"required"`
	CompanyEmail string        `json:"company_email,omitempty" binding:"required"`
	Instragram   string        `json:"instagram,omitempty" binding:"required"`
	ButtonText   string        `json:"button_text,omitempty" binding:"required"`
	CreatedAt    string        `json:"-"`
	UpdatedAt    string        `json:"-"`
	DeletedAt    string        `json:"-"`
}
