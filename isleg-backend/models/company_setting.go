package models

import (
	"errors"
	"github/abbgo/isleg/isleg-backend/helpers"
)

type CompanySetting struct {
	ID        string `json:"id,omitempty"`
	Logo      string `json:"logo,omitempty"`
	Favicon   string `json:"favicon,omitempty"`
	Email     string `json:"email,omitempty"`
	Instagram string `json:"instagram,omitempty"`
	Imo       string `json:"imo,omitempty"`
	CreatedAt string `json:"-"`
	UpdatedAt string `json:"-"`
	DeletedAt string `json:"-"`
}

func ValidateCompanySettingData(email, instagram, imo string) error {

	emailResult := helpers.IsEmailValid(email)
	if email == "" || !emailResult {
		return errors.New("email address is required or it doesn't match")
	}

	if instagram == "" {
		return errors.New("instagram is required")
	}

	if imo == "" {
		return errors.New("imo is required")
	}

	return nil
}
