package controllers

import (
	"github/abbgo/isleg/isleg-backend/config"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateCompanyPhone(c *gin.Context) {

	db, err := config.ConnDB()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer db.Close()

	// GET DATA FROM REQUEST
	phone := c.PostForm("phone")

	// validate data
	if phone == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "phone is required",
		})
		return
	}

	_, err = strconv.Atoi(phone)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	if len(phone) != 8 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "the length of the phone number must be 8",
		})
		return
	}

	// create company phone
	resultComPhone, err := db.Query("INSERT INTO company_phone (phone) VALUES ($1)", phone)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer resultComPhone.Close()

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "company phone successfully added",
	})

}

func UpdateCompanyPhoneByID(c *gin.Context) {

	db, err := config.ConnDB()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer db.Close()

	ID := c.Param("id")

	rowCompanyPhone, err := db.Query("SELECT id FROM company_phone WHERE id = $1 AND deleted_at IS NULL", ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer rowCompanyPhone.Close()

	var comPhoneID string

	for rowCompanyPhone.Next() {
		if err := rowCompanyPhone.Scan(&comPhoneID); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}

	if comPhoneID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "record not found",
		})
		return
	}

	phone := c.PostForm("phone")

	// validate data
	if phone == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "phone is required",
		})
		return
	}

	_, err = strconv.Atoi(phone)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	if len(phone) != 8 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "the length of the phone number must be 8",
		})
		return
	}

	currentTime := time.Now()

	resultComPhone, err := db.Query("UPDATE company_phone SET phone = $1 , updated_at = $3 WHERE id = $2", phone, ID, currentTime)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer resultComPhone.Close()

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "company phone successfully updated",
	})

}

func GetCompanyPhoneByID(c *gin.Context) {

	db, err := config.ConnDB()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer db.Close()

	ID := c.Param("id")

	rowComPhone, err := db.Query("SELECT phone FROM company_phone WHERE id = $1 AND deleted_at IS NULL", ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer rowComPhone.Close()

	var phoneNumber string

	for rowComPhone.Next() {
		if err := rowComPhone.Scan(&phoneNumber); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}

	if phoneNumber == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "record not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":        true,
		"company_phone": phoneNumber,
	})

}

func GetCompanyPhones(c *gin.Context) {

	db, err := config.ConnDB()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer db.Close()

	var companyPhones []string

	// get all company phone number
	rows, err := db.Query("SELECT phone FROM company_phone WHERE deleted_at IS NULL")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var companyPhone string
		if err := rows.Scan(&companyPhone); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
		companyPhones = append(companyPhones, companyPhone)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":         true,
		"company_phones": companyPhones,
	})

}

func DeleteCompanyPhoneByID(c *gin.Context) {

	db, err := config.ConnDB()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer db.Close()

	ID := c.Param("id")

	rowComPhone, err := db.Query("SELECT id FROM company_phone WHERE id = $1 AND deleted_at IS NULL", ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer rowComPhone.Close()

	var comPhoneID string

	for rowComPhone.Next() {
		if err := rowComPhone.Scan(&comPhoneID); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}

	if comPhoneID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "record not found",
		})
		return
	}

	currentTime := time.Now()

	resultComPhone, err := db.Query("UPDATE company_phone SET deleted_at = $1 WHERE id = $2", currentTime, ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer resultComPhone.Close()

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "company phone successfully deleted",
	})

}

func RestoreCompanyPhoneByID(c *gin.Context) {

	db, err := config.ConnDB()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer db.Close()

	ID := c.Param("id")

	rowComPhone, err := db.Query("SELECT id FROM company_phone WHERE id = $1 AND deleted_at IS NOT NULL", ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer rowComPhone.Close()

	var comPhoneID string

	for rowComPhone.Next() {
		if err := rowComPhone.Scan(&comPhoneID); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}

	if comPhoneID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "record not found",
		})
		return
	}

	resultComPhone, err := db.Query("UPDATE company_phone SET deleted_at = NULL WHERE id = $1", ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer resultComPhone.Close()

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "company phone successfully restored",
	})

}

func DeletePermanentlyCompanyPhoneByID(c *gin.Context) {

	db, err := config.ConnDB()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer db.Close()

	ID := c.Param("id")

	rowCompanyPhone, err := db.Query("SELECT id FROM company_phone WHERE id = $1 AND deleted_at IS NOT NULL", ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer rowCompanyPhone.Close()

	var comPhoneID string

	for rowCompanyPhone.Next() {
		if err := rowCompanyPhone.Scan(&comPhoneID); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}

	if comPhoneID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "record not found",
		})
		return
	}

	resultComPhone, err := db.Query("DELETE FROM company_phone WHERE id = $1", ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer resultComPhone.Close()

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "company phone successfully deleted",
	})

}
