package controllers

import (
	"context"
	"github/abbgo/isleg/isleg-backend/config"
	"github/abbgo/isleg/isleg-backend/helpers"
	"net/http"
	"time"

	// backController "github/abbgo/isleg/isleg-backend/controllers/back"

	"github.com/gin-gonic/gin"
)

type DailyStatistic struct {
	OrderNumber   uint    `json:"order_number"`
	Debit         float32 `json:"debit"`
	Credit        float32 `json:"credit"`
	ShippingPrice float32 `json:"shipping_price"`
	Leftover      float32 `json:"leftover"`
}

func GetDailyStatistics(c *gin.Context) {
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	currentTime := time.Now()
	dayBegin := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 0, 0, 0, 0, time.Local)
	dayEnd := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 24, 0, 0, 0, time.Local)

	rowsOrder, err := db.Query(context.Background(), "SELECT id,order_number,total_price,shipping_price FROM orders WHERE deleted_at IS NOT NULL AND updated_at >= $1 AND updated_at < $2 ", dayBegin, dayEnd)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer rowsOrder.Close()

	var dailyStatistics []DailyStatistic
	for rowsOrder.Next() {
		var dailyStatistic DailyStatistic
		var order_id string
		if err := rowsOrder.Scan(&order_id, &dailyStatistic.OrderNumber, &dailyStatistic.Credit, &dailyStatistic.ShippingPrice); err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}

		if err := db.QueryRow(context.Background(), "SELECT DISTINCT ON (p.id) SUM(p.price*op.quantity_of_product) FROM products p INNER JOIN ordered_products op ON op.product_id = p.id WHERE op.order_id = $1 GROUP BY p.id", order_id).Scan(&dailyStatistic.Debit); err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}

		dailyStatistic.Leftover = dailyStatistic.Credit - dailyStatistic.Debit
		dailyStatistics = append(dailyStatistics, dailyStatistic)
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"daily":  dailyStatistics,
	})

}
