package models

import (
	"context"
	"errors"
	"github/abbgo/isleg/isleg-backend/config"
	"strconv"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type Customer struct {
	ID              string            `json:"id,omitempty"`
	FullName        string            `json:"full_name,omitempty" binding:"required,min=3"`
	PhoneNumber     string            `json:"phone_number,omitempty" binding:"required,e164,len=12"`
	Password        string            `json:"password,omitempty"`
	Birthday        string            `json:"birthday,omitempty"`
	Gender          string            `json:"gender,omitempty"`
	Email           string            `json:"email,omitempty" binding:"email"`
	IsRegister      bool              `json:"is_register,omitempty"`
	CreatedAt       string            `json:"-"`
	UpdatedAt       string            `json:"-"`
	DeletedAt       string            `json:"-"`
	CustomerAddress []CustomerAddress `json:"customer_address,omitempty"` // one to many
	OtpSecretKey    string            `json:"-"`
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

func ValidateCustomerRegister(phoneNumber, email string) error {
	db, err := config.ConnDB()
	if err != nil {
		return err
	}
	defer db.Close()

	if phoneNumber != "" {
		if !strings.HasPrefix(phoneNumber, "+993") {
			return errors.New("phone number must start with +993")
		}

		_, err := strconv.Atoi(strings.Trim(phoneNumber, "+"))
		if err != nil {
			return err
		}

		if len(phoneNumber) != 12 {
			return errors.New("phone number must be 12 in length")
		}

		var phone_number string
		db.QueryRow(context.Background(), "SELECT phone_number FROM customers WHERE phone_number = $1 AND is_register = true AND deleted_at IS NULL", phoneNumber).Scan(&phone_number)

		if phone_number != "" {
			return errors.New("this customer already exists")
		}
	}

	if email != "" {
		var email_address string
		db.QueryRow(context.Background(), "SELECT email FROM customers WHERE email = $1 AND is_register = true AND deleted_at IS NULL", email).Scan(&email_address)

		if email_address != "" {
			return errors.New("this customer already exists")
		}
	}

	// if gender != "" {
	// 	if gender != "1" && gender != "0" {
	// 		return errors.New("gender must be 0 or 1")
	// 	}
	// }

	// if len(addresses) == 0 {
	// 	return errors.New("address is required")
	// }

	// if len(addresses) != 0 {
	// 	for _, v := range addresses {
	// 		if v == "" {
	// 			return errors.New("address is required")
	// 		}
	// 	}
	// }

	return nil
}

func ValidateCustomerLogin(phoneNumber string) error {
	if !strings.HasPrefix(phoneNumber, "+993") {
		return errors.New("phone number must start with +993")
	}

	_, err := strconv.Atoi(strings.Trim(phoneNumber, "+"))
	if err != nil {
		return err
	}

	if len(phoneNumber) != 12 {
		return errors.New("phone number must be 12 in length")
	}

	return nil
}
