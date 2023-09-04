package models

import (
	"context"
	"errors"
	"github/abbgo/isleg/isleg-backend/config"
)

type Admin struct {
	ID          string `json:"id,omitempty"`
	FullName    string `json:"full_name,omitempty" binding:"required,min=3"`
	PhoneNumber string `json:"phone_number,omitempty" binding:"required,e164,len=12"`
	Password    string `json:"password,omitempty" binding:"required,min=3,max=25"`
	Type        string `json:"type,omitempty" binding:"required"`
	CreatedAt   string `json:"-"`
	UpdatedAt   string `json:"-"`
	DeletedAt   string `json:"-"`
}

func ValidateRegisterAdmin(phoneNumber, adminType string) error {

	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		return err
	}
	defer db.Close()

	if adminType != "admin" && adminType != "super_admin" {
		return errors.New("admin type must be admin or super_admin")
	}

	rowAdmin, err := db.Query(context.Background(), "SELECT phone_number FROM admins WHERE phone_number = $1 AND deleted_at IS NULL", phoneNumber)
	if err != nil {
		return err
	}

	var phone_number string
	for rowAdmin.Next() {
		if err := rowAdmin.Scan(&phone_number); err != nil {
			return err
		}
	}

	if phone_number != "" {
		return errors.New("this admin already exists")
	}

	return nil

}
