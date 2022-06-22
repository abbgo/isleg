package routes

import (
	backController "github/abbgo/isleg/isleg-backend/controllers/back"
	frontController "github/abbgo/isleg/isleg-backend/controllers/front"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Routes() *gin.Engine {

	routes := gin.Default()
	routes.Use(cors.Default())

	back := routes.Group("/admin")
	{

		back.POST("/language", backController.CreateLanguage)

		back.GET("/company-setting", backController.CreateCompanySetting)

		back.POST("/translation-header", backController.CreateTranslationHeader)

		back.POST("/translation-footer", backController.CreateTranslationFooter)

		back.POST("/category", backController.CreateCategory)

		back.POST("/brend", backController.CreateBrend)

		back.POST("/product", backController.CreateProduct)

	}

	front := routes.Group("/api/:lang")
	{
		front.GET("/header", frontController.GetHeaderData)
		front.GET("/footer", frontController.GetFooterData)
		front.GET("/brends", frontController.GetBrends)
		front.GET("/homepage-categories", frontController.GetHomePageCategories)
		front.GET("/:category_id/:limit/:page", backController.GetOneCategoryWithProducts)

	}

	return routes

}
