package models

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Afisa struct {
	ID        uuid.UUID `json:"id"`
	Image     string    `json:"image"`
	CreatedAt string    `json:"-"`
	UpdatedAt string    `json:"-"`
	DeletedAt string    `json:"-"`
}

type TranslationAfisa struct {
	ID          uuid.UUID `json:"id"`
	AfisaID     uuid.UUID `json:"afisa_id"`
	LangID      uuid.UUID `json:"lang_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   string    `json:"-"`
	UpdatedAt   string    `json:"-"`
	DeletedAt   string    `json:"-"`
}

func ValidateAfisaData(languages []Language, dataNames []string, context *gin.Context) error {

	for _, dataName := range dataNames {
		for _, v := range languages {
			if context.PostForm(dataName+"_"+v.NameShort) == "" {
				return errors.New(dataName + "_" + v.NameShort + " is required")
			}
		}
	}

	return nil

}
