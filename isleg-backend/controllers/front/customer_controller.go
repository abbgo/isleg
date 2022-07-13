package controllers

import (
	"github/abbgo/isleg/isleg-backend/auth"
	"github/abbgo/isleg/isleg-backend/config"

	backController "github/abbgo/isleg/isleg-backend/controllers/back"
	"github/abbgo/isleg/isleg-backend/models"
	"net/http"

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
	phoneNumber := c.PostForm("phone_number")
	password := c.PostForm("password")
	gender := c.PostForm("gender")
	birthday := c.PostForm("birthday")
	addresses := c.PostFormArray("addresses")

	err = models.ValidateCustomerData(fullName, phoneNumber, password, gender, addresses)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
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

	_, err = config.ConnDB().Exec("INSERT INTO customers (full_name,phone_number,password,birthday,gender,addresses) VALUES ($1,$2,$3,$4,$5,$6)", fullName, "+993"+phoneNumber, hashPassword, birthday, gender, pq.StringArray(addresses))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	accessTokenString, err := auth.GenerateAccessToken(phoneNumber)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	refreshTokenString, err := auth.GenerateRefreshToken(phoneNumber)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":        true,
		"message":       "customer successfully added",
		"access_token":  accessTokenString,
		"refresh_token": refreshTokenString,
	})

}
