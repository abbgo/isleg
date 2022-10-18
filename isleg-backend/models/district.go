package models

type District struct {
	ID        string  `json:"id,omitempty"`
	Price     float64 `json:"price,omitempty"`
	CreatedAt string  `json:"-"`
	UpdatedAt string  `json:"-"`
	DeletedAt string  `json:"-"`
}

type TranslationDistrict struct {
	ID         string `json:"id,omitempty"`
	LangID     string `json:"lang_id,omitempty"`
	DistrictID string `json:"district_id,omitempty"`
	CreatedAt  string `json:"-"`
	UpdatedAt  string `json:"-"`
	DeletedAt  string `json:"-"`
}
