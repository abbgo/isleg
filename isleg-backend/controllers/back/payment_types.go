package controllers

import (
	"github/abbgo/isleg/isleg-backend/config"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func CreatePaymentType(c *gin.Context) {

	db, err := config.ConnDB()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer db.Close()

	// GET ALL LANGUAGE
	languages, err := GetAllLanguageWithIDAndNameShort()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	// VALIDATE DATA
	for _, v := range languages {
		if c.PostForm("type_"+v.NameShort) == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": "type_" + v.NameShort + " is required",
			})
			return
		}
	}

	// create company address
	for _, v := range languages {
		resultComAddres, err := db.Query("INSERT INTO payment_types (lang_id,type) VALUES ($1,$2)", v.ID, c.PostForm("type_"+v.NameShort))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
		defer resultComAddres.Close()
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "payment type successfully added",
	})

}

func UpdatePaymentTypeByID(c *gin.Context) {

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

	rowPaymentType, err := db.Query("SELECT id FROM payment_types WHERE id = $1 AND deleted_at IS NULL", ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer rowPaymentType.Close()

	var id string

	for rowPaymentType.Next() {
		if err := rowPaymentType.Scan(&id); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}

	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "record not found",
		})
		return
	}

	paymentType := c.PostForm("type")
	if paymentType == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "paymentType required",
		})
		return
	}

	currentTime := time.Now()

	resultPaymentType, err := db.Query("UPDATE payment_types SET type = $1, updated_at = $2 WHERE id = $3", paymentType, currentTime, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer resultPaymentType.Close()

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "payment types successfully updated",
	})

}
