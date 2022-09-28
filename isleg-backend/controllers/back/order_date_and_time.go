package controllers

import (
	"github/abbgo/isleg/isleg-backend/config"
	"github/abbgo/isleg/isleg-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

func CreateOrderTime(c *gin.Context) {

	db, err := config.ConnDB()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer db.Close()

	languages, err := GetAllLanguageWithIDAndNameShort()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	date := c.PostForm("date")
	times := c.PostFormArray("times")

	dataNames := []string{"translation_date"}

	if err := models.ValidateOrderDateAndTime(date, times, languages, dataNames, c); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	resultOrderDates, err := db.Query("INSERT INTO order_dates (date) VALUES ($1)", date)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer resultOrderDates.Close()

	lastOrderDateID, err := db.Query("SELECT id FROM order_dates ORDER BY created_at DESC LIMIT 1")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer lastOrderDateID.Close()

	var orderDateID string

	for lastOrderDateID.Next() {
		if err := lastOrderDateID.Scan(&orderDateID); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}

	lastID, err := uuid.Parse(orderDateID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	resultOrderTimes, err := db.Query("INSERT INTO order_times (order_date_id,time) VALUES ($1,unnest($2::varchar[]))", lastID, pq.Array(times))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer resultOrderTimes.Close()

	for _, v := range languages {

		resultTrOrderDates, err := db.Query("INSERT INTO translation_order_dates (lang_id,order_date_id,date) VALUES ($1,$2,$3)", v.ID, lastID, c.PostForm("translation_date_"+v.NameShort))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
		defer resultTrOrderDates.Close()

	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "order date and time successfully added",
	})

}
