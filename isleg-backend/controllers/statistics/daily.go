package controllers

import (
	"context"
	"github/abbgo/isleg/isleg-backend/config"
	"github/abbgo/isleg/isleg-backend/helpers"
	"net/http"
	"time"

	backController "github/abbgo/isleg/isleg-backend/controllers/back"

	"github.com/gin-gonic/gin"
)

type DailyStatistic struct {
	PaymentType        string      `json:"payment_type"`
	TotalDebit         float32     `json:"total_debit"`
	TotalCredit        float32     `json:"total_credit"`
	TotalShippingPrice float32     `json:"total_shipping_price"`
	TotalLeftover      float32     `json:"total_leftover"`
	Statistics         []Statistic `json:"statistics"`
}

type Statistic struct {
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

	langID, err := backController.GetLangID("tm")
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	currentTime := time.Now()
	dayBegin := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 0, 0, 0, 0, time.Local)
	dayEnd := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 24, 0, 0, 0, time.Local)

	rowsPaymentTypes, err := db.Query(context.Background(), "SELECT name,value FROM payment_types WHERE deleted_at IS NULL AND lang_id = $1 ORDER BY value ASC", langID)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer rowsPaymentTypes.Close()

	var dailyStatistics []DailyStatistic
	for rowsPaymentTypes.Next() {
		var value uint8
		var dailyStatistic DailyStatistic
		if err := rowsPaymentTypes.Scan(&dailyStatistic.PaymentType, &value); err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}

		rowsOrder, err := db.Query(context.Background(), "SELECT id,order_number,total_price,shipping_price FROM orders WHERE deleted_at IS NOT NULL AND updated_at >= $1 AND updated_at < $2 AND payment_type = $3", dayBegin, dayEnd, value)
		if err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}
		defer rowsOrder.Close()

		var statistics []Statistic
		for rowsOrder.Next() {
			var order_id string
			var statistic Statistic
			if err := rowsOrder.Scan(&order_id, &statistic.OrderNumber, &statistic.Credit, &statistic.ShippingPrice); err != nil {
				helpers.HandleError(c, 400, err.Error())
				return
			}

			if err := db.QueryRow(context.Background(), "SELECT DISTINCT ON (p.id) SUM(p.price*op.quantity_of_product) FROM products p INNER JOIN ordered_products op ON op.product_id = p.id WHERE op.order_id = $1 GROUP BY p.id", order_id).Scan(&statistic.Debit); err != nil {
				helpers.HandleError(c, 400, err.Error())
				return
			}
			statistic.Leftover = statistic.Credit - statistic.Debit
			statistics = append(statistics, statistic)
			dailyStatistic.Statistics = statistics
			dailyStatistic.TotalCredit += statistic.Credit
			dailyStatistic.TotalDebit += statistic.Debit
			dailyStatistic.TotalLeftover += statistic.Leftover
			dailyStatistic.TotalShippingPrice += statistic.ShippingPrice
		}
		dailyStatistics = append(dailyStatistics, dailyStatistic)
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"daily":  dailyStatistics,
	})

}
