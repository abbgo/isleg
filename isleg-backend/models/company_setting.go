package models

import (
	"errors"
	"github/abbgo/isleg/isleg-backend/pkg"

	"github.com/google/uuid"
)

type CompanySetting struct {
	ID        uuid.UUID `json:"id,omitempty"`
	Logo      string    `json:"logo,omitempty"`
	Favicon   string    `json:"favicon,omitempty"`
	Email     string    `json:"email,omitempty"`
	Instagram string    `json:"instagram,omitempty"`
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
