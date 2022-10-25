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
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "RefreshToken", "Authorization"},
		AllowCredentials: true,
		AllowAllOrigins:  true,
		MaxAge:           12 * time.Hour,
	}))

	// routes belong to admin panel
	back := routes.Group("/admin")
	{
		back.POST("/language", backController.CreateLanguage)                             // fully ready
		back.PUT("/language/:id", backController.UpdateLanguageByID)                      // fully ready
		back.GET("/language/:id", backController.GetLanguageByID)                         // fully ready
		back.GET("/languages", backController.GetLanguages)                               // fully ready
		back.DELETE("/language/:id", backController.DeleteLanguageByID)                   // fully ready
		back.GET("/restore-language/:id", backController.RestoreLanguageByID)             // fully ready
		back.DELETE("/delete-language/:id", backController.DeletePermanentlyLanguageByID) // fully ready

		back.POST("/company-setting", backController.CreateCompanySetting) // fully ready
		back.PUT("/company-setting", backController.UpdateCompanySetting)  // fully ready
		back.GET("/company-setting", backController.GetCompanySetting)     // fully ready

		back.POST("/order-time", backController.CreateOrderTime) // fully ready

		back.POST("/translation-header", backController.CreateTranslationHeader)        // fully ready
		back.PUT("/translation-header/:id", backController.UpdateTranslationHeaderByID) // fully reade
		back.GET("/translation-header/:id", backController.GetTranslationHeaderByID)    // fully ready

		back.POST("/translation-footer", backController.CreateTranslationFooter)        // fully ready
		back.PUT("/translation-footer/:id", backController.UpdateTranslationFooterByID) // fully ready
		back.GET("/translation-footer/:id", backController.GetTranslationFooterByID)    // fully ready

		back.POST("/translation-secure", backController.CreateTranslationSecure)        // fully ready
		back.PUT("/translation-secure/:id", backController.UpdateTranslationSecureByID) // fully ready
		back.GET("/translation-secure/:id", backController.GetTranslationSecureByID)    // fully ready

		back.POST("/translation-payment", backController.CreateTranslationPayment)        // fully ready
		back.PUT("/translation-payment/:id", backController.UpdateTranslationPaymentByID) // fully ready
		back.GET("/translation-payment/:id", backController.GetTranslationPaymentByID)    // fully ready

		back.POST("/translation-about", backController.CreateTranslationAbout)        // fully ready
		back.PUT("/translation-about/:id", backController.UpdateTranslationAboutByID) // fully ready
		back.GET("/translation-about/:id", backController.GetTranslationAboutByID)    // fully ready

		back.POST("/translation-contact", backController.CreateTranslationContact)        // fully ready
		back.PUT("/translation-contact/:id", backController.UpdateTranslationContactByID) // fully ready
		back.GET("/translation-contact/:id", backController.GetTranslationContactByID)    // fully raedy

		back.POST("/translation-my-information-page", backController.CreateTranslationMyInformationPage)        // fully ready
		back.PUT("/translation-my-information-page/:id", backController.UpdateTranslationMyInformationPageByID) // fully ready
		back.GET("/translation-my-information-page/:id", backController.GetTranslationMyInformationPageByID)    // fully ready

		back.POST("/translation-update-password-page", backController.CreateTranslationUpdatePasswordPage)        // fully ready
		back.PUT("/translation-update-password-page/:id", backController.UpdateTranslationUpdatePasswordPageByID) // fully ready
		back.GET("/translation-update-password-page/:id", backController.GetTranslationUpdatePasswordPageByID)    // fully ready

		back.POST("/translation-basket-page", backController.CreateTranslationBasketPage)        // fully ready
		back.PUT("/translation-basket-page/:id", backController.UpdateTranslationBasketPageByID) // fully ready
		back.GET("/translation-basket-page/:id", backController.GetTranslationBasketPageByID)    // fully ready

		back.POST("/translation-order-page", backController.CreateTranslationOrderPage)        // fully ready
		back.PUT("/translation-order-page/:id", backController.UpdateTranslationOrderPageByID) // fully ready
		back.GET("/translation-order-page/:id", backController.GetTranslationOrderPageByID)    // fully ready

		back.POST("/translation-my-order-page", backController.CreateTranslationMyOrderPage)        // fully ready
		back.PUT("/translation-my-order-page/:id", backController.UpdateTranslationMyOrderPageByID) // fully ready
		back.GET("/translation-my-order-page/:id", backController.GetTranslationMyOrderPageByID)    // fully ready

		back.POST("/category", backController.CreateCategory)                             // fully ready
		back.PUT("/category/:id", backController.UpdateCategoryByID)                      // fully ready
		back.GET("/category/:id", backController.GetCategoryByID)                         // fully ready
		back.GET("/categories", backController.GetCategories)                             // fully ready
		back.DELETE("/category/:id", backController.DeleteCategoryByID)                   // fully ready
		back.GET("/restore-category/:id", backController.RestoreCategoryByID)             // fully ready
		back.DELETE("/delete-category/:id", backController.DeletePermanentlyCategoryByID) // fully ready

		back.POST("/brend", backController.CreateBrend)                             // fully ready
		back.PUT("/brend/:id", backController.UpdateBrendByID)                      // fully ready
		back.GET("/brend/:id", backController.GetBrendByID)                         // fully ready
		back.GET("/brends", backController.GetBrends)                               // fully ready
		back.DELETE("/brend/:id", backController.DeleteBrendByID)                   // fully ready
		back.GET("/restore-brend/:id", backController.RestoreBrendByID)             // fully ready
		back.DELETE("/delete-brend/:id", backController.DeletePermanentlyBrendByID) // fully ready

		back.POST("/product", backController.CreateProduct)                             // fully ready
		back.PUT("/product/:id", backController.UpdateProductByID)                      // fully ready
		back.GET("/product/:id", backController.GetProductByID)                         // fully ready
		back.GET("/products", backController.GetProducts)                               // fully ready
		back.DELETE("/product/:id", backController.DeleteProductByID)                   // fully ready
		back.GET("/restore-product/:id", backController.RestoreProductByID)             // fully ready
		back.DELETE("/delete-product/:id", backController.DeletePermanentlyProductByID) // fully ready

		back.POST("/company-phone", backController.CreateCompanyPhone)                             // fully ready
		back.PUT("/company-phone/:id", backController.UpdateCompanyPhoneByID)                      // fully ready
		back.GET("/company-phone/:id", backController.GetCompanyPhoneByID)                         // fully ready
		back.DELETE("/company-phone/:id", backController.DeleteCompanyPhoneByID)                   // fully ready
		back.GET("/restore-company-phone/:id", backController.RestoreCompanyPhoneByID)             // fully ready
		back.DELETE("/delete-company-phone/:id", backController.DeletePermanentlyCompanyPhoneByID) // fully ready

		back.POST("/company-address", backController.CreateCompanyAddress)        // fully ready
		back.PUT("/company-address/:id", backController.UpdateCompanyAddressByID) // fully ready
		back.GET("/company-address/:id", backController.GetCompanyAddressByID)    // fully ready

		back.POST("/payment-type", backController.CreatePaymentType)        // fully ready
		back.PUT("/payment-type/:id", backController.UpdatePaymentTypeByID) // fully ready
		back.GET("/payment-type/:id", backController.GetPaymentTypeByID)    // fully ready
		back.GET("/payment-types", backController.GetPaymentTypes)          // fully ready

		back.POST("/afisa", backController.CreateAfisa)                             // fully ready
		back.PUT("/afisa/:id", backController.UpdateAfisaByID)                      // fully ready
		back.GET("/afisa/:id", backController.GetAfisaByID)                         // fully ready
		back.GET("/afisas", backController.GetAfisas)                               // fully ready
		back.DELETE("/afisa/:id", backController.DeleteAfisaByID)                   // fully ready
		back.GET("/restore-afisa/:id", backController.RestoreAfisaByID)             // fully ready
		back.DELETE("/delete-afisa/:id", backController.DeletePermanentlyAfisaByID) // fully ready

		back.POST("/district", backController.CreateDistrict) // fully ready

		back.POST("/shop", backController.CreateShop)                             // fully ready
		back.PUT("/shop/:id", backController.UpdateShopByID)                      // fully ready
		back.GET("/shop/:id", backController.GetShopByID)                         // fully ready
		back.GET("/shops", backController.GetShops)                               // fully ready
		back.DELETE("/shop/:id", backController.DeleteShopByID)                   // fully ready
		back.GET("/restore-shop/:id", backController.RestoreShopByID)             // fully ready
		back.DELETE("/delete-shop/:id", backController.DeletePermanentlyShopByID) // fully ready

	}

	// customer routes
	customer := routes.Group("/api/auth")
	{
		customer.POST("/register", frontController.RegisterCustomer) // fully ready
		customer.POST("/login", frontController.LoginCustomer)       // fully ready
		customer.POST("/refresh", auth.Refresh)                      // fully ready
	}

	// routes belong to front
	front := routes.Group("/api/:lang")
	{
		// get header data
		front.GET("/header", frontController.GetHeaderData) // fully ready

		// get footer data
		front.GET("/footer", frontController.GetFooterData) // fully ready

		// get all brend
		front.GET("/brends", frontController.GetBrends) // fully ready

		// get company phone numbers
		front.GET("/company-phones", backController.GetCompanyPhones) // fully ready

		// get company address
		front.GET("/company-address", backController.GetCompanyAddress) // fully ready

		// get Terms of Service and Privacy Policy page translation
		front.GET("/translation-secure", backController.GetTranslationSecureByLangID) // fully ready

		// get Delivery and payment order page translation
		front.GET("/translation-payment", backController.GetTranslationPaymentByLangID) // fully ready

		// get about us page translation
		front.GET("/translation-about", backController.GetTranslationAboutByLangID) // fully ready

		// get contact us page translation
		front.GET("/translation-contact", backController.GetTranslationContactByLangID) // fully ready

		// get update password page translation
		front.GET("/translation-update-password-page", backController.GetTranslationUpdatePasswordPageByLangID) // fully ready

		// get basket page translation
		front.GET("/translation-basket-page", backController.GetTranslationBasketPageByLangID) // fully ready

		// get order page translation
		front.GET("/translation-order-page", backController.GetTranslationOrderPageByLangID) // fully ready

		// get my order page translation
		front.GET("/translation-my-order-page", backController.GetTranslationMyOrderPageByLangID) // fully ready

		// get payment ttype by lang id
		front.GET("/payment-types", backController.GetPaymentTypesByLangID) // fully ready

		// homepage categories
		front.GET("/homepage-categories", frontController.GetHomePageCategories) // fully ready

		// // get one category with products
		front.GET("/:category_id/:limit/:page", backController.GetOneCategoryWithProducts) // funksiyany gowy optimize etmeli

		// get order time
		front.GET("/order-time", backController.GetOrderTime) // funksiyany optimize etmeli

		// search
		front.POST("/search", frontController.Search) // funksiyany optimize etmeli

		// get my information page translation
		front.GET("/translation-my-information-page", backController.GetTranslationMyInformationPageByLangID)

		// to order
		front.POST("/to-order", frontController.ToOrder) // fully ready

		// to order
		front.POST("/send-mail", frontController.SendMail) // fully ready

		securedCustomer := front.Group("/").Use(middlewares.Auth())
		{
			// add like if customer exists
			securedCustomer.POST("/like", frontController.AddLike) // funksiyany optimize etmeli

			// remove like if customer exists
			securedCustomer.DELETE("/like/:product_id", frontController.RemoveLike) // funksiyany optimize etmeli

			// get like products if customer exists
			// securedCustomer.GET("/likes/:customer_id", frontController.GetLikes)

			// get like products without customer by product id
			// securedCustomer.GET("/likes-without-customer", frontController.GetLikedProductsWithoutCustomer) // funksiyany optimize etmeli

			// add product to cart
			securedCustomer.POST("/add-cart", frontController.AddCart) // funksiyany optimize etmeli

			// get product of cart
			// securedCustomer.GET("/get-cart/:customer_id", frontController.GetCartProducts)

			// remove product from cart
			securedCustomer.DELETE("/remove-cart", frontController.RemoveCart) // funksiyany optimize etmeli

			// get customer orders
			securedCustomer.GET("/orders", frontController.GetCustomerOrders) // funksiyany optimize etmeli

			// get customer orders
			securedCustomer.GET("/addresses", frontController.GetCustomerAddresses) // funksiyany optimize etmeli

			// get customer informations
			securedCustomer.GET("/my-information", frontController.GetCustomerInformation) // funksiyany optimize etmeli

			// get customer informations
			securedCustomer.PUT("/my-information", frontController.UpdateCustomerInformation) // funksiyany optimize etmeli

			// update customer address status
			securedCustomer.PUT("/address", frontController.UpdateCustomerAddressStatus) // funksiyany optimize etmeli

			// update customer password
			securedCustomer.PUT("/customer-password", frontController.UpdateCustomerPassword) //+

		}

	}

	return routes

}
