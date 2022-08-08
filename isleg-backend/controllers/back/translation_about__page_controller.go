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

func UpdateTranslationAbout(c *gin.Context) {

	ID := c.Param("id")

	rowFlag, err := config.ConnDB().Query("SELECT id FROM translation_about WHERE id = $1 AND deleted_at IS NULL", ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	var id string

	for rowFlag.Next() {
		if err := rowFlag.Scan(&id); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}

	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "record not found",
		})
		return
	}

	dataNames := []string{"title", "content"}

	// VALIDATE DATA
	err = models.ValidateTranslationAboutUpdate(dataNames, c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	_, err = config.ConnDB().Exec("UPDATE translation_about SET title = $1, content = $2  WHERE id = $3", c.PostForm("title"), c.PostForm("content"), id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "translation about successfully updated",
	})

}

func GetOneTranslationAbout(c *gin.Context) {

	ID := c.Param("id")

	rowFlag, err := config.ConnDB().Query("SELECT title,content FROM translation_about WHERE id = $1 AND deleted_at IS NULL", ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	var t TrAbout

	for rowFlag.Next() {
		if err := rowFlag.Scan(&t.Title, &t.Content); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}

	if t.Title == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "record not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":            true,
		"translation_about": t,
	})

}

func GetTranslationAbout(c *gin.Context) {

	langID, err := CheckLanguage(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	// get translation about where lang_id equal langID
	aboutRow, err := config.ConnDB().Query("SELECT title,content FROM translation_about WHERE lang_id = $1 AND deleted_at IS NULL", langID)
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
