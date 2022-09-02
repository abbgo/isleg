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
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "RefershToken", "Authorization"},
		AllowCredentials: true,
		AllowAllOrigins:  true,
		MaxAge:           12 * time.Hour,
	}))

	// routes belong to admin panel
	back := routes.Group("/admin")
	{
		back.POST("/language", backController.CreateLanguage)
		back.PUT("/language/:id", backController.UpdateLanguageByID)
		back.GET("/language/:id", backController.GetLanguageByID)
		back.GET("/languages", backController.GetLanguages)
		back.DELETE("/language/:id", backController.DeleteLanguageByID)
		back.GET("/restore-language/:id", backController.RestoreLanguageByID)
		back.DELETE("/delete-language/:id", backController.DeletePermanentlyLanguageByID)

		back.POST("/company-setting", backController.CreateCompanySetting)
		back.PUT("/company-setting", backController.UpdateCompanySetting)
		back.GET("/company-setting", backController.GetCompanySetting)

		back.POST("/translation-header", backController.CreateTranslationHeader)
		back.PUT("/translation-header/:id", backController.UpdateTranslationHeaderByID)
		back.GET("/translation-header/:id", backController.GetTranslationHeaderByID)

		back.POST("/translation-footer", backController.CreateTranslationFooter)
		back.PUT("/translation-footer/:id", backController.UpdateTranslationFooterByID)
		back.GET("/translation-footer/:id", backController.GetTranslationFooterByID)

		back.POST("/translation-secure", backController.CreateTranslationSecure)
		back.PUT("/translation-secure/:id", backController.UpdateTranslationSecureByID)
		back.GET("/translation-secure/:id", backController.GetTranslationSecureByID)

		back.POST("/translation-payment", backController.CreateTranslationPayment)
		back.PUT("/translation-payment/:id", backController.UpdateTranslationPaymentByID)
		back.GET("/translation-payment/:id", backController.GetTranslationPaymentByID)

		back.POST("/translation-about", backController.CreateTranslationAbout)
		back.PUT("/translation-about/:id", backController.UpdateTranslationAboutByID)
		back.GET("/translation-about/:id", backController.GetTranslationAboutByID)

		back.POST("/translation-contact", backController.CreateTranslationContact)
		back.PUT("/translation-contact/:id", backController.UpdateTranslationContactByID)
		back.GET("/translation-contact/:id", backController.GetTranslationContactByID)

		back.POST("/translation-my-information-page", backController.CreateTranslationMyInformationPage)
		back.PUT("/translation-my-information-page/:id", backController.UpdateTranslationMyInformationPageByID)
		back.GET("/translation-my-information-page/:id", backController.GetTranslationMyInformationPageByID)

		back.POST("/translation-update-password-page", backController.CreateTranslationUpdatePasswordPage)
		back.PUT("/translation-update-password-page/:id", backController.UpdateTranslationUpdatePasswordPageByID)
		back.GET("/translation-update-password-page/:id", backController.GetTranslationUpdatePasswordPageByID)

		back.POST("/translation-basket-page", backController.CreateTranslationBasketPage)
		back.PUT("/translation-basket-page/:id", backController.UpdateTranslationBasketPageByID)
		back.GET("/translation-basket-page/:id", backController.GetTranslationBasketPageByID)

		back.POST("/translation-order-page", backController.CreateTranslationOrderPage)
		back.PUT("/translation-order-page/:id", backController.UpdateTranslationOrderPageByID)
		back.GET("/translation-order-page/:id", backController.GetTranslationOrderPageByID)

		back.POST("/translation-my-order-page", backController.CreateTranslationMyOrderPage)
		// back.PUT("/translation-order-page/:id", backController.UpdateTranslationOrderPageByID)
		// back.GET("/translation-order-page/:id", backController.GetTranslationOrderPageByID)

		back.POST("/category", backController.CreateCategory)
		back.PUT("/category/:id", backController.UpdateCategoryByID)
		back.GET("/category/:id", backController.GetCategoryByID)
		back.GET("/categories", backController.GetCategories)
		back.DELETE("/category/:id", backController.DeleteCategoryByID)
		back.GET("/restore-category/:id", backController.RestoreCategoryByID)
		back.DELETE("/delete-category/:id", backController.DeletePermanentlyCategoryByID)

		back.POST("/brend", backController.CreateBrend)
		back.PUT("/brend/:id", backController.UpdateBrendByID)
		back.GET("/brend/:id", backController.GetBrendByID)
		back.GET("/brends", backController.GetBrends)
		back.DELETE("/brend/:id", backController.DeleteBrendByID)
		back.GET("/restore-brend/:id", backController.RestoreBrendByID)
		back.DELETE("/delete-brend/:id", backController.DeletePermanentlyBrendByID)

		back.POST("/product", backController.CreateProduct)
		back.PUT("/product/:id", backController.UpdateProductByID)
		back.GET("/product/:id", backController.GetProductByID)
		back.GET("/products", backController.GetProducts)
		back.DELETE("/product/:id", backController.DeleteProductByID)
		back.GET("/restore-product/:id", backController.RestoreProductByID)
		back.DELETE("/delete-product/:id", backController.DeletePermanentlyProductByID)

		back.POST("/company-phone", backController.CreateCompanyPhone)
		back.PUT("/company-phone/:id", backController.UpdateCompanyPhoneByID)
		back.GET("/company-phone/:id", backController.GetCompanyPhoneByID)
		back.DELETE("/company-phone/:id", backController.DeleteCompanyPhoneByID)
		back.GET("/restore-company-phone/:id", backController.RestoreCompanyPhoneByID)
		back.DELETE("/delete-company-phone/:id", backController.DeletePermanentlyCompanyPhoneByID)

		back.POST("/company-address", backController.CreateCompanyAddress)
		back.PUT("/company-address/:id", backController.UpdateCompanyAddressByID)
		back.GET("/company-address/:id", backController.GetCompanyAddressByID)

		back.POST("/afisa", backController.CreateAfisa)
		back.PUT("/afisa/:id", backController.UpdateAfisaByID)
		back.GET("/afisa/:id", backController.GetAfisaByID)
		back.GET("/afisas", backController.GetAfisas)
		back.DELETE("/afisa/:id", backController.DeleteAfisaByID)
		back.GET("/restore-afisa/:id", backController.RestoreAfisaByID)
		back.DELETE("/delete-afisa/:id", backController.DeletePermanentlyAfisaByID)

		back.POST("/district", backController.CreateDistrict)

		back.POST("/shop", backController.CreateShop)
		back.PUT("/shop/:id", backController.UpdateShopByID)
		back.GET("/shop/:id", backController.GetShopByID)
		back.GET("/shops", backController.GetShops)
		back.DELETE("/shop/:id", backController.DeleteShopByID)
		back.GET("/restore-shop/:id", backController.RestoreShopByID)
		back.DELETE("/delete-shop/:id", backController.DeletePermanentlyShopByID)

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
		front.GET("/translation-secure", backController.GetTranslationSecureByLangID)

		// get Delivery and payment order page translation
		front.GET("/translation-payment", backController.GetTranslationPaymentByLangID)

		// get about us page translation
		front.GET("/translation-about", backController.GetTranslationAboutByLangID)

		// get contact us page translation
		front.GET("/translation-contact", backController.GetTranslationContactByLangID)

		// get update password page translation
		front.GET("/translation-update-password-page", backController.GetTranslationUpdatePasswordPageByLangID)

		// get basket page translation
		front.GET("/translation-basket-page", backController.GetTranslationBasketPageByLangID)

		// get order page translation
		front.GET("/translation-order-page", backController.GetTranslationOrderPageByLangID)

		// homepage categories
		front.GET("/homepage-categories", frontController.GetHomePageCategories)

		// get one category with products
		front.GET("/:category_id/:limit/:page", backController.GetOneCategoryWithProducts)

		securedCustomer := front.Group("/").Use(middlewares.Auth())
		{
			// get my information page translation
			securedCustomer.GET("/translation-my-information-page", backController.GetTranslationMyInformationPageByLangID)

			// add favourite products of customer
			securedCustomer.POST("/like", frontController.AddLike)

			// get all favourite products of customer
			securedCustomer.GET("/likes/:customer_id", frontController.GetCustomerLikes)

			// remove favourite products of customer
			securedCustomer.DELETE("/like/:customer_id/:product_id", frontController.RemoveLike)

			// add product to cart
			securedCustomer.POST("/basket", frontController.AddProductToBasket)

		}

	}

	return routes

}
