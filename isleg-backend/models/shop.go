package models

import (
	"context"
	"errors"
	"github/abbgo/isleg/isleg-backend/config"
)

type Shop struct {
	ID          string `json:"id,omitempty"`
	OwnerName   string `json:"owner_name,omitempty" validate:"required,min=3"`
	Address     string `json:"address,omitempty" validate:"required,min=10"`
	PhoneNumber string `json:"phone_number,omitempty" validate:"required"`
	RunningTime string `json:"running_time,omitempty" validate:"required,min=3"`
	Image       string `json:"image,omitempty" validate:"required,min=3"`
	CategoryID  string `json:"category_id,omitempty" validate:"required,min=3"`
	Name        string `json:"name,omitempty" validate:"required,min=3"`
	CreatedAt   string `json:"-"`
	UpdatedAt   string `json:"-"`
	DeletedAt   string `json:"-"`
}

func ValidateShop(shopName string) error {
	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		return err
	}
	defer db.Close()

	var name string
	db.QueryRow(context.Background(), "SELECT id FROM shops WHERE name = $1 AND deleted_at IS NULL", shopName).Scan(&name)
	if name != "" {
		return errors.New("this shop already exists")
	}
	return nil
}
