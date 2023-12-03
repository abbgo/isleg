package models

type CountOfCustomer struct {
	ID    string `json:"id,omitempty"`
	Count uint   `json:"owner_name,omitempty"`
	Date  string `json:"date,omitempty"`
}
