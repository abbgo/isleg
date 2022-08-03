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

func UpdateTranslationFooter(c *gin.Context) {

	trFootID := c.Param("id")

	rowFlag, err := config.ConnDB().Query("SELECT id FROM translation_footer WHERE id = $1 AND deleted_at IS NULL", trFootID)
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

	dataNames := []string{"about", "payment", "contact", "secure", "word"}

	// VALIDATE DATA
	err = models.ValidateTranslationFooterUpdate(dataNames, c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	_, err = config.ConnDB().Exec("UPDATE translation_footer SET about = $1, payment = $2, contact = $3, secure = $4, word = $5  WHERE id = $6", c.PostForm("about"), c.PostForm("payment"), c.PostForm("contact"), c.PostForm("secure"), c.PostForm("word"), id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "translation footer successfully updated",
	})

}

func GetOneTranslationFooter(c *gin.Context) {

	trFootID := c.Param("id")

	rowFlag, err := config.ConnDB().Query("SELECT about,payment,contact,secure,word FROM translation_footer WHERE id = $1 AND deleted_at IS NULL", trFootID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	var t TranslationFooterForFooter

	for rowFlag.Next() {
		if err := rowFlag.Scan(&t.About, &t.Payment, &t.Contact, &t.Secure, &t.Word); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}

	if t.About == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "record not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":             true,
		"translation_footer": t,
	})

}

func GetTranslationFooter(langID string) (TranslationFooterForFooter, error) {

	var t TranslationFooterForFooter

	// get translation footer where lang_id equal langID
	row, err := config.ConnDB().Query("SELECT about,payment,contact,secure,word FROM translation_footer WHERE lang_id = $1 AND deleted_at IS NULL", langID)
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
