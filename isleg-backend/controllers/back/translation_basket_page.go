package controllers

import (
	"github/abbgo/isleg/isleg-backend/config"
	"github/abbgo/isleg/isleg-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateTranslationBasketPage(c *gin.Context) {

	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer func() {
		if err := db.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

	// get data from request
	var trBasketPages []models.TranslationBasketPage

	if err := c.BindJSON(&trBasketPages); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	//check lsng_id
	for _, v := range trBasketPages {

		rowLang, err := db.Query("SELECT id FROM languages WHERE id = $1 AND deleted_at IS NULL", v.LangID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
		defer func() {
			if err := rowLang.Close(); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  false,
					"message": err.Error(),
				})
				return
			}
		}()

		var langID string

		for rowLang.Next() {
			if err := rowLang.Scan(&langID); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  false,
					"message": err.Error(),
				})
				return
			}
		}

		if langID == "" {
			c.JSON(http.StatusNotFound, gin.H{
				"status":  false,
				"message": "language not found",
			})
			return
		}

	}

	// create translation_basket_page
	for _, v := range trBasketPages {

		resultTrBasketPage, err := db.Query("INSERT INTO translation_basket_page (lang_id,quantity_of_goods,total_price,discount,delivery,total,currency,to_order,your_basket,empty_the_basket) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)", v.LangID, v.QuantityOfGoods, v.TotalPrice, v.Discount, v.Delivery, v.Total, v.Currency, v.ToOrder, v.YourBasket, v.EmptyTheBasket)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
		defer func() {
			if err := resultTrBasketPage.Close(); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  false,
					"message": err.Error(),
				})
				return
			}
		}()

	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "data successfully added",
	})

}

func UpdateTranslationBasketPageByID(c *gin.Context) {

	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer func() {
		if err := db.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

	// get id from request parameter
	var trBasketPage models.TranslationBasketPage

	if err := c.BindJSON(&trBasketPage); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	// check id
	rowTRBasketPage, err := db.Query("SELECT id FROM translation_basket_page WHERE id = $1 AND deleted_at IS NULL", trBasketPage.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer func() {
		if err := rowTRBasketPage.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

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
		c.JSON(http.StatusNotFound, gin.H{
			"status":  false,
			"message": "record not found",
		})
		return
	}

	resultTrBasketPage, err := db.Query("UPDATE translation_basket_page SET quantity_of_goods = $1, total_price = $2 , discount = $3, delivery = $4 , total = $5 , currency = $6 , to_order = $7, your_basket = $8 , empty_the_basket = $9, lang_id = $11  WHERE id = $10", trBasketPage.QuantityOfGoods, trBasketPage.TotalPrice, trBasketPage.Discount, trBasketPage.Delivery, trBasketPage.Total, trBasketPage.Currency, trBasketPage.ToOrder, trBasketPage.YourBasket, trBasketPage.EmptyTheBasket, trBasketPage.ID, trBasketPage.LangID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer func() {
		if err := resultTrBasketPage.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "data successfully updated",
	})

}

func GetTranslationBasketPageByID(c *gin.Context) {

	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer func() {
		if err := db.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

	// get id from request parameter
	ID := c.Param("id")

	// check id and get data from database
	rowTRBasketPage, err := db.Query("SELECT id,quantity_of_goods,total_price,discount,delivery,total,currency,to_order,your_basket,empty_the_basket FROM translation_basket_page WHERE id = $1 AND deleted_at IS NULL", ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer func() {
		if err := rowTRBasketPage.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

	var t models.TranslationBasketPage

	for rowTRBasketPage.Next() {
		if err := rowTRBasketPage.Scan(&t.ID, &t.QuantityOfGoods, &t.TotalPrice, &t.Discount, &t.Delivery, &t.Total, &t.Currency, &t.ToOrder, &t.YourBasket, &t.EmptyTheBasket); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}

	if t.ID == "" {
		c.JSON(http.StatusNotFound, gin.H{
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

func GetTranslationBasketPageByLangID(c *gin.Context) {

	db, err := config.ConnDB()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer func() {
		if err := db.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

	// GET DATA FROM ROUTE PARAMETER
	langShortName := c.Param("lang")

	// GET language id
	langID, err := GetLangID(langShortName)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	// get translation-basket-page where lang_id equal langID
	rowTRBasketPage, err := db.Query("SELECT quantity_of_goods,total_price,discount,delivery,total,currency,to_order,your_basket,empty_the_basket FROM translation_basket_page WHERE lang_id = $1 AND deleted_at IS NULL", langID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer func() {
		if err := rowTRBasketPage.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

	var t models.TranslationBasketPage

	for rowTRBasketPage.Next() {
		if err := rowTRBasketPage.Scan(&t.QuantityOfGoods, &t.TotalPrice, &t.Discount, &t.Delivery, &t.Total, &t.Currency, &t.ToOrder, &t.YourBasket, &t.EmptyTheBasket); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}

	if t.QuantityOfGoods == "" {
		c.JSON(http.StatusNotFound, gin.H{
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
