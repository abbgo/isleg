package controllers

import (
	"github/abbgo/isleg/isleg-backend/config"
	"github/abbgo/isleg/isleg-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TrAbout struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func CreateTranslationAbout(c *gin.Context) {

	// GET ALL LANGUAGE
	languages, err := GetAllLanguageWithIDAndNameShort()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	dataNames := []string{"title", "content"}

	// VALIDATE DATA
	if err = models.ValidateTranslationAboutData(languages, dataNames, c); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	// create translation about
	for _, v := range languages {
		_, err := config.ConnDB().Exec("INSERT INTO translation_about (lang_id,title,content) VALUES ($1,$2,$3)", v.ID, c.PostForm("title_"+v.NameShort), c.PostForm("content_"+v.NameShort))
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
		"message": "translation about successfully added",
	})

}

func GetTranslationAbout(c *gin.Context) {

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

	// get translation about where lang_id equal langID
	aboutRow, err := config.ConnDB().Query("SELECT title,content FROM translation_about WHERE lang_id = $1", langID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	var translationAbout TrAbout

	for aboutRow.Next() {
		if err := aboutRow.Scan(&translationAbout.Title, &translationAbout.Content); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"status":            true,
		"translation_about": translationAbout,
	})

}
