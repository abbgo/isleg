package models

import "github.com/google/uuid"

type Shop struct {
	ID            uuid.UUID `json:"id"`
	OwnerName     string    `json:"owner_name"`
	Address       string    `json:"address"`
	PhoneNumber   string    `json:"phone_number"`
	NumberOfGoods uint      `json:"number_of_goods"`
	RunningTime   string    `json:"running_time"`
	CreatedAt     string    `json:"-"`
	UpdatedAt     string    `json:"-"`
	DeletedAt     string    `json:"-"`
}

type CategoryShop struct {
	ID         uuid.UUID `json:"id"`
	CategoryID uuid.UUID `json:"category_id"`
	ShopID     uuid.UUID `json:"shop_id"`
	CreatedAt  string    `json:"-"`
	UpdatedAt  string    `json:"-"`
	DeletedAt  string    `json:"-"`
}
