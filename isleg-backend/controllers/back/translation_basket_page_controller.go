package controllers

import (
	"github/abbgo/isleg/isleg-backend/config"
	"github/abbgo/isleg/isleg-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateTranslationBasketPage(c *gin.Context) {

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

	dataNames := []string{"quantity_of_goods", "total_price", "discount", "delivery", "total", "currency", "to_order", "your_basket"}

	// VALIDATE DATA
	if err = models.ValidateTranslationBasketPageData(languages, dataNames, c); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	// create translation_my_information_page
	for _, v := range languages {
		resultTrBasketPage, err := db.Query("INSERT INTO translation_basket_page (lang_id,quantity_of_goods,total_price,discount,delivery,total,currency,to_order,your_basket) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9)", v.ID, c.PostForm("quantity_of_goods_"+v.NameShort), c.PostForm("total_price_"+v.NameShort), c.PostForm("discount_"+v.NameShort), c.PostForm("delivery_"+v.NameShort), c.PostForm("total_"+v.NameShort), c.PostForm("currency_"+v.NameShort), c.PostForm("to_order_"+v.NameShort), c.PostForm("your_basket_"+v.NameShort))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
		defer resultTrBasketPage.Close()
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "translation basket page successfully added",
	})

}
