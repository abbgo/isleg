package controllers

import (
	"github/abbgo/isleg/isleg-backend/config"
	"github/abbgo/isleg/isleg-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateTranslationHeader(c *gin.Context) {

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
		if c.PostForm("research_"+v.NameShort) == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": "research_" + v.NameShort + " is required",
			})
			return
		}
	}
	for _, v := range languages {
		if c.PostForm("phone_"+v.NameShort) == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": "phone_" + v.NameShort + " is required",
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
		if c.PostForm("forgot_password_"+v.NameShort) == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": "forgot_password_" + v.NameShort + " is required",
			})
			return
		}
	}
	for _, v := range languages {
		if c.PostForm("sign_in_"+v.NameShort) == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": "sign_in_" + v.NameShort + " is required",
			})
			return
		}
	}
	for _, v := range languages {
		if c.PostForm("sign_up_"+v.NameShort) == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": "sign_up_" + v.NameShort + " is required",
			})
			return
		}
	}
	for _, v := range languages {
		if c.PostForm("name_"+v.NameShort) == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": "name_" + v.NameShort + " is required",
			})
			return
		}
	}
	for _, v := range languages {
		if c.PostForm("password_verification_"+v.NameShort) == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": "password_verification_" + v.NameShort + " is required",
			})
			return
		}
	}
	for _, v := range languages {
		if c.PostForm("verify_secure_"+v.NameShort) == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": "verify_secure_" + v.NameShort + " is required",
			})
			return
		}
	}
	for _, v := range languages {
		if c.PostForm("my_information_"+v.NameShort) == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": "my_information_" + v.NameShort + " is required",
			})
			return
		}
	}
	for _, v := range languages {
		if c.PostForm("my_favorites_"+v.NameShort) == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": "my_favorites_" + v.NameShort + " is required",
			})
			return
		}
	}
	for _, v := range languages {
		if c.PostForm("my_orders_"+v.NameShort) == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": "my_orders_" + v.NameShort + " is required",
			})
			return
		}
	}
	for _, v := range languages {
		if c.PostForm("log_out_"+v.NameShort) == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": "log_out_" + v.NameShort + " is required",
			})
			return
		}
	}

	// CREATE TRANSLATION HEADER
	for _, v := range languages {
		_, err := config.ConnDB().Exec("INSERT INTO translation_header (lang_id,research,phone,password,forgot_password,sign_in,sign_up,name,password_verification,verify_secure,my_information,my_favorites,my_orders,log_out) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14)", v.ID, c.PostForm("research_"+v.NameShort), c.PostForm("phone_"+v.NameShort), c.PostForm("password_"+v.NameShort), c.PostForm("forgot_password_"+v.NameShort), c.PostForm("sign_in_"+v.NameShort), c.PostForm("sign_up_"+v.NameShort), c.PostForm("name_"+v.NameShort), c.PostForm("password_verification_"+v.NameShort), c.PostForm("verify_secure_"+v.NameShort), c.PostForm("my_information_"+v.NameShort), c.PostForm("my_favorites_"+v.NameShort), c.PostForm("my_orders_"+v.NameShort), c.PostForm("log_out_"+v.NameShort))
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
		"message": "translation header successfully added",
	})

}
