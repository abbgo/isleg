package models

import (
	"context"
	"errors"
	"github/abbgo/isleg/isleg-backend/config"

	"github.com/google/uuid"
	"gopkg.in/guregu/null.v4"
)

type Category struct {
	ID                  string                `json:"id,omitempty"`
	ParentCategoryID    null.String           `json:"parent_category_id,omitempty"`
	Image               string                `json:"image,omitempty"`
	IsHomeCategory      bool                  `json:"is_home_category,omitempty"`
	CreatedAt           string                `json:"-"`
	UpdatedAt           string                `json:"-"`
	DeletedAt           string                `json:"-"`
	TranslationCategory []TranslationCategory `json:"translation_category,omitempty" binding:"required"` // one to many
}

type TranslationCategory struct {
	ID         string        `json:"id,omitempty"`
	LangID     uuid.NullUUID `json:"lang_id,omitempty" binding:"required"`
	CategoryID uuid.NullUUID `json:"category_id,omitempty"`
	Name       string        `json:"name,omitempty" binding:"required"`
	Slug       string        `json:"slug,omitempty"`
	CreatedAt  string        `json:"-"`
	UpdatedAt  string        `json:"-"`
	DeletedAt  string        `json:"-"`
}

func ValidateCategory(categoryID, parentCategoryID, fileName, metod string) error {

	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		return err
	}
	defer db.Close()

	if categoryID != "" { // validate id and get image of category
		var category_id string
		db.QueryRow(context.Background(), "SELECT id FROM categories WHERE id = $1 AND deleted_at IS NULL", categoryID).Scan(&category_id)

		if category_id == "" {
			return errors.New("category not found")
		}
	}

	// validate parentCategoryID
	if parentCategoryID != "" {

		if metod == "create" {
			if fileName != "" {
				return errors.New("child cannot be an image of the category")
			}
		}

		var parentCategory string
		db.QueryRow(context.Background(), "SELECT id FROM categories WHERE id = $1 AND deleted_at IS NULL", parentCategoryID).Scan(&parentCategory)

		if parentCategory == "" {
			return errors.New("parent category not found")
		}

		return nil
	} else {
		if metod == "create" {
			if fileName == "" {
				return errors.New("parent category image is required")
			}
		}
	}

	return nil
}
