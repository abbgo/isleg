package models

type HelperImages struct {
	ID        string `json:"id,omitempty"`
	Image     string `json:"image,omitempty"`
	CreatedAt string `json:"-"`
	UpdatedAt string `json:"-"`
	DeletedAt string `json:"-"`
}
