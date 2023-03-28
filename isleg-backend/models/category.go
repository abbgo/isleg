package models

import (
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
	CreatedAt  string        `json:"-"`
	UpdatedAt  string        `json:"-"`
	DeletedAt  string        `json:"-"`
}

func ValidateCategory(parentCategoryID, fileName string) (string, error) {

	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		return "", err
	}
	defer func() (string, error) {
		if err := db.Close(); err != nil {
			return "", err
		}
		return "", nil
	}()

	// validate parentCategoryID
	if parentCategoryID != "" {

		if fileName != "" {
			return "", errors.New("child cannot be an image of the category")
		}

		rowCategory, err := db.Query("SELECT id FROM categories WHERE id = $1 AND deleted_at IS NULL", parentCategoryID)
		if err != nil {
			return "", err
		}
		defer func() (string, string, error) {
			if err := rowCategory.Close(); err != nil {
				return "", "", err
			}
			return "", "", nil
		}()
		var parentCategory string
		for rowCategory.Next() {
			if err := rowCategory.Scan(&parentCategory); err != nil {
				return "", err
			}
		}

		if parentCategory == "" {
			return "", errors.New("parent category not found")
		}

		return parentCategory, nil
	} else {
		if fileName == "" {
			return "", errors.New("parent category image is required")
		}
	}

	return "", nil

}

func ValidateCategoryForUpdate(categoryID, parentCategoryID string) error {

	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		return err
	}
	defer func() error {
		if err := db.Close(); err != nil {
			return err
		}
		return nil
	}()

	// validate id and get image of category
	rowCategor, err := db.Query("SELECT id FROM categories WHERE id = $1 AND deleted_at IS NULL", categoryID)
	if err != nil {
		return err
	}
	defer func() error {
		if err := rowCategor.Close(); err != nil {
			return err
		}
		return nil
	}()

	var category_id string

	for rowCategor.Next() {
		if err := rowCategor.Scan(&category_id); err != nil {
			return err
		}
	}

	if category_id == "" {
		return errors.New("category not found")
	}

	// validate parentCategoryID
	if parentCategoryID != "" {

		rowCategory, err := db.Query("SELECT id FROM categories WHERE id = $1 AND deleted_at IS NULL", parentCategoryID)
		if err != nil {
			return err
		}
		defer func() error {
			if err := rowCategory.Close(); err != nil {
				return err
			}
			return nil
		}()

		var parentCategory string

		for rowCategory.Next() {
			if err := rowCategory.Scan(&parentCategory); err != nil {
				return err
			}
		}

		if parentCategory == "" {
			return errors.New("parent category not found")
		}

		return nil
	}

	return nil

}
