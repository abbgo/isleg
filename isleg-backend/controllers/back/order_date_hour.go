package controllers

import (
	"context"
	"github/abbgo/isleg/isleg-backend/config"
	"github/abbgo/isleg/isleg-backend/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

type DateHour struct {
	Hour   uint     `json:"hour,omitempty"`
	DateID string   `json:"date_id,omitempty" binding:"required"`
	Times  []string `json:"times,omitempty" binding:"required"`
}

func CreateOrderDateHour(c *gin.Context) {

	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	var dateHours []DateHour

	if err := c.BindJSON(&dateHours); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	// validate data
	for _, v := range dateHours {
		if v.Hour > 23 {
			helpers.HandleError(c, 400, "hour cannot be older than 23")
			return
		}

		resultDateHour, err := db.Query(context.Background(), "INSERT INTO date_hours (hour,date_id) VALUES ($1,$2) RETURNING id", v.Hour, v.DateID)
		if err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}

		var dateHourID string
		for resultDateHour.Next() {
			if err := resultDateHour.Scan(&dateHourID); err != nil {
				helpers.HandleError(c, 400, err.Error())
				return
			}
		}

		_, err = db.Exec(context.Background(), "INSERT INTO date_hour_times (date_hour_id,time_id) VALUES ($1,unnest($2::uuid[])) RETURNING id", dateHourID, pq.Array(v.Times))
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
