package models

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Afisa struct {
	ID               string             `json:"id,omitempty"`
	Image            string             `json:"image,omitempty"`
	CreatedAt        string             `json:"-"`
	UpdatedAt        string             `json:"-"`
	DeletedAt        string             `json:"-"`
	TranslationAfisa []TranslationAfisa `json:"translation_afisa,omitempty"` // one to many
}

type TranslationAfisa struct {
	ID          string        `json:"id,omitempty"`
	AfisaID     uuid.NullUUID `json:"afisa_id,omitempty"`
	LangID      uuid.NullUUID `json:"lang_id,omitempty"`
	Title       string        `json:"title,omitempty"`
	Description string        `json:"description,omitempty"`
	CreatedAt   string        `json:"-"`
	UpdatedAt   string        `json:"-"`
	DeletedAt   string        `json:"-"`
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
