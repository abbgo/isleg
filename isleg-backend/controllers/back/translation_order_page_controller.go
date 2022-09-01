package controllers

import (
	"github/abbgo/isleg/isleg-backend/config"
	"github/abbgo/isleg/isleg-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateTranslationOrderPage(c *gin.Context) {

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

	dataNames := []string{"content", "type_of_payment", "choose_a_delivery_time", "your_address", "mark", "to_order", "tomorrow", "cash", "payment_terminal"}

	// VALIDATE DATA
	if err = models.ValidateTranslationOrderPageData(languages, dataNames, c); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	// create translation_my_information_page
	for _, v := range languages {
		resultTrOrderPage, err := db.Query("INSERT INTO translation_order_page (lang_id,content,type_of_payment,choose_a_delivery_time,your_address,mark,to_order,tomorrow,cash,payment_terminal) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)", v.ID, c.PostForm("content_"+v.NameShort), c.PostForm("type_of_payment_"+v.NameShort), c.PostForm("choose_a_delivery_time_"+v.NameShort), c.PostForm("your_address_"+v.NameShort), c.PostForm("mark_"+v.NameShort), c.PostForm("to_order_"+v.NameShort), c.PostForm("tomorrow_"+v.NameShort), c.PostForm("cash_"+v.NameShort), c.PostForm("payment_terminal_"+v.NameShort))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
		defer resultTrOrderPage.Close()
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "translation order page successfully added",
	})

}
