package controllers

import (
	"github/abbgo/isleg/isleg-backend/config"
	"github/abbgo/isleg/isleg-backend/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
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

	var orderDate models.OrderDates

	if err := c.BindJSON(&orderDate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	// validate data
	if err := models.ValidateOrderDateAndTime(orderDate.Date, orderDate.OrderTimes); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	for _, v := range orderDate.TranslationOrderDates {

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
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": "language not found",
			})
			return
		}

	}

	// add data to order_dates table and return last id
	resultOrderDates, err := db.Query("INSERT INTO order_dates (date) VALUES ($1) RETURNING id", orderDate.Date)
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

	var orderDateID string

	for resultOrderDates.Next() {
		if err := resultOrderDates.Scan(&orderDateID); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}

	// add data to order_times table
	for _, v := range orderDate.OrderTimes {

		resultOrderTimes, err := db.Query("INSERT INTO order_times (order_date_id,time) VALUES ($1,$2)", orderDateID, v.Time)
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

	}

	// add translation order date to database
	for _, v := range orderDate.TranslationOrderDates {

		resultTrOrderDates, err := db.Query("INSERT INTO translation_order_dates (lang_id,order_date_id,date) VALUES ($1,$2,$3)", v.LangID, orderDateID, v.Date)
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
		"message": "data successfully added",
	})

}

func UpdateOrderTime(c *gin.Context) {

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

	var orderDate models.OrderDates

	if err := c.BindJSON(&orderDate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	rowOrderDate, err := db.Query("SELECT id FROM order_dates WHERE id = $1 AND deleted_at IS NULL", orderDate.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer func() {
		if err := rowOrderDate.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

	var order_date_id string

	for rowOrderDate.Next() {
		if err := rowOrderDate.Scan(&order_date_id); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}

	if order_date_id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "order date not found",
		})
		return
	}

	resultOrderDate, err := db.Query("UPDATE order_dates SET date = $1 WHERE id = $2", orderDate.Date, orderDate.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer func() {
		if err := resultOrderDate.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

	resultDeleteOrderTime, err := db.Query("DELETE FROM order_times WHERE order_date_id = $1", orderDate.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer func() {
		if err := resultDeleteOrderTime.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

	for _, v := range orderDate.OrderTimes {

		resultInsertOrderTime, err := db.Query("INSERT INTO order_times (order_date_id,time) VALUES ($1,$2)", orderDate.ID, v.Time)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
		defer func() {
			if err := resultInsertOrderTime.Close(); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  false,
					"message": err.Error(),
				})
				return
			}
		}()

	}

	for _, v := range orderDate.TranslationOrderDates {

		restultTrOrderDate, err := db.Query("UPDATE translation_order_dates SET date = $1 WHERE lang_id = $2 AND order_date_id = $3", v.Date, v.LangID, orderDate.ID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
		defer func() {
			if err := restultTrOrderDate.Close(); err != nil {
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

func GetOrderTimeByID(c *gin.Context) {

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

	ID := c.Param("id")

	rowOrderDate, err := db.Query("SELECT id,date FROM order_dates WHERE id = $1 AND deleted_at IS NULL", ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer func() {
		if err := rowOrderDate.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

	var orderDate models.OrderDates

	for rowOrderDate.Next() {
		if err := rowOrderDate.Scan(&orderDate.ID, &orderDate.Date); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}

		if orderDate.ID == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": "order date not found",
			})
			return
		}

		rowsOrderTimes, err := db.Query("SELECT time FROM order_times WHERE order_date_id = $1", orderDate.ID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
		defer func() {
			if err := rowsOrderTimes.Close(); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  false,
					"message": err.Error(),
				})
				return
			}
		}()

		var orderTimes []models.OrderTimes

		for rowsOrderTimes.Next() {
			var orderTime models.OrderTimes

			if err := rowsOrderTimes.Scan(&orderTime.Time); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  false,
					"message": err.Error(),
				})
				return
			}

			orderTimes = append(orderTimes, orderTime)
		}

		orderDate.OrderTimes = orderTimes

		rowsTrOrderDate, err := db.Query("SELECT date FROM translation_order_dates WHERE order_date_id = $1", orderDate.ID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
		defer func() {
			if err := rowsTrOrderDate.Close(); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  false,
					"message": err.Error(),
				})
				return
			}
		}()

		var trsOrderDate []models.TranslationOrderDates

		for rowsTrOrderDate.Next() {
			var trOrderDate models.TranslationOrderDates

			if err := rowsTrOrderDate.Scan(&trOrderDate.Date); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  false,
					"message": err.Error(),
				})
				return
			}

			trsOrderDate = append(trsOrderDate, trOrderDate)
		}

		orderDate.TranslationOrderDates = trsOrderDate

	}

	c.JSON(http.StatusOK, gin.H{
		"status":     true,
		"order_date": orderDate,
	})

}

