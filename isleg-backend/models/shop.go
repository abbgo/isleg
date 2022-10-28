package models

import (
	"errors"
	"github/abbgo/isleg/isleg-backend/config"
)

type Shop struct {
	ID          string   `json:"id,omitempty"`
	OwnerName   string   `json:"owner_name,omitempty" binding:"required"`
	Address     string   `json:"address,omitempty" binding:"required"`
	PhoneNumber string   `json:"phone_number,omitempty" binding:"required"`
	RunningTime string   `json:"running_time,omitempty" binding:"required"`
	CreatedAt   string   `json:"-"`
	UpdatedAt   string   `json:"-"`
	DeletedAt   string   `json:"-"`
	Categories  []string `json:"categories,omitempty" binding:"required"`
}

type CategoryShop struct {
	ID         string `json:"id,omitempty"`
	CategoryID string `json:"category_id,omitempty" binding:"required"`
	ShopID     string `json:"shop_id,omitempty" binding:"required"`
	CreatedAt  string `json:"-"`
	UpdatedAt  string `json:"-"`
	DeletedAt  string `json:"-"`
}

func ValidateShopData(categories []string) error {

	db, err := config.ConnDB()
	if err != nil {
		return nil
	}
	defer db.Close()

	if len(categories) == 0 {
		return errors.New("shop category is required")
	} else {
		for _, v := range categories {
			rawCategory, err := db.Query("SELECT id FROM categories WHERE id = $1 AND deleted_at IS NULL", v)
			if err != nil {
				return err
			}

			var categoryID string

			for rawCategory.Next() {
				if err := rawCategory.Scan(&categoryID); err != nil {
					return err
				}
			}

			if categoryID == "" {
				return errors.New("category not found")
			}
		}
	}

	return nil

}
