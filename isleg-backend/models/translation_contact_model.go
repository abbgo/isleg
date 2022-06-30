package models

import "github.com/google/uuid"

type TranslationContact struct {
	ID           uuid.UUID `json:"id"`
	LangID       uuid.UUID `json:"lang_id"`
	FullName     string    `json:"full_name"`
	Email        string    `json:"email"`
	Phone        string    `json:"phone"`
	Letter       string    `json:"letter"`
	CompanyPhone string    `json:"company_phone"`
	Imo          string    `json:"imo"`
	CompanyEmail string    `json:"company_email"`
	Instragram   string    `json:"instagram"`
	CreatedAt    string    `json:"-"`
	UpdatedAt    string    `json:"-"`
	DeletedAt    string    `json:"-"`
}
