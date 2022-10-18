package models

import (
	"errors"

	"github.com/gin-gonic/gin"
)

type TranslationBasketPage struct {
	ID              string `json:"id,omitempty"`
	LangID          string `json:"lang_id,omitempty"`
	QuantityOfGoods string `json:"quantity_of_goods,omitempty"`
	TotalPrice      string `json:"total_price,omitempty"`
	Discount        string `json:"discount,omitempty"`
	Delivery        string `json:"delivery,omitempty"`
	Total           string `json:"total,omitempty"`
	Currency        string `json:"currency,omitempty"`
	ToOrder         string `json:"to_order,omitempty"`
	YourBasket      string `json:"your_basket,omitempty"`
	EmptyTheBasket  string `json:"empty_the_basket,omitempty"`
	CreatedAt       string `json:"-"`
	UpdatedAt       string `json:"-"`
	DeletedAt       string `json:"-"`
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
