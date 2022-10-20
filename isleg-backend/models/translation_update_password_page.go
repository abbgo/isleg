package models

import (
	"errors"

	"github.com/gin-gonic/gin"
)

type TranslationUpdatePasswordPage struct {
	ID             string `json:"id,omitempty"`
	LangID         string `json:"lang_id,omitempty"`
	Title          string `json:"title,omitempty"`
	Password       string `json:"password,omitempty"`
	VerifyPassword string `json:"verify_password,omitempty"`
	Explanation    string `json:"explanation,omitempty"`
	Save           string `json:"save,omitempty"`
	CreatedAt      string `json:"-"`
	UpdatedAt      string `json:"-"`
	DeletedAt      string `json:"-"`
}

func ValidateTranslationUpdatePasswordPageUpdate(dataNames []string, context *gin.Context) error {

	for _, dataName := range dataNames {
		if context.PostForm(dataName) == "" {
			return errors.New(dataName + " is required")
		}
	}

	return nil

}
