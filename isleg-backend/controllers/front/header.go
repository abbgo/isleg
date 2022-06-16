package controllers

import (
	backController "github/abbgo/isleg/isleg-backend/controllers/back"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HeaderData struct {
	LogoFavicon backController.LogoFavicon
}

func GetHeaderData(c *gin.Context) {

	logoFavicon, err := backController.GetCompanySettingForHeader()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	headerData := HeaderData{
		LogoFavicon: logoFavicon,
	}

}
