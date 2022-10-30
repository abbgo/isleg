package models

type TranslationSecure struct {
	ID        string `json:"id,omitempty"`
	LangID    string `json:"lang_id,omitempty" binding:"required"`
	Title     string `json:"title,omitempty" binding:"required"`
	Content   string `json:"content,omitempty" binding:"required"`
	CreatedAt string `json:"-"`
	UpdatedAt string `json:"-"`
	DeletedAt string `json:"-"`
}
