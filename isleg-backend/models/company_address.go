package models

type CompanyAddress struct {
	ID        string `json:"id,omitempty"`
	LangID    string `json:"lang_id,omitempty" binding:"required"`
	Address   string `json:"address,omitempty" binding:"required"`
	CreatedAt string `json:"-"`
	UpdatedAt string `json:"-"`
	DeletedAt string `json:"-"`
}
