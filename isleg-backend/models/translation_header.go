package models

type TranslationHeader struct {
	ID                   string `json:"id,omitempty"`
	LangID               string `json:"lang_id,omitempty"`
	Research             string `json:"research,omitempty"`
	Phone                string `json:"phone,omitempty"`
	Password             string `json:"password,omitempty"`
	ForgotPassword       string `json:"forgot_password,omitempty"`
	SignIn               string `json:"sign_in,omitempty"`
	SignUp               string `json:"sign_up,omitempty"`
	Name                 string `json:"name,omitempty"`
	PasswordVerification string `json:"password_verification,omitempty"`
	VerifySecure         string `json:"verify_secure,omitempty"`
	MyInformation        string `json:"my_information,omitempty"`
	MyFavorites          string `json:"my_favorites,omitempty"`
	MyOrders             string `json:"my_orders,omitempty"`
	LogOut               string `json:"log_out,omitempty"`
	Basket               string `json:"basket,omitempty"`
	Email                string `json:"email,omitempty"`
	AddToBasket          string `json:"add_to_basket,omitempty"`
	CreatedAt            string `json:"-"`
	UpdatedAt            string `json:"-"`
	DeletedAt            string `json:"-"`
}
