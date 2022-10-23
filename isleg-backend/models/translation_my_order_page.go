package models

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
