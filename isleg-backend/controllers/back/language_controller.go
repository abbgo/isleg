package controllers

import (
	"github/abbgo/isleg/isleg-backend/config"
	"github/abbgo/isleg/isleg-backend/models"
	"github/abbgo/isleg/isleg-backend/pkg"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type LanguageForHeader struct {
	NameShort string `json:"name_short"`
	Flag      string `json:"flag"`
}

func CreateLanguage(c *gin.Context) {

	db, err := config.ConnDB()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer db.Close()

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
	resultLang, err := db.Query("INSERT INTO languages (name_short,flag) VALUES ($1,$2)", strings.ToLower(nameShort), "uploads/language/"+newFileName)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer resultLang.Close()

	// GET ID OF ADDED LANGUAGE
	lastLandID, err := db.Query("SELECT id FROM languages WHERE deleted_at IS NULL ORDER BY created_at DESC LIMIT 1")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer lastLandID.Close()

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
	resultTrHeader, err := db.Query("INSERT INTO translation_header (lang_id) VALUES ($1)", langID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer resultTrHeader.Close()

	// CREATE TRANSLATION FOOTER
	resultTrFooter, err := db.Query("INSERT INTO translation_footer (lang_id) VALUES ($1)", langID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer resultTrFooter.Close()

	// CREATE TRANSLATION secure
	resultTRSecure, err := db.Query("INSERT INTO translation_secure (lang_id) VALUES ($1)", langID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer resultTRSecure.Close()

	// CREATE TRANSLATION payment
	resultTRPayment, err := db.Query("INSERT INTO translation_payment (lang_id) VALUES ($1)", langID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer resultTRPayment.Close()

	// CREATE TRANSLATION about
	resultTRABout, err := db.Query("INSERT INTO translation_about (lang_id) VALUES ($1)", langID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer resultTRABout.Close()

	// CREATE company address
	resultComAddress, err := db.Query("INSERT INTO company_address (lang_id) VALUES ($1)", langID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer resultComAddress.Close()

	// CREATE translation_contact
	resultTRContact, err := db.Query("INSERT INTO translation_contact (lang_id) VALUES ($1)", langID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer resultTRContact.Close()

	// CREATE translation_my_information_page
	resultTRMyInfPage, err := db.Query("INSERT INTO translation_my_information_page (lang_id) VALUES ($1)", langID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer resultTRMyInfPage.Close()

	// CREATE translation_update_password_page
	resultUpdMyPass, err := db.Query("INSERT INTO translation_update_password_page (lang_id) VALUES ($1)", langID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer resultUpdMyPass.Close()

	// GET ID OF ALL CATEGORIES
	var categoryIDs []string
	categoryRows, err := db.Query("SELECT id FROM categories ORDER BY created_at ASC")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer categoryRows.Close()

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
		resultTRCategory, err := db.Query("INSERT INTO translation_category (lang_id,category_id) VALUES ($1,$2)", langID, v)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
		defer resultTRCategory.Close()
	}

	// GET ID OF ALL PRODUCTS
	var productIDs []string
	productRows, err := db.Query("SELECT id FROM products ORDER BY created_at ASC")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer productRows.Close()

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
		resultTRProduct, err := db.Query("INSERT INTO translation_product (lang_id,product_id) VALUES ($1,$2)", langID, v)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
		defer resultTRProduct.Close()
	}

	// GET ID OF ALL AFISAS
	var afisaIDs []string
	afisaRows, err := db.Query("SELECT id FROM afisa ORDER BY created_at ASC")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer afisaRows.Close()

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
		resultTRAfisa, err := db.Query("INSERT INTO translation_afisa (lang_id,afisa_id) VALUES ($1,$2)", langID, v)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
		defer resultTRAfisa.Close()
	}

	// GET ID  OF ALL DISTRICTS
	var districtIDs []string
	districtRows, err := db.Query("SELECT id FROM district ORDER BY created_at ASC")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer districtRows.Close()

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
		resultTRDistrict, err := db.Query("INSERT INTO translation_district (lang_id,district_id) VALUES ($1,$2)", langID, v)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
		defer resultTRDistrict.Close()
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "language successfully added",
	})

}

func UpdateLanguageByID(c *gin.Context) {

	db, err := config.ConnDB()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer db.Close()

	langID := c.Param("id")
	nameShort := c.PostForm("name_short")
	var fileName string

	rowFlag, err := db.Query("SELECT flag FROM languages WHERE id = $1 AND deleted_at IS NULL", langID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer rowFlag.Close()

	var flag string

	for rowFlag.Next() {
		if err := rowFlag.Scan(&flag); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}

	if flag == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "language not found",
		})
		return
	}

	if nameShort == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "language name_short is required",
		})
		return
	}

	file, err := c.FormFile("flag")
	if err != nil {
		fileName = flag
	} else {
		extensionFile := filepath.Ext(file.Filename)

		if extensionFile != ".jpg" && extensionFile != ".jpeg" && extensionFile != ".png" && extensionFile != ".gif" {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": "the file must be an image",
			})
			return
		}

		newFileName := uuid.New().String() + extensionFile
		c.SaveUploadedFile(file, "./uploads/language/"+newFileName)

		if err := os.Remove("./" + flag); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}

		fileName = "uploads/language/" + newFileName
	}

	currentTime := time.Now()

	resultLang, err := db.Query("UPDATE languages SET name_short = $1 , flag = $2 , updated_at = $4 WHERE id = $3", nameShort, fileName, langID, currentTime)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer resultLang.Close()

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "language successfully updated",
	})

}

