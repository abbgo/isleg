package controllers

import (
	"github/abbgo/isleg/isleg-backend/config"
	"github/abbgo/isleg/isleg-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TrSecure struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func CreateTranslationSecure(c *gin.Context) {

	// GET ALL LANGUAGE
	languageRows, err := config.ConnDB().Query("SELECT id,name_short FROM languages ORDER BY created_at ASC")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	var languages []models.Language

	for languageRows.Next() {
		var language models.Language
		if err := languageRows.Scan(&language.ID, &language.NameShort); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
		languages = append(languages, language)
	}

	// VALIDATE DATA
	for _, v := range languages {
		if c.PostForm("title_"+v.NameShort) == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": "title_" + v.NameShort + " is required",
			})
			return
		}
	}

	for _, v := range languages {
		if c.PostForm("content_"+v.NameShort) == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": "content_" + v.NameShort + " is required",
			})
			return
		}
	}

	// create translation secure
	for _, v := range languages {
		_, err := config.ConnDB().Exec("INSERT INTO translation_secure (lang_id,title,content) VALUES ($1,$2,$3)", v.ID, c.PostForm("title_"+v.NameShort), c.PostForm("content_"+v.NameShort))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "translation secure successfully added",
	})

}

func GetTranslationSecure(c *gin.Context) {

	// GET DATA FROM ROUTE PARAMETER
	langShortName := c.Param("lang")

	// GET language id
	var langID string
	row, err := config.ConnDB().Query("SELECT id FROM languages WHERE name_short = $1", langShortName)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	for row.Next() {
		if err := row.Scan(&langID); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}

	// get translation secure where lang_id equal langID
	secureRow, err := config.ConnDB().Query("SELECT title,content FROM translation_secure WHERE lang_id = $1", langID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	var translationSecure TrSecure

	for secureRow.Next() {
		if err := secureRow.Scan(&translationSecure.Title, &translationSecure.Content); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"status":             true,
		"translation_secure": translationSecure,
	})
}
