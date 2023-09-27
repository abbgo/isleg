package models

type SearchsOfCustomers struct {
	ID        string `json:"id,omitempty"`
	Search    string `json:"search,omitempty" binding:"required"`
	CreatedAt string `json:"-"`
	UpdatedAt string `json:"-"`
	DeletedAt string `json:"-"`
}
