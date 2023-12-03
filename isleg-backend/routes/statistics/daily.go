package dailyStatisticsApi

import (
	dailyStatisticController "github/abbgo/isleg/isleg-backend/controllers/statistics"
	"github/abbgo/isleg/isleg-backend/middlewares"

	"github.com/gin-gonic/gin"
)

func DailyStatisticsRoutes(back *gin.RouterGroup) {

	dailyStatistics := back.Group("/daily")
	{
		// GetDailyStatistics - gunlik statistika ucin
		dailyStatistics.GET("", middlewares.IsSuperAdmin(), dailyStatisticController.GetDailyStatistics)

		// GetDailyStatistics - gunlik statistika ucin
		dailyStatistics.GET("count-of-customers", middlewares.IsSuperAdmin(), dailyStatisticController.GetDailyCountOfCustomers)
	}

}
