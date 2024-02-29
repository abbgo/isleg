package models

import (
	"context"
	"errors"
	"github/abbgo/isleg/isleg-backend/config"

	"github.com/google/uuid"
	"gopkg.in/guregu/null.v4"
)

type Category struct {
	ID                       string                `json:"id,omitempty"`
	ParentCategoryID         null.String           `json:"parent_category_id,omitempty"`
	Image                    string                `json:"image,omitempty"`
	IsHomeCategory           bool                  `json:"is_home_category,omitempty"`
	CreatedAt                string                `json:"-"`
	UpdatedAt                string                `json:"-"`
	DeletedAt                string                `json:"-"`
	OrderNumber              uint                  `json:"order_number,omitempty"`
	OldOrderNumber           uint                  `json:"old_order_number,omitempty"`
	OrderNumberInHomePage    uint                  `json:"order_number_in_home_page,omitempty"`
	OldOrderNumberInHomePage uint                  `json:"old_order_number_in_home_page,omitempty"`
	IsVisible                bool                  `json:"is_visible,omitempty"`
	TranslationCategory      []TranslationCategory `json:"translation_category,omitempty" binding:"required"` // one to many
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

func ValidateCategory(categoryID, parentCategoryID, fileName, metod string, orderNumber, OldOrderNumber, orderNumberInHomePage, oldOrderNumberInHomePage uint, isHomeCategory bool) (uint, uint, error) {

	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		return 0, 0, err
	}
	defer db.Close()

	if categoryID != "" { // validate id and get image of category
		var category_id string
		db.QueryRow(context.Background(), "SELECT id FROM categories WHERE id = $1 AND deleted_at IS NULL", categoryID).Scan(&category_id)

		if category_id == "" {
			return 0, 0, errors.New("category not found")
		}
	}

	// validate parentCategoryID
	if parentCategoryID != "" {
		if metod == "create" {
			if fileName != "" {
				return 0, 0, errors.New("child cannot be an image of the category")
			}
		}

		var parentCategory string
		db.QueryRow(context.Background(), "SELECT id FROM categories WHERE id = $1 AND deleted_at IS NULL", parentCategoryID).Scan(&parentCategory)

		if parentCategory == "" {
			return 0, 0, errors.New("parent category not found")
		}

		return 0, 0, nil
	} else {
		if metod == "create" {
			if fileName == "" {
				return 0, 0, errors.New("parent category image is required")
			}

			if orderNumber != 0 {
				var category_id string
				db.QueryRow(context.Background(), "SELECT id FROM categories WHERE order_number = $1 AND deleted_at IS NULL AND parent_category_id IS NULL", orderNumber).Scan(&category_id)
				if category_id != "" {
					return 0, 0, errors.New("this order_number already exists")
				}
			} else {
				var order_umber null.Int
				if err := db.QueryRow(context.Background(), "SELECT MAX(order_number) FROM categories WHERE deleted_at IS NULL AND parent_category_id IS NULL").Scan(&order_umber); err != nil {
					return 0, 0, err
				}
				orderNumber = uint(order_umber.Int64) + 1
			}
		} else {
			if orderNumber == 0 {
				return 0, 0, errors.New("order_number is required")
			}

			var order_number uint
			db.QueryRow(context.Background(), "SELECT order_number FROM categories WHERE parent_category_id IS NULL AND deleted_at IS NULL AND id = $1", categoryID).Scan(&order_number)
			if order_number == 0 {
				return 0, 0, errors.New("order_number not found")
			}

			if order_number != orderNumber {
				if orderNumber == OldOrderNumber {
					return 0, 0, errors.New("order_number don't equal old_order_number")
				}

				if orderNumber < OldOrderNumber {
					_, err = db.Exec(context.Background(), "UPDATE categories SET order_number = order_number + 1 WHERE parent_category_id IS NULL AND deleted_at IS NULL AND order_number >= $1 AND order_number < $2", orderNumber, OldOrderNumber)
					if err != nil {
						return 0, 0, err
					}
				} else if orderNumber > OldOrderNumber {
					_, err = db.Exec(context.Background(), "UPDATE categories SET order_number = order_number - 1 WHERE parent_category_id IS NULL AND deleted_at IS NULL AND order_number > $1 AND order_number <= $2", OldOrderNumber, orderNumber)
					if err != nil {
						return 0, 0, err
					}
				}
			}
		}
	}

	if isHomeCategory {
		if metod == "create" {
			if orderNumberInHomePage != 0 {
				var category_id string
				db.QueryRow(context.Background(), "SELECT id FROM categories WHERE is_home_category = true AND order_number_in_home_page = $1 AND deleted_at IS NULL", orderNumberInHomePage).Scan(&category_id)
				if category_id != "" {
					return 0, 0, errors.New("this order_number already exists")
				}
			} else {
				var order_number_in_home_page null.Int
				if err := db.QueryRow(context.Background(), "SELECT MAX(order_number_in_home_page) FROM categories WHERE deleted_at IS NULL AND is_home_category = true").Scan(&order_number_in_home_page); err != nil {
					return 0, 0, err
				}
				orderNumberInHomePage = uint(order_number_in_home_page.Int64) + 1
			}
		} else {
			var old_is_home_category bool
			db.QueryRow(context.Background(), "SELECT is_home_category FROM categories WHERE id = $1 AND deleted_at IS NULL", categoryID).Scan(&old_is_home_category)
			if old_is_home_category {
				if orderNumberInHomePage == 0 {
					return 0, 0, errors.New("order_number_in_home_page is required")
				}

				var order_number_in_home_page uint
				db.QueryRow(context.Background(), "SELECT order_number FROM categories WHERE deleted_at IS NULL AND id = $1 AND is_home_category = true", categoryID).Scan(&order_number_in_home_page)
				if order_number_in_home_page == 0 {
					return 0, 0, errors.New("order_number_in_home_page not found")
				}

				if order_number_in_home_page != orderNumberInHomePage {
					if orderNumberInHomePage == oldOrderNumberInHomePage {
						return 0, 0, errors.New("order_number_in_home_page don't equal old_order_number_in_home_page")
					}

					if orderNumberInHomePage < oldOrderNumberInHomePage {
						_, err = db.Exec(context.Background(), "UPDATE categories SET order_number_in_home_page = order_number_in_home_page + 1 WHERE is_home_category = true AND deleted_at IS NULL AND order_number_in_home_page >= $1 AND order_number_in_home_page < $2", orderNumberInHomePage, oldOrderNumberInHomePage)
						if err != nil {
							return 0, 0, err
						}
					} else if orderNumberInHomePage > oldOrderNumberInHomePage {
						_, err = db.Exec(context.Background(), "UPDATE categories SET order_number_in_home_page = order_number_in_home_page - 1 WHERE is_home_category = true IS NULL AND deleted_at IS NULL AND order_number_in_home_page > $1 AND order_number_in_home_page <= $2", oldOrderNumberInHomePage, orderNumberInHomePage)
						if err != nil {
							return 0, 0, err
						}
					}
				}
			} else {
				var minOrderNumber, maxOrderNumber uint
				db.QueryRow(context.Background(), "SELECT MIN(order_number_in_home_page),MAX(order_number_in_home_page) FROM categories WHERE is_home_category = true AND deleted_at IS NULL").Scan(&minOrderNumber, &maxOrderNumber)

				if orderNumberInHomePage == minOrderNumber {
					_, err = db.Exec(context.Background(), "UPDATE categories SET order_number_in_home_page = order_number_in_home_page + 1 WHERE is_home_category = true AND deleted_at IS NULL AND order_number_in_home_page >= $1", orderNumberInHomePage)
					if err != nil {
						return 0, 0, err
					}
				} else if orderNumberInHomePage == maxOrderNumber {
					_, err = db.Exec(context.Background(), "UPDATE categories SET order_number_in_home_page = order_number_in_home_page + 1 WHERE order_number_in_home_page = $1", maxOrderNumber)
					if err != nil {
						return 0, 0, err
					}
				} else if minOrderNumber < orderNumberInHomePage && orderNumberInHomePage < maxOrderNumber {
					_, err = db.Exec(context.Background(), "UPDATE categories SET order_number_in_home_page = order_number_in_home_page + 1 WHERE is_home_category = true AND deleted_at IS NULL AND order_number_in_home_page >= $1", orderNumberInHomePage)
					if err != nil {
						return 0, 0, err
					}
				}
			}
		}
	}

	return orderNumberInHomePage, orderNumber, nil
}
