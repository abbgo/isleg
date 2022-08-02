package controllers

import (
	"fmt"
	"github/abbgo/isleg/isleg-backend/config"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateCompanyPhone(c *gin.Context) {

	// GET DATA FROM REQUEST
	phone := c.PostForm("phone")

	fmt.Println(phone)

	// validate data
	if phone == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "phone is required",
		})
		return
	}

	_, err := strconv.Atoi(phone)
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
	_, err = config.ConnDB().Exec("INSERT INTO company_phone (phone) VALUES ($1)", phone)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "company phone successfully added",
	})

}

func GetCompanyPhones(c *gin.Context) {

	var companyPhones []string

	// get all company phone number
	rows, err := config.ConnDB().Query("SELECT phone FROM company_phone WHERE deleted_at IS NULL")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

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
