package models

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type TranslationContact struct {
	ID           uuid.UUID `json:"id"`
	LangID       uuid.UUID `json:"lang_id"`
	FullName     string    `json:"full_name"`
	Email        string    `json:"email"`
	Phone        string    `json:"phone"`
	Letter       string    `json:"letter"`
	CompanyPhone string    `json:"company_phone"`
	Imo          string    `json:"imo"`
	CompanyEmail string    `json:"company_email"`
	Instragram   string    `json:"instagram"`
	ButtonText   string    `json:"button_text"`
	CreatedAt    string    `json:"-"`
	UpdatedAt    string    `json:"-"`
	DeletedAt    string    `json:"-"`
}

func ValidateTranslationContactData(languages []Language, dataNames []string, context *gin.Context) error {

	for _, dataName := range dataNames {
		for _, v := range languages {
			if context.PostForm(dataName+"_"+v.NameShort) == "" {
				return errors.New(dataName + "_" + v.NameShort + " is required")
			}
		}
	}

	return nil

}
