package controllers

import (
	"github/abbgo/isleg/isleg-backend/config"
	"github/abbgo/isleg/isleg-backend/models"
	"github/abbgo/isleg/isleg-backend/pkg"
	"net/http"
	"os"
	"path/filepath"
	"time"

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

	db, err := config.ConnDB()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer func() {
		if err := db.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

	// GET DATA FROM REQUEST
	email := c.PostForm("email")
	instagram := c.PostForm("instagram")

	// VALIDATE DATA
	err = models.ValidateCompanySettingData(email, instagram)
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
	resultComSetting, err := db.Query("INSERT INTO company_setting (logo,favicon,email,instagram) VALUES ($1,$2,$3,$4)", "uploads/setting/"+newFileNameLogo, "uploads/setting/"+newFileNameFavicon, email, instagram)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer func() {
		if err := resultComSetting.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "company setting successfully added",
	})

}

func UpdateCompanySetting(c *gin.Context) {

	db, err := config.ConnDB()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer func() {
		if err := db.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

	email := c.PostForm("email")
	instagram := c.PostForm("instagram")
	var logoName, faviconName string

	rowComSet, err := db.Query("SELECT logo,favicon FROM company_setting WHERE deleted_at IS NULL ORDER BY created_at ASC LIMIT 1")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer func() {
		if err := rowComSet.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

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

	currentTime := time.Now()

	resultComPSETTING, err := db.Query("UPDATE company_setting SET logo = $1,favicon=$2,email=$3,instagram=$4,updated_at=$5", logoName, faviconName, email, instagram, currentTime)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer func() {
		if err := resultComPSETTING.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "company setting successfully updated",
	})

}

func GetCompanySetting(c *gin.Context) {

	db, err := config.ConnDB()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer db.Close()

	rowComSet, err := db.Query("SELECT logo,favicon,email,instagram FROM company_setting WHERE deleted_at IS NULL ORDER BY created_at ASC LIMIT 1")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer rowComSet.Close()

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

	db, err := config.ConnDB()
	if err != nil {

		return LogoFavicon{}, nil
	}
	defer db.Close()

	var logoFavicon LogoFavicon

	// GET LOGO AND FAVICON
	row, err := db.Query("SELECT logo,favicon FROM company_setting WHERE deleted_at IS NULL ORDER BY created_at ASC LIMIT 1")
	if err != nil {
		return LogoFavicon{}, err
	}
	defer row.Close()

	for row.Next() {
		if err := row.Scan(&logoFavicon.Logo, &logoFavicon.Favicon); err != nil {
			return LogoFavicon{}, err
		}
	}

	return logoFavicon, nil

}
