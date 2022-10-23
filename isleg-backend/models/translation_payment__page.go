package models

type TranslationPayment struct {
	ID        string `json:"id,omitempty"`
	LangID    string `json:"lang_id,omitempty"`
	Title     string `json:"title,omitempty"`
	Content   string `json:"content,omitempty"`
	CreatedAt string `json:"-"`
	UpdatedAt string `json:"-"`
	DeletedAt string `json:"-"`
}
