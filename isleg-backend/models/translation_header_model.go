package models

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type TranslationHeader struct {
	ID                   uuid.UUID `json:"id"`
	LangID               uuid.UUID `json:"lang_id"`
	Research             string    `json:"research"`
	Phone                string    `json:"phone"`
	Password             string    `json:"password"`
	ForgotPassword       string    `json:"forgot_password"`
	SignIn               string    `json:"sign_in"`
	SignUp               string    `json:"sign_up"`
	Name                 string    `json:"name"`
	PasswordVerification string    `json:"password_verification"`
	VerifySecure         string    `json:"verify_secure"`
	MyInformation        string    `json:"my_information"`
	MyFavorites          string    `json:"my_favorites"`
	MyOrders             string    `json:"my_orders"`
	LogOut               string    `json:"log_out"`
	Basket               string    `json:"basket"`
	Email                string    `json:"email"`
	AddToBasket          string    `json:"add_to_basket"`
	CreatedAt            string    `json:"-"`
	UpdatedAt            string    `json:"-"`
	DeletedAt            string    `json:"-"`
}

func ValidateTranslationHeaderCreate(languages []Language, dataNames []string, context *gin.Context) error {

	for _, dataName := range dataNames {
		for _, v := range languages {
			if context.PostForm(dataName+"_"+v.NameShort) == "" {
				return errors.New(dataName + "_" + v.NameShort + " is required")
			}
		}
	}

	return nil

}

func ValidateTranslationHeaderUpdate(dataNames []string, context *gin.Context) error {

	for _, dataName := range dataNames {
		if context.PostForm(dataName) == "" {
			return errors.New(dataName + " is required")
		}
	}

	return nil

}
