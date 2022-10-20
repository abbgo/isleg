package models

import (
	"errors"

	"github.com/gin-gonic/gin"
)

type TranslationContact struct {
	ID           string `json:"id,omitempty"`
	LangID       string `json:"lang_id,omitempty"`
	FullName     string `json:"full_name,omitempty"`
	Email        string `json:"email,omitempty"`
	Phone        string `json:"phone,omitempty"`
	Letter       string `json:"letter,omitempty"`
	CompanyPhone string `json:"company_phone,omitempty"`
	Imo          string `json:"imo,omitempty"`
	CompanyEmail string `json:"company_email,omitempty"`
	Instragram   string `json:"instagram,omitempty"`
	ButtonText   string `json:"button_text,omitempty"`
	CreatedAt    string `json:"-"`
	UpdatedAt    string `json:"-"`
	DeletedAt    string `json:"-"`
}

func ValidateTranslationContactUpdate(dataNames []string, context *gin.Context) error {

	for _, dataName := range dataNames {
		if context.PostForm(dataName) == "" {
			return errors.New(dataName + " is required")
		}
	}

	return nil

}
