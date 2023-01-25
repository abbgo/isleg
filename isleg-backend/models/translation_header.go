package models

import "github.com/google/uuid"

type TranslationHeader struct {
	ID                   string        `json:"id,omitempty"`
	LangID               uuid.NullUUID `json:"lang_id,omitempty" binding:"required"`
	Research             string        `json:"research,omitempty" binding:"required"`
	Phone                string        `json:"phone,omitempty" binding:"required"`
	Password             string        `json:"password,omitempty" binding:"required"`
	ForgotPassword       string        `json:"forgot_password,omitempty" binding:"required"`
	SignIn               string        `json:"sign_in,omitempty" binding:"required"`
	SignUp               string        `json:"sign_up,omitempty" binding:"required"`
	Name                 string        `json:"name,omitempty" binding:"required"`
	PasswordVerification string        `json:"password_verification,omitempty" binding:"required"`
	VerifySecure         string        `json:"verify_secure,omitempty" binding:"required"`
	MyInformation        string        `json:"my_information,omitempty" binding:"required"`
	MyFavorites          string        `json:"my_favorites,omitempty" binding:"required"`
	MyOrders             string        `json:"my_orders,omitempty" binding:"required"`
	LogOut               string        `json:"log_out,omitempty" binding:"required"`
	Basket               string        `json:"basket,omitempty" binding:"required"`
	Email                string        `json:"email,omitempty" binding:"required"`
	AddToBasket          string        `json:"add_to_basket,omitempty" binding:"required"`
	CreatedAt            string        `json:"-"`
	UpdatedAt            string        `json:"-"`
	DeletedAt            string        `json:"-"`
}
