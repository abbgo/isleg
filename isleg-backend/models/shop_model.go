package models

import (
	"errors"
	"github/abbgo/isleg/isleg-backend/config"
	"strconv"

	"github.com/google/uuid"
)

type Shop struct {
	ID          uuid.UUID `json:"id"`
	OwnerName   string    `json:"owner_name"`
	Address     string    `json:"address"`
	PhoneNumber string    `json:"phone_number"`
	RunningTime string    `json:"running_time"`
	CreatedAt   string    `json:"-"`
	UpdatedAt   string    `json:"-"`
	DeletedAt   string    `json:"-"`
}

type CategoryShop struct {
	ID         uuid.UUID `json:"id"`
	CategoryID uuid.UUID `json:"category_id"`
	ShopID     uuid.UUID `json:"shop_id"`
	CreatedAt  string    `json:"-"`
	UpdatedAt  string    `json:"-"`
	DeletedAt  string    `json:"-"`
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
