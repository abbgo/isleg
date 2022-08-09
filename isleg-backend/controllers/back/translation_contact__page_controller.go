package controllers

import (
	"github/abbgo/isleg/isleg-backend/config"
	"github/abbgo/isleg/isleg-backend/models"
	"net/http"
	"time"

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

func UpdateTranslationContactByID(c *gin.Context) {

	ID := c.Param("id")

	rowFlag, err := config.ConnDB().Query("SELECT id FROM translation_contact WHERE id = $1 AND deleted_at IS NULL", ID)
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

	dataNames := []string{"full_name", "email", "phone", "letter", "company_phone", "imo", "company_email", "instagram", "button_text"}

	// VALIDATE DATA
	err = models.ValidateTranslationContactUpdate(dataNames, c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	currentTime := time.Now()

	_, err = config.ConnDB().Exec("UPDATE translation_contact SET full_name = $1, email = $2 , phone = $3 , letter = $4 , company_phone = $5 , imo = $6, company_email = $7, instagram = $8, button_text = $9 , updated_at = $11 WHERE id = $10", c.PostForm("full_name"), c.PostForm("email"), c.PostForm("phone"), c.PostForm("letter"), c.PostForm("company_phone"), c.PostForm("imo"), c.PostForm("company_email"), c.PostForm("instagram"), c.PostForm("button_text"), id, currentTime)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "translation contact successfully updated",
	})

}

func GetTranslationContactByID(c *gin.Context) {

	ID := c.Param("id")

	rowFlag, err := config.ConnDB().Query("SELECT full_name,email,phone,letter,company_phone,imo,company_email,instagram,button_text FROM translation_contact WHERE id = $1 AND deleted_at IS NULL", ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	var t TrContact

	for rowFlag.Next() {
		if err := rowFlag.Scan(&t.FullName, &t.Email, &t.Phone, &t.Letter, &t.CompanyPhone, &t.Imo, &t.CompanyEmail, &t.Instragram, &t.ButtonText); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}

	if t.FullName == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "record not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":              true,
		"translation_contact": t,
	})

}

func GetTranslationContactByLangID(c *gin.Context) {

	langID, err := CheckLanguage(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	// get translation contact where lang_id equal langID
	aboutRow, err := config.ConnDB().Query("SELECT full_name,email,phone,letter,company_phone,imo,company_email,instagram,button_text FROM translation_contact WHERE lang_id = $1 AND deleted_at IS NULL", langID)
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
