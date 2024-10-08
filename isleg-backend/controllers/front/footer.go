package controllers

import (
	"net/http"

	backController "github/abbgo/isleg/isleg-backend/controllers/back"
	"github/abbgo/isleg/isleg-backend/helpers"
	"github/abbgo/isleg/isleg-backend/pkg"

	"github.com/gin-gonic/gin"
)

// GetFooterData funksiya asakdaky maglumatlary getirip beryar:
// footer - in terjimesini
func GetFooterData(c *gin.Context) {
	langID, err := pkg.ValidateMiddlewareData(c, "lang_id")
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	// get translation footer from translation footer controller
	translationFooter, err := backController.GetTranslationFooter(langID)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":      true,
		"footer_data": translationFooter,
	})
}
