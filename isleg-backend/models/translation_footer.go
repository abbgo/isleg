package models

import (
	"errors"

	"github.com/gin-gonic/gin"
)

type TranslationFooter struct {
	ID        string `json:"id,omitempty"`
	LangID    string `json:"lang_id,omitempty"`
	About     string `json:"about,omitempty"`
	Payment   string `json:"payment,omitempty"`
	Contact   string `json:"contact,omitempty"`
	Secure    string `json:"secure,omitempty"`
	Word      string `json:"word,omitempty"`
	CreatedAt string `json:"-"`
	UpdatedAt string `json:"-"`
	DeletedAt string `json:"-"`
}

func ValidateTranslationFooterData(languages []Language, dataNames []string, context *gin.Context) error {

	for _, dataName := range dataNames {
		for _, v := range languages {
			if context.PostForm(dataName+"_"+v.NameShort) == "" {
				return errors.New(dataName + "_" + v.NameShort + " is required")
			}
		}
	}

	return nil

}

func ValidateTranslationFooterUpdate(dataNames []string, context *gin.Context) error {

	for _, dataName := range dataNames {
		if context.PostForm(dataName) == "" {
			return errors.New(dataName + " is required")
		}
	}

	return nil

}
