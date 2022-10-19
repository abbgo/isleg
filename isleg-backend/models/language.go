package models

type Language struct {
	ID                            string                          `json:"id,omitempty"`
	NameShort                     string                          `json:"name_short,omitempty"`
	Flag                          string                          `json:"flag,omitempty"`
	CreatedAt                     string                          `json:"-"`
	UpdatedAt                     string                          `json:"-"`
	DeletedAt                     string                          `json:"-"`
	TranslationAfisa              []TranslationAfisa              `json:"translation_afisa,omitempty"`
	TranslationProduct            []TranslationProduct            `json:"translation_product,omitempty"`
	TranslationCategory           []TranslationCategory           `json:"translation_category,omitempty"`
	CompanyAddress                []CompanyAddress                `json:"company_address,omitempty"`
	TranslationDistrict           []TranslationDistrict           `json:"translation_district,omitempty"`
	TranslationOrderDates         []TranslationOrderDates         `json:"translation_order_dates,omitempty"`
	PaymentTypes                  []PaymentTypes                  `json:"payment_types,omitempty"`
	TranslationAbout              []TranslationAbout              `json:"translation_about,omitempty"`
	TranslationBasketPage         []TranslationBasketPage         `json:"translation_basket_page,omitempty"`
	TranslationContact            []TranslationContact            `json:"translation_contact,omitempty"`
	TranslationFooter             []TranslationFooter             `json:"translation_footer,omitempty"`
	TranslationHeader             []TranslationHeader             `json:"translation_header,omitempty"`
	TranslationMyInformationPage  []TranslationMyInformationPage  `json:"translation_my_information_page,omitempty"`
	TranslationMyOrderPage        []TranslationMyOrderPage        `json:"translation_my_order_page,omitempty"`
	TranslationOrderPage          []TranslationOrderPage          `json:"translation_oder_page,omitempty"`
	TranslationPayment            []TranslationPayment            `json:"translation_payment,omitempty"`
	TranslationSecure             []TranslationSecure             `json:"translation_secure,omitempty"`
	TranslationUpdatePasswordPage []TranslationUpdatePasswordPage `json:"translation_update_password_page,omitempty"`
}
