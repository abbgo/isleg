package models

import (
	"errors"

	"github.com/gin-gonic/gin"
)

type TranslationMyInformationPage struct {
	ID             string `json:"id,omitempty"`
	LangID         string `json:"lang_id,omitempty"`
	Birthday       string `json:"birthday,omitempty"`
	Address        string `json:"address,omitempty"`
	UpdatePassword string `json:"update_password,omitempty"`
	Save           string `json:"save,omitempty"`
	CreatedAt      string `json:"-"`
	UpdatedAt      string `json:"-"`
	DeletedAt      string `json:"-"`
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
