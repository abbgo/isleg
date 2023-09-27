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

func CreateTranslationOrderPage(c *gin.Context) {
	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	// get data from request
	var trOrderPages []models.TranslationOrderPage
	if err := c.BindJSON(&trOrderPages); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	// check lang_id
	for _, v := range trOrderPages {
		var langID string
		db.QueryRow(context.Background(), "SELECT id FROM languages WHERE id = $1 AND deleted_at IS NULL", v.LangID).Scan(&langID)

		if langID == "" {
			helpers.HandleError(c, 404, "language not found")
			return
		}
	}

	// create translation_my_information_page
	for _, v := range trOrderPages {
		_, err := db.Exec(context.Background(), "INSERT INTO translation_order_page (lang_id,content,type_of_payment,choose_a_delivery_time,your_address,mark,to_order) VALUES ($1,$2,$3,$4,$5,$6,$7)", v.LangID, v.Content, v.TypeOfPayment, v.ChooseADeliveryTime, v.YourAddress, v.Mark, v.ToOrder)
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

func UpdateTranslationOrderPageByID(c *gin.Context) {
	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	// get id from request parameter
	var trOrderPage models.TranslationOrderPage
	if err := c.BindJSON(&trOrderPage); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	// check id
	var id string
	db.QueryRow(context.Background(), "SELECT id FROM translation_order_page WHERE id = $1 AND deleted_at IS NULL", trOrderPage.ID).Scan(&id)

	if id == "" {
		helpers.HandleError(c, 404, "record not found")
		return
	}

	_, err = db.Exec(context.Background(), "UPDATE translation_order_page SET content = $1, type_of_payment = $2 , choose_a_delivery_time = $3, your_address = $4 , mark = $5 , to_order = $6, lang_id = $8 WHERE id = $7", trOrderPage.Content, trOrderPage.TypeOfPayment, trOrderPage.ChooseADeliveryTime, trOrderPage.YourAddress, trOrderPage.Mark, trOrderPage.ToOrder, trOrderPage.ID, trOrderPage.LangID)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "data successfully updated",
	})
}

func GetTranslationOrderPageByID(c *gin.Context) {
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
	var t models.TranslationOrderPage
	db.QueryRow(context.Background(), "SELECT id,content,type_of_payment,choose_a_delivery_time,your_address,mark,to_order FROM translation_order_page WHERE id = $1 AND deleted_at IS NULL", ID).Scan(&t.ID, &t.Content, &t.TypeOfPayment, &t.ChooseADeliveryTime, &t.YourAddress, &t.Mark, &t.ToOrder)

	if t.ID == "" {
		helpers.HandleError(c, 404, "record not found")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":                 true,
		"translation_order_page": t,
	})
}

func GetTranslationOrderPageByLangID(c *gin.Context) {
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
	var t models.TranslationOrderPage
	db.QueryRow(context.Background(), "SELECT content,type_of_payment,choose_a_delivery_time,your_address,mark,to_order FROM translation_order_page WHERE lang_id = $1 AND deleted_at IS NULL", langID).Scan(&t.Content, &t.TypeOfPayment, &t.ChooseADeliveryTime, &t.YourAddress, &t.Mark, &t.ToOrder)

	if t.Content == "" {
		helpers.HandleError(c, 404, "record not found")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":                 true,
		"translation_order_page": t,
	})
}
