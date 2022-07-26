package models

import (
	"errors"
	"strconv"

	"github.com/google/uuid"
)

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

func ValidateShopData(ownerName, address, phoneNumber, runningTime string, categories []string) error {

	if ownerName == "" {
		return errors.New("owner name is required")
	}

	if address == "" {
		return errors.New("address is required")
	}

	if phoneNumber == "" {
		return errors.New("phoneNumber is required")
	}

	_, err := strconv.Atoi(phoneNumber)
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
	}

	return nil

}
