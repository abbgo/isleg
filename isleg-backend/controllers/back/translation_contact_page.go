package controllers

import (
	"context"
	"github/abbgo/isleg/isleg-backend/config"
	"github/abbgo/isleg/isleg-backend/helpers"
	"github/abbgo/isleg/isleg-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateTranslationContact(c *gin.Context) {
	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	// get data from request
	var trContacts []models.TranslationContact
	if err := c.BindJSON(&trContacts); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	// check lang_id
	for _, v := range trContacts {
		var langID string
		db.QueryRow(context.Background(), "SELECT id FROM languages WHERE id = $1 AND deleted_at IS NULL", v.LangID).Scan(&langID)

		if langID == "" {
			helpers.HandleError(c, 404, "language not found")
			return
		}
	}

	// create translation contact
	for _, v := range trContacts {
		_, err := db.Exec(context.Background(), "INSERT INTO translation_contact (lang_id,full_name,email,phone,letter,company_phone,imo,company_email,instagram,button_text) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)", v.LangID, v.FullName, v.Email, v.Phone, v.Letter, v.CompanyPhone, v.Imo, v.CompanyEmail, v.Instragram, v.ButtonText)
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

func UpdateTranslationContactByID(c *gin.Context) {
	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	// get id of translation contact from request parameter
	var trContact models.TranslationContact
	if err := c.BindJSON(&trContact); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	// check id
	var id string
	db.QueryRow(context.Background(), "SELECT id FROM translation_contact WHERE id = $1 AND deleted_at IS NULL", trContact.ID).Scan(&id)

	if id == "" {
		helpers.HandleError(c, 404, "record not found")
		return
	}

	// update data
	_, err = db.Exec(context.Background(), "UPDATE translation_contact SET full_name = $1, email = $2 , phone = $3 , letter = $4 , company_phone = $5 , imo = $6, company_email = $7, instagram = $8, button_text = $9, lang_id = $11 WHERE id = $10", trContact.FullName, trContact.Email, trContact.Phone, trContact.Letter, trContact.CompanyPhone, trContact.Imo, trContact.CompanyEmail, trContact.Instragram, trContact.ButtonText, trContact.ID, trContact.LangID)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "data successfully updated",
	})
}

func GetTranslationContactByID(c *gin.Context) {
	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()
	// get id of translation contact from request parameter
	ID := c.Param("id")

	// check id and get data from database
	var t models.TranslationContact
	db.QueryRow(context.Background(), "SELECT id,full_name,email,phone,letter,company_phone,imo,company_email,instagram,button_text FROM translation_contact WHERE id = $1 AND deleted_at IS NULL", ID).Scan(&t.ID, &t.FullName, &t.Email, &t.Phone, &t.Letter, &t.CompanyPhone, &t.Imo, &t.CompanyEmail, &t.Instragram, &t.ButtonText)

	if t.ID == "" {
		helpers.HandleError(c, 404, "record not found")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":              true,
		"translation_contact": t,
	})
}

func GetTranslationContactByLangID(c *gin.Context) {
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	langID, err := CheckLanguage(c)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	// get translation contact where lang_id equal langID
	var translationContact models.TranslationContact
	db.QueryRow(context.Background(), "SELECT full_name,email,phone,letter,company_phone,imo,company_email,instagram,button_text FROM translation_contact WHERE lang_id = $1 AND deleted_at IS NULL", langID).Scan(&translationContact.FullName, &translationContact.Email, &translationContact.Phone, &translationContact.Letter, &translationContact.CompanyPhone, &translationContact.Imo, &translationContact.CompanyEmail, &translationContact.Instragram, &translationContact.ButtonText)

	if translationContact.FullName == "" {
		helpers.HandleError(c, 404, "record not found")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":              true,
		"translation_contact": translationContact,
	})
}
