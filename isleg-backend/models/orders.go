package models

type Orders struct {
	ID            string  `json:"id,omitempty"`
	CustomerID    string  `json:"customer_id,omitempty"`
	CustomerMark  string  `json:"customer_mark,omitempty"`
	OrderTime     string  `json:"order_time,omitempty"`
	PaymentType   string  `json:"payment_type,omitempty"`
	TotalPrice    float64 `json:"total_price,omitempty"`
	OrderNumber   int     `json:"order_number,omitempty"`
	ShippingPrice float64 `json:"shipping_price,omitempty"`
	CreatedAt     string  `json:"-"`
	UpdatedAt     string  `json:"-"`
	DeletedAt     string  `json:"-"`
}
