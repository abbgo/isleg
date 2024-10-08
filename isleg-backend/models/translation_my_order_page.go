package models

import "github.com/google/uuid"

type TranslationMyOrderPage struct {
	ID           string        `json:"id,omitempty"`
	LangID       uuid.NullUUID `json:"lang_id,omitempty" binding:"required"`
	Orders       string        `json:"orders,omitempty" binding:"required"`
	Date         string        `json:"date,omitempty" binding:"required"`
	Price        string        `json:"price,omitempty" binding:"required"`
	Image        string        `json:"image,omitempty" binding:"required"`
	Name         string        `json:"name,omitempty" binding:"required"`
	Brend        string        `json:"brend,omitempty" binding:"required"`
	ProcuctPrice string        `json:"product_price,omitempty" binding:"required"`
	Amount       string        `json:"amount,omitempty" binding:"required"`
	TotalPrice   string        `json:"total_price,omitempty" binding:"required"`
	CreatedAt    string        `json:"-"`
	UpdatedAt    string        `json:"-"`
	DeletedAt    string        `json:"-"`
}
