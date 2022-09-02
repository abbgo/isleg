package controllers

import (
	"github/abbgo/isleg/isleg-backend/config"
	"github/abbgo/isleg/isleg-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateTranslationMyOrderPage(c *gin.Context) {

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

	dataNames := []string{"orders", "date", "price", "currency", "image", "name", "brend", "code", "amount", "total_price"}

	// VALIDATE DATA
	if err = models.ValidateTranslationMyOrderPageData(languages, dataNames, c); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	// create translation_my_information_page
	for _, v := range languages {
		resultTrMyOrderPage, err := db.Query("INSERT INTO translation_my_order_page (lang_id,orders,date,price,currency,image,name,brend,code,amount,total_price) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)", v.ID, c.PostForm("orders_"+v.NameShort), c.PostForm("date_"+v.NameShort), c.PostForm("price_"+v.NameShort), c.PostForm("currency_"+v.NameShort), c.PostForm("image_"+v.NameShort), c.PostForm("name_"+v.NameShort), c.PostForm("brend_"+v.NameShort), c.PostForm("code_"+v.NameShort), c.PostForm("amount_"+v.NameShort), c.PostForm("total_price_"+v.NameShort))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
		defer resultTrMyOrderPage.Close()
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "translation my order page successfully added",
	})

}
