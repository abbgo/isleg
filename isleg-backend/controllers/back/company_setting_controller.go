package controllers

import (
	"github/abbgo/isleg/isleg-backend/config"
	"github/abbgo/isleg/isleg-backend/models"
	"github/abbgo/isleg/isleg-backend/pkg"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LogoFavicon struct {
	Logo    string `json:"logo"`
	Favicon string `json:"favicon"`
}

func CreateCompanySetting(c *gin.Context) {

	// GET DATA FROM REQUEST
	email := c.PostForm("email")
	instagram := c.PostForm("instagram")

	// VALIDATE DATA
	err := models.ValidateCompanySettingData(email, instagram)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	// FILE UPLOAD

	// LOGO
	newFileNameLogo, err := pkg.FileUpload("logo", "logo", c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	// FAVICON
	newFileNameFavicon, err := pkg.FileUpload("favicon", "favicon", c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	// CREATE COMPANY SETTING
	_, err = config.ConnDB().Exec("INSERT INTO company_setting (logo,favicon,email,instagram) VALUES ($1,$2,$3,$4)", "uploads/"+newFileNameLogo, "uploads/"+newFileNameFavicon, email, instagram)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "company setting successfully added",
	})

}

func GetCompanySettingForHeader() (LogoFavicon, error) {

	var logoFavicon LogoFavicon

	// GET LOGO AND FAVICON
	row, err := config.ConnDB().Query("SELECT logo,favicon FROM company_setting WHERE deleted_at IS NULL ORDER BY created_at ASC LIMIT 1")
	if err != nil {
		return LogoFavicon{}, err
	}
	for row.Next() {
		if err := row.Scan(&logoFavicon.Logo, &logoFavicon.Favicon); err != nil {
			return LogoFavicon{}, err
		}
	}

	return logoFavicon, nil

}
