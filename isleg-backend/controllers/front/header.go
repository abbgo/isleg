package controllers

import (
	backController "github/abbgo/isleg/isleg-backend/controllers/back"
	"github/abbgo/isleg/isleg-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HeaderData struct {
	LogoFavicon       models.CompanySetting           `json:"logo_favicon"`
	TranslationHeader models.TranslationHeader        `json:"translation_header"`
	Languages         []models.Language               `json:"languages"`
	Categories        []backController.ResultCategory `json:"categories"`
}

// GetHeaderData funksiyadan asakdaky maglumatlar alynyar:
// firmanyn logasy we favicony , yagny sazlamalaryn maglumatlary
// header - in terjimesi
// dillerin gysga ady we suraty
// kategoriyalar
func GetHeaderData(c *gin.Context) {

	langID, err := backController.CheckLanguage(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	// get logo and favicon from company setting controller
	logoFavicon, err := backController.GetCompanySettingForHeader()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	// get translation header from translation header controller
	translationHeader, err := backController.GetTranslationHeaderForHeader(langID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	// get all language from language controller
	languages, err := backController.GetAllLanguageForHeader()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	// get all category from category controller
	categories, _, err := backController.GetAllCategoryForHeader(langID, "", "", 0, 0)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	headerData := HeaderData{
		LogoFavicon:       logoFavicon,
		TranslationHeader: translationHeader,
		Languages:         languages,
		Categories:        categories,
	}

	c.JSON(http.StatusOK, gin.H{
		"status":      true,
		"header_data": headerData,
	})

}
