package controllers

import (
	"github/abbgo/isleg/isleg-backend/config"
	"github/abbgo/isleg/isleg-backend/models"
	"github/abbgo/isleg/isleg-backend/pkg"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateTranslationAbout(c *gin.Context) {

	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer func() {
		if err := db.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

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
	if err = pkg.ValidateTranslations(languages, dataNames, c); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	// create translation about
	for _, v := range languages {
		resultTRAbout, err := db.Query("INSERT INTO translation_about (lang_id,title,content) VALUES ($1,$2,$3)", v.ID, c.PostForm("title_"+v.NameShort), c.PostForm("content_"+v.NameShort))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
		defer func() {
			if err := resultTRAbout.Close(); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  false,
					"message": err.Error(),
				})
				return
			}
		}()
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "data successfully added",
	})

}

func UpdateTranslationAboutByID(c *gin.Context) {

	db, err := config.ConnDB()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer func() {
		if err := db.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

	ID := c.Param("id")

	rowFlag, err := db.Query("SELECT id FROM translation_about WHERE id = $1 AND deleted_at IS NULL", ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer func() {
		if err := rowFlag.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

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

	currentTime := time.Now()

	resultTRABout, err := db.Query("UPDATE translation_about SET title = $1, content = $2 , updated_at = $4 WHERE id = $3", c.PostForm("title"), c.PostForm("content"), id, currentTime)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer func() {
		if err := resultTRABout.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "translation about successfully updated",
	})

}

func GetTranslationAboutByID(c *gin.Context) {

	db, err := config.ConnDB()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer func() {
		if err := db.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

	ID := c.Param("id")

	rowFlag, err := db.Query("SELECT title,content FROM translation_about WHERE id = $1 AND deleted_at IS NULL", ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer func() {
		if err := rowFlag.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

	var t models.TranslationAbout

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

func GetTranslationAboutByLangID(c *gin.Context) {

	db, err := config.ConnDB()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer func() {
		if err := db.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

	langID, err := CheckLanguage(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	// get translation about where lang_id equal langID
	aboutRow, err := db.Query("SELECT title,content FROM translation_about WHERE lang_id = $1 AND deleted_at IS NULL", langID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer func() {
		if err := aboutRow.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

	var translationAbout models.TranslationAbout

	for aboutRow.Next() {
		if err := aboutRow.Scan(&translationAbout.Title, &translationAbout.Content); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}

	if translationAbout.Title == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "record not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":            true,
		"translation_about": translationAbout,
	})

}
