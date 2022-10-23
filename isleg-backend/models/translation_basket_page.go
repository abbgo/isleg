package models

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
