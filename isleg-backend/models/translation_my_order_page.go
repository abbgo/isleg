package models

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type TranslationMyOrderPage struct {
	ID         uuid.UUID `json:"id"`
	LangID     uuid.UUID `json:"lang_id"`
	Orders     string    `json:"orders"`
	Date       string    `json:"date"`
	Price      string    `json:"price"`
	Currency   string    `json:"currency"`
	Image      string    `json:"image"`
	Name       string    `json:"name"`
	Brend      string    `json:"brend"`
	Code       string    `json:"code"`
	Amount     string    `json:"amount"`
	TotalPrice string    `json:"total_price"`
	CreatedAt  string    `json:"-"`
	UpdatedAt  string    `json:"-"`
	DeletedAt  string    `json:"-"`
}

func ValidateTranslationMyOrderPageData(languages []Language, dataNames []string, context *gin.Context) error {

	for _, dataName := range dataNames {
		for _, v := range languages {
			if context.PostForm(dataName+"_"+v.NameShort) == "" {
				return errors.New(dataName + "_" + v.NameShort + " is required")
			}
		}
	}

	return nil

}

func ValidateTranslationMyOrderPageUpdate(dataNames []string, context *gin.Context) error {

	for _, dataName := range dataNames {
		if context.PostForm(dataName) == "" {
			return errors.New(dataName + " is required")
		}
	}

	return nil

}
