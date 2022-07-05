package controllers

import (
	"github/abbgo/isleg/isleg-backend/config"
	"github/abbgo/isleg/isleg-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateTranslationUpdatePasswordPage(c *gin.Context) {

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
		if c.PostForm("password_"+v.NameShort) == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": "password_" + v.NameShort + " is required",
			})
			return
		}
	}

	for _, v := range languages {
		if c.PostForm("verify_password_"+v.NameShort) == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": "verify_password_" + v.NameShort + " is required",
			})
			return
		}
	}

	for _, v := range languages {
		if c.PostForm("explanation_"+v.NameShort) == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": "explanation_" + v.NameShort + " is required",
			})
			return
		}
	}

	for _, v := range languages {
		if c.PostForm("save_"+v.NameShort) == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": "save_" + v.NameShort + " is required",
			})
			return
		}
	}

	for _, v := range languages {
		_, err := config.ConnDB().Exec("INSERT INTO translation_update_password_page (lang_id,title,password,verify_password,explanation,save) VALUES ($1,$2,$3,$4,$5,$6)", v.ID, c.PostForm("title_"+v.NameShort), c.PostForm("password_"+v.NameShort), c.PostForm("verify_password_"+v.NameShort), c.PostForm("explanation_"+v.NameShort), c.PostForm("save_"+v.NameShort))
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
		"message": "translation update password page successfully added",
	})

}
