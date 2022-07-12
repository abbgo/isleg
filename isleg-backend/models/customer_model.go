package models

import (
	"errors"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

type Customer struct {
	ID          uuid.UUID      `json:"id"`
	FullName    string         `json:"full_name"`
	PhoneNumber string         `json:"phone_number"`
	Password    string         `json:"password"`
	Birthday    time.Time      `json:"birthday"`
	Gender      string         `json:"gender"`
	Addresses   pq.StringArray `json:"addresses"`
	CreatedAt   string         `json:"-"`
	UpdatedAt   string         `json:"-"`
	DeletedAt   string         `json:"-"`
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}
	hashPassword := string(bytes)
	return hashPassword, nil
}

func CheckPassword(providedPassword, oldPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(oldPassword), []byte(providedPassword))
	if err != nil {
		return err
	}
	return nil
}

func ValidateCustomerData(fullName, phoneNumber, password, gender string, addresses []string) error {

	if fullName == "" {
		return errors.New("full name is required")
	}

	if phoneNumber == "" {
		return errors.New("phone number is required")
	}

	_, err := strconv.Atoi(phoneNumber)
	if err != nil {
		return err
	}

	if len(phoneNumber) != 7 {
		return errors.New("the length of the phone number must be 7")
	}

	if password == "" {
		return errors.New("password is required")
	}

	if len(password) < 5 || len(password) > 25 {
		return errors.New("password length must be between 5 and 25")
	}

	if gender != "" {
		if gender != "1" && gender != "0" {
			return errors.New("gender must be 0 or 1")
		}
	}

	if len(addresses) == 0 {
		return errors.New("address is required")
	}
	for _, v := range addresses {
		if v == "" {
			return errors.New("address is required")
		}
	}

	return nil

}
