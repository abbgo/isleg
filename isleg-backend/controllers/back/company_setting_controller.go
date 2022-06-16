package controllers

import (
	"github/abbgo/isleg/isleg-backend/config"
	"github/abbgo/isleg/isleg-backend/pkg"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
	emailResult := pkg.IsEmailValid(email)
	if email == "" || !emailResult {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "email address is required or it doesn't match",
		})
		return
	}
	if instagram == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "instagram is required",
		})
		return
	}
	// FILE UPLOAD
	// LOGO
	fileLogo, err := c.FormFile("logo_path")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	extensionLogo := filepath.Ext(fileLogo.Filename)
	// VALIDATE IMAGE
	if extensionLogo != ".jpg" && extensionLogo != ".jpeg" && extensionLogo != ".png" && extensionLogo != ".gif" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "the file must be an image.",
		})
		return
	}
	newFileNameLogo := "logo" + uuid.New().String() + extensionLogo
	c.SaveUploadedFile(fileLogo, "./uploads/"+newFileNameLogo)

	// FAVICON
	fileFavicon, err := c.FormFile("favicon_path")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	extensionFavicon := filepath.Ext(fileFavicon.Filename)
	// VALIDATE IMAGE
	if extensionFavicon != ".jpg" && extensionFavicon != ".jpeg" && extensionFavicon != ".png" && extensionFavicon != ".gif" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "the file must be an image.",
		})
		return
	}
	newFileNameFavicon := "favicon" + uuid.New().String() + extensionFavicon
	c.SaveUploadedFile(fileFavicon, "./uploads/"+newFileNameFavicon)

	// CREATE COMPANY SETTING
	_, err = config.ConnDB().Exec("INSERT INTO company_setting (logo_path,favicon_path,email,instagram) VALUES ($1,$2,$3,$4)", "uploads/"+newFileNameLogo, "uploads/"+newFileNameFavicon, email, instagram)
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
	row, err := config.ConnDB().Query("SELECT logo_path,favicon_path FROM company_setting ORDER BY created_at ASC LIMIT 1")
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
