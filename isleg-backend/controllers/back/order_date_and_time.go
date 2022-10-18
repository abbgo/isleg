package controllers

import (
	"fmt"
	"github/abbgo/isleg/isleg-backend/config"
	"github/abbgo/isleg/isleg-backend/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

type OrderDateAndTime struct {
	ID              string      `json:"-"`
	Date            string      `json:"date"`
	Times           []OrderTime `json:"times"`
	TranslationDate string      `json:"translation_date"`
}

type OrderTime struct {
	Time string `json:"time"`
}

func CreateOrderTime(c *gin.Context) {

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
	defer func() {
		if err := resultOrderDates.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

	lastOrderDateID, err := db.Query("SELECT id FROM order_dates ORDER BY created_at DESC LIMIT 1")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer func() {
		if err := lastOrderDateID.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

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
	defer func() {
		if err := resultOrderTimes.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

	for _, v := range languages {

		resultTrOrderDates, err := db.Query("INSERT INTO translation_order_dates (lang_id,order_date_id,date) VALUES ($1,$2,$3)", v.ID, lastID, c.PostForm("translation_date_"+v.NameShort))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
		defer func() {
			if err := resultTrOrderDates.Close(); err != nil {
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
		"message": "order date and time successfully added",
	})

}

func GetOrderTime(c *gin.Context) {

	db, err := config.ConnDB()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer db.Close()

	langID, err := CheckLanguage(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	currentHour := time.Now().Hour()
	dates := []string{}
	times := []string{}

	if 0 <= currentHour && currentHour < 6 {

		dates = append(dates, "today")

	} else if 6 <= currentHour && currentHour < 14 {

		dates = append(dates, "today", "tomorrow")

		times = append(times, "18:00 - 21:00", "09:00 - 12:00")

	} else if 14 <= currentHour && currentHour < 24 {

		dates = append(dates, "tomorrow")

	}

	fmt.Println(dates, times)

	rowsOrderDate, err := db.Query("select od.id , od.date , tod.date from order_dates od inner join translation_order_dates tod on tod.order_date_id = od.id where tod.lang_id = $1 and od.deleted_at is null and tod.deleted_at is null and od.date = any($2)", langID, pq.Array(dates))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer rowsOrderDate.Close()

	var orderDateAndTimes []OrderDateAndTime

	for rowsOrderDate.Next() {
		var orderDateAndTime OrderDateAndTime

		if err := rowsOrderDate.Scan(&orderDateAndTime.ID, &orderDateAndTime.Date, &orderDateAndTime.TranslationDate); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}

		if len(times) == 0 {

			rowsOrderTime, err := db.Query("SELECT time FROM order_times WHERE order_date_id = $1 AND deleted_at IS NULL", orderDateAndTime.ID)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  false,
					"message": err.Error(),
				})
				return
			}
			defer rowsOrderTime.Close()

			var orderTimes []OrderTime

			for rowsOrderTime.Next() {
				var orderTime OrderTime

				if err := rowsOrderTime.Scan(&orderTime.Time); err != nil {
					c.JSON(http.StatusBadRequest, gin.H{
						"status":  false,
						"message": err.Error(),
					})
					return
				}

				orderTimes = append(orderTimes, orderTime)

			}

			orderDateAndTime.Times = orderTimes

		} else {

			if orderDateAndTime.Date == "today" {

				rowsOrderTime, err := db.Query("SELECT time FROM order_times WHERE order_date_id = $1 AND deleted_at IS NULL AND time = $2", orderDateAndTime.ID, times[0])
				if err != nil {
					c.JSON(http.StatusBadRequest, gin.H{
						"status":  false,
						"message": err.Error(),
					})
					return
				}
				defer rowsOrderTime.Close()

				var orderTimes []OrderTime

				for rowsOrderTime.Next() {
					var orderTime OrderTime

					if err := rowsOrderTime.Scan(&orderTime.Time); err != nil {
						c.JSON(http.StatusBadRequest, gin.H{
							"status":  false,
							"message": err.Error(),
						})
						return
					}

					orderTimes = append(orderTimes, orderTime)

				}

				orderDateAndTime.Times = orderTimes

			} else if orderDateAndTime.Date == "tomorrow" {

				rowsOrderTime, err := db.Query("SELECT time FROM order_times WHERE order_date_id = $1 AND deleted_at IS NULL AND time = $2", orderDateAndTime.ID, times[1])
				if err != nil {
					c.JSON(http.StatusBadRequest, gin.H{
						"status":  false,
						"message": err.Error(),
					})
					return
				}
				defer rowsOrderTime.Close()

				var orderTimes []OrderTime

				for rowsOrderTime.Next() {
					var orderTime OrderTime

					if err := rowsOrderTime.Scan(&orderTime.Time); err != nil {
						c.JSON(http.StatusBadRequest, gin.H{
							"status":  false,
							"message": err.Error(),
						})
						return
					}

					orderTimes = append(orderTimes, orderTime)

				}

				orderDateAndTime.Times = orderTimes

			}

		}

		orderDateAndTimes = append(orderDateAndTimes, orderDateAndTime)

	}

	c.JSON(http.StatusOK, gin.H{
		"status":      true,
		"order_times": orderDateAndTimes,
	})

}
