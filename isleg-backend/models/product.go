package models

import (
	"errors"
	"github/abbgo/isleg/isleg-backend/config"
	"strconv"
)

type Product struct {
	ID                 string               `json:"id,omitempty"`
	BrendID            string               `json:"brend_id,omitempty"`
	Price              float64              `json:"price,omitempty"`
	OldPrice           float64              `json:"old_price,omitempty"`
	Amount             uint                 `json:"amount,omitempty"`
	ProductCode        string               `json:"product_code,omitempty"`
	LimitAmount        uint                 `json:"limit_amount,omitempty"`
	IsNew              bool                 `json:"is_new,omitempty"`
	CreatedAt          string               `json:"-"`
	UpdatedAt          string               `json:"-"`
	DeletedAt          string               `json:"-"`
	MainImage          MainImage            `json:"main_image,omitempty"`          // one to one
	Images             []Images             `json:"images,omitempty"`              // one to many
	TranslationProduct []TranslationProduct `json:"translation_product,omitempty"` // one to many
}

type MainImage struct {
	ID        string `json:"id,omitempty"`
	ProductID string `json:"product_id,omitempty"`
	Small     string `json:"small,omitempty"`
	Medium    string `json:"medium,omitempty"`
	Large     string `json:"large,omitempty"`
	CreatedAt string `json:"-"`
	UpdatedAt string `json:"-"`
	DeletedAt string `json:"-"`
}

type Images struct {
	ID        string `json:"id,omitempty"`
	ProductID string `json:"product_id,omitempty"`
	Small     string `json:"small,omitempty"`
	Large     string `json:"large,omitempty"`
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

func ValidateProductModel(brendID, priceStr, oldPriceStr, amountStr, limitAmountStr, productCode, isNewStr string, categories []string) (float64, float64, uint64, uint64, bool, error) {

	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		return 0, 0, 0, 0, false, err
	}
	defer func() (float64, float64, uint64, uint64, bool, error) {
		if err := db.Close(); err != nil {
			return 0, 0, 0, 0, false, err
		}
		return 0, 0, 0, 0, false, nil
	}()

	// validate categies
	if len(categories) == 0 {
		return 0, 0, 0, 0, false, errors.New("product category is required")
	}

	// check catrgory id
	for _, v := range categories {
		rawCategory, err := db.Query("SELECT id FROM categories WHERE id = $1 AND deleted_at IS NULL", v)
		if err != nil {
			return 0, 0, 0, 0, false, err
		}
		defer func() (float64, float64, uint64, uint64, bool, error) {
			if err := rawCategory.Close(); err != nil {
				return 0, 0, 0, 0, false, err
			}
			return 0, 0, 0, 0, false, nil
		}()

		var categoryID string

		for rawCategory.Next() {
			if err := rawCategory.Scan(&categoryID); err != nil {
				return 0, 0, 0, 0, false, err
			}
		}

		if categoryID == "" {
			return 0, 0, 0, 0, false, errors.New("category not found")
		}
	}

	// validate is new
	isNew, err := strconv.ParseBool(isNewStr)
	if err != nil {
		return 0, 0, 0, 0, false, err
	}

	// validatte code of product
	if productCode == "" {
		return 0, 0, 0, 0, false, errors.New("product code is required")
	}

	// validate limit amount of product
	limitAmount, err := strconv.ParseUint(limitAmountStr, 10, 64)
	if err != nil {
		return 0, 0, 0, 0, false, err
	}

	// validate amount of product
	amount, err := strconv.ParseUint(amountStr, 10, 64)
	if err != nil {
		return 0, 0, 0, 0, false, err
	}

	if limitAmount > amount {
		return 0, 0, 0, 0, false, errors.New("cannot be less than limit_amount amount")
	}

	// validatr price of product
	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		return 0, 0, 0, 0, false, err
	}

	// validate old_price
	var oldPrice float64

	if oldPriceStr != "" {
		oldPrice, err = strconv.ParseFloat(oldPriceStr, 64)
		if err != nil {
			return 0, 0, 0, 0, false, err
		}

		if oldPrice < price {
			return 0, 0, 0, 0, false, errors.New("cannot be less than oldPrice Price")
		}

	} else {
		oldPrice = 0
	}

	// validate brend_id
	rowBrend, err := db.Query("SELECT id FROM brends WHERE id = $1 AND deleted_at IS NULL", brendID)
	if err != nil {
		return 0, 0, 0, 0, false, err
	}
	defer func() (float64, float64, uint64, uint64, bool, error) {
		if err := rowBrend.Close(); err != nil {
			return 0, 0, 0, 0, false, err
		}
		return 0, 0, 0, 0, false, nil
	}()

	var brend_id string

	for rowBrend.Next() {
		if err := rowBrend.Scan(&brend_id); err != nil {
			return 0, 0, 0, 0, false, err
		}
	}

	if brend_id == "" {
		return 0, 0, 0, 0, false, errors.New("brend not found")
	}

	return price, oldPrice, amount, limitAmount, isNew, nil

}
