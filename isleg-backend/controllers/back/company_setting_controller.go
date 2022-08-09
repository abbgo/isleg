package controllers

import (
	"github/abbgo/isleg/isleg-backend/config"
	"github/abbgo/isleg/isleg-backend/models"
	"github/abbgo/isleg/isleg-backend/pkg"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type LogoFavicon struct {
	Logo    string `json:"logo"`
	Favicon string `json:"favicon"`
}

type ComSet struct {
	Logo      string `json:"logo"`
	Favicon   string `json:"favicon"`
	Email     string `json:"email"`
	Instagram string `json:"instagram"`
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
	newFileNameLogo, err := pkg.FileUpload("logo", "setting", c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	// FAVICON
	newFileNameFavicon, err := pkg.FileUpload("favicon", "setting", c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	// CREATE COMPANY SETTING
	_, err = config.ConnDB().Exec("INSERT INTO company_setting (logo,favicon,email,instagram) VALUES ($1,$2,$3,$4)", "uploads/setting/"+newFileNameLogo, "uploads/setting/"+newFileNameFavicon, email, instagram)
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

func UpdateCompanySetting(c *gin.Context) {

	email := c.PostForm("email")
	instagram := c.PostForm("instagram")
	var logoName, faviconName string

	rowComSet, err := config.ConnDB().Query("SELECT logo,favicon FROM company_setting WHERE deleted_at IS NULL ORDER BY created_at ASC LIMIT 1")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	var logo, favicon string

	for rowComSet.Next() {
		if err := rowComSet.Scan(&logo, &favicon); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}

	if logo == "" || favicon == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "record not found",
		})
		return
	}

	err = models.ValidateCompanySettingData(email, instagram)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	fileLogo, err := c.FormFile("logo")
	if err != nil {
		logoName = logo
	} else {
		extensionFile := filepath.Ext(fileLogo.Filename)

		if extensionFile != ".jpg" && extensionFile != ".jpeg" && extensionFile != ".png" && extensionFile != ".gif" {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": "the file must be an image",
			})
			return
		}

		newFileName := uuid.New().String() + extensionFile
		c.SaveUploadedFile(fileLogo, "./uploads/setting/"+newFileName)

		if err := os.Remove("./" + logo); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}

		logoName = "uploads/setting/" + newFileName
	}

	fileFavicon, err := c.FormFile("favicon")
	if err != nil {
		faviconName = favicon
	} else {
		extensionFile := filepath.Ext(fileFavicon.Filename)

		if extensionFile != ".jpg" && extensionFile != ".jpeg" && extensionFile != ".png" && extensionFile != ".gif" {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": "the file must be an image",
			})
			return
		}

		newFileName := uuid.New().String() + extensionFile
		c.SaveUploadedFile(fileFavicon, "./uploads/setting/"+newFileName)

		if err := os.Remove("./" + favicon); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}

		faviconName = "uploads/setting/" + newFileName
	}

	_, err = config.ConnDB().Exec("UPDATE company_setting SET logo = $1,favicon=$2,email=$3,instagram=$4", logoName, faviconName, email, instagram)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "company setting successfully updated",
	})

}

func GetCompanySetting(c *gin.Context) {

	rowComSet, err := config.ConnDB().Query("SELECT logo,favicon,email,instagram FROM company_setting WHERE deleted_at IS NULL ORDER BY created_at ASC LIMIT 1")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	var comSet ComSet

	for rowComSet.Next() {
		if err := rowComSet.Scan(&comSet.Logo, &comSet.Favicon, &comSet.Email, &comSet.Instagram); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}

	if comSet.Logo == "" || comSet.Favicon == "" || comSet.Email == "" || comSet.Instagram == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "record not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":          true,
		"company_setting": comSet,
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
