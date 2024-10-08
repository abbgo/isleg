package controllers

import (
	"context"
	"github/abbgo/isleg/isleg-backend/config"
	"github/abbgo/isleg/isleg-backend/helpers"
	"github/abbgo/isleg/isleg-backend/models"
	"github/abbgo/isleg/isleg-backend/pkg"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateTranslationBasketPage(c *gin.Context) {
	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	// get data from request
	var trBasketPages []models.TranslationBasketPage
	if err := c.BindJSON(&trBasketPages); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	//check lsng_id
	for _, v := range trBasketPages {
		var langID string
		db.QueryRow(context.Background(), "SELECT id FROM languages WHERE id = $1 AND deleted_at IS NULL", v.LangID).Scan(&langID)

		if langID == "" {
			helpers.HandleError(c, 404, "language not found")
			return
		}
	}

	// create translation_basket_page
	for _, v := range trBasketPages {
		_, err := db.Exec(context.Background(), "INSERT INTO translation_basket_page (lang_id,quantity_of_goods,total_price,discount,delivery,total,to_order,your_basket,empty_the_basket,empty_the_like_page) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)", v.LangID, v.QuantityOfGoods, v.TotalPrice, v.Discount, v.Delivery, v.Total, v.ToOrder, v.YourBasket, v.EmptyTheBasket, v.EmptyTheLikePage)
		if err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}
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
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	// get id from request parameter
	var trBasketPage models.TranslationBasketPage
	if err := c.BindJSON(&trBasketPage); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	// check id
	var id string
	db.QueryRow(context.Background(), "SELECT id FROM translation_basket_page WHERE id = $1 AND deleted_at IS NULL", trBasketPage.ID).Scan(&id)

	if id == "" {
		helpers.HandleError(c, 404, "record not found")
		return
	}

	_, err = db.Exec(context.Background(), "UPDATE translation_basket_page SET quantity_of_goods = $1, total_price = $2 , discount = $3, delivery = $4 , total = $5, to_order = $6, your_basket = $7 , empty_the_basket = $8, lang_id = $10 , empty_the_like_page = $11  WHERE id = $9", trBasketPage.QuantityOfGoods, trBasketPage.TotalPrice, trBasketPage.Discount, trBasketPage.Delivery, trBasketPage.Total, trBasketPage.ToOrder, trBasketPage.YourBasket, trBasketPage.EmptyTheBasket, trBasketPage.ID, trBasketPage.LangID, trBasketPage.EmptyTheLikePage)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "data successfully updated",
	})
}

func GetTranslationBasketPageByID(c *gin.Context) {
	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	// get id from request parameter
	ID := c.Param("id")

	// check id and get data from database
	var t models.TranslationBasketPage
	db.QueryRow(context.Background(), "SELECT id,quantity_of_goods,total_price,discount,delivery,total,to_order,your_basket,empty_the_basket,empty_the_like_page FROM translation_basket_page WHERE id = $1 AND deleted_at IS NULL", ID).Scan(&t.ID, &t.QuantityOfGoods, &t.TotalPrice, &t.Discount, &t.Delivery, &t.Total, &t.ToOrder, &t.YourBasket, &t.EmptyTheBasket, &t.EmptyTheLikePage)

	if t.ID == "" {
		helpers.HandleError(c, 404, "record not found")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":                  true,
		"translation_basket_page": t,
	})
}

func GetTranslationBasketPageByLangID(c *gin.Context) {
	langID, err := pkg.ValidateMiddlewareData(c, "lang_id")
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	// get translation-basket-page where lang_id equal langID
	var t models.TranslationBasketPage
	db.QueryRow(context.Background(), "SELECT quantity_of_goods,total_price,discount,delivery,total,to_order,your_basket,empty_the_basket,empty_the_like_page FROM translation_basket_page WHERE lang_id = $1 AND deleted_at IS NULL", langID).Scan(&t.QuantityOfGoods, &t.TotalPrice, &t.Discount, &t.Delivery, &t.Total, &t.ToOrder, &t.YourBasket, &t.EmptyTheBasket, &t.EmptyTheLikePage)

	if t.QuantityOfGoods == "" {
		helpers.HandleError(c, 404, "record not found")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":                  true,
		"translation_basket_page": t,
	})
}
