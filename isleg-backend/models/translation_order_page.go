package models

import (
	"errors"

	"github.com/gin-gonic/gin"
)

type TranslationOrderPage struct {
	ID                  string `json:"id,omitempty"`
	LangID              string `json:"lang_id,omitempty"`
	Content             string `json:"content,omitempty"`
	TypeOfPayment       string `json:"type_of_payment,omitempty"`
	ChooseADeliveryTime string `json:"choose_a_delivery_time,omitempty"`
	YourAddress         string `json:"your_address,omitempty"`
	Mark                string `json:"mark,omitempty"`
	ToOrder             string `json:"to_order,omitempty"`
	Tomorrow            string `json:"tomorrow,omitempty"`
	Cash                string `json:"cash,omitempty"`
	PaymentTerminal     string `json:"payment_terminal,omitempty"`
	CreatedAt           string `json:"-"`
	UpdatedAt           string `json:"-"`
	DeletedAt           string `json:"-"`
}

func ValidateTranslationOrderPageData(languages []Language, dataNames []string, context *gin.Context) error {

	for _, dataName := range dataNames {
		for _, v := range languages {
			if context.PostForm(dataName+"_"+v.NameShort) == "" {
				return errors.New(dataName + "_" + v.NameShort + " is required")
			}
		}
	}

	return nil

}

func ValidateTranslationOrderPageUpdate(dataNames []string, context *gin.Context) error {

	for _, dataName := range dataNames {
		if context.PostForm(dataName) == "" {
			return errors.New(dataName + " is required")
		}
	}

	return nil

}
