package models

import (
	"github.com/google/uuid"
)

type Product struct {
	ID          uuid.UUID `json:"id,omitempty"`
	BrendID     uuid.UUID `json:"brend_id,omitempty"`
	Price       float64   `json:"price,omitempty"`
	OldPrice    float64   `json:"old_price,omitempty"`
	Amount      uint      `json:"amount,omitempty"`
	ProductCode string    `json:"product_code,omitempty"`
	LimitAmount uint      `json:"limit_amount,omitempty"`
	IsNew       bool      `json:"is_new,omitempty"`
	CreatedAt   string    `json:"-"`
	UpdatedAt   string    `json:"-"`
	DeletedAt   string    `json:"-"`
}

type MainImage struct {
	ID        uuid.UUID `json:"id,omitempty"`
	ProductID uuid.UUID `json:"product_id,omitempty"`
	Small     string    `json:"small,omitempty"`
	Medium    string    `json:"medium,omitempty"`
	Large     string    `json:"large,omitempty"`
	CreatedAt string    `json:"-"`
	UpdatedAt string    `json:"-"`
	DeletedAt string    `json:"-"`
}

type Images struct {
	ID        uuid.UUID `json:"id,omitempty"`
	ProductID uuid.UUID `json:"product_id,omitempty"`
	Small     string    `json:"small,omitempty"`
	Large     string    `json:"large,omitempty"`
	CreatedAt string    `json:"-"`
	UpdatedAt string    `json:"-"`
	DeletedAt string    `json:"-"`
}

type TranslationProduct struct {
	ID          uuid.UUID `json:"id,omitempty"`
	LangID      uuid.UUID `json:"lang_id,omitempty"`
	ProductID   uuid.UUID `json:"product_id,omitempty"`
	Name        string    `json:"name,omitempty"`
	Description string    `json:"description,omitempty"`
	Slug        string    `json:"slug,omitempty"`
	CreatedAt   string    `json:"-"`
	UpdatedAt   string    `json:"-"`
	DeletedAt   string    `json:"-"`
}

type CategoryProduct struct {
	ID         uuid.UUID `json:"id,omitempty"`
	CategoryID uuid.UUID `json:"category_id,omitempty"`
	ProductID  uuid.UUID `json:"product_id,omitempty"`
	CreatedAt  string    `json:"-"`
	UpdatedAt  string    `json:"-"`
	DeletedAt  string    `json:"-"`
}
