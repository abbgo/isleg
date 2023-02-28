package controllers

import (
	"github/abbgo/isleg/isleg-backend/config"
	"github/abbgo/isleg/isleg-backend/models"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// type OrderDateAndTime struct {
// 	ID              string      `json:"-"`
// 	Date            string      `json:"date"`
// 	Times           []OrderTime `json:"times"`
// 	TranslationDate string      `json:"translation_date"`
// }

// type OrderTime struct {
// 	Time string `json:"time"`
// }

type OrderDateAndTime struct {
	Date        string `json:"date"`
	Translation string `json:"translation"`
	Time        string `json:"time"`
}

type OrderTimes struct {
	Translation string `json:"translation"`
	Times       []Wagt `json:"times"`
}

type Wagt struct {
	Time string `json:"time"`
}

func CreateOrderDate(c *gin.Context) {

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
	if err := models.ValidateOrderDate(orderDate.Date); err != nil {
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
			c.JSON(http.StatusNotFound, gin.H{
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

// func UpdateOrderTimeByID(c *gin.Context) {

// 	// initialize database connection
// 	db, err := config.ConnDB()
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"status":  false,
// 			"message": err.Error(),
// 		})
// 		return
// 	}
// 	defer func() {
// 		if err := db.Close(); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{
// 				"status":  false,
// 				"message": err.Error(),
// 			})
// 			return
// 		}
// 	}()

// 	var orderDate models.OrderDates

// 	if err := c.BindJSON(&orderDate); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"status":  false,
// 			"message": err.Error(),
// 		})
// 		return
// 	}

// 	rowOrderDate, err := db.Query("SELECT id FROM order_dates WHERE id = $1 AND deleted_at IS NULL", orderDate.ID)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"status":  false,
// 			"message": err.Error(),
// 		})
// 		return
// 	}
// 	defer func() {
// 		if err := rowOrderDate.Close(); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{
// 				"status":  false,
// 				"message": err.Error(),
// 			})
// 			return
// 		}
// 	}()

// 	var order_date_id string

// 	for rowOrderDate.Next() {
// 		if err := rowOrderDate.Scan(&order_date_id); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{
// 				"status":  false,
// 				"message": err.Error(),
// 			})
// 			return
// 		}
// 	}

// 	if order_date_id == "" {
// 		c.JSON(http.StatusNotFound, gin.H{
// 			"status":  false,
// 			"message": "order date not found",
// 		})
// 		return
// 	}

// 	resultOrderDate, err := db.Query("UPDATE order_dates SET date = $1 WHERE id = $2", orderDate.Date, orderDate.ID)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"status":  false,
// 			"message": err.Error(),
// 		})
// 		return
// 	}
// 	defer func() {
// 		if err := resultOrderDate.Close(); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{
// 				"status":  false,
// 				"message": err.Error(),
// 			})
// 			return
// 		}
// 	}()

// 	resultDeleteOrderTime, err := db.Query("DELETE FROM order_times WHERE order_date_id = $1", orderDate.ID)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"status":  false,
// 			"message": err.Error(),
// 		})
// 		return
// 	}
// 	defer func() {
// 		if err := resultDeleteOrderTime.Close(); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{
// 				"status":  false,
// 				"message": err.Error(),
// 			})
// 			return
// 		}
// 	}()

// 	for _, v := range orderDate.OrderTimes {

// 		resultInsertOrderTime, err := db.Query("INSERT INTO order_times (order_date_id,time) VALUES ($1,$2)", orderDate.ID, v.Time)
// 		if err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{
// 				"status":  false,
// 				"message": err.Error(),
// 			})
// 			return
// 		}
// 		defer func() {
// 			if err := resultInsertOrderTime.Close(); err != nil {
// 				c.JSON(http.StatusBadRequest, gin.H{
// 					"status":  false,
// 					"message": err.Error(),
// 				})
// 				return
// 			}
// 		}()

// 	}

// 	for _, v := range orderDate.TranslationOrderDates {

// 		restultTrOrderDate, err := db.Query("UPDATE translation_order_dates SET date = $1 WHERE lang_id = $2 AND order_date_id = $3", v.Date, v.LangID, orderDate.ID)
// 		if err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{
// 				"status":  false,
// 				"message": err.Error(),
// 			})
// 			return
// 		}
// 		defer func() {
// 			if err := restultTrOrderDate.Close(); err != nil {
// 				c.JSON(http.StatusBadRequest, gin.H{
// 					"status":  false,
// 					"message": err.Error(),
// 				})
// 				return
// 			}
// 		}()

// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"status":  true,
// 		"message": "data successfully updated",
// 	})

// }

// func GetOrderTimeByID(c *gin.Context) {

// 	// initialize database connection
// 	db, err := config.ConnDB()
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"status":  false,
// 			"message": err.Error(),
// 		})
// 		return
// 	}
// 	defer func() {
// 		if err := db.Close(); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{
// 				"status":  false,
// 				"message": err.Error(),
// 			})
// 			return
// 		}
// 	}()

// 	ID := c.Param("id")

// 	rowOrderDate, err := db.Query("SELECT id,date FROM order_dates WHERE id = $1 AND deleted_at IS NULL", ID)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"status":  false,
// 			"message": err.Error(),
// 		})
// 		return
// 	}
// 	defer func() {
// 		if err := rowOrderDate.Close(); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{
// 				"status":  false,
// 				"message": err.Error(),
// 			})
// 			return
// 		}
// 	}()

// 	var orderDate models.OrderDates

// 	for rowOrderDate.Next() {
// 		if err := rowOrderDate.Scan(&orderDate.ID, &orderDate.Date); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{
// 				"status":  false,
// 				"message": err.Error(),
// 			})
// 			return
// 		}

// 		if orderDate.ID == "" {
// 			c.JSON(http.StatusNotFound, gin.H{
// 				"status":  false,
// 				"message": "order date not found",
// 			})
// 			return
// 		}

// 		rowsOrderTimes, err := db.Query("SELECT time FROM order_times WHERE order_date_id = $1", orderDate.ID)
// 		if err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{
// 				"status":  false,
// 				"message": err.Error(),
// 			})
// 			return
// 		}
// 		defer func() {
// 			if err := rowsOrderTimes.Close(); err != nil {
// 				c.JSON(http.StatusBadRequest, gin.H{
// 					"status":  false,
// 					"message": err.Error(),
// 				})
// 				return
// 			}
// 		}()

// 		var orderTimes []models.OrderTimes

// 		for rowsOrderTimes.Next() {
// 			var orderTime models.OrderTimes

// 			if err := rowsOrderTimes.Scan(&orderTime.Time); err != nil {
// 				c.JSON(http.StatusBadRequest, gin.H{
// 					"status":  false,
// 					"message": err.Error(),
// 				})
// 				return
// 			}

// 			orderTimes = append(orderTimes, orderTime)
// 		}

// 		orderDate.OrderTimes = orderTimes

// 		rowsTrOrderDate, err := db.Query("SELECT date FROM translation_order_dates WHERE order_date_id = $1", orderDate.ID)
// 		if err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{
// 				"status":  false,
// 				"message": err.Error(),
// 			})
// 			return
// 		}
// 		defer func() {
// 			if err := rowsTrOrderDate.Close(); err != nil {
// 				c.JSON(http.StatusBadRequest, gin.H{
// 					"status":  false,
// 					"message": err.Error(),
// 				})
// 				return
// 			}
// 		}()

// 		var trsOrderDate []models.TranslationOrderDates

// 		for rowsTrOrderDate.Next() {
// 			var trOrderDate models.TranslationOrderDates

// 			if err := rowsTrOrderDate.Scan(&trOrderDate.Date); err != nil {
// 				c.JSON(http.StatusBadRequest, gin.H{
// 					"status":  false,
// 					"message": err.Error(),
// 				})
// 				return
// 			}

// 			trsOrderDate = append(trsOrderDate, trOrderDate)
// 		}

// 		orderDate.TranslationOrderDates = trsOrderDate

// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"status":     true,
// 		"order_date": orderDate,
// 	})

// }

// func GetOrderTimes(c *gin.Context) {

// 	// initialize database connection
// 	db, err := config.ConnDB()
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"status":  false,
// 			"message": err.Error(),
// 		})
// 		return
// 	}
// 	defer func() {
// 		if err := db.Close(); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{
// 				"status":  false,
// 				"message": err.Error(),
// 			})
// 			return
// 		}
// 	}()

// 	rowsOrderDate, err := db.Query("SELECT id,date FROM order_dates WHERE deleted_at IS NULL")
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"status":  false,
// 			"message": err.Error(),
// 		})
// 		return
// 	}
// 	defer func() {
// 		if err := rowsOrderDate.Close(); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{
// 				"status":  false,
// 				"message": err.Error(),
// 			})
// 			return
// 		}
// 	}()

// 	var orderDates []models.OrderDates

// 	for rowsOrderDate.Next() {
// 		var orderDate models.OrderDates

// 		if err := rowsOrderDate.Scan(&orderDate.ID, &orderDate.Date); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{
// 				"status":  false,
// 				"message": err.Error(),
// 			})
// 			return
// 		}

// 		rowsOrderTimes, err := db.Query("SELECT time FROM order_times WHERE order_date_id = $1", orderDate.ID)
// 		if err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{
// 				"status":  false,
// 				"message": err.Error(),
// 			})
// 			return
// 		}
// 		defer func() {
// 			if err := rowsOrderTimes.Close(); err != nil {
// 				c.JSON(http.StatusBadRequest, gin.H{
// 					"status":  false,
// 					"message": err.Error(),
// 				})
// 				return
// 			}
// 		}()

// 		var orderTimes []models.OrderTimes

// 		for rowsOrderTimes.Next() {
// 			var orderTime models.OrderTimes

// 			if err := rowsOrderTimes.Scan(&orderTime.Time); err != nil {
// 				c.JSON(http.StatusBadRequest, gin.H{
// 					"status":  false,
// 					"message": err.Error(),
// 				})
// 				return
// 			}

// 			orderTimes = append(orderTimes, orderTime)
// 		}

// 		orderDate.OrderTimes = orderTimes

// 		rowsTrOrderDate, err := db.Query("SELECT date FROM translation_order_dates WHERE order_date_id = $1", orderDate.ID)
// 		if err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{
// 				"status":  false,
// 				"message": err.Error(),
// 			})
// 			return
// 		}
// 		defer func() {
// 			if err := rowsTrOrderDate.Close(); err != nil {
// 				c.JSON(http.StatusBadRequest, gin.H{
// 					"status":  false,
// 					"message": err.Error(),
// 				})
// 				return
// 			}
// 		}()

// 		var trsOrderDate []models.TranslationOrderDates

// 		for rowsTrOrderDate.Next() {
// 			var trOrderDate models.TranslationOrderDates

// 			if err := rowsTrOrderDate.Scan(&trOrderDate.Date); err != nil {
// 				c.JSON(http.StatusBadRequest, gin.H{
// 					"status":  false,
// 					"message": err.Error(),
// 				})
// 				return
// 			}

// 			trsOrderDate = append(trsOrderDate, trOrderDate)
// 		}

// 		orderDate.TranslationOrderDates = trsOrderDate

// 		orderDates = append(orderDates, orderDate)

// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"status":      true,
// 		"order_dates": orderDates,
// 	})

// }

// func DeleteOrderTimeByID(c *gin.Context) {

// 	// initialize database connection
// 	db, err := config.ConnDB()
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"status":  false,
// 			"message": err.Error(),
// 		})
// 		return
// 	}
// 	defer func() {
// 		if err := db.Close(); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{
// 				"status":  false,
// 				"message": err.Error(),
// 			})
// 			return
// 		}
// 	}()

// 	ID := c.Param("id")

// 	rowOrderDate, err := db.Query("SELECT id FROM order_dates WHERE id = $1 AND deleted_at IS NULL", ID)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"status":  false,
// 			"message": err.Error(),
// 		})
// 		return
// 	}
// 	defer func() {
// 		if err := rowOrderDate.Close(); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{
// 				"status":  false,
// 				"message": err.Error(),
// 			})
// 			return
// 		}
// 	}()

// 	var order_date_id string

// 	for rowOrderDate.Next() {
// 		if err := rowOrderDate.Scan(&order_date_id); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{
// 				"status":  false,
// 				"message": err.Error(),
// 			})
// 			return
// 		}
// 	}

// 	if order_date_id == "" {
// 		c.JSON(http.StatusNotFound, gin.H{
// 			"status":  false,
// 			"message": "order date not found",
// 		})
// 		return
// 	}

// 	resultProc, err := db.Query("CALL delete_order_date($1)", ID)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"status":  false,
// 			"message": err.Error(),
// 		})
// 		return
// 	}
// 	defer func() {
// 		if err := resultProc.Close(); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{
// 				"status":  false,
// 				"message": err.Error(),
// 			})
// 			return
// 		}
// 	}()

// 	c.JSON(http.StatusOK, gin.H{
// 		"status":  true,
// 		"message": "data successfully deleted",
// 	})

// }

// func RestoreOrderTimeByID(c *gin.Context) {

// 	// initialize database connection
// 	db, err := config.ConnDB()
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"status":  false,
// 			"message": err.Error(),
// 		})
// 		return
// 	}
// 	defer func() {
// 		if err := db.Close(); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{
// 				"status":  false,
// 				"message": err.Error(),
// 			})
// 			return
// 		}
// 	}()

// 	ID := c.Param("id")

// 	rowOrderDate, err := db.Query("SELECT id FROM order_dates WHERE id = $1 AND deleted_at IS NOT NULL", ID)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"status":  false,
// 			"message": err.Error(),
// 		})
// 		return
// 	}
// 	defer func() {
// 		if err := rowOrderDate.Close(); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{
// 				"status":  false,
// 				"message": err.Error(),
// 			})
// 			return
// 		}
// 	}()

// 	var order_date_id string

// 	for rowOrderDate.Next() {
// 		if err := rowOrderDate.Scan(&order_date_id); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{
// 				"status":  false,
// 				"message": err.Error(),
// 			})
// 			return
// 		}
// 	}

// 	if order_date_id == "" {
// 		c.JSON(http.StatusNotFound, gin.H{
// 			"status":  false,
// 			"message": "order date not found",
// 		})
// 		return
// 	}

// 	resultPROC, err := db.Query("CALL restore_order_date($1)", ID)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"status":  false,
// 			"message": err.Error(),
// 		})
// 		return
// 	}
// 	defer func() {
// 		if err := resultPROC.Close(); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{
// 				"status":  false,
// 				"message": err.Error(),
// 			})
// 			return
// 		}
// 	}()

// 	c.JSON(http.StatusOK, gin.H{
// 		"status":  true,
// 		"message": "data successfully restored",
// 	})

// }

// func DeletePermanentlyOrderTimeByID(c *gin.Context) {

// 	// initialize database connection
// 	db, err := config.ConnDB()
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"status":  false,
// 			"message": err.Error(),
// 		})
// 		return
// 	}
// 	defer func() {
// 		if err := db.Close(); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{
// 				"status":  false,
// 				"message": err.Error(),
// 			})
// 			return
// 		}
// 	}()

// 	ID := c.Param("id")

// 	rowOrderDate, err := db.Query("SELECT id FROM order_dates WHERE id = $1 AND deleted_at IS NOT NULL", ID)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"status":  false,
// 			"message": err.Error(),
// 		})
// 		return
// 	}
// 	defer func() {
// 		if err := rowOrderDate.Close(); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{
// 				"status":  false,
// 				"message": err.Error(),
// 			})
// 			return
// 		}
// 	}()

// 	var order_date_id string

// 	for rowOrderDate.Next() {
// 		if err := rowOrderDate.Scan(&order_date_id); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{
// 				"status":  false,
// 				"message": err.Error(),
// 			})
// 			return
// 		}
// 	}

// 	if order_date_id == "" {
// 		c.JSON(http.StatusNotFound, gin.H{
// 			"status":  false,
// 			"message": "order date not found",
// 		})
// 		return
// 	}

// 	resultOrderDate, err := db.Query("DELETE FROM order_dates WHERE id = $1", ID)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"status":  false,
// 			"message": err.Error(),
// 		})
// 		return
// 	}
// 	defer func() {
// 		if err := resultOrderDate.Close(); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{
// 				"status":  false,
// 				"message": err.Error(),
// 			})
// 			return
// 		}
// 	}()

// 	c.JSON(http.StatusOK, gin.H{
// 		"status":  true,
// 		"message": "data successfully deleted",
// 	})

// }

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

	currentHour := 19

	rowsOrderDate, err := db.Query("select distinct on (ot.time) od.date, tod.date , ot.time from order_dates od inner join translation_order_dates tod on tod.order_date_id = od.id inner join date_hours dh on dh.date_id = od.id inner join date_hour_times dht on dht.date_hour_id = dh.id inner join order_times ot on ot.id = dht.time_id where ot.deleted_at is null and dht.deleted_at is null and dh.deleted_at is null and tod.lang_id = $1 and od.deleted_at is null and tod.deleted_at is null and dh.hour = $2", langID, currentHour)
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

	var orderDates []OrderDateAndTime

	for rowsOrderDate.Next() {
		var orderDate OrderDateAndTime

		if err := rowsOrderDate.Scan(&orderDate.Date, &orderDate.Translation, &orderDate.Time); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
		orderDates = append(orderDates, orderDate)
	}

	var orderTimes []OrderTimes
	var orderTime OrderTimes
	var todays, tomorrows []Wagt
	var trToday, trTomorrow string

	for _, v := range orderDates {
		var today, tomorrow Wagt
		if v.Date == "today" {
			today.Time = v.Time
			todays = append(todays, today)
			trToday = v.Translation
			continue
		}
		if v.Date == "tomorrow" {
			tomorrow.Time = v.Time

			tomorrows = append(tomorrows, tomorrow)
			trTomorrow = v.Translation
		}
	}

	lenTodayArr := len(todays)
	for i := 0; i < lenTodayArr; i++ {
		for j := i; j < lenTodayArr; j++ {
			iInt, _ := strconv.Atoi(strings.Split(strings.Split(todays[i].Time, " - ")[0], ":")[0])
			jInt, _ := strconv.Atoi(strings.Split(strings.Split(todays[j].Time, " - ")[0], ":")[0])
			var str string
			if iInt > jInt {
				str = todays[i].Time
				todays[i].Time = todays[j].Time
				todays[j].Time = str
			}
		}
	}

	lenTomorrowArr := len(tomorrows)
	for i := 0; i < lenTomorrowArr; i++ {
		for j := i; j < lenTomorrowArr; j++ {
			iInt, _ := strconv.Atoi(strings.Split(strings.Split(tomorrows[i].Time, " - ")[0], ":")[0])
			jInt, _ := strconv.Atoi(strings.Split(strings.Split(tomorrows[j].Time, " - ")[0], ":")[0])
			var str string
			if iInt > jInt {
				str = tomorrows[i].Time
				tomorrows[i].Time = tomorrows[j].Time
				tomorrows[j].Time = str
			}
		}
	}

	orderTime.Translation = trToday
	orderTime.Times = todays
	orderTimes = append(orderTimes, orderTime)

	orderTime.Translation = trTomorrow
	orderTime.Times = tomorrows
	orderTimes = append(orderTimes, orderTime)

	c.JSON(http.StatusOK, gin.H{
		"status":      true,
		"order_times": orderTimes,
	})

}
