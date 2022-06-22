package controllers

import (
	"github/abbgo/isleg/isleg-backend/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateCompanyPhone(c *gin.Context) {

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

	// create company phone
	_, err := config.ConnDB().Exec("INSERT INTO company_phone (phone) VALUES ($1)", phone)
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
