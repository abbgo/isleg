package models

import (
	"errors"
	"github/abbgo/isleg/isleg-backend/pkg"

	"github.com/google/uuid"
)

type CompanySetting struct {
	ID        uuid.UUID `json:"id"`
	Logo      string    `json:"logo"`
	Favicon   string    `json:"favicon"`
	Email     string    `json:"email"`
	Instagram string    `json:"instagram"`
	CreatedAt string    `json:"-"`
	UpdatedAt string    `json:"-"`
	DeletedAt string    `json:"-"`
}

func ValidateCompanySettingData(email, instagram string) error {

	emailResult := pkg.IsEmailValid(email)
	if email == "" || !emailResult {
		return errors.New("email address is required or it doesn't match")
	}

	if instagram == "" {
		return errors.New("instagram is required")
	}

	return nil
}
