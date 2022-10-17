package models

type Language struct {
	ID        string `json:"id,omitempty"`
	NameShort string `json:"name_short,omitempty"`
	Flag      string `json:"flag,omitempty"`
	CreatedAt string `json:"-"`
	UpdatedAt string `json:"-"`
	DeletedAt string `json:"-"`
}
