package models

type Notification struct {
	ID        string `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	CreatedAt string `json:"-"`
	UpdatedAt string `json:"-"`
	DeletedAt string `json:"-"`
}

type TranslationNotification struct {
	ID             string `json:"id,omitempty"`
	NotificationID string `json:"notification_id,omitempty"`
	LangID         string `json:"lang_id,omitempty"`
	Translations   string `json:"translation,omitempty"`
	CreatedAt      string `json:"-"`
	UpdatedAt      string `json:"-"`
	DeletedAt      string `json:"-"`
}
