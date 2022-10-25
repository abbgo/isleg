package models

type TranslationOrderPage struct {
	ID                  string `json:"id,omitempty"`
	LangID              string `json:"lang_id,omitempty"`
	Content             string `json:"content,omitempty"`
	TypeOfPayment       string `json:"type_of_payment,omitempty"`
	ChooseADeliveryTime string `json:"choose_a_delivery_time,omitempty"`
	YourAddress         string `json:"your_address,omitempty"`
	Mark                string `json:"mark,omitempty"`
	ToOrder             string `json:"to_order,omitempty"`
	CreatedAt           string `json:"-"`
	UpdatedAt           string `json:"-"`
	DeletedAt           string `json:"-"`
}
