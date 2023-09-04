package controllers

import (
	"context"
	"github/abbgo/isleg/isleg-backend/config"
	"github/abbgo/isleg/isleg-backend/helpers"
	"github/abbgo/isleg/isleg-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateTranslationPayment(c *gin.Context) {

	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	var trPaymentPages []models.TranslationPayment

	// get data from request
	if err := c.BindJSON(&trPaymentPages); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	// check lang_id
	for _, v := range trPaymentPages {

		rowLang, err := db.Query(context.Background(), "SELECT id FROM languages WHERE id = $1 AND deleted_at IS NULL", v.LangID)
		if err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}

		var langID string
		for rowLang.Next() {
			if err := rowLang.Scan(&langID); err != nil {
				helpers.HandleError(c, 400, err.Error())
				return
			}
		}

		if langID == "" {
			helpers.HandleError(c, 404, "language not found")
			return
		}

	}

	// create translation payment
	for _, v := range trPaymentPages {
		_, err := db.Exec(context.Background(), "INSERT INTO translation_payment (lang_id,title,content) VALUES ($1,$2,$3)", v.LangID, v.Title, v.Content)
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

func UpdateTranslationPaymentByID(c *gin.Context) {

	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	// get id of translation payment from request parameter
	var trPaymentPage models.TranslationPayment

	if err := c.BindJSON(&trPaymentPage); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	// check id
	rowFlag, err := db.Query(context.Background(), "SELECT id FROM translation_payment WHERE id = $1 AND deleted_at IS NULL", trPaymentPage.ID)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	var id string
	for rowFlag.Next() {
		if err := rowFlag.Scan(&id); err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}
	}

	if id == "" {
		helpers.HandleError(c, 404, "record not found")
		return
	}

	// update data
	_, err = db.Exec(context.Background(), "UPDATE translation_payment SET title = $1, content = $2, lang_id = $4 WHERE id = $3", trPaymentPage.Title, trPaymentPage.Content, trPaymentPage.ID, trPaymentPage.LangID)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "data successfully updated",
	})

}

func GetTranslationPaymentByID(c *gin.Context) {

	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	// get id of translation payment from request parameter
	ID := c.Param("id")

	// check id and get data from database
	rowFlag, err := db.Query(context.Background(), "SELECT id,title,content FROM translation_payment WHERE id = $1 AND deleted_at IS NULL", ID)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	var t models.TranslationPayment
	for rowFlag.Next() {
		if err := rowFlag.Scan(&t.ID, &t.Title, &t.Content); err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}
	}

	if t.ID == "" {
		helpers.HandleError(c, 404, "record not found")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":              true,
		"translation_payment": t,
	})

}

func GetTranslationPaymentByLangID(c *gin.Context) {

	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	langID, err := CheckLanguage(c)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	// get translation payment where lang_id equal langID
	paymentRow, err := db.Query(context.Background(), "SELECT title,content FROM translation_payment WHERE lang_id = $1 AND deleted_at IS NULL", langID)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	var translationPayment models.TranslationPayment
	for paymentRow.Next() {
		if err := paymentRow.Scan(&translationPayment.Title, &translationPayment.Content); err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}
	}

	if translationPayment.Title == "" {
		helpers.HandleError(c, 404, "record not found")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":              true,
		"translation_payment": translationPayment,
	})

}
