package controllers

import (
	"github/abbgo/isleg/isleg-backend/config"
	"github/abbgo/isleg/isleg-backend/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type TrBasketPage struct {
	QuantityOfGoods string `json:"quantity_of_goods"`
	TotalPrice      string `json:"total_price"`
	Discount        string `json:"discount"`
	Delivery        string `json:"delivery"`
	Total           string `json:"total"`
	Currency        string `json:"currency"`
	ToOrder         string `json:"to_order"`
	YourBasket      string `json:"your_basket"`
}

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

func UpdateTranslationBasketPageByID(c *gin.Context) {

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

	rowTRBasketPage, err := db.Query("SELECT id FROM translation_basket_page WHERE id = $1 AND deleted_at IS NULL", ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer rowTRBasketPage.Close()

	var id string

	for rowTRBasketPage.Next() {
		if err := rowTRBasketPage.Scan(&id); err != nil {
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

	dataNames := []string{"quantity_of_goods", "total_price", "discount", "delivery", "total", "currency", "to_order", "your_basket"}

	// VALIDATE DATA
	err = models.ValidateTranslationBasketPageUpdate(dataNames, c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	currentTime := time.Now()

	resultTrBasketPage, err := db.Query("UPDATE translation_basket_page SET quantity_of_goods = $1, total_price = $2 , discount = $3, delivery = $4 , total = $5 , currency = $6 , to_order = $7, your_basket = $8 , updated_at = $9 WHERE id = $10", c.PostForm("quantity_of_goods"), c.PostForm("total_price"), c.PostForm("discount"), c.PostForm("delivery"), c.PostForm("total"), c.PostForm("currency"), c.PostForm("to_order"), c.PostForm("your_basket"), currentTime, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer resultTrBasketPage.Close()

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "translation basket page successfully updated",
	})

}

func GetTranslationBasketPageByID(c *gin.Context) {

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

	rowTRBasketPage, err := db.Query("SELECT quantity_of_goods,total_price,discount,delivery,total,currency,to_order,your_basket FROM translation_basket_page WHERE id = $1 AND deleted_at IS NULL", ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer rowTRBasketPage.Close()

	var t TrBasketPage

	for rowTRBasketPage.Next() {
		if err := rowTRBasketPage.Scan(&t.QuantityOfGoods, &t.TotalPrice, &t.Discount, &t.Delivery, &t.Total, &t.Currency, &t.ToOrder, &t.YourBasket); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}

	if t.QuantityOfGoods == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "record not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":                  true,
		"translation_basket_page": t,
	})

}
