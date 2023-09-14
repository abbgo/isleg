package controllers

import (
	"context"
	"github/abbgo/isleg/isleg-backend/config"
	"github/abbgo/isleg/isleg-backend/helpers"
	"github/abbgo/isleg/isleg-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateCompanyPhone(c *gin.Context) {
	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	// GET DATA FROM REQUEST
	var companyPhone models.CompanyPhone
	if err := c.BindJSON(&companyPhone); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	// create company phone
	_, err = db.Exec(context.Background(), "INSERT INTO company_phone (phone) VALUES ($1)", companyPhone.Phone)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "data successfully added",
	})
}

func UpdateCompanyPhoneByID(c *gin.Context) {
	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	// get data from request
	var companyPhone models.CompanyPhone
	if err := c.BindJSON(&companyPhone); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	// check id
	var comPhoneID string
	if err := db.QueryRow(context.Background(), "SELECT id FROM company_phone WHERE id = $1 AND deleted_at IS NULL", companyPhone.ID).Scan(&comPhoneID); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	if comPhoneID == "" {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  false,
			"message": "record not found",
		})
		return
	}

	_, err = db.Exec(context.Background(), "UPDATE company_phone SET phone = $1 WHERE id = $2", companyPhone.Phone, companyPhone.ID)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "data successfully updated",
	})
}

func GetCompanyPhoneByID(c *gin.Context) {
	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	// get id from request parameter
	ID := c.Param("id")

	// check id and get data from database
	var phoneNumber string
	if err := db.QueryRow(context.Background(), "SELECT phone FROM company_phone WHERE id = $1 AND deleted_at IS NULL", ID).Scan(&phoneNumber); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	if phoneNumber == "" {
		helpers.HandleError(c, 404, "record not found")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":        true,
		"company_phone": phoneNumber,
	})
}

// GetCompanyPhones funksiya firmanyn ahli telefon belgilerini alyar
func GetCompanyPhones(c *gin.Context) {
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	var companyPhones []string

	// get all company phone number
	rows, err := db.Query(context.Background(), "SELECT phone FROM company_phone WHERE deleted_at IS NULL")
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	for rows.Next() {
		var companyPhone string
		if err := rows.Scan(&companyPhone); err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}
		companyPhones = append(companyPhones, companyPhone)
	}

	rowCompanySetting, err := db.Query(context.Background(), "SELECT email,instagram,imo FROM company_setting ORDER BY created_at LIMIT 1")
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	var companySetting models.CompanySetting
	for rowCompanySetting.Next() {
		if err := rowCompanySetting.Scan(&companySetting.Email, &companySetting.Instagram, &companySetting.Imo); err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"status":          true,
		"company_phones":  companyPhones,
		"company_setting": companySetting,
	})
}

func DeleteCompanyPhoneByID(c *gin.Context) {
	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	// get id from request parameter
	ID := c.Param("id")

	// check id
	var comPhoneID string
	if err := db.QueryRow(context.Background(), "SELECT id FROM company_phone WHERE id = $1 AND deleted_at IS NULL", ID).Scan(&comPhoneID); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	if comPhoneID == "" {
		helpers.HandleError(c, 404, "record not found")
		return
	}

	_, err = db.Exec(context.Background(), "UPDATE company_phone SET deleted_at = now() WHERE id = $2", ID)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "data successfully deleted",
	})
}

func RestoreCompanyPhoneByID(c *gin.Context) {
	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	// get id from request parameter
	ID := c.Param("id")

	// check id
	var comPhoneID string
	if err := db.QueryRow(context.Background(), "SELECT id FROM company_phone WHERE id = $1 AND deleted_at IS NOT NULL", ID).Scan(&comPhoneID); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	if comPhoneID == "" {
		helpers.HandleError(c, 404, "record not found")
		return
	}

	_, err = db.Exec(context.Background(), "UPDATE company_phone SET deleted_at = NULL WHERE id = $1", ID)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "data successfully restored",
	})
}

func DeletePermanentlyCompanyPhoneByID(c *gin.Context) {
	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	// get id from request parameter
	ID := c.Param("id")

	//check id
	var comPhoneID string
	if err := db.QueryRow(context.Background(), "SELECT id FROM company_phone WHERE id = $1 AND deleted_at IS NOT NULL", ID).Scan(&comPhoneID); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	if comPhoneID == "" {
		helpers.HandleError(c, 404, "record not found")
		return
	}

	// delete category phone
	_, err = db.Exec(context.Background(), "DELETE FROM company_phone WHERE id = $1", ID)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "data successfully deleted",
	})
}
