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
	ButtonText   string `json:"button_text"`
}

func CreateTranslationContact(c *gin.Context) {

	// GET ALL LANGUAGE
	languages, err := GetAllLanguageWithIDAndNameShort()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	dataNames := []string{"full_name", "email", "phone", "letter", "company_phone", "imo", "company_email", "instagram", "button_text"}

	// VALIDATE DATA
	if err = models.ValidateTranslationContactData(languages, dataNames, c); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	// create translation contact
	for _, v := range languages {
		_, err := config.ConnDB().Exec("INSERT INTO translation_contact (lang_id,full_name,email,phone,letter,company_phone,imo,company_email,instagram,button_text) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)", v.ID, c.PostForm("full_name_"+v.NameShort), c.PostForm("email_"+v.NameShort), c.PostForm("phone_"+v.NameShort), c.PostForm("letter_"+v.NameShort), c.PostForm("company_phone_"+v.NameShort), c.PostForm("imo_"+v.NameShort), c.PostForm("company_email_"+v.NameShort), c.PostForm("instagram_"+v.NameShort), c.PostForm("button_text_"+v.NameShort))
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
	langID, err := GetLangID(langShortName)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	// get translation contact where lang_id equal langID
	aboutRow, err := config.ConnDB().Query("SELECT full_name,email,phone,letter,company_phone,imo,company_email,instagram,button_text FROM translation_contact WHERE lang_id = $1", langID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	var translationContact TrContact

	for aboutRow.Next() {
		if err := aboutRow.Scan(&translationContact.FullName, &translationContact.Email, &translationContact.Phone, &translationContact.Letter, &translationContact.CompanyPhone, &translationContact.Imo, &translationContact.CompanyEmail, &translationContact.Instragram, &translationContact.ButtonText); err != nil {
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
