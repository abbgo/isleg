package models

type OrderedProducts struct {
	ID                string `json:"id,omitempty"`
	ProductID         string `json:"product_id,omitempty"`
	QuantityOfProduct uint   `json:"quantity_of_product,omitempty"`
	OrderID           string `json:"order_id,omitempty"`
	CreatedAt         string `json:"-"`
	UpdatedAt         string `json:"-"`
	DeletedAt         string `json:"-"`
}
