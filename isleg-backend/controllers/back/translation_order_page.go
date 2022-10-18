package controllers

import (
	"github/abbgo/isleg/isleg-backend/config"
	"github/abbgo/isleg/isleg-backend/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateTranslationOrderPage(c *gin.Context) {

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

	dataNames := []string{"content", "type_of_payment", "choose_a_delivery_time", "your_address", "mark", "to_order", "tomorrow", "cash", "payment_terminal"}

	// VALIDATE DATA
	if err = models.ValidateTranslationOrderPageData(languages, dataNames, c); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	// create translation_my_information_page
	for _, v := range languages {
		resultTrOrderPage, err := db.Query("INSERT INTO translation_order_page (lang_id,content,type_of_payment,choose_a_delivery_time,your_address,mark,to_order,tomorrow,cash,payment_terminal) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)", v.ID, c.PostForm("content_"+v.NameShort), c.PostForm("type_of_payment_"+v.NameShort), c.PostForm("choose_a_delivery_time_"+v.NameShort), c.PostForm("your_address_"+v.NameShort), c.PostForm("mark_"+v.NameShort), c.PostForm("to_order_"+v.NameShort), c.PostForm("tomorrow_"+v.NameShort), c.PostForm("cash_"+v.NameShort), c.PostForm("payment_terminal_"+v.NameShort))
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
		"message": "translation order page successfully added",
	})

}

func UpdateTranslationOrderPageByID(c *gin.Context) {

	db, err := config.ConnDB()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer db.Close()

	ID := c.Param("id")

	rowTrOrderPage, err := db.Query("SELECT id FROM translation_order_page WHERE id = $1 AND deleted_at IS NULL", ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer rowTrOrderPage.Close()

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
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "record not found",
		})
		return
	}

	dataNames := []string{"content", "type_of_payment", "choose_a_delivery_time", "your_address", "mark", "to_order", "tomorrow", "cash", "payment_terminal"}

	// VALIDATE DATA
	err = models.ValidateTranslationOrderPageUpdate(dataNames, c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	currentTime := time.Now()

	resultTrOrderPage, err := db.Query("UPDATE translation_order_page SET content = $1, type_of_payment = $2 , choose_a_delivery_time = $3, your_address = $4 , mark = $5 , to_order = $6 , tomorrow = $7, cash = $8 , payment_terminal = $9, updated_at = $10 WHERE id = $11", c.PostForm("content"), c.PostForm("type_of_payment"), c.PostForm("choose_a_delivery_time"), c.PostForm("your_address"), c.PostForm("mark"), c.PostForm("to_order"), c.PostForm("tomorrow"), c.PostForm("cash"), c.PostForm("payment_terminal"), currentTime, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer resultTrOrderPage.Close()

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "translation order page successfully updated",
	})

}

func GetTranslationOrderPageByID(c *gin.Context) {

	db, err := config.ConnDB()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer db.Close()

	ID := c.Param("id")

	rowTrOrderPage, err := db.Query("SELECT content,type_of_payment,choose_a_delivery_time,your_address,mark,to_order,tomorrow,cash,payment_terminal FROM translation_order_page WHERE id = $1 AND deleted_at IS NULL", ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer rowTrOrderPage.Close()

	var t TrOrderPage

	for rowTrOrderPage.Next() {
		if err := rowTrOrderPage.Scan(&t.Content, &t.TypeOfPayment, &t.ChooseADeliveryTime, &t.YourAddress, &t.Mark, &t.ToOrder, &t.Tomorrow, &t.Cash, &t.PaymentTerminal); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}

	if t.Content == "" {
		c.JSON(http.StatusBadRequest, gin.H{
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
	defer db.Close()

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
	rowTrOrderPage, err := db.Query("SELECT content,type_of_payment,choose_a_delivery_time,your_address,mark,to_order,tomorrow,cash,payment_terminal FROM translation_order_page WHERE lang_id = $1 AND deleted_at IS NULL", langID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer rowTrOrderPage.Close()

	var t TrOrderPage

	for rowTrOrderPage.Next() {
		if err := rowTrOrderPage.Scan(&t.Content, &t.TypeOfPayment, &t.ChooseADeliveryTime, &t.YourAddress, &t.Mark, &t.ToOrder, &t.Tomorrow, &t.Cash, &t.PaymentTerminal); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}

	if t.Content == "" {
		c.JSON(http.StatusBadRequest, gin.H{
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
