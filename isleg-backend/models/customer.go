package models

import (
	"errors"
	"github/abbgo/isleg/isleg-backend/config"

	"golang.org/x/crypto/bcrypt"
	"gopkg.in/guregu/null.v4"
)

type Customer struct {
	ID              string            `json:"id,omitempty"`
	FullName        string            `json:"full_name,omitempty" binding:"required,min=3"`
	PhoneNumber     string            `json:"phone_number,omitempty" binding:"required,e164,len=12"`
	Password        string            `json:"password,omitempty" binding:"required,min=5,max=25"`
	Birthday        null.Time         `json:"birthday,omitempty"`
	Gender          string            `json:"gender,omitempty"`
	Email           string            `json:"email,omitempty" binding:"email"`
	IsRegister      bool              `json:"is_register,omitempty"`
	CreatedAt       string            `json:"-"`
	UpdatedAt       string            `json:"-"`
	DeletedAt       string            `json:"-"`
	CustomerAddress []CustomerAddress `json:"customer_address,omitempty"` // one to many
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
	defer func() error {
		if err := db.Close(); err != nil {
			return err
		}
		return nil
	}()

	if phoneNumber != "" {
		// if !strings.HasPrefix(phoneNumber, "+993") {
		// 	return errors.New("phone number must start with +993")
		// }

		// _, err := strconv.Atoi(strings.Trim(phoneNumber, "+"))
		// if err != nil {
		// 	return err
		// }

		row, err := db.Query("SELECT phone_number FROM customers WHERE phone_number = $1 AND is_register = true AND deleted_at IS NULL", phoneNumber)
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

	if email != "" {
		rowEmail, err := db.Query("SELECT email FROM customers WHERE email = $1 AND is_register = true AND deleted_at IS NULL", email)
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

// func ValidateCustomerLogin(phoneNumber string) error {

// 	if phoneNumber != "" {
// 		if !strings.HasPrefix(phoneNumber, "+993") {
// 			return errors.New("phone number must start with +993")
// 		}
// 	}

// 	return nil

// }
