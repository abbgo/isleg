package models

import (
	"errors"
	"github/abbgo/isleg/isleg-backend/config"
	"strconv"

	"github.com/google/uuid"
)

type Category struct {
	ID                  string                `json:"id,omitempty"`
	ParentCategoryID    uuid.NullUUID         `json:"parent_category_id,omitempty"`
	Image               string                `json:"image,omitempty"`
	IsHomeCategory      bool                  `json:"is_home_category,omitempty"`
	CreatedAt           string                `json:"-"`
	UpdatedAt           string                `json:"-"`
	DeletedAt           string                `json:"-"`
	TranslationCategory []TranslationCategory `json:"translation_category,omitempty"` // one to many
}

type TranslationCategory struct {
	ID         string        `json:"id,omitempty"`
	LangID     uuid.NullUUID `json:"lang_id,omitempty"`
	CategoryID uuid.NullUUID `json:"category_id,omitempty"`
	Name       string        `json:"name,omitempty"`
	CreatedAt  string        `json:"-"`
	UpdatedAt  string        `json:"-"`
	DeletedAt  string        `json:"-"`
}

func ValidateCategory(isHomeCategory, parentCategoryID, fileName string) (bool, string, error) {

	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		return false, "", err
	}
	defer func() (bool, string, error) {
		if err := db.Close(); err != nil {
			return false, "", err
		}
		return false, "", nil
	}()

	// validate is home category
	is_home_category, err := strconv.ParseBool(isHomeCategory)
	if err != nil {
		return false, "", err
	}

	// validate parentCategoryID
	if parentCategoryID != "" {

		if fileName != "" {
			return false, "", errors.New("child cannot be an image of the category")
		}

		rowCategory, err := db.Query("SELECT id FROM categories WHERE id = $1 AND deleted_at IS NULL", parentCategoryID)
		if err != nil {
			return false, "", err
		}
		defer func() (bool, string, string, error) {
			if err := rowCategory.Close(); err != nil {
				return false, "", "", err
			}
			return false, "", "", nil
		}()

		var parentCategory string

		for rowCategory.Next() {
			if err := rowCategory.Scan(&parentCategory); err != nil {
				return false, "", err
			}
		}

		if parentCategory == "" {
			return false, "", errors.New("parent category not found")
		}

		return is_home_category, parentCategory, nil
	} else {
		if fileName == "" {
			return false, "", errors.New("parent category image is required")
		}
	}

	return is_home_category, "", nil

}

func ValidateCategoryForUpdate(isHomeCategory, categoryID, parentCategoryID string) (bool, string, string, error) {

	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		return false, "", "", err
	}
	defer func() (bool, string, string, error) {
		if err := db.Close(); err != nil {
			return false, "", "", err
		}
		return false, "", "", nil
	}()

	// validate is home category
	is_home_category, err := strconv.ParseBool(isHomeCategory)
	if err != nil {
		return false, "", "", err
	}

	// validate id and get image of category
	rowCategor, err := db.Query("SELECT id,image FROM categories WHERE id = $1 AND deleted_at IS NULL", categoryID)
	if err != nil {
		return false, "", "", err
	}
	defer func() (bool, string, string, error) {
		if err := rowCategor.Close(); err != nil {
			return false, "", "", err
		}
		return false, "", "", nil
	}()

	var category_id, image string

	for rowCategor.Next() {
		if err := rowCategor.Scan(&category_id, &image); err != nil {
			return false, "", "", err
		}
	}

	if category_id == "" {
		return false, "", "", errors.New("category not found")
	}

	// validate parentCategoryID
	if parentCategoryID != "" {

		rowCategory, err := db.Query("SELECT id FROM categories WHERE id = $1 AND deleted_at IS NULL", parentCategoryID)
		if err != nil {
			return false, "", "", err
		}
		defer func() (bool, string, string, error) {
			if err := rowCategory.Close(); err != nil {
				return false, "", "", err
			}
			return false, "", "", nil
		}()

		var parentCategory string

		for rowCategory.Next() {
			if err := rowCategory.Scan(&parentCategory); err != nil {
				return false, "", "", err
			}
		}

		if parentCategory == "" {
			return false, "", "", errors.New("parent category not found")
		}

		return is_home_category, parentCategory, image, nil
	}

	return is_home_category, "", image, nil

}
