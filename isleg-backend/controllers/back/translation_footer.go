package controllers

import (
	"context"
	"github/abbgo/isleg/isleg-backend/config"
	"github/abbgo/isleg/isleg-backend/helpers"
	"github/abbgo/isleg/isleg-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateTranslationFooter(c *gin.Context) {
	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	var trFooters []models.TranslationFooter
	if err := c.BindJSON(&trFooters); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	// check lang_id
	for _, v := range trFooters {
		var langID string
		db.QueryRow(context.Background(), "SELECT id FROM languages WHERE id = $1 AND deleted_atr IS NULL", v.LangID).Scan(&langID)
		if langID == "" {
			helpers.HandleError(c, 404, "lamguage not found")
			return
		}
	}

	// create translation footer
	for _, v := range trFooters {
		_, err := db.Exec(context.Background(), "INSERT INTO translation_footer (lang_id,about,payment,contact,secure,word) VALUES ($1,$2,$3,$4,$5,$6)", v.LangID, v.About, v.Payment, v.Contact, v.Contact, v.Word)
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

func UpdateTranslationFooterByID(c *gin.Context) {

	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	// get id of translation footer from request parameter
	var trFooter models.TranslationFooter
	if err := c.BindJSON(&trFooter); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	// check id
	var id string
	db.QueryRow(context.Background(), "SELECT id FROM translation_footer WHERE id = $1 AND deleted_at IS NULL", trFooter.ID).Scan(&id)

	if id == "" {
		helpers.HandleError(c, 404, "record not found")
		return
	}

	// update data of translation footer
	_, err = db.Exec(context.Background(), "UPDATE translation_footer SET about = $1, payment = $2, contact = $3, secure = $4, word = $5, lang_id = $7 WHERE id = $6", trFooter.About, trFooter.Payment, trFooter.Contact, trFooter.Secure, trFooter.Word, trFooter.ID, trFooter.LangID)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "data successfully updated",
	})
}

func GetTranslationFooterByID(c *gin.Context) {
	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	// get id of translation footer from request parameter
	trFootID := c.Param("id")

	//check id and get data from table
	var t models.TranslationFooter
	db.QueryRow(context.Background(), "SELECT about,payment,contact,secure,word FROM translation_footer WHERE id = $1 AND deleted_at IS NULL", trFootID).Scan(&t.About, &t.Payment, &t.Contact, &t.Secure, &t.Word)

	if t.About == "" {
		helpers.HandleError(c, 404, "record not found")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":             true,
		"translation_footer": t,
	})
}

func GetTranslationFooter(langID string) (models.TranslationFooter, error) {
	db, err := config.ConnDB()
	if err != nil {
		return models.TranslationFooter{}, err
	}
	defer db.Close()

	// get translation footer where lang_id equal langID
	var t models.TranslationFooter
	db.QueryRow(context.Background(), "SELECT about,payment,contact,secure,word FROM translation_footer WHERE lang_id = $1 AND deleted_at IS NULL", langID).Scan(&t.About, &t.Payment, &t.Contact, &t.Secure, &t.Word)

	return t, nil
}
