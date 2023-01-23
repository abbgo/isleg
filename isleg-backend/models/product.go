package models

import (
	"database/sql"
	"errors"
	"github/abbgo/isleg/isleg-backend/config"
	"strconv"
)

type Product struct {
	ID                 string               `json:"id,omitempty"`
	BrendID            sql.NullString       `json:"brend_id,omitempty"`
	ShopID             sql.NullString       `json:"shop_id,omitempty"`
	Price              float64              `json:"price,omitempty"`
	OldPrice           float64              `json:"old_price"`
	Percentage         float64              `json:"percentage"`
	Amount             uint                 `json:"amount,omitempty"`
	LimitAmount        uint                 `json:"limit_amount,omitempty"`
	IsNew              bool                 `json:"is_new,omitempty"`
	CreatedAt          string               `json:"-"`
	UpdatedAt          string               `json:"-"`
	DeletedAt          string               `json:"-"`
	MainImage          string               `json:"main_image,omitempty"`
	Images             []Images             `json:"images,omitempty"`              // one to many
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

func ValidateProductModel(productID, brendID, shopID, priceStr, oldPriceStr, amountStr, limitAmountStr, isNewStr string, categories []string) ([]Images, string, float64, float64, uint64, uint64, bool, error) {

	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		return []Images{}, "", 0, 0, 0, 0, false, err
	}
	defer func() ([]Images, string, float64, float64, uint64, uint64, bool, error) {
		if err := db.Close(); err != nil {
			return []Images{}, "", 0, 0, 0, 0, false, err
		}
		return []Images{}, "", 0, 0, 0, 0, false, nil
	}()

	// validate categies
	if len(categories) == 0 {
		return []Images{}, "", 0, 0, 0, 0, false, errors.New("product category is required")
	}

	// check catrgory id
	for _, v := range categories {
		rawCategory, err := db.Query("SELECT id FROM categories WHERE id = $1 AND deleted_at IS NULL", v)
		if err != nil {
			return []Images{}, "", 0, 0, 0, 0, false, err
		}
		defer func() ([]Images, string, float64, float64, uint64, uint64, bool, error) {
			if err := rawCategory.Close(); err != nil {
				return []Images{}, "", 0, 0, 0, 0, false, err
			}
			return []Images{}, "", 0, 0, 0, 0, false, nil
		}()

		var categoryID string

		for rawCategory.Next() {
			if err := rawCategory.Scan(&categoryID); err != nil {
				return []Images{}, "", 0, 0, 0, 0, false, err
			}
		}

		if categoryID == "" {
			return []Images{}, "", 0, 0, 0, 0, false, errors.New("category not found")
		}
	}

	// validate is new
	isNew, err := strconv.ParseBool(isNewStr)
	if err != nil {
		return []Images{}, "", 0, 0, 0, 0, false, err
	}

	// validate limit amount of product
	limitAmount, err := strconv.ParseUint(limitAmountStr, 10, 64)
	if err != nil {
		return []Images{}, "", 0, 0, 0, 0, false, err
	}

	// validate amount of product
	amount, err := strconv.ParseUint(amountStr, 10, 64)
	if err != nil {
		return []Images{}, "", 0, 0, 0, 0, false, err
	}

	if limitAmount > amount {
		return []Images{}, "", 0, 0, 0, 0, false, errors.New("cannot be less than limit_amount amount")
	}

	// validatr price of product
	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		return []Images{}, "", 0, 0, 0, 0, false, err
	}

	// validate old_price
	var oldPrice float64

	if oldPriceStr != "" {
		oldPrice, err = strconv.ParseFloat(oldPriceStr, 64)
		if err != nil {
			return []Images{}, "", 0, 0, 0, 0, false, err
		}

		if oldPrice < price {
			return []Images{}, "", 0, 0, 0, 0, false, errors.New("cannot be less than oldPrice Price")
		}

	} else {
		oldPrice = 0
	}

	// validate shop_id
	if shopID != "" {
		rowShop, err := db.Query("SELECT id FROM shops WHERE id = $1 AND deleted_at IS NULL", shopID)
		if err != nil {
			return []Images{}, "", 0, 0, 0, 0, false, err
		}
		defer func() ([]Images, string, float64, float64, uint64, uint64, bool, error) {
			if err := rowShop.Close(); err != nil {
				return []Images{}, "", 0, 0, 0, 0, false, err
			}
			return []Images{}, "", 0, 0, 0, 0, false, nil
		}()

		var shop_id string

		for rowShop.Next() {
			if err := rowShop.Scan(&shop_id); err != nil {
				return []Images{}, "", 0, 0, 0, 0, false, err
			}
		}

		if shop_id == "" {
			return []Images{}, "", 0, 0, 0, 0, false, errors.New("shop not found")
		}
	}

	// validate brend_id
	if brendID != "" {
		rowBrend, err := db.Query("SELECT id FROM brends WHERE id = $1 AND deleted_at IS NULL", brendID)
		if err != nil {
			return []Images{}, "", 0, 0, 0, 0, false, err
		}
		defer func() ([]Images, string, float64, float64, uint64, uint64, bool, error) {
			if err := rowBrend.Close(); err != nil {
				return []Images{}, "", 0, 0, 0, 0, false, err
			}
			return []Images{}, "", 0, 0, 0, 0, false, nil
		}()

		var brend_id string

		for rowBrend.Next() {
			if err := rowBrend.Scan(&brend_id); err != nil {
				return []Images{}, "", 0, 0, 0, 0, false, err
			}
		}

		if brend_id == "" {
			return []Images{}, "", 0, 0, 0, 0, false, errors.New("brend not found")
		}
	}

	if productID != "" {

		rowProduct, err := db.Query("SELECT id FROM products WHERE id = $1 AND deleted_at IS NULL", productID)
		if err != nil {
			return []Images{}, "", 0, 0, 0, 0, false, err
		}
		defer func() ([]Images, string, float64, float64, uint64, uint64, bool, error) {
			if err := rowProduct.Close(); err != nil {
				return []Images{}, "", 0, 0, 0, 0, false, err
			}
			return []Images{}, "", 0, 0, 0, 0, false, nil
		}()

		var id string

		for rowProduct.Next() {
			if err := rowProduct.Scan(&id); err != nil {
				return []Images{}, "", 0, 0, 0, 0, false, err
			}
		}

		if id == "" {
			return []Images{}, "", 0, 0, 0, 0, false, errors.New("record not found")
		}

		rowMainImage, err := db.Query("SELECT main_image FROM products WHERE deleted_at IS NULL AND id = $1", productID)
		if err != nil {
			return []Images{}, "", 0, 0, 0, 0, false, err
		}
		defer func() ([]Images, string, float64, float64, uint64, uint64, bool, error) {
			if err := rowMainImage.Close(); err != nil {
				return []Images{}, "", 0, 0, 0, 0, false, err
			}
			return []Images{}, "", 0, 0, 0, 0, false, nil
		}()

		var mainImage string

		for rowMainImage.Next() {
			if err := rowMainImage.Scan(&mainImage); err != nil {
				return []Images{}, "", 0, 0, 0, 0, false, err
			}
		}

		if mainImage == "" {
			return []Images{}, "", 0, 0, 0, 0, false, errors.New("main image of product not found")
		}

		rowImages, err := db.Query("SELECT image FROM images WHERE deleted_at IS NULL AND product_id = $1", productID)
		if err != nil {
			return []Images{}, "", 0, 0, 0, 0, false, err
		}
		defer func() ([]Images, string, float64, float64, uint64, uint64, bool, error) {
			if err := rowImages.Close(); err != nil {
				return []Images{}, "", 0, 0, 0, 0, false, err
			}
			return []Images{}, "", 0, 0, 0, 0, false, nil
		}()

		var images []Images

		for rowImages.Next() {
			var image Images

			if err := rowImages.Scan(&image.Image); err != nil {
				return []Images{}, "", 0, 0, 0, 0, false, err
			}

			images = append(images, image)
		}

		return images, mainImage, price, oldPrice, amount, limitAmount, isNew, nil

	}

	return []Images{}, "", price, oldPrice, amount, limitAmount, isNew, nil

}
