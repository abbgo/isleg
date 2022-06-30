package controllers

import (
	"github/abbgo/isleg/isleg-backend/config"
	"github/abbgo/isleg/isleg-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TrContact struct {
	FullName     string `json:"full_name"`
	Email        string `json:"email"`
	Phone        string `json:"phone"`
	Letter       string `json:"letter"`
	CompanyPhone string `json:"company_phone"`
	Imo          string `json:"imo"`
	CompanyEmail string `json:"company_email"`
	Instragram   string `json:"instagram"`
}

func CreateTranslationContact(c *gin.Context) {

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
		if c.PostForm("full_name_"+v.NameShort) == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": "full_name_" + v.NameShort + " is required",
			})
			return
		}
	}

	for _, v := range languages {
		if c.PostForm("email_"+v.NameShort) == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": "email_" + v.NameShort + " is required",
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
		if c.PostForm("letter_"+v.NameShort) == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": "letter_" + v.NameShort + " is required",
			})
			return
		}
	}

	for _, v := range languages {
		if c.PostForm("company_phone_"+v.NameShort) == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": "company_phone_" + v.NameShort + " is required",
			})
			return
		}
	}

	for _, v := range languages {
		if c.PostForm("imo_"+v.NameShort) == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": "imo_" + v.NameShort + " is required",
			})
			return
		}
	}

	for _, v := range languages {
		if c.PostForm("company_email_"+v.NameShort) == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": "company_email_" + v.NameShort + " is required",
			})
			return
		}
	}

	for _, v := range languages {
		if c.PostForm("instagram_"+v.NameShort) == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": "instagram_" + v.NameShort + " is required",
			})
			return
		}
	}

	// create translation contact
	for _, v := range languages {
		_, err := config.ConnDB().Exec("INSERT INTO translation_contact (lang_id,full_name,email,phone,letter,company_phone,imo,company_email,instagram) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9)", v.ID, c.PostForm("full_name_"+v.NameShort), c.PostForm("email_"+v.NameShort), c.PostForm("phone_"+v.NameShort), c.PostForm("letter_"+v.NameShort), c.PostForm("company_phone_"+v.NameShort), c.PostForm("imo_"+v.NameShort), c.PostForm("company_email_"+v.NameShort), c.PostForm("instagram_"+v.NameShort))
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
		"message": "translation contact successfully added",
	})

}

func GetTranslationContact(c *gin.Context) {

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

	// get translation contact where lang_id equal langID
	aboutRow, err := config.ConnDB().Query("SELECT full_name,email,phone,letter,company_phone,imo,company_email,instagram FROM translation_contact WHERE lang_id = $1", langID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	var translationContact TrContact

	for aboutRow.Next() {
		if err := aboutRow.Scan(&translationContact.FullName, &translationContact.Email, &translationContact.Phone, &translationContact.Letter, &translationContact.CompanyPhone, &translationContact.Imo, &translationContact.CompanyEmail, &translationContact.Instragram); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"status":              true,
		"translation_contact": translationContact,
	})

}
