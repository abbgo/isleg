package models

type CompanyPhone struct {
	ID        string `json:"id,omitempty"`
	Phone     string `json:"phone,omitempty" binding:"required,e164"`
	CreatedAt string `json:"-"`
	UpdatedAt string `json:"-"`
	DeletedAt string `json:"-"`
}
