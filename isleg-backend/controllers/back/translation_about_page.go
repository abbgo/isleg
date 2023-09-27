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

func CreateTranslationAbout(c *gin.Context) {
	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	// get data from request
	var trAbouts []models.TranslationAbout
	if err := c.BindJSON(&trAbouts); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	// check lang_id
	for _, v := range trAbouts {
		var langID string
		db.QueryRow(context.Background(), "SELECT id FROM languages WHERE id = $1 AND deleted_at IS NULL", v.LangID).Scan(&langID)

		if langID == "" {
			helpers.HandleError(c, 404, "language not found")
			return
		}
	}

	// create translation about
	for _, v := range trAbouts {
		_, err := db.Exec(context.Background(), "INSERT INTO translation_about (lang_id,title,content) VALUES ($1,$2,$3)", v.LangID, v.Title, v.Content)
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

func UpdateTranslationAboutByID(c *gin.Context) {
	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	// get id of translation about from request parameter
	var trAbout models.TranslationAbout
	//get data from request
	if err := c.BindJSON(&trAbout); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	// check id
	var id string
	db.QueryRow(context.Background(), "SELECT id FROM translation_about WHERE id = $1 AND deleted_at IS NULL", trAbout.ID).Scan(&id)

	if id == "" {
		helpers.HandleError(c, 404, "record not found")
		return
	}

	// update data
	_, err = db.Exec(context.Background(), "UPDATE translation_about SET title = $1, content = $2, lang_id = $4 WHERE id = $3", trAbout.Title, trAbout.Content, trAbout.ID, trAbout.LangID)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "data successfully updated",
	})
}

func GetTranslationAboutByID(c *gin.Context) {
	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	// get id of translation about from request parameter
	ID := c.Param("id")

	// check id and get data from database
	var t models.TranslationAbout
	db.QueryRow(context.Background(), "SELECT id,title,content FROM translation_about WHERE id = $1 AND deleted_at IS NULL", ID).Scan(&t.ID, &t.Title, &t.Content)

	if t.ID == "" {
		helpers.HandleError(c, 404, "record not found")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":            true,
		"translation_about": t,
	})
}

func GetTranslationAboutByLangID(c *gin.Context) {
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

	// get translation about where lang_id equal langID
	var translationAbout models.TranslationAbout
	db.QueryRow(context.Background(), "SELECT title,content FROM translation_about WHERE lang_id = $1 AND deleted_at IS NULL", langID).Scan(&translationAbout.Title, &translationAbout.Content)

	if translationAbout.Title == "" {
		helpers.HandleError(c, 404, "record not found")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":            true,
		"translation_about": translationAbout,
	})
}
