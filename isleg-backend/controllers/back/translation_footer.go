package controllers

import (
	"github/abbgo/isleg/isleg-backend/config"
	"github/abbgo/isleg/isleg-backend/models"
	"net/http"

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

	var trFooters []models.TranslationFooter

	if err := c.BindJSON(&trFooters); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	// check lang_id
	for _, v := range trFooters {

		rowLang, err := db.Query("SELECT id FROM languages WHERE id = $1 AND deleted_atr IS NULL", v.LangID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
		defer func() {
			if err := rowLang.Close(); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  false,
					"message": err.Error(),
				})
				return
			}
		}()

		var langID string

		for rowLang.Next() {
			if err := rowLang.Scan(&langID); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  false,
					"message": err.Error(),
				})
				return
			}
		}

		if langID == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": "lamguage not found",
			})
			return
		}

	}

	// create translation footer
	for _, v := range trFooters {

		resultTRFooter, err := db.Query("INSERT INTO translation_footer (lang_id,about,payment,contact,secure,word) VALUES ($1,$2,$3,$4,$5,$6)", v.LangID, v.About, v.Payment, v.Contact, v.Contact, v.Word)
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

	// get id of translation footer from request parameter
	var trFooter models.TranslationFooter

	// check id
	rowFlag, err := db.Query("SELECT id FROM translation_footer WHERE id = $1 AND deleted_at IS NULL", trFooter.ID)
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

	// update data of translation footer
	rsultTRFooter, err := db.Query("UPDATE translation_footer SET about = $1, payment = $2, contact = $3, secure = $4, word = $5, lang_id = $7 WHERE id = $6", trFooter.About, trFooter.Payment, trFooter.Contact, trFooter.Secure, trFooter.Word, trFooter.ID, trFooter.LangID)
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
		"message": "data successfully updated",
	})

}

func GetTranslationFooterByID(c *gin.Context) {

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

	// get id of translation footer from request parameter
	trFootID := c.Param("id")

	//check id and get data from table
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
