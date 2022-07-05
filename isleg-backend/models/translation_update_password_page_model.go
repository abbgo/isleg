package models

import "github.com/google/uuid"

type TranslationUpdatePasswordPage struct {
	ID             uuid.UUID `json:"id"`
	LangID         uuid.UUID `json:"lang_id"`
	Title          string    `json:"title"`
	VerifyPassword string    `json:"verify_password"`
	Explanation    string    `json:"explanation"`
	Save           string    `json:"save"`
	CreatedAt      string    `json:"-"`
	UpdatedAt      string    `json:"-"`
	DeletedAt      string    `json:"-"`
}
