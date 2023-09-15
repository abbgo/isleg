package controllers

import (
	"context"
	"github/abbgo/isleg/isleg-backend/config"
	"github/abbgo/isleg/isleg-backend/helpers"
	"github/abbgo/isleg/isleg-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateTranslationMyOrderPage(c *gin.Context) {
	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	// get data from reuqest

	var trMyOrderPages []models.TranslationMyOrderPage
	if err := c.BindJSON(&trMyOrderPages); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	// check lang_id
	for _, v := range trMyOrderPages {
		var langID string
		db.QueryRow(context.Background(), "SELECT id FROM languages WHERE id = $1 AND deleted_at IS NULL", v.LangID).Scan(&langID)

		if langID == "" {
			helpers.HandleError(c, 404, "language not found")
			return
		}
	}

	// create translation_my_order_page
	for _, v := range trMyOrderPages {
		_, err := db.Exec(context.Background(), "INSERT INTO translation_my_order_page (lang_id,orders,date,price,image,name,brend,product_price,amount,total_price) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)", v.LangID, v.Orders, v.Date, v.Price, v.Image, v.Name, v.Brend, v.ProcuctPrice, v.Amount, v.TotalPrice)
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

func UpdateTranslationMyOrderPageByID(c *gin.Context) {
	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	// get id from request parameter
	var trMyOrderPage models.TranslationMyOrderPage
	if err := c.BindJSON(&trMyOrderPage); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	// check id
	var id string
	db.QueryRow(context.Background(), "SELECT id FROM translation_my_order_page WHERE id = $1 AND deleted_at IS NULL", trMyOrderPage.ID).Scan(&id)

	if id == "" {
		helpers.HandleError(c, 404, "record not found")
		return
	}

	// update data
	_, err = db.Exec(context.Background(), "UPDATE translation_my_order_page SET orders = $1, date = $2 , price = $3 , image = $4 , name = $5 , brend = $6, product_price = $7 , amount = $8, total_price = $9, lang_id = $11 WHERE id = $10", trMyOrderPage.Orders, trMyOrderPage.Date, trMyOrderPage.Price, trMyOrderPage.Image, trMyOrderPage.Name, trMyOrderPage.Brend, trMyOrderPage.ProcuctPrice, trMyOrderPage.Amount, trMyOrderPage.TotalPrice, trMyOrderPage.ID, trMyOrderPage.LangID)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "data successfully updated",
	})
}

func GetTranslationMyOrderPageByID(c *gin.Context) {
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
	var t models.TranslationMyOrderPage
	db.QueryRow(context.Background(), "SELECT id,orders,date,price,image,name,brend,product_price,amount,total_price FROM translation_my_order_page WHERE id = $1 AND deleted_at IS NULL", ID).Scan(&t.ID, &t.Orders, &t.Date, &t.Price, &t.Image, &t.Name, &t.Brend, &t.ProcuctPrice, &t.Amount, &t.TotalPrice)

	if t.ID == "" {
		helpers.HandleError(c, 404, "record not found")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":                    true,
		"translation_my_order_page": t,
	})
}

func GetTranslationMyOrderPageByLangID(c *gin.Context) {
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	// GET DATA FROM ROUTE PARAMETER
	langShortName := c.Param("lang")

	// GET language id
	langID, err := GetLangID(langShortName)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	// get translation-basket-page where lang_id equal langID
	var t models.TranslationMyOrderPage
	db.QueryRow(context.Background(), "SELECT orders,date,price,image,name,brend,product_price,amount,total_price FROM translation_my_order_page WHERE lang_id = $1 AND deleted_at IS NULL", langID).Scan(&t.Orders, &t.Date, &t.Price, &t.Image, &t.Name, &t.Brend, &t.ProcuctPrice, &t.Amount, &t.TotalPrice)

	if t.Orders == "" {
		helpers.HandleError(c, 404, "record not found")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":                    true,
		"translation_my_order_page": t,
	})
}
