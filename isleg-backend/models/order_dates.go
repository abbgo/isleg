package models

import "github.com/google/uuid"

type OrderDates struct {
	ID        uuid.UUID `json:"id"`
	Date      string    `json:"date"`
	CreatedAt string    `json:"-"`
	UpdatedAt string    `json:"-"`
	DeletedAt string    `json:"-"`
}
