package controllers

import (
	"github/abbgo/isleg/isleg-backend/config"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

type DateHour struct {
	Hour   uint     `json:"hour,omitempty" binding:"required"`
	DateID string   `json:"date_id,omitempty" binding:"required"`
	Times  []string `json:"times,omitempty" binding:"required"`
}

func CreateOrderDateHour(c *gin.Context) {

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

	var dateHour DateHour

	if err := c.BindJSON(&dateHour); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	// validate data
	if dateHour.Hour > 23 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "hour cannot be older than 23",
		})
		return
	}

	resultDateHour, err := db.Query("INSERT INTO date_hours (hour,date_id) VALUES ($1,$2) RETURNING id", dateHour.Hour, dateHour.DateID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer func() {
		if err := resultDateHour.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

	var dateHourID string

	for resultDateHour.Next() {
		if err := resultDateHour.Scan(&dateHourID); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}

	resultDateHourTime, err := db.Query("INSERT INTO date_hour_times (date_hour_id,time_id) VALUES ($1,unnest($2::uuid[])) RETURNING id", dateHourID, pq.Array(dateHour.Times))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer func() {
		if err := resultDateHourTime.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "data successfully added",
	})

}
