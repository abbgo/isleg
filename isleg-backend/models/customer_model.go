package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type Customer struct {
	ID          uuid.UUID      `json:"id"`
	Name        string         `json:"name"`
	PhoneNumber string         `json:"phone_number"`
	Password    string         `json:"password"`
	Birthday    time.Time      `json:"birthday"`
	Gender      string         `json:"gender"`
	Addresses   pq.StringArray `json:"addresses"`
	CreatedAt   string         `json:"-"`
	UpdatedAt   string         `json:"-"`
	DeletedAt   string         `json:"-"`
}
