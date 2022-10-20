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

func ValidateTranslationFooterUpdate(dataNames []string, context *gin.Context) error {

	for _, dataName := range dataNames {
		if context.PostForm(dataName) == "" {
			return errors.New(dataName + " is required")
		}
	}

	return nil

}
