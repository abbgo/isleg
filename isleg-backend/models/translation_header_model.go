package models

import "github.com/google/uuid"

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
	CreatedAt            string    `json:"-"`
	UpdatedAt            string    `json:"-"`
	DeletedAt            string    `json:"-"`
}
