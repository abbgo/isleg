package models

import "github.com/google/uuid"

type TranslationBasketPage struct {
	ID               string        `json:"id,omitempty"`
	LangID           uuid.NullUUID `json:"lang_id,omitempty" binding:"required"`
	QuantityOfGoods  string        `json:"quantity_of_goods,omitempty" binding:"required"`
	TotalPrice       string        `json:"total_price,omitempty" binding:"required"`
	Discount         string        `json:"discount,omitempty" binding:"required"`
	Delivery         string        `json:"delivery,omitempty" binding:"required"`
	Total            string        `json:"total,omitempty" binding:"required"`
	ToOrder          string        `json:"to_order,omitempty" binding:"required"`
	YourBasket       string        `json:"your_basket,omitempty" binding:"required"`
	EmptyTheBasket   string        `json:"empty_the_basket,omitempty" binding:"required"`
	EmptyTheLikePage string        `json:"empty_the_like_page,omitempty" binding:"required"`
	CreatedAt        string        `json:"-"`
	UpdatedAt        string        `json:"-"`
	DeletedAt        string        `json:"-"`
}
