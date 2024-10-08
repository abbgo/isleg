package routes

import (
	backApi "github/abbgo/isleg/isleg-backend/routes/back"
	frontApi "github/abbgo/isleg/isleg-backend/routes/front"
	statisticApi "github/abbgo/isleg/isleg-backend/routes/statistics"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

func Routes() *gin.Engine {

	routes := gin.Default()

	// cors
	// routes.Use(cors.Default())

	routes.Use(gzip.Gzip(gzip.DefaultCompression))

	routes.Use(cors.New(cors.Config{
		// AllowOrigins:     []string{"https://foo.com"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "RefreshToken", "Authorization"},
		AllowCredentials: true,
		AllowAllOrigins:  true,
		// MaxAge:           12 * time.Hour,
	}))

	// routes belong to admin panel
	back := routes.Group("/api/admin")
	{
		backApi.AdminRoutes(back)

		// bu group - daki ahli yazylan funksiyalary dine ahli adminler isledip biler
		// CheckAdmin middleware sony kesgitleyar
		backApi.BackRoutes(back)

	}

	front := routes.Group("/api")
	{
		// customer routes
		frontApi.CustomerRoutes(front)

		// bu group - a degisli api - lerden maglumat alynanda ( :lang ) paramter boyunca uytgedilip
		// terjime alynyar
		frontApi.FrontRoutes(front)
	}

	statistics := routes.Group("/api/statistics")
	{
		// statistics routes
		statisticApi.DailyStatisticsRoutes(statistics)

	}

	return routes

}
