package models

type TranslationMyInformationPage struct {
	ID             string `json:"id,omitempty"`
	LangID         string `json:"lang_id,omitempty"`
	Birthday       string `json:"birthday,omitempty"`
	Address        string `json:"address,omitempty"`
	UpdatePassword string `json:"update_password,omitempty"`
	Save           string `json:"save,omitempty"`
	CreatedAt      string `json:"-"`
	UpdatedAt      string `json:"-"`
	DeletedAt      string `json:"-"`
}
