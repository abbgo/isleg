package models

type PaymentTypes struct {
	ID        string `json:"id,omitempty"`
	LangID    string `json:"lang_id,omitempty"`
	Type      string `json:"type,omitempty"`
	CreatedAt string `json:"-"`
	UpdatedAt string `json:"-"`
	DeletedAt string `json:"-"`
}
