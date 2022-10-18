package models

import (
	"errors"

	"github.com/gin-gonic/gin"
)

type TranslationMyOrderPage struct {
	ID         string `json:"id,omitempty"`
	LangID     string `json:"lang_id,omitempty"`
	Orders     string `json:"orders,omitempty"`
	Date       string `json:"date,omitempty"`
	Price      string `json:"price,omitempty"`
	Currency   string `json:"currency,omitempty"`
	Image      string `json:"image,omitempty"`
	Name       string `json:"name,omitempty"`
	Brend      string `json:"brend,omitempty"`
	Code       string `json:"code,omitempty"`
	Amount     string `json:"amount,omitempty"`
	TotalPrice string `json:"total_price,omitempty"`
	CreatedAt  string `json:"-"`
	UpdatedAt  string `json:"-"`
	DeletedAt  string `json:"-"`
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
