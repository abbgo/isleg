package controllers

import (
	"github/abbgo/isleg/isleg-backend/config"
	"github/abbgo/isleg/isleg-backend/models"
	"github/abbgo/isleg/isleg-backend/pkg"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateTranslationFooter(c *gin.Context) {

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

	dataNames := []string{"about", "payment", "contact", "secure", "word"}

	// VALIDATE DATA
	err = pkg.ValidateTranslations(languages, dataNames, c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	// create translation footer
	for _, v := range languages {
		resultTRFooter, err := db.Query("INSERT INTO translation_footer (lang_id,about,payment,contact,secure,word) VALUES ($1,$2,$3,$4,$5,$6)", v.ID, c.PostForm("about_"+v.NameShort), c.PostForm("payment_"+v.NameShort), c.PostForm("contact_"+v.NameShort), c.PostForm("secure_"+v.NameShort), c.PostForm("word_"+v.NameShort))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
		defer func() {
			if err := resultTRFooter.Close(); err != nil {
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

func UpdateTranslationFooterByID(c *gin.Context) {

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

	trFootID := c.Param("id")

	rowFlag, err := db.Query("SELECT id FROM translation_footer WHERE id = $1 AND deleted_at IS NULL", trFootID)
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

	currentTime := time.Now()

	rsultTRFooter, err := db.Query("UPDATE translation_footer SET about = $1, payment = $2, contact = $3, secure = $4, word = $5 , updated_at = $7 WHERE id = $6", c.PostForm("about"), c.PostForm("payment"), c.PostForm("contact"), c.PostForm("secure"), c.PostForm("word"), id, currentTime)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer func() {
		if err := rsultTRFooter.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "translation footer successfully updated",
	})

}

func GetTranslationFooterByID(c *gin.Context) {

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

	trFootID := c.Param("id")

	rowFlag, err := db.Query("SELECT about,payment,contact,secure,word FROM translation_footer WHERE id = $1 AND deleted_at IS NULL", trFootID)
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

	var t models.TranslationFooter

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

func GetTranslationFooter(langID string) (models.TranslationFooter, error) {

	db, err := config.ConnDB()
	if err != nil {
		return models.TranslationFooter{}, err
	}
	defer func() (models.TranslationFooter, error) {
		if err := db.Close(); err != nil {
			return models.TranslationFooter{}, err
		}
		return models.TranslationFooter{}, nil
	}()

	var t models.TranslationFooter

	// get translation footer where lang_id equal langID
	row, err := db.Query("SELECT about,payment,contact,secure,word FROM translation_footer WHERE lang_id = $1 AND deleted_at IS NULL", langID)
	if err != nil {
		return models.TranslationFooter{}, err
	}
	defer func() (models.TranslationFooter, error) {
		if err := row.Close(); err != nil {
			return models.TranslationFooter{}, err
		}
		return models.TranslationFooter{}, nil
	}()

	for row.Next() {
		if err := row.Scan(&t.About, &t.Payment, &t.Contact, &t.Secure, &t.Word); err != nil {
			return models.TranslationFooter{}, err
		}
	}

	return t, nil

}
