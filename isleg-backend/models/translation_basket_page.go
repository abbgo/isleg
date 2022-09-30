package models

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type TranslationBasketPage struct {
	ID              uuid.UUID `json:"id"`
	LangID          uuid.UUID `json:"lang_id"`
	QuantityOfGoods string    `json:"quantity_of_goods"`
	TotalPrice      string    `json:"total_price"`
	Discount        string    `json:"discount"`
	Delivery        string    `json:"delivery"`
	Total           string    `json:"total"`
	Currency        string    `json:"currency"`
	ToOrder         string    `json:"to_order"`
	YourBasket      string    `json:"your_basket"`
	EmptyTheBasket  string    `json:"empty_the_basket"`
	CreatedAt       string    `json:"-"`
	UpdatedAt       string    `json:"-"`
	DeletedAt       string    `json:"-"`
}

func ValidateTranslationBasketPageData(languages []Language, dataNames []string, context *gin.Context) error {

	for _, dataName := range dataNames {
		for _, v := range languages {
			if context.PostForm(dataName+"_"+v.NameShort) == "" {
				return errors.New(dataName + "_" + v.NameShort + " is required")
			}
		}
	}

	return nil

}

func ValidateTranslationBasketPageUpdate(dataNames []string, context *gin.Context) error {

	for _, dataName := range dataNames {
		if context.PostForm(dataName) == "" {
			return errors.New(dataName + " is required")
		}
	}

	return nil

}
