package models

type CompanyAddress struct {
	ID        string `json:"id,omitempty"`
	LangID    string `json:"lang_id,omitempty"`
	Address   string `json:"address,omitempty"`
	CreatedAt string `json:"-"`
	UpdatedAt string `json:"-"`
	DeletedAt string `json:"-"`
}
