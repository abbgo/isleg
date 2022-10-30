package models

type TranslationFooter struct {
	ID        string `json:"id,omitempty"`
	LangID    string `json:"lang_id,omitempty" binding:"required"`
	About     string `json:"about,omitempty" binding:"required"`
	Payment   string `json:"payment,omitempty" binding:"required"`
	Contact   string `json:"contact,omitempty" binding:"required"`
	Secure    string `json:"secure,omitempty" binding:"required"`
	Word      string `json:"word,omitempty" binding:"required"`
	CreatedAt string `json:"-"`
	UpdatedAt string `json:"-"`
	DeletedAt string `json:"-"`
}
