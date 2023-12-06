package controllers

import (
	"context"
	"github/abbgo/isleg/isleg-backend/config"
	"github/abbgo/isleg/isleg-backend/helpers"
	"github/abbgo/isleg/isleg-backend/models"
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
	var (
		dailyStatistics       []DailyStatistic
		TotalDayDebit         float64
		TotalDayCredit        float64
		TotalDayShippingPrice float64
		TotalDayLeftover      float64
	)
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

		TotalDayDebit += float64(dailyStatistic.TotalDebit)
		TotalDayCredit += float64(dailyStatistic.TotalCredit)
		TotalDayShippingPrice += float64(dailyStatistic.TotalShippingPrice)
		TotalDayLeftover += float64(dailyStatistic.TotalLeftover)

		dailyStatistics = append(dailyStatistics, dailyStatistic)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":                   true,
		"daily":                    dailyStatistics,
		"total_day_debit":          TotalDayDebit,
		"total_day_credit":         TotalDayCredit,
		"total_day_shipping_price": TotalDayShippingPrice,
		"total_day_leftover":       TotalDayLeftover,
	})

}

func GetDailyCountOfCustomers(c *gin.Context) {
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	rows, err := db.Query(context.Background(), "SELECT count,day FROM count_of_customers")
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer rows.Close()

	var counts []models.CountOfCustomer
	for rows.Next() {
		var count models.CountOfCustomer
		var day time.Time
		if err := rows.Scan(&count.Count, &day); err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}
		count.Date = day.Format("02.01.2006")
		counts = append(counts, count)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":                true,
		"daily_customers_count": counts,
	})

}

type SearchOfCustomers struct {
	Date  string   `json:"date"`
	Texts []string `json:"texts"`
}

func GetDailySearchOfCustomers(c *gin.Context) {
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	rowsGroupBy, err := db.Query(context.Background(), "SELECT day FROM search_of_customers GROUP BY day ORDER BY day ASC")
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer rowsGroupBy.Close()

	var searchs []SearchOfCustomers
	for rowsGroupBy.Next() {
		var search SearchOfCustomers
		var day time.Time
		if err := rowsGroupBy.Scan(&day); err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}

		search.Date = day.Format("02.01.2006")

		rows, err := db.Query(context.Background(), "SELECT search_text FROM search_of_customers WHERE day = $1", day.Format("2006-01-02"))
		if err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}
		defer rows.Close()

		for rows.Next() {
			var text string
			if err := rows.Scan(&text); err != nil {
				helpers.HandleError(c, 400, err.Error())
				return
			}
			search.Texts = append(search.Texts, text)
		}
		searchs = append(searchs, search)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":                 true,
		"daily_customers_search": searchs,
	})

}
