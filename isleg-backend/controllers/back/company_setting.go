package controllers

import (
	"context"
	"github/abbgo/isleg/isleg-backend/config"
	"github/abbgo/isleg/isleg-backend/helpers"
	"github/abbgo/isleg/isleg-backend/models"
	"github/abbgo/isleg/isleg-backend/pkg"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateCompanySetting(c *gin.Context) {
	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	// get email and instagram of company_setting from request
	email := c.PostForm("email")
	instagram := c.PostForm("instagram")
	imo := c.PostForm("imo")

	// validate email and instagram
	err = models.ValidateCompanySettingData(email, instagram, imo)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	// upload logo
	newFileNameLogo, err := pkg.FileUpload("logo", "setting", "image", c)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	// upload favicon
	newFileNameFavicon, err := pkg.FileUpload("favicon", "setting", "image", c)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	// add data to database
	_, err = db.Query(context.Background(), "INSERT INTO company_setting (logo,favicon,email,instagram,imo) VALUES ($1,$2,$3,$4,$5)", "uploads/setting/"+newFileNameLogo, "uploads/setting/"+newFileNameFavicon, email, instagram, imo)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "company setting successfully added",
	})
}

func UpdateCompanySetting(c *gin.Context) {
	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	// get data from request
	email := c.PostForm("email")
	instagram := c.PostForm("instagram")
	imo := c.PostForm("imo")

	var logoName, faviconName string

	// Check if there is a company_setting and get logo and favicon
	var logo, favicon string
	db.QueryRow(context.Background(), "SELECT logo,favicon FROM company_setting WHERE deleted_at IS NULL ORDER BY created_at ASC LIMIT 1").Scan(&logo, &favicon)
	if logo == "" || favicon == "" {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  false,
			"message": "record not found",
		})
		return
	}

	// validate email and instagram
	err = models.ValidateCompanySettingData(email, instagram, imo)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	// upload logo
	logoName, err = pkg.FileUploadForUpdate("logo", "setting", logo, c)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	// upload favicon
	faviconName, err = pkg.FileUploadForUpdate("favicon", "setting", favicon, c)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	// update data in database
	_, err = db.Exec(context.Background(), "UPDATE company_setting SET logo = $1, favicon=$2, email=$3, instagram=$4, imo=$5", logoName, faviconName, email, instagram, imo)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "company setting successfully updated",
	})
}

func GetCompanySetting(c *gin.Context) {
	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	// get data from database
	var comSet models.CompanySetting
	db.QueryRow(context.Background(), "SELECT id,logo,favicon,email,instagram,imo FROM company_setting WHERE deleted_at IS NULL ORDER BY created_at ASC LIMIT 1").Scan(&comSet.ID, &comSet.Logo, &comSet.Favicon, &comSet.Email, &comSet.Instagram, &comSet.Imo)
	if comSet.ID == "" {
		helpers.HandleError(c, 404, "record not found")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":          true,
		"company_setting": comSet,
	})
}

func GetCompanySettingForHeader() (models.CompanySetting, error) {
	db, err := config.ConnDB()
	if err != nil {
		return models.CompanySetting{}, nil
	}
	defer db.Close()

	var logoFavicon models.CompanySetting

	// GET LOGO AND FAVICON
	db.QueryRow(context.Background(), "SELECT logo,favicon FROM company_setting WHERE deleted_at IS NULL ORDER BY created_at ASC LIMIT 1").Scan(&logoFavicon.Logo, &logoFavicon.Favicon)
	return logoFavicon, nil
}
