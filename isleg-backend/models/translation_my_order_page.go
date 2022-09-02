package models

import "github.com/google/uuid"

type TranslationMyOrderPage struct {
	ID         uuid.UUID `json:"id"`
	LangID     uuid.UUID `json:"lang_id"`
	Orders     string    `json:"orders"`
	Date       string    `json:"date"`
	Price      string    `json:"price"`
	Currency   string    `json:"currency"`
	Image      string    `json:"image"`
	Name       string    `json:"name"`
	Brend      string    `json:"brend"`
	Code       string    `json:"code"`
	Amount     string    `json:"amount"`
	TotalPrice string    `json:"total_price"`
	CreatedAt  string    `json:"-"`
	UpdatedAt  string    `json:"-"`
	DeletedAt  string    `json:"-"`
}
