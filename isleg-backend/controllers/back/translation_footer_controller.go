package controllers

import (
	"github/abbgo/isleg/isleg-backend/config"
	"github/abbgo/isleg/isleg-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TranslationFooterForFooter struct {
	About   string `json:"about"`
	Payment string `json:"payment"`
	Contact string `json:"contact"`
	Secure  string `json:"secure"`
	Word    string `json:"word"`
}

func CreateTranslationFooter(c *gin.Context) {

	// GET ALL LANGUAGE
	languages, err := GetAllLanguageWithIDAndNameShort()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	dataNames := []string{"about", "payment", "contact", "secure", "word"}

	// VALIDATE DATA
	err = models.ValidateTranslationFooterData(languages, dataNames, c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	// create translation footer
	for _, v := range languages {
		_, err := config.ConnDB().Exec("INSERT INTO translation_footer (lang_id,about,payment,contact,secure,word) VALUES ($1,$2,$3,$4,$5,$6)", v.ID, c.PostForm("about_"+v.NameShort), c.PostForm("payment_"+v.NameShort), c.PostForm("contact_"+v.NameShort), c.PostForm("secure_"+v.NameShort), c.PostForm("word_"+v.NameShort))
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
		"message": "translation footer successfully added",
	})

}

func GetTranslationFooter(langID string) (TranslationFooterForFooter, error) {

	var t TranslationFooterForFooter

	// get translation footer where lang_id equal langID
	row, err := config.ConnDB().Query("SELECT about,payment,contact,secure,word FROM translation_footer WHERE lang_id = $1", langID)
	if err != nil {
		return TranslationFooterForFooter{}, err
	}

	for row.Next() {
		if err := row.Scan(&t.About, &t.Payment, &t.Contact, &t.Secure, &t.Word); err != nil {
			return TranslationFooterForFooter{}, err
		}
	}

	return t, nil

}
