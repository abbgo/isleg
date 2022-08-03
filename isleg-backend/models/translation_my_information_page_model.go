package models

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type TranslationMyInformationPage struct {
	ID             uuid.UUID `json:"id"`
	LangID         uuid.UUID `json:"lang_id"`
	Birthday       string    `json:"birthday"`
	Address        string    `json:"address"`
	UpdatePassword string    `json:"update_password"`
	Save           string    `json:"save"`
	CreatedAt      string    `json:"-"`
	UpdatedAt      string    `json:"-"`
	DeletedAt      string    `json:"-"`
}

func ValidateTranslationMyInformationPageData(languages []Language, dataNames []string, context *gin.Context) error {

	for _, dataName := range dataNames {
		for _, v := range languages {
			if context.PostForm(dataName+"_"+v.NameShort) == "" {
				return errors.New(dataName + "_" + v.NameShort + " is required")
			}
		}
	}

	return nil

}

func ValidateTranslationMyInformationPageUpdate(dataNames []string, context *gin.Context) error {

	for _, dataName := range dataNames {
		if context.PostForm(dataName) == "" {
			return errors.New(dataName + " is required")
		}
	}

	return nil

}
