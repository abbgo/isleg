package models

import "github.com/google/uuid"

type TranslationOrderDates struct {
	ID          uuid.UUID `json:"id"`
	LangID      uuid.UUID `json:"lang_id"`
	OrderDateID uuid.UUID `json:"order_date_id"`
	CreatedAt   string    `json:"-"`
	UpdatedAt   string    `json:"-"`
	DeletedAt   string    `json:"-"`
}
