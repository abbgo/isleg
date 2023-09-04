package controllers

import (
	"context"
	"github/abbgo/isleg/isleg-backend/config"
	"github/abbgo/isleg/isleg-backend/helpers"
	"github/abbgo/isleg/isleg-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateTranslationSecure(c *gin.Context) {

	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	var trSecures []models.TranslationSecure

	// get data from request
	if err := c.BindJSON(&trSecures); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	// check lang_id
	for _, v := range trSecures {

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

	// create translation secure
	for _, v := range trSecures {
		_, err := db.Exec(context.Background(), "INSERT INTO translation_secure (lang_id,title,content) VALUES ($1,$2,$3)", v.LangID, v.Title, v.Content)
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

func UpdateTranslationSecureByID(c *gin.Context) {

	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	// get id of translation secure from request parameter
	var trSecure models.TranslationSecure

	if err := c.BindJSON(&trSecure); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	// check id of translation secure table
	rowFlag, err := db.Query(context.Background(), "SELECT id FROM translation_secure WHERE id = $1 AND deleted_at IS NULL", trSecure.ID)
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

	// update data of table
	_, err = db.Exec(context.Background(), "UPDATE translation_secure SET title = $1, content = $2, lang_id = $4 WHERE id = $3", trSecure.Title, trSecure.Content, trSecure.ID, trSecure.LangID)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "data successfully updated",
	})

}

func GetTranslationSecureByID(c *gin.Context) {

	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	// get id of translation secure table from request parameter
	ID := c.Param("id")

	// check id and get data
	rowFlag, err := db.Query(context.Background(), "SELECT id,title,content FROM translation_secure WHERE id = $1 AND deleted_at IS NULL", ID)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	var t models.TranslationSecure
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
		"status":             true,
		"translation_secure": t,
	})

}

func GetTranslationSecureByLangID(c *gin.Context) {

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

	// get translation secure where lang_id equal langID
	secureRow, err := db.Query(context.Background(), "SELECT title,content FROM translation_secure WHERE lang_id = $1 AND deleted_at IS NULL", langID)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	var translationSecure models.TranslationSecure
	for secureRow.Next() {
		if err := secureRow.Scan(&translationSecure.Title, &translationSecure.Content); err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}
	}

	if translationSecure.Title == "" {
		helpers.HandleError(c, 404, "record not found")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":             true,
		"translation_secure": translationSecure,
	})

}
