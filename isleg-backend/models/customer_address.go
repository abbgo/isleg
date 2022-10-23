package models

type CustomerAddress struct {
	ID         string `json:"id,omitempty"`
	CustomerID string `json:"customer_id,omitempty"`
	Address    string `json:"address,omitempty"`
	IsActive   bool   `json:"is_active,omitempty"`
	CreatedAt  string `json:"-"`
	UpdatedAt  string `json:"-"`
	DeletedAt  string `json:"-"`
}
