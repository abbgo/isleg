package models

import "github.com/google/uuid"

type CompanySetting struct {
	ID          uuid.UUID `json:"id"`
	LogoPath    string    `json:"logo_path"`
	FaviconPath string    `json:"favicon_path"`
	Email       string    `json:"email"`
	Instagram   string    `json:"instagram"`
	CreatedAt   string    `json:"-"`
	UpdatedAt   string    `json:"-"`
	DeletedAt   string    `json:"-"`
}
