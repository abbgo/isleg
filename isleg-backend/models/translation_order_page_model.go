package models

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type TranslationOrderPage struct {
	ID                  uuid.UUID `json:"id"`
	LangID              uuid.UUID `json:"lang_id"`
	Content             string    `json:"content"`
	TypeOfPayment       string    `json:"type_of_payment"`
	ChooseADeliveryTime string    `json:"choose_a_delivery_time"`
	YourAddress         string    `json:"your_address"`
	Mark                string    `json:"mark"`
	ToOrder             string    `json:"to_order"`
	Tomorrow            string    `json:"tomorrow"`
	Cash                string    `json:"cash"`
	PaymentTerminal     string    `json:"payment_terminal"`
	CreatedAt           string    `json:"-"`
	UpdatedAt           string    `json:"-"`
	DeletedAt           string    `json:"-"`
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
