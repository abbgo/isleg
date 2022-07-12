package controllers

import (
	backController "github/abbgo/isleg/isleg-backend/controllers/back"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HeaderData struct {
	LogoFavicon       backController.LogoFavicon                `json:"logo_favicon"`
	TranslationHeader backController.TranslationHeaderForHeader `json:"translation_header"`
	Languages         []backController.LanguageForHeader        `json:"languages"`
	Categories        []backController.ResultCategory           `json:"categories"`
}

func GetHeaderData(c *gin.Context) {

	// GET DATA FROM ROUTE PARAMETER
	langShortName := c.Param("lang")

	// GET language id
	langID, err := backController.GetLangID(langShortName)
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
	categories, err := backController.GetAllCategoryForHeader(langID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
			// "message": "yalnys bar",
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
