package models

import (
	"errors"
	"github/abbgo/isleg/isleg-backend/config"

	"github.com/google/uuid"
	"gopkg.in/guregu/null.v4"
)

type Product struct {
	ID                 string               `json:"id,omitempty"`
	BrendID            uuid.NullUUID        `json:"brend_id,omitempty"`
	ShopID             uuid.NullUUID        `json:"shop_id,omitempty"`
	Price              float64              `json:"price,omitempty"`
	OldPrice           float64              `json:"old_price"`
	Percentage         float64              `json:"percentage"`
	Benefit            null.Float           `json:"-"`
	Amount             uint                 `json:"amount,omitempty"`
	LimitAmount        uint                 `json:"limit_amount,omitempty"`
	IsNew              bool                 `json:"is_new,omitempty"`
	CreatedAt          string               `json:"-"`
	UpdatedAt          string               `json:"-"`
	DeletedAt          string               `json:"-"`
	MainImage          string               `json:"main_image,omitempty"`
	Images             []string             `json:"images,omitempty"`              // one to many
	TranslationProduct []TranslationProduct `json:"translation_product,omitempty"` // one to many
	Categories         []string             `json:"categories,omitempty"`
	Brend              Brend                `json:"brend,omitempty"`
	Shop               Shop                 `json:"shop,omitempty"`
}

type Images struct {
	ID        string `json:"id,omitempty"`
	ProductID string `json:"product_id,omitempty"`
	Image     string `json:"image,omitempty"`
	CreatedAt string `json:"-"`
	UpdatedAt string `json:"-"`
	DeletedAt string `json:"-"`
}

type TranslationProduct struct {
	ID          string `json:"id,omitempty"`
	LangID      string `json:"lang_id,omitempty"`
	ProductID   string `json:"product_id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Slug        string `json:"slug,omitempty"`
	CreatedAt   string `json:"-"`
	UpdatedAt   string `json:"-"`
	DeletedAt   string `json:"-"`
}

type CategoryProduct struct {
	ID         string `json:"id,omitempty"`
	CategoryID string `json:"category_id,omitempty"`
	ProductID  string `json:"product_id,omitempty"`
	CreatedAt  string `json:"-"`
	UpdatedAt  string `json:"-"`
	DeletedAt  string `json:"-"`
}

func ValidateProductModel(currentImages []string, benefitStr float64, productID, brendID, shopID string, priceStr float64, oldPriceStr float64, amountStr, limitAmountStr uint, isNewStr bool, categories []string) (float64, []Images, string, float64, float64, uint, uint, bool, error) {

	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		return 0, []Images{}, "", 0, 0, 0, 0, false, err
	}
	defer func() (float64, []Images, string, float64, float64, uint64, uint64, bool, error) {
		if err := db.Close(); err != nil {
			return 0, []Images{}, "", 0, 0, 0, 0, false, err
		}
		return 0, []Images{}, "", 0, 0, 0, 0, false, nil
	}()

	// validate categies
	if len(categories) == 0 {
		return 0, []Images{}, "", 0, 0, 0, 0, false, errors.New("product category is required")
	}

	// check catrgory id
	for _, v := range categories {
		rawCategory, err := db.Query("SELECT id FROM categories WHERE id = $1 AND deleted_at IS NULL", v)
		if err != nil {
			return 0, []Images{}, "", 0, 0, 0, 0, false, err
		}
		defer func() (float64, []Images, string, float64, float64, uint64, uint64, bool, error) {
			if err := rawCategory.Close(); err != nil {
				return 0, []Images{}, "", 0, 0, 0, 0, false, err
			}
			return 0, []Images{}, "", 0, 0, 0, 0, false, nil
		}()

		var categoryID string

		for rawCategory.Next() {
			if err := rawCategory.Scan(&categoryID); err != nil {
				return 0, []Images{}, "", 0, 0, 0, 0, false, err
			}
		}

		if categoryID == "" {
			return 0, []Images{}, "", 0, 0, 0, 0, false, errors.New("category not found")
		}
	}

	if limitAmountStr > amountStr {
		return 0, []Images{}, "", 0, 0, 0, 0, false, errors.New("cannot be less than limit_amount amount")
	}
	if priceStr < 0 {
		return 0, []Images{}, "", 0, 0, 0, 0, false, errors.New("price cannot be less than zero")
	}

	if benefitStr < 0 {
		return 0, []Images{}, "", 0, 0, 0, 0, false, errors.New("benefit cannot be less than zero")
	}

	// validate old_price
	if oldPriceStr != 0 {
		if oldPriceStr < 0 {
			return 0, []Images{}, "", 0, 0, 0, 0, false, errors.New("old price cannot be less than zero")
		}

		if oldPriceStr < priceStr {
			return 0, []Images{}, "", 0, 0, 0, 0, false, errors.New("cannot be less than oldPrice Price")
		}

	}

	if productID != "" {

		rowMainImage, err := db.Query("SELECT main_image FROM products WHERE deleted_at IS NULL AND id = $1", productID)
		if err != nil {
			return 0, []Images{}, "", 0, 0, 0, 0, false, err
		}
		defer func() (float64, []Images, string, float64, float64, uint64, uint64, bool, error) {
			if err := rowMainImage.Close(); err != nil {
				return 0, []Images{}, "", 0, 0, 0, 0, false, err
			}
			return 0, []Images{}, "", 0, 0, 0, 0, false, nil
		}()

		var mainImage string

		for rowMainImage.Next() {
			if err := rowMainImage.Scan(&mainImage); err != nil {
				return 0, []Images{}, "", 0, 0, 0, 0, false, err
			}
		}

		if mainImage == "" {
			return 0, []Images{}, "", 0, 0, 0, 0, false, errors.New("main image of product not found")
		}

		rowImages, err := db.Query("SELECT id,image FROM images WHERE deleted_at IS NULL AND product_id = $1", productID)
		if err != nil {
			return 0, []Images{}, "", 0, 0, 0, 0, false, err
		}
		defer func() (float64, []Images, string, float64, float64, uint64, uint64, bool, error) {
			if err := rowImages.Close(); err != nil {
				return 0, []Images{}, "", 0, 0, 0, 0, false, err
			}
			return 0, []Images{}, "", 0, 0, 0, 0, false, nil
		}()

		var images []Images

		for rowImages.Next() {
			var image Images

			if err := rowImages.Scan(&image.ID, &image.Image); err != nil {
				return 0, []Images{}, "", 0, 0, 0, 0, false, err
			}

			if len(currentImages) != 0 {
				var foundImage bool
				for _, v1 := range currentImages {
					if image.Image == v1 {
						foundImage = true
					}
				}
				if !foundImage {
					images = append(images, image)
				}
			} else {
				images = append(images, image)
			}

		}

		return benefitStr, images, mainImage, priceStr, oldPriceStr, amountStr, limitAmountStr, isNewStr, nil

	}

	return benefitStr, []Images{}, "", priceStr, oldPriceStr, amountStr, limitAmountStr, isNewStr, nil

}
