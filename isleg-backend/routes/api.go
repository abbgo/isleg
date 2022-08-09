package routes

import (
	"github/abbgo/isleg/isleg-backend/auth"
	backController "github/abbgo/isleg/isleg-backend/controllers/back"
	frontController "github/abbgo/isleg/isleg-backend/controllers/front"
	"github/abbgo/isleg/isleg-backend/middlewares"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Routes() *gin.Engine {

	routes := gin.Default()

	// cors
	// routes.Use(cors.Default())

	routes.Use(cors.New(cors.Config{
		// AllowOrigins:     []string{"https://foo.com"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "RefershToken", "Authorization"},
		AllowCredentials: true,
		AllowAllOrigins:  true,
		MaxAge:           12 * time.Hour,
	}))

	// routes belong to admin panel
	back := routes.Group("/admin")
	{

		back.POST("/language", backController.CreateLanguage)
		back.PUT("/language/:id", backController.UpdateLanguage)
		back.GET("/language/:id", backController.GetOneLanguage)
		back.GET("/languages", backController.GetAllLanguage)
		back.DELETE("/language/:id", backController.DeleteLanguage)
		back.GET("/restore-language/:id", backController.RestoreLanguage)
		back.DELETE("/delete-language/:id", backController.DeletePermanentlyLanguage)

		back.POST("/company-setting", backController.CreateCompanySetting)
		back.PUT("/company-setting", backController.UpdateCompanySetting)
		back.GET("/company-setting", backController.GetOneCompanySetting)

		back.POST("/translation-header", backController.CreateTranslationHeader)
		back.PUT("/translation-header/:id", backController.UpdateTranslationHeader)
		back.GET("/translation-header/:id", backController.GetOneTranslationHeader)

		back.POST("/translation-footer", backController.CreateTranslationFooter)
		back.PUT("/translation-footer/:id", backController.UpdateTranslationFooter)
		back.GET("/translation-footer/:id", backController.GetOneTranslationFooter)

		back.POST("/translation-secure", backController.CreateTranslationSecure)
		back.PUT("/translation-secure/:id", backController.UpdateTranslationSecure)
		back.GET("/translation-secure/:id", backController.GetOneTranslationSecure)

		back.POST("/translation-payment", backController.CreateTranslationPayment)
		back.PUT("/translation-payment/:id", backController.UpdateTranslationPayment)
		back.GET("/translation-payment/:id", backController.GetOneTranslationPayment)

		back.POST("/translation-about", backController.CreateTranslationAbout)
		back.PUT("/translation-about/:id", backController.UpdateTranslationAbout)
		back.GET("/translation-about/:id", backController.GetOneTranslationAbout)

		back.POST("/translation-contact", backController.CreateTranslationContact)
		back.PUT("/translation-contact/:id", backController.UpdateTranslationContact)
		back.GET("/translation-contact/:id", backController.GetOneTranslationContact)

		back.POST("/translation-my-information-page", backController.CreateTranslationMyInformationPage)
		back.PUT("/translation-my-information-page/:id", backController.UpdateTranslationMyInformationPage)
		back.GET("/translation-my-information-page/:id", backController.GetOneTranslationMyInformationPage)

		back.POST("/translation-update-password-page", backController.CreateTranslationUpdatePasswordPage)
		back.PUT("/translation-update-password-page/:id", backController.UpdateTranslationUpdatePasswordPage)
		back.GET("/translation-update-password-page/:id", backController.GetOneTranslationUpdatePasswordPage)

		back.POST("/category", backController.CreateCategory)
		back.PUT("/category/:id", backController.UpdateCategory)
		back.GET("/category/:id", backController.GetOneCategory)
		back.GET("/categories", backController.GetAllCategory)
		back.DELETE("/category/:id", backController.DeleteCategory)
		back.GET("/restore-category/:id", backController.RestoreCategory)
		back.DELETE("/delete-category/:id", backController.DeletePermanentlyCategory)

		back.POST("/brend", backController.CreateBrend)
		back.PUT("/brend/:id", backController.UpdateBrend)
		back.GET("/brend/:id", backController.GetBrend)
		back.GET("/brend", backController.GetBrends)
		back.DELETE("/brend/:id", backController.DeleteBrend)
		back.GET("/restore-brend/:id", backController.RestoreBrend)
		back.DELETE("/delete-brend/:id", backController.DeletePermanentlyBrend)

		back.POST("/product", backController.CreateProduct)

		back.POST("/company-phone", backController.CreateCompanyPhone)

		back.POST("/company-address", backController.CreateCompanyAddress)

		back.POST("/afisa", backController.CreateAfisa)

		back.POST("/district", backController.CreateDistrict)

		back.POST("/shop", backController.CreateShop)

	}

	// customer routes
	customer := routes.Group("/api/auth")
	{
		customer.POST("/register", frontController.RegisterCustomer)
		customer.POST("/login", frontController.LoginCustomer)
		customer.POST("/refresh", auth.Refresh)
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

		// get update password page translation
		front.GET("/translation-update-password-page", backController.GetTranslationUpdatePasswordPage)

		// homepage categories
		front.GET("/homepage-categories", frontController.GetHomePageCategories)

		// get one category with products
		front.GET("/:category_id/:limit/:page", backController.GetOneCategoryWithProducts)

		securedCustomer := front.Group("/").Use(middlewares.Auth())
		{
			// get my information page translation
			securedCustomer.GET("/translation-my-information-page", backController.GetTranslationMyInformationPage)

			// get all favourite products of customer
			securedCustomer.POST("/like", frontController.AddLike)
		}

	}

	return routes

}
