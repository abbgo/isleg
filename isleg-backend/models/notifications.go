package models

type Notification struct {
	ID           string                    `json:"id,omitempty"`
	Name         string                    `json:"name,omitempty" binding:"required"`
	CreatedAt    string                    `json:"-"`
	UpdatedAt    string                    `json:"-"`
	DeletedAt    string                    `json:"-"`
	Translations []TranslationNotification `json:"translations,omitempty" binding:"required"`
}

type TranslationNotification struct {
	ID             string `json:"id,omitempty"`
	NotificationID string `json:"notification_id,omitempty"`
	LangID         string `json:"lang_id,omitempty" binding:"required"`
	Translation    string `json:"translation,omitempty" binding:"required"`
	CreatedAt      string `json:"-"`
	UpdatedAt      string `json:"-"`
	DeletedAt      string `json:"-"`
}
