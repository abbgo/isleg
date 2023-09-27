package controllers

import (
	backController "github/abbgo/isleg/isleg-backend/controllers/back"
	"github/abbgo/isleg/isleg-backend/helpers"
	"github/abbgo/isleg/isleg-backend/models"
	"github/abbgo/isleg/isleg-backend/pkg"
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
	langID, err := pkg.ValidateMiddlewareData(c, "lang_id")
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	// get logo and favicon from company setting controller
	logoFavicon, err := backController.GetCompanySettingForHeader()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	// get translation header from translation header controller
	translationHeader, err := backController.GetTranslationHeaderForHeader(langID)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	// get all language from language controller
	languages, err := backController.GetAllLanguageForHeader()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	// get all category from category controller
	categories, _, err := backController.GetAllCategoryForHeader(langID, "", "", 0, 0, true)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
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
