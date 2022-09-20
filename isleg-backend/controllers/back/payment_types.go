package controllers

import (
	"github/abbgo/isleg/isleg-backend/config"
	"net/http"

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
