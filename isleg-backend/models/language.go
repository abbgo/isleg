package models

type Language struct {
	ID                            string                          `json:"id,omitempty"`
	NameShort                     string                          `json:"name_short,omitempty"`
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
