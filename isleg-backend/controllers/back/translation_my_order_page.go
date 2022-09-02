package controllers

import (
	"github/abbgo/isleg/isleg-backend/config"
	"github/abbgo/isleg/isleg-backend/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type TrMyOrderPage struct {
	Orders     string `json:"orders"`
	Date       string `json:"date"`
	Price      string `json:"price"`
	Currency   string `json:"currency"`
	Image      string `json:"image"`
	Name       string `json:"name"`
	Brend      string `json:"brend"`
	Code       string `json:"code"`
	Amount     string `json:"amount"`
	TotalPrice string `json:"total_price"`
}

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

func UpdateTranslationMyOrderPageByID(c *gin.Context) {

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

	rowTrMyOrderPage, err := db.Query("SELECT id FROM translation_my_order_page WHERE id = $1 AND deleted_at IS NULL", ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer rowTrMyOrderPage.Close()

	var id string

	for rowTrMyOrderPage.Next() {
		if err := rowTrMyOrderPage.Scan(&id); err != nil {
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

	dataNames := []string{"orders", "date", "price", "currency", "image", "name", "brend", "code", "amount", "total_price"}

	// VALIDATE DATA
	err = models.ValidateTranslationMyOrderPageUpdate(dataNames, c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	currentTime := time.Now()

	resultTrMyOrderPage, err := db.Query("UPDATE translation_my_order_page SET orders = $1, date = $2 , price = $3, currency = $4 , image = $5 , name = $6 , brend = $7, code = $8 , amount = $9, total_price = $10, updated_at = $11 WHERE id = $12", c.PostForm("orders"), c.PostForm("date"), c.PostForm("price"), c.PostForm("currency"), c.PostForm("image"), c.PostForm("name"), c.PostForm("brend"), c.PostForm("code"), c.PostForm("amount"), c.PostForm("total_price"), currentTime, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer resultTrMyOrderPage.Close()

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "translation my order page successfully updated",
	})

}

func GetTranslationMyOrderPageByID(c *gin.Context) {

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

	rowTrMyOrderPage, err := db.Query("SELECT orders,date,price,currency,image,name,brend,code,amount,total_price FROM translation_my_order_page WHERE id = $1 AND deleted_at IS NULL", ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer rowTrMyOrderPage.Close()

	var t TrMyOrderPage

	for rowTrMyOrderPage.Next() {
		if err := rowTrMyOrderPage.Scan(&t.Orders, &t.Date, &t.Price, &t.Currency, &t.Image, &t.Name, &t.Brend, &t.Code, &t.Amount, &t.TotalPrice); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}

	if t.Orders == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "record not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":                    true,
		"translation_my_order_page": t,
	})

}