func GetOrderTimes(c *gin.Context) {

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

	rowsOrderDate, err := db.Query("SELECT id,date FROM order_dates WHERE deleted_at IS NULL")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer func() {
		if err := rowsOrderDate.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

	var orderDates []models.OrderDates

	for rowsOrderDate.Next() {
		var orderDate models.OrderDates

		if err := rowsOrderDate.Scan(&orderDate.ID, &orderDate.Date); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}

		rowsOrderTimes, err := db.Query("SELECT time FROM order_times WHERE order_date_id = $1", orderDate.ID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
		defer func() {
			if err := rowsOrderTimes.Close(); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  false,
					"message": err.Error(),
				})
				return
			}
		}()

		var orderTimes []models.OrderTimes

		for rowsOrderTimes.Next() {
			var orderTime models.OrderTimes

			if err := rowsOrderTimes.Scan(&orderTime.Time); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  false,
					"message": err.Error(),
				})
				return
			}

			orderTimes = append(orderTimes, orderTime)
		}

		orderDate.OrderTimes = orderTimes

		rowsTrOrderDate, err := db.Query("SELECT date FROM translation_order_dates WHERE order_date_id = $1", orderDate.ID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
		defer func() {
			if err := rowsTrOrderDate.Close(); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  false,
					"message": err.Error(),
				})
				return
			}
		}()

		var trsOrderDate []models.TranslationOrderDates

		for rowsTrOrderDate.Next() {
			var trOrderDate models.TranslationOrderDates

			if err := rowsTrOrderDate.Scan(&trOrderDate.Date); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  false,
					"message": err.Error(),
				})
				return
			}

			trsOrderDate = append(trsOrderDate, trOrderDate)
		}

		orderDate.TranslationOrderDates = trsOrderDate

		orderDates = append(orderDates, orderDate)

	}

	c.JSON(http.StatusOK, gin.H{
		"status":      true,
		"order_dates": orderDates,
	})

}

func DeleteOrderTimeByID(c *gin.Context) {

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

	ID := c.Param("id")

	rowOrderDate, err := db.Query("SELECT id FROM order_dates WHERE id = $1 AND deleted_at IS NULL", ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer func() {
		if err := rowOrderDate.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

	var order_date_id string

	for rowOrderDate.Next() {
		if err := rowOrderDate.Scan(&order_date_id); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}

	if order_date_id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "order date not found",
		})
		return
	}

	resultProc, err := db.Query("CALL delete_order_date($1)", ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer func() {
		if err := resultProc.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "data successfully deleted",
	})

}

func RestoreOrderTimeByID(c *gin.Context) {

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

	ID := c.Param("id")

	rowOrderDate, err := db.Query("SELECT id FROM order_dates WHERE id = $1 AND deleted_at IS NOT NULL", ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer func() {
		if err := rowOrderDate.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

	var order_date_id string

	for rowOrderDate.Next() {
		if err := rowOrderDate.Scan(&order_date_id); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}

	if order_date_id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "order date not found",
		})
		return
	}

	resultPROC, err := db.Query("CALL restore_order_date($1)", ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer func() {
		if err := resultPROC.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "data successfully restored",
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
	defer func() {
		if err := db.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

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

	rowsOrderDate, err := db.Query("select od.id , od.date , tod.date from order_dates od inner join translation_order_dates tod on tod.order_date_id = od.id where tod.lang_id = $1 and od.deleted_at is null and tod.deleted_at is null and od.date = any($2)", langID, pq.Array(dates))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer func() {
		if err := rowsOrderDate.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

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
			defer func() {
				if err := rowsOrderTime.Close(); err != nil {
					c.JSON(http.StatusBadRequest, gin.H{
						"status":  false,
						"message": err.Error(),
					})
					return
				}
			}()

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
				defer func() {
					if err := rowsOrderTime.Close(); err != nil {
						c.JSON(http.StatusBadRequest, gin.H{
							"status":  false,
							"message": err.Error(),
						})
						return
					}
				}()

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
				defer func() {
					if err := rowsOrderTime.Close(); err != nil {
						c.JSON(http.StatusBadRequest, gin.H{
							"status":  false,
							"message": err.Error(),
						})
						return
					}
				}()

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
