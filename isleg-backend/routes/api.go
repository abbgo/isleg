package routes

import (
	backController "github/abbgo/isleg/isleg-backend/controllers/back"
	frontController "github/abbgo/isleg/isleg-backend/controllers/front"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Routes() *gin.Engine {

	routes := gin.Default()

	// cors
	routes.Use(cors.Default())

	// routes belong to admin panel
	back := routes.Group("/admin")
	{

		back.POST("/language", backController.CreateLanguage)

		back.GET("/company-setting", backController.CreateCompanySetting)

		back.POST("/translation-header", backController.CreateTranslationHeader)

		back.POST("/translation-footer", backController.CreateTranslationFooter)

		back.POST("/translation-secure", backController.CreateTranslationSecure)

		back.POST("/category", backController.CreateCategory)

		back.POST("/brend", backController.CreateBrend)

		back.POST("/product", backController.CreateProduct)

		back.POST("/company-phone", backController.CreateCompanyPhone)

		back.POST("/company-address", backController.CreateCompanyAddress)

		back.POST("/afisa", backController.CreateAfisa)

		back.POST("/district", backController.CreateDistrict)

	}

	// routes belong to front
	front := routes.Group("/api/:lang")
	{
		front.GET("/header", frontController.GetHeaderData)
		front.GET("/footer", frontController.GetFooterData)
		front.GET("/brends", frontController.GetBrends)
		front.GET("/company-phones", backController.GetCompanyPhones)
		front.GET("/company-address", backController.GetCompanyAddress)
		front.GET("/translation-secure", backController.GetTranslationSecure)
		front.GET("/homepage-categories", frontController.GetHomePageCategories)
		front.GET("/:category_id/:limit/:page", backController.GetOneCategoryWithProducts)
	}

	return routes

}
