package controllers

import (
	"net/http"

	backController "github/abbgo/isleg/isleg-backend/controllers/back"

	"github.com/gin-gonic/gin"
)

// GetFooterData funksiya asakdaky maglumatlary getirip beryar:
// footer - in terjimesini
func GetFooterData(c *gin.Context) {

	langID, err := backController.CheckLanguage(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	// get translation footer from translation footer controller
	translationFooter, err := backController.GetTranslationFooter(langID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":      true,
		"footer_data": translationFooter,
	})

}
