package models

type TranslationMyInformationPage struct {
	ID             string `json:"id,omitempty"`
	LangID         string `json:"lang_id,omitempty" binding:"required"`
	Birthday       string `json:"birthday,omitempty" binding:"required"`
	Address        string `json:"address,omitempty" binding:"required"`
	UpdatePassword string `json:"update_password,omitempty" binding:"required"`
	Save           string `json:"save,omitempty" binding:"required"`
	CreatedAt      string `json:"-"`
	UpdatedAt      string `json:"-"`
	DeletedAt      string `json:"-"`
}
