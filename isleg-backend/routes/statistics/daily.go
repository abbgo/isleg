package dailyStatisticsApi

import (
	dailyStatisticController "github/abbgo/isleg/isleg-backend/controllers/statistics"
	"github/abbgo/isleg/isleg-backend/middlewares"

	"github.com/gin-gonic/gin"
)

func DailyStatisticsRoutes(back *gin.RouterGroup) {

	dailyStatistics := back.Group("/daily")
	{
		// GetDailyStatistics - gunlik edilen sowdalar barada maglumat
		dailyStatistics.GET("", middlewares.IsSuperAdmin(), dailyStatisticController.GetDailyStatistics)

		// GetDailyCountOfCustomers - gunlik musderilerin giren sany
		dailyStatistics.GET("count-of-customers", middlewares.IsSuperAdmin(), dailyStatisticController.GetDailyCountOfCustomers)

		// GetDailySearchOfCustomers - gunlik musderilerin gozleg eden harylary
		dailyStatistics.GET("search-of-customers", middlewares.IsSuperAdmin(), dailyStatisticController.GetDailySearchOfCustomers)
	}

}
