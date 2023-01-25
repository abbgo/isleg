package models

import "github.com/google/uuid"

type TranslationOrderPage struct {
	ID                  string        `json:"id,omitempty"`
	LangID              uuid.NullUUID `json:"lang_id,omitempty" binding:"required"`
	Content             string        `json:"content,omitempty" binding:"required"`
	TypeOfPayment       string        `json:"type_of_payment,omitempty" binding:"required"`
	ChooseADeliveryTime string        `json:"choose_a_delivery_time,omitempty" binding:"required"`
	YourAddress         string        `json:"your_address,omitempty" binding:"required"`
	Mark                string        `json:"mark,omitempty" binding:"required"`
	ToOrder             string        `json:"to_order,omitempty" binding:"required"`
	CreatedAt           string        `json:"-"`
	UpdatedAt           string        `json:"-"`
	DeletedAt           string        `json:"-"`
}
