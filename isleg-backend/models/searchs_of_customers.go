package models

type SearchsOfCustomers struct {
	ID          string `json:"id,omitempty"`
	Search      string `json:"search,omitempty"`
	Count       uint   `json:"count,omitempty"`
	HasProducts bool   `json:"has_products,omitempty"`
	Slug        string `json:"slug,omitempty"`
	CreatedAt   string `json:"-"`
	UpdatedAt   string `json:"-"`
	DeletedAt   string `json:"-"`
}
