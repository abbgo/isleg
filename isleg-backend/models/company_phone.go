package models

type CompanyPhone struct {
	ID        string `json:"id,omitempty"`
	Phone     string `json:"phone,omitempty"`
	CreatedAt string `json:"-"`
	UpdatedAt string `json:"-"`
	DeletedAt string `json:"-"`
}
