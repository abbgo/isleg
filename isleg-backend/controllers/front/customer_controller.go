package controllers

import (
	"github/abbgo/isleg/isleg-backend/config"
	backController "github/abbgo/isleg/isleg-backend/controllers/back"
	"github/abbgo/isleg/isleg-backend/models"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

func RegisterCustomer(c *gin.Context) {

	// GET DATA FROM ROUTE PARAMETER
	langShortName := c.Param("lang")

	// GET language id
	_, err := backController.GetLangID(langShortName)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	fullName := c.PostForm("full_name")
	if fullName == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "full name is required",
		})
		return
	}

	phoneNumber := c.PostForm("phone_number")
	if phoneNumber == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "phone number is required",
		})
		return
	}

	phoneNumberValid := strings.HasPrefix(phoneNumber, "+9936")
	if !phoneNumberValid {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "phone number must start with +9936",
		})
		return
	}

	if len(phoneNumber) != 12 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "the length of the phone number must be 12",
		})
		return
	}

	password := c.PostForm("password")
	if password == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "password is required",
		})
		return
	}

	if len(password) < 5 || len(password) > 25 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "password length must be between 5 and 25",
		})
		return
	}

	hashPassword, err := models.HashPassword(password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	gender := c.PostForm("gender")
	if gender != "" {
		if gender != "1" && gender != "0" {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": "gender must be 0 or 1",
			})
			return
		}
	}

	birthday := c.PostForm("birthday")

	addresses := c.PostFormArray("addresses")
	if len(addresses) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "address is required",
		})
		return
	}
	for _, v := range addresses {
		if v == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": "address is required",
			})
			return
		}
	}

	_, err = config.ConnDB().Exec("INSERT INTO customers (full_name,phone_number,password,birthday,gender,addresses) VALUES ($1,$2,$3,$4,$5,$6)", fullName, phoneNumber, hashPassword, birthday, gender, pq.StringArray(addresses))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "customer successfully added",
	})

}
