package models

import (
	"github.com/google/uuid"
)

type Product struct {
	ID          uuid.UUID `json:"id"`
	BrendID     uuid.UUID `json:"brend_id"`
	Price       float64   `json:"price"`
	OldPrice    float64   `json:"old_price"`
	Amount      uint      `json:"amount"`
	ProductCode string    `json:"product_code"`
	LimitAmount uint      `json:"limit_amount"`
	IsNew       bool      `json:"is_new"`
	CreatedAt   string    `json:"-"`
	UpdatedAt   string    `json:"-"`
	DeletedAt   string    `json:"-"`
}

type MainImage struct {
	ID        uuid.UUID `json:"id"`
	ProductID uuid.UUID `json:"product_id"`
	Small     string    `json:"small"`
	Medium    string    `json:"medium"`
	Large     string    `json:"large"`
	CreatedAt string    `json:"-"`
	UpdatedAt string    `json:"-"`
	DeletedAt string    `json:"-"`
}

type Images struct {
	ID        uuid.UUID `json:"id"`
	ProductID uuid.UUID `json:"product_id"`
	Small     string    `json:"small"`
	Large     string    `json:"large"`
	CreatedAt string    `json:"-"`
	UpdatedAt string    `json:"-"`
	DeletedAt string    `json:"-"`
}

type TranslationProduct struct {
	ID          uuid.UUID `json:"id"`
	LangID      uuid.UUID `json:"lang_id"`
	ProductID   uuid.UUID `json:"product_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   string    `json:"-"`
	UpdatedAt   string    `json:"-"`
	DeletedAt   string    `json:"-"`
}

type CategoryProduct struct {
	ID         uuid.UUID `json:"id"`
	CategoryID uuid.UUID `json:"category_id"`
	ProductID  uuid.UUID `json:"product_id"`
	CreatedAt  string    `json:"-"`
	UpdatedAt  string    `json:"-"`
	DeletedAt  string    `json:"-"`
}
