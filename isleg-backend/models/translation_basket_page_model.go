package models

import "github.com/google/uuid"

type TranslationBasketPage struct {
	ID              uuid.UUID `json:"id"`
	LangID          uuid.UUID `json:"lang_id"`
	QuantityOfGoods string    `json:"quantity_of_goods"`
	TotalPrice      string    `json:"total_price"`
	Discount        string    `json:"discount"`
	Delivery        string    `json:"delivery"`
	Total           string    `json:"total"`
	Currency        string    `json:"currency"`
	ToOrder         string    `json:"to_order"`
	YourOrder       string    `json:"your_basket"`
	CreatedAt       string    `json:"-"`
	UpdatedAt       string    `json:"-"`
	DeletedAt       string    `json:"-"`
}
