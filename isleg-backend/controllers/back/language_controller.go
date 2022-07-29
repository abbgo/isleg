package controllers

import (
	"github/abbgo/isleg/isleg-backend/config"
	"github/abbgo/isleg/isleg-backend/models"
	"github/abbgo/isleg/isleg-backend/pkg"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type LanguageForHeader struct {
	NameShort string `json:"name_short"`
	Flag      string `json:"flag"`
}

func CreateLanguage(c *gin.Context) {

	// GET DATA FROM REQUEST
	nameShort := c.PostForm("name_short")

	// VALIDATE DATA
	if nameShort == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "language name_short is required",
		})
		return
	}

	// FILE UPLOAD
	newFileName, err := pkg.FileUpload("flag", "language", c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	// CREATE LANGUAGE
	_, err = config.ConnDB().Exec("INSERT INTO languages (name_short,flag) VALUES ($1,$2)", strings.ToLower(nameShort), "uploads/"+newFileName)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	// GET ID OF ADDED LANGUAGE
	lastLandID, err := config.ConnDB().Query("SELECT id FROM languages WHERE deleted_at IS NULL ORDER BY created_at DESC LIMIT 1")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	var langID string
	for lastLandID.Next() {
		if err := lastLandID.Scan(&langID); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}

	// CREATE TRANSLATION HEADER
	_, err = config.ConnDB().Exec("INSERT INTO translation_header (lang_id) VALUES ($1)", langID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	// CREATE TRANSLATION FOOTER
	_, err = config.ConnDB().Exec("INSERT INTO translation_footer (lang_id) VALUES ($1)", langID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	// CREATE TRANSLATION secure
	_, err = config.ConnDB().Exec("INSERT INTO translation_secure (lang_id) VALUES ($1)", langID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	// CREATE TRANSLATION payment
	_, err = config.ConnDB().Exec("INSERT INTO translation_payment (lang_id) VALUES ($1)", langID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	// CREATE TRANSLATION about
	_, err = config.ConnDB().Exec("INSERT INTO translation_about (lang_id) VALUES ($1)", langID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	// CREATE company address
	_, err = config.ConnDB().Exec("INSERT INTO company_address (lang_id) VALUES ($1)", langID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	// CREATE translation_contact
	_, err = config.ConnDB().Exec("INSERT INTO translation_contact (lang_id) VALUES ($1)", langID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	// CREATE translation_my_information_page
	_, err = config.ConnDB().Exec("INSERT INTO translation_my_information_page (lang_id) VALUES ($1)", langID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	// CREATE translation_update_password_page
	_, err = config.ConnDB().Exec("INSERT INTO translation_update_password_page (lang_id) VALUES ($1)", langID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	// GET ID OF ALL CATEGORIES
	var categoryIDs []string
	categoryRows, err := config.ConnDB().Query("SELECT id FROM categories ORDER BY created_at ASC")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	for categoryRows.Next() {
		var categoryID string
		if err := categoryRows.Scan(&categoryID); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
		categoryIDs = append(categoryIDs, categoryID)
	}

	// CREATE TRANSLATION CATEGORY
	for _, v := range categoryIDs {
		_, err = config.ConnDB().Exec("INSERT INTO translation_category (lang_id,category_id) VALUES ($1,$2)", langID, v)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}

	// GET ID OF ALL PRODUCTS
	var productIDs []string
	productRows, err := config.ConnDB().Query("SELECT id FROM products ORDER BY created_at ASC")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	for productRows.Next() {
		var productID string
		if err := productRows.Scan(&productID); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
		productIDs = append(productIDs, productID)
	}

	// CREATE TRANSLATION product
	for _, v := range productIDs {
		_, err = config.ConnDB().Exec("INSERT INTO translation_product (lang_id,product_id) VALUES ($1,$2)", langID, v)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}

	// GET ID OF ALL AFISAS
	var afisaIDs []string
	afisaRows, err := config.ConnDB().Query("SELECT id FROM afisa ORDER BY created_at ASC")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	for afisaRows.Next() {
		var afisaID string
		if err := afisaRows.Scan(&afisaID); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
		afisaIDs = append(afisaIDs, afisaID)
	}

	// CREATE TRANSLATION afisa
	for _, v := range afisaIDs {
		_, err = config.ConnDB().Exec("INSERT INTO translation_afisa (lang_id,afisa_id) VALUES ($1,$2)", langID, v)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}

	// GET ID  OF ALL DISTRICTS
	var districtIDs []string
	districtRows, err := config.ConnDB().Query("SELECT id FROM district ORDER BY created_at ASC")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	for districtRows.Next() {
		var districtID string
		if err := districtRows.Scan(&districtID); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
		districtIDs = append(districtIDs, districtID)
	}

	// CREATE TRANSLATION district
	for _, v := range districtIDs {
		_, err = config.ConnDB().Exec("INSERT INTO translation_district (lang_id,district_id) VALUES ($1,$2)", langID, v)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "language successfully added",
	})

}

func GetAllLanguageForHeader() ([]LanguageForHeader, error) {

	var ls []LanguageForHeader

	// GET Language For Header
	rows, err := config.ConnDB().Query("SELECT name_short,flag FROM languages WHERE deleted_at IS NULL")
	if err != nil {
		return []LanguageForHeader{}, err
	}
	for rows.Next() {
		var l LanguageForHeader
		if err := rows.Scan(&l.NameShort, &l.Flag); err != nil {
			return []LanguageForHeader{}, err
		}
		ls = append(ls, l)
	}

	return ls, nil

}

func GetAllLanguageWithIDAndNameShort() ([]models.Language, error) {

	languageRows, err := config.ConnDB().Query("SELECT id,name_short FROM languages WHERE deleted_at IS NULL ORDER BY created_at ASC")
	if err != nil {
		return []models.Language{}, err
	}

	var languages []models.Language

	for languageRows.Next() {
		var language models.Language
		if err := languageRows.Scan(&language.ID, &language.NameShort); err != nil {
			return []models.Language{}, err
		}
		languages = append(languages, language)
	}

	return languages, nil

}

func GetLangID(langShortName string) (string, error) {

	var langID string

	row, err := config.ConnDB().Query("SELECT id FROM languages WHERE name_short = $1 AND deleted_at IS NULL", langShortName)
	if err != nil {
		return "", err
	}

	for row.Next() {
		if err := row.Scan(&langID); err != nil {
			return "", err
		}
	}

	return langID, nil

}
