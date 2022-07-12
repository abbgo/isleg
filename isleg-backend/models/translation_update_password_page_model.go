package models

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type TranslationUpdatePasswordPage struct {
	ID             uuid.UUID `json:"id"`
	LangID         uuid.UUID `json:"lang_id"`
	Title          string    `json:"title"`
	Password       string    `json:"password"`
	VerifyPassword string    `json:"verify_password"`
	Explanation    string    `json:"explanation"`
	Save           string    `json:"save"`
	CreatedAt      string    `json:"-"`
	UpdatedAt      string    `json:"-"`
	DeletedAt      string    `json:"-"`
}

func ValidateTranslationUpdatePasswordPageData(languages []Language, dataNames []string, context *gin.Context) error {

	for _, dataName := range dataNames {
		for _, v := range languages {
			if context.PostForm(dataName+"_"+v.NameShort) == "" {
				return errors.New(dataName + "_" + v.NameShort + " is required")
			}
		}
	}

	return nil

}
