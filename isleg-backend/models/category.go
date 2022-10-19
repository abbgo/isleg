package models

type Category struct {
	ID                  string                `json:"id,omitempty"`
	ParentCategoryID    string                `json:"parent_category_id,omitempty"`
	Image               string                `json:"image,omitempty"`
	IsHomeCategory      bool                  `json:"is_home_category,omitempty"`
	CreatedAt           string                `json:"-"`
	UpdatedAt           string                `json:"-"`
	DeletedAt           string                `json:"-"`
	TranslationCategory []TranslationCategory `json:"translation_category,omitempty"`
}

type TranslationCategory struct {
	ID         string `json:"id,omitempty"`
	LangID     string `json:"lang_id,omitempty"`
	CategoryID string `json:"category_id,omitempty"`
	Name       string `json:"name,omitempty"`
	CreatedAt  string `json:"-"`
	UpdatedAt  string `json:"-"`
	DeletedAt  string `json:"-"`
}
