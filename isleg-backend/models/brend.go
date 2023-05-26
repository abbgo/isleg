package models

type Brend struct {
	ID        string    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty" binding:"required"`
	Image     string    `json:"image,omitempty"`
	CreatedAt string    `json:"-"`
	UpdatedAt string    `json:"-"`
	DeletedAt string    `json:"-"`
	Products  []Product `json:"products,omitempty"` // one to many
	Slug      string    `json:"slug,omitempty"`
}