func GetLanguageByID(c *gin.Context) {

	db, err := config.ConnDB()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer db.Close()

	langID := c.Param("id")

	rowLanguage, err := db.Query("SELECT name_short,flag FROM languages WHERE id = $1 AND deleted_at IS NULL", langID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer rowLanguage.Close()

	var lang LanguageForHeader

	for rowLanguage.Next() {
		if err := rowLanguage.Scan(&lang.NameShort, &lang.Flag); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}

	if lang.NameShort == "" || lang.Flag == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "language not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":   true,
		"language": lang,
	})

}

func GetLanguages(c *gin.Context) {
	languages, err := GetAllLanguageForHeader()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":   true,
		"language": languages,
	})

}

func DeleteLanguageByID(c *gin.Context) {

	db, err := config.ConnDB()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer db.Close()

	langID := c.Param("id")

	rowFlag, err := db.Query("SELECT flag FROM languages WHERE id = $1 AND deleted_at IS NULL", langID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer rowFlag.Close()

	var flag string

	for rowFlag.Next() {
		if err := rowFlag.Scan(&flag); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}

	if flag == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "language not found",
		})
		return
	}

	currentTime := time.Now()

	resutlTRLand, err := db.Query("UPDATE languages SET deleted_at = $1 WHERE id = $2", currentTime, langID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer resutlTRLand.Close()

	resultTRHeader, err := db.Query("UPDATE translation_header SET deleted_at = $1 WHERE lang_id = $2", currentTime, langID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer resultTRHeader.Close()

	resultTRFooter, err := db.Query("UPDATE translation_footer SET deleted_at = $1 WHERE lang_id = $2", currentTime, langID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer resultTRFooter.Close()

	resulttTrSECURE, err := db.Query("UPDATE translation_secure SET deleted_at = $1 WHERE lang_id = $2", currentTime, langID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer resulttTrSECURE.Close()

	resultTRPayment, err := db.Query("UPDATE translation_payment SET deleted_at = $1 WHERE lang_id = $2", currentTime, langID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer resultTRPayment.Close()

	resultTRABout, err := db.Query("UPDATE translation_about SET deleted_at = $1 WHERE lang_id = $2", currentTime, langID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer resultTRABout.Close()

	resultComAdddres, err := db.Query("UPDATE company_address SET deleted_at = $1 WHERE lang_id = $2", currentTime, langID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer resultComAdddres.Close()

	resultTRContact, err := db.Query("UPDATE translation_contact SET deleted_at = $1 WHERE lang_id = $2", currentTime, langID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer resultTRContact.Close()

	resultMyInfPage, err := db.Query("UPDATE translation_my_information_page SET deleted_at = $1 WHERE lang_id = $2", currentTime, langID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer resultMyInfPage.Close()

	resultTraUpdatePassPage, err := db.Query("UPDATE translation_update_password_page SET deleted_at = $1 WHERE lang_id = $2", currentTime, langID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer resultTraUpdatePassPage.Close()

	resultTRCategory, err := db.Query("UPDATE translation_category SET deleted_at = $1 WHERE lang_id = $2", currentTime, langID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer resultTRCategory.Close()

	resultTRProduct, err := db.Query("UPDATE translation_product SET deleted_at = $1 WHERE lang_id = $2", currentTime, langID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer resultTRProduct.Close()

	resultTRAfisa, err := db.Query("UPDATE translation_afisa SET deleted_at = $1 WHERE lang_id = $2", currentTime, langID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer resultTRAfisa.Close()

	resultTrDistrict, err := db.Query("UPDATE translation_district SET deleted_at = $1 WHERE lang_id = $2", currentTime, langID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer resultTrDistrict.Close()

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "language successfully deleted",
	})

}

func RestoreLanguageByID(c *gin.Context) {

	db, err := config.ConnDB()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer db.Close()

	langID := c.Param("id")

	rowFlag, err := db.Query("SELECT flag FROM languages WHERE id = $1 AND deleted_at IS NOT NULL", langID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer rowFlag.Close()

	var flag string

	for rowFlag.Next() {
		if err := rowFlag.Scan(&flag); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}

	if flag == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "language not found",
		})
		return
	}

	resultLang, err := db.Query("UPDATE languages SET deleted_at = NULL WHERE id = $1", langID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer resultLang.Close()

	resultTRHeader, err := db.Query("UPDATE translation_header SET deleted_at = NULL WHERE lang_id = $1", langID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer resultTRHeader.Close()

	resultTRFooter, err := db.Query("UPDATE translation_footer SET deleted_at = NULL WHERE lang_id = $1", langID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer resultTRFooter.Close()

	resultTRSecure, err := db.Query("UPDATE translation_secure SET deleted_at = NULL WHERE lang_id = $1", langID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer resultTRSecure.Close()

	resultTrPayment, err := db.Query("UPDATE translation_payment SET deleted_at = NULL WHERE lang_id = $1", langID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer resultTrPayment.Close()

	resultTRABut, err := db.Query("UPDATE translation_about SET deleted_at = NULL WHERE lang_id = $1", langID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer resultTRABut.Close()

	resultComAddres, err := db.Query("UPDATE company_address SET deleted_at = NULL WHERE lang_id = $1", langID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer resultComAddres.Close()

	resultTRContact, err := db.Query("UPDATE translation_contact SET deleted_at = NULL WHERE lang_id = $1", langID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer resultTRContact.Close()

	resultMyInfPage, err := db.Query("UPDATE translation_my_information_page SET deleted_at = NULL WHERE lang_id = $1", langID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer resultMyInfPage.Close()

	resultUpPassPage, err := db.Query("UPDATE translation_update_password_page SET deleted_at = NULL WHERE lang_id = $1", langID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer resultUpPassPage.Close()

	resultTrCategory, err := db.Query("UPDATE translation_category SET deleted_at = NULL WHERE lang_id = $1", langID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer resultTrCategory.Close()

	resultTRProduct, err := db.Query("UPDATE translation_product SET deleted_at = NULL WHERE lang_id = $1", langID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer resultTRProduct.Close()

	resultTRAfisa, err := db.Query("UPDATE translation_afisa SET deleted_at = NULL WHERE lang_id = $1", langID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer resultTRAfisa.Close()

	resultTrDistrcit, err := db.Query("UPDATE translation_district SET deleted_at = NULL WHERE lang_id = $1", langID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer resultTrDistrcit.Close()

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "language successfully restored",
	})

}

func DeletePermanentlyLanguageByID(c *gin.Context) {

	db, err := config.ConnDB()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer db.Close()

	langID := c.Param("id")

	rowFlag, err := db.Query("SELECT flag FROM languages WHERE id = $1 AND deleted_at IS NOT NULL", langID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer rowFlag.Close()

	var flag string

	for rowFlag.Next() {
		if err := rowFlag.Scan(&flag); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}

	if flag == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "language not found",
		})
		return
	}

	if err := os.Remove("./" + flag); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	resultLang, err := db.Query("DELETE FROM languages WHERE id = $1", langID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer resultLang.Close()

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "language successfully deleted",
	})

}

func GetAllLanguageForHeader() ([]LanguageForHeader, error) {

	db, err := config.ConnDB()
	if err != nil {
		return []LanguageForHeader{}, nil
	}
	defer db.Close()

	var ls []LanguageForHeader

	// GET Language For Header
	rows, err := db.Query("SELECT name_short,flag FROM languages WHERE deleted_at IS NULL")
	if err != nil {
		return []LanguageForHeader{}, err
	}
	defer rows.Close()

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

	db, err := config.ConnDB()
	if err != nil {
		return []models.Language{}, nil
	}
	defer db.Close()

	languageRows, err := db.Query("SELECT id,name_short FROM languages WHERE deleted_at IS NULL ORDER BY created_at ASC")
	if err != nil {
		return []models.Language{}, err
	}
	defer languageRows.Close()

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

	db, err := config.ConnDB()
	if err != nil {
		return "", nil
	}
	defer db.Close()

	var langID string

	row, err := db.Query("SELECT id FROM languages WHERE name_short = $1 AND deleted_at IS NULL", langShortName)
	if err != nil {
		return "", err
	}
	defer row.Close()

	for row.Next() {
		if err := row.Scan(&langID); err != nil {
			return "", err
		}
	}

	return langID, nil

}

func CheckLanguage(c *gin.Context) (string, error) {

	// GET DATA FROM ROUTE PARAMETER
	langShortName := c.Param("lang")

	// GET ID OFF LANGUAGE
	langID, err := GetLangID(langShortName)
	if err != nil {
		return "", err
	}

	return langID, nil

}
