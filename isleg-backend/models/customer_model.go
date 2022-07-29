package models

import (
	"errors"
	"github/abbgo/isleg/isleg-backend/config"
	"github/abbgo/isleg/isleg-backend/pkg"
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
	Email       string         `json:"email"`
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

func ValidateCustomerData(fullName, phoneNumber, password, email string) error {

	if fullName == "" {
		return errors.New("full name is required")
	}

	if phoneNumber == "" {
		return errors.New("phone number is required")
	} else {
		row, err := config.ConnDB().Query("SELECT phone_number FROM customers WHERE phone_number = $1 AND deleted_at IS NULL", phoneNumber)
		if err != nil {
			return err
		}

		var phone_number string

		for row.Next() {
			if err := row.Scan(&phone_number); err != nil {
				return err
			}
		}

		if phone_number != "" {
			return errors.New("this customer already exists")
		}
	}

	_, err := strconv.Atoi(phoneNumber)
	if err != nil {
		return err
	}

	if len(phoneNumber) != 8 {
		return errors.New("the length of the phone number must be 8")
	}

	if password == "" {
		return errors.New("password is required")
	}

	if len(password) < 5 || len(password) > 25 {
		return errors.New("password length must be between 5 and 25")
	}

	if email != "" {
		if !pkg.IsEmailValid(email) {
			return errors.New("email it doesn't match")
		}

		rowEmail, err := config.ConnDB().Query("SELECT email FROM customers WHERE email = $1 AND deleted_at IS NULL", email)
		if err != nil {
			return err
		}

		var email_address string

		for rowEmail.Next() {
			if err := rowEmail.Scan(&email_address); err != nil {
				return err
			}
		}

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
