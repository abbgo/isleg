package models

import (
	"errors"
	"github/abbgo/isleg/isleg-backend/config"
	"strconv"
)

type Shop struct {
	ID          string `json:"id,omitempty"`
	OwnerName   string `json:"owner_name,omitempty"`
	Address     string `json:"address,omitempty"`
	PhoneNumber string `json:"phone_number,omitempty"`
	RunningTime string `json:"running_time,omitempty"`
	CreatedAt   string `json:"-"`
	UpdatedAt   string `json:"-"`
	DeletedAt   string `json:"-"`
}

type CategoryShop struct {
	ID         string `json:"id,omitempty"`
	CategoryID string `json:"category_id,omitempty"`
	ShopID     string `json:"shop_id,omitempty"`
	CreatedAt  string `json:"-"`
	UpdatedAt  string `json:"-"`
	DeletedAt  string `json:"-"`
}

func ValidateShopData(ownerName, address, phoneNumber, runningTime string, categories []string) error {

	db, err := config.ConnDB()
	if err != nil {
		return nil
	}
	defer db.Close()

	if ownerName == "" {
		return errors.New("owner name is required")
	}

	if address == "" {
		return errors.New("address is required")
	}

	if phoneNumber == "" {
		return errors.New("phoneNumber is required")
	}

	_, err = strconv.Atoi(phoneNumber)
	if err != nil {
		return err
	}

	if len(phoneNumber) != 8 {
		return errors.New("the length of the phone number must be 8")
	}

	if runningTime == "" {
		return errors.New("running time is required")
	}

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
