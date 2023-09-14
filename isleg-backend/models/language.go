package models

import (
	"context"
	"errors"
	"github/abbgo/isleg/isleg-backend/config"
	"strings"
)

type Language struct {
	ID                            string                          `json:"id,omitempty"`
	NameShort                     string                          `json:"name_short,omitempty" binding:"required"`
	Flag                          string                          `json:"flag,omitempty"`
	CreatedAt                     string                          `json:"-"`
	UpdatedAt                     string                          `json:"-"`
	DeletedAt                     string                          `json:"-"`
	TranslationAfisa              []TranslationAfisa              `json:"translation_afisa,omitempty"`                // one to many
	TranslationProduct            []TranslationProduct            `json:"translation_product,omitempty"`              // one to many
	TranslationCategory           []TranslationCategory           `json:"translation_category,omitempty"`             // one to many
	CompanyAddress                []CompanyAddress                `json:"company_address,omitempty"`                  // one to many
	TranslationDistrict           []TranslationDistrict           `json:"translation_district,omitempty"`             // one to many
	TranslationOrderDates         []TranslationOrderDates         `json:"translation_order_dates,omitempty"`          // one to many
	PaymentTypes                  []PaymentTypes                  `json:"payment_types,omitempty"`                    // one to many
	TranslationAbout              []TranslationAbout              `json:"translation_about,omitempty"`                // one to many
	TranslationBasketPage         []TranslationBasketPage         `json:"translation_basket_page,omitempty"`          // one to many
	TranslationContact            []TranslationContact            `json:"translation_contact,omitempty"`              // one to many
	TranslationFooter             []TranslationFooter             `json:"translation_footer,omitempty"`               // one to many
	TranslationHeader             []TranslationHeader             `json:"translation_header,omitempty"`               // one to many
	TranslationMyInformationPage  []TranslationMyInformationPage  `json:"translation_my_information_page,omitempty"`  // one to many
	TranslationMyOrderPage        []TranslationMyOrderPage        `json:"translation_my_order_page,omitempty"`        // one to many
	TranslationOrderPage          []TranslationOrderPage          `json:"translation_oder_page,omitempty"`            // one to many
	TranslationPayment            []TranslationPayment            `json:"translation_payment,omitempty"`              // one to many
	TranslationSecure             []TranslationSecure             `json:"translation_secure,omitempty"`               // one to many
	TranslationUpdatePasswordPage []TranslationUpdatePasswordPage `json:"translation_update_password_page,omitempty"` // one to many
}

func ValidateLanguage(nameShort, functionType, langID string) (string, error) {
	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		return "", err
	}
	defer db.Close()

	if functionType == "create" {
		if nameShort == "" {
			return "", errors.New("short name is required")
		}

		var oldNameShort string
		if err := db.QueryRow(context.Background(), "SELECT name_short FROM languages WHERE name_short = $1 AND deleted_at IS NULL", strings.ToLower(nameShort)).Scan(&oldNameShort); err != nil {
			return "", err
		}
		if oldNameShort != "" {
			return "", errors.New("short name already exists")
		}
	} else if functionType == "update" {
		// Check if there is a language, id equal to langID
		var id, flag string
		if err := db.QueryRow(context.Background(), "SELECT id,flag FROM languages WHERE id = $1 AND deleted_at IS NULL", langID).Scan(&id, &flag); err != nil {
			return "", err
		}

		if flag == "" {
			return "", errors.New("language not found")
		}

		var oldID string
		if err := db.QueryRow(context.Background(), "SELECT id FROM languages WHERE name_short = $1 AND deleted_at IS NULL", strings.ToLower(nameShort)).Scan(&oldID); err != nil {
			return "", err
		}

		if oldID != "" && id != oldID {
			return "", errors.New("short name already exists")
		}
		return flag, nil
	} else {
		return "", errors.New("invalid function type")
	}
	return "", nil
}
