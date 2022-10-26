package models

type TranslationUpdatePasswordPage struct {
	ID             string `json:"id,omitempty"`
	LangID         string `json:"lang_id,omitempty" binding:"required"`
	Title          string `json:"title,omitempty" binding:"required"`
	Password       string `json:"password,omitempty" binding:"required"`
	VerifyPassword string `json:"verify_password,omitempty" binding:"required"`
	Explanation    string `json:"explanation,omitempty" binding:"required"`
	Save           string `json:"save,omitempty" binding:"required"`
	CreatedAt      string `json:"-"`
	UpdatedAt      string `json:"-"`
	DeletedAt      string `json:"-"`
}
