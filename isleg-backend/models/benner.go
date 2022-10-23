package models

type Banner struct {
	ID        string `json:"id,omitempty"`
	Image     string `json:"image,omitempty"`
	Url       string `json:"url,omitempty"`
	CreatedAt string `json:"-"`
	UpdatedAt string `json:"-"`
	DeletedAt string `json:"-"`
}
