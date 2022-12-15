package controllers

import (
	"github/abbgo/isleg/isleg-backend/config"
	"github/abbgo/isleg/isleg-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateTranslationOrderPage(c *gin.Context) {

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

	// get data from request
	var trOrderPages []models.TranslationOrderPage

	if err := c.BindJSON(&trOrderPages); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	// check lang_id
	for _, v := range trOrderPages {

		rowLang, err := db.Query("SELECT id FROM languages WHERE id = $1 AND deleted_at IS NULL", v.LangID)
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
			c.JSON(http.StatusNotFound, gin.H{
				"status":  false,
				"message": "language not found",
			})
			return
		}

	}

	// create translation_my_information_page
	for _, v := range trOrderPages {

		resultTrOrderPage, err := db.Query("INSERT INTO translation_order_page (lang_id,content,type_of_payment,choose_a_delivery_time,your_address,mark,to_order) VALUES ($1,$2,$3,$4,$5,$6,$7)", v.LangID, v.Content, v.TypeOfPayment, v.ChooseADeliveryTime, v.YourAddress, v.Mark, v.ToOrder)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
		defer func() {
			if err := resultTrOrderPage.Close(); err != nil {
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

func UpdateTranslationOrderPageByID(c *gin.Context) {

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

	// get id from request parameter
	var trOrderPage models.TranslationOrderPage

	if err := c.BindJSON(&trOrderPage); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	// check id
	rowTrOrderPage, err := db.Query("SELECT id FROM translation_order_page WHERE id = $1 AND deleted_at IS NULL", trOrderPage.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer func() {
		if err := rowTrOrderPage.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

	var id string

	for rowTrOrderPage.Next() {
		if err := rowTrOrderPage.Scan(&id); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}

	if id == "" {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  false,
			"message": "record not found",
		})
		return
	}

	resultTrOrderPage, err := db.Query("UPDATE translation_order_page SET content = $1, type_of_payment = $2 , choose_a_delivery_time = $3, your_address = $4 , mark = $5 , to_order = $6, lang_id = $8 WHERE id = $7", trOrderPage.Content, trOrderPage.TypeOfPayment, trOrderPage.ChooseADeliveryTime, trOrderPage.YourAddress, trOrderPage.Mark, trOrderPage.ToOrder, trOrderPage.ID, trOrderPage.LangID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer func() {
		if err := resultTrOrderPage.Close(); err != nil {
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

func GetTranslationOrderPageByID(c *gin.Context) {

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

	// get id from request parameter
	ID := c.Param("id")

	// check id and get data from database
	rowTrOrderPage, err := db.Query("SELECT id,content,type_of_payment,choose_a_delivery_time,your_address,mark,to_order FROM translation_order_page WHERE id = $1 AND deleted_at IS NULL", ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer func() {
		if err := rowTrOrderPage.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

	var t models.TranslationOrderPage

	for rowTrOrderPage.Next() {
		if err := rowTrOrderPage.Scan(&t.ID, &t.Content, &t.TypeOfPayment, &t.ChooseADeliveryTime, &t.YourAddress, &t.Mark, &t.ToOrder); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}

	if t.ID == "" {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  false,
			"message": "record not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":                 true,
		"translation_order_page": t,
	})

}

func GetTranslationOrderPageByLangID(c *gin.Context) {

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

	// get translation-basket-page where lang_id equal langID
	rowTrOrderPage, err := db.Query("SELECT content,type_of_payment,choose_a_delivery_time,your_address,mark,to_order FROM translation_order_page WHERE lang_id = $1 AND deleted_at IS NULL", langID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer func() {
		if err := rowTrOrderPage.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

	var t models.TranslationOrderPage

	for rowTrOrderPage.Next() {
		if err := rowTrOrderPage.Scan(&t.Content, &t.TypeOfPayment, &t.ChooseADeliveryTime, &t.YourAddress, &t.Mark, &t.ToOrder); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}

	if t.Content == "" {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  false,
			"message": "record not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":                 true,
		"translation_order_page": t,
	})

}
