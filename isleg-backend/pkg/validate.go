package pkg

import (
	"errors"
	"github/abbgo/isleg/isleg-backend/models"

	"github.com/gin-gonic/gin"
)

func ValidateTranslations(languages []models.Language, dataNames []string, context *gin.Context) error {

	for _, dataName := range dataNames {
		for _, v := range languages {
			if context.PostForm(dataName+"_"+v.NameShort) == "" {
				return errors.New(dataName + "_" + v.NameShort + " is required")
			}
		}
	}

	return nil

}
