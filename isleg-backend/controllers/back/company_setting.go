package controllers

import (
	"github/abbgo/isleg/isleg-backend/config"
	"github/abbgo/isleg/isleg-backend/models"
	"github/abbgo/isleg/isleg-backend/pkg"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateCompanySetting(c *gin.Context) {

	// initialize database connection
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

	// get email and instagram of company_setting from request
	email := c.PostForm("email")
	instagram := c.PostForm("instagram")
	imo := c.PostForm("imo")

	// validate email and instagram
	err = models.ValidateCompanySettingData(email, instagram, imo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	// upload logo
	newFileNameLogo, err := pkg.FileUpload("logo", "setting", c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	// upload favicon
	newFileNameFavicon, err := pkg.FileUpload("favicon", "setting", c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	// add data to database
	resultComSetting, err := db.Query("INSERT INTO company_setting (logo,favicon,email,instagram,imo) VALUES ($1,$2,$3,$4,$5)", "uploads/setting/"+newFileNameLogo, "uploads/setting/"+newFileNameFavicon, email, instagram, imo)
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

	// initialize database connection
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

	// get data from request
	email := c.PostForm("email")
	instagram := c.PostForm("instagram")
	imo := c.PostForm("imo")

	var logoName, faviconName string

	// Check if there is a company_setting and get logo and favicon
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

	// validate email and instagram
	err = models.ValidateCompanySettingData(email, instagram, imo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	// upload logo
	logoName, err = pkg.FileUploadForUpdate("logo", "setting", logo, c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	// upload favicon
	faviconName, err = pkg.FileUploadForUpdate("favicon", "setting", favicon, c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	// update data in database
	resultComPSETTING, err := db.Query("UPDATE company_setting SET logo = $1, favicon=$2, email=$3, instagram=$4, imo=$5", logoName, faviconName, email, instagram, imo)
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

	// initialize database connection
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

	// get data from database
	rowComSet, err := db.Query("SELECT id,logo,favicon,email,instagram,imo FROM company_setting WHERE deleted_at IS NULL ORDER BY created_at ASC LIMIT 1")
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

	var comSet models.CompanySetting

	for rowComSet.Next() {
		if err := rowComSet.Scan(&comSet.ID, &comSet.Logo, &comSet.Favicon, &comSet.Email, &comSet.Instagram, &comSet.Imo); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}

	if comSet.ID == "" {
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

func GetCompanySettingForHeader() (models.CompanySetting, error) {

	db, err := config.ConnDB()
	if err != nil {
		return models.CompanySetting{}, nil
	}
	defer func() (models.CompanySetting, error) {
		if err := db.Close(); err != nil {
			return models.CompanySetting{}, err
		}
		return models.CompanySetting{}, nil
	}()

	var logoFavicon models.CompanySetting

	// GET LOGO AND FAVICON
	row, err := db.Query("SELECT logo,favicon FROM company_setting WHERE deleted_at IS NULL ORDER BY created_at ASC LIMIT 1")
	if err != nil {
		return models.CompanySetting{}, err
	}
	defer row.Close()

	for row.Next() {
		if err := row.Scan(&logoFavicon.Logo, &logoFavicon.Favicon); err != nil {
			return models.CompanySetting{}, err
		}
	}

	return logoFavicon, nil

}
