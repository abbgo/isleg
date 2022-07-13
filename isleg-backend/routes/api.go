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

		back.POST("/translation-payment", backController.CreateTranslationPayment)

		back.POST("/translation-about", backController.CreateTranslationAbout)

		back.POST("/translation-contact", backController.CreateTranslationContact)

		back.POST("/translation-my-information-page", backController.CreateTranslationMyInformationPage)

		back.POST("/translation-update-password-page", backController.CreateTranslationUpdatePasswordPage)

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
		// get header data
		front.GET("/header", frontController.GetHeaderData)

		// get footer data
		front.GET("/footer", frontController.GetFooterData)

		// get all brend
		front.GET("/brends", frontController.GetBrends)

		// get company phone numbers
		front.GET("/company-phones", backController.GetCompanyPhones)

		// get company address
		front.GET("/company-address", backController.GetCompanyAddress)

		// get Terms of Service and Privacy Policy page translation
		front.GET("/translation-secure", backController.GetTranslationSecure)

		// get Delivery and payment order page translation
		front.GET("/translation-payment", backController.GetTranslationPayment)

		// get about us page translation
		front.GET("/translation-about", backController.GetTranslationAbout)

		// get contact us page translation
		front.GET("/translation-contact", backController.GetTranslationContact)

		// get my information page translation
		front.GET("/translation-my-information-page", backController.GetTranslationMyInformationPage)

		// get update password page translation
		front.GET("/translation-update-password-page", backController.GetTranslationUpdatePasswordPage)

		// homepage categories
		front.GET("/homepage-categories", frontController.GetHomePageCategories)

		// get one category with products
		front.GET("/:category_id/:limit/:page", backController.GetOneCategoryWithProducts)

		// customer routes
		front.POST("/register", frontController.RegisterCustomer)
		front.POST("/login", frontController.LoginCustomer)

	}

	return routes

}
