package routes

import (
	"github/abbgo/isleg/isleg-backend/auth"

	adminController "github/abbgo/isleg/isleg-backend/controllers/admin"
	backController "github/abbgo/isleg/isleg-backend/controllers/back"
	frontController "github/abbgo/isleg/isleg-backend/controllers/front"

	"github/abbgo/isleg/isleg-backend/middlewares"
	"time"

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
		MaxAge:           12 * time.Hour,
	}))

	// routes belong to admin panel
	back := routes.Group("/admin")
	{

		admin := back.Group("/auth")
		{
			admin.POST("/register", middlewares.IsSuperAdmin(), adminController.RegisterAdmin)
			admin.POST("/login", adminController.LoginAdmin)
			admin.POST("/refresh", auth.RefreshTokenForAdmin)
		}

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

		back.POST("/order-time", backController.CreateOrderTime)
		back.PUT("/order-time", backController.UpdateOrderTime)
		back.GET("/order-time/:id", backController.GetOrderTimeByID)
		back.GET("/order-times", backController.GetOrderTimes)
		back.DELETE("/order-time/:id", backController.DeleteOrderTimeByID)
		back.GET("/restore-order-time/:id", backController.RestoreOrderTimeByID)
		back.DELETE("/delete-order-time/:id", backController.DeletePermanentlyOrderTimeByID)

		back.POST("/banner", backController.CreateBanner)
		back.PUT("/banner/:id", backController.UpdateBannerByID)
		back.GET("/banner/:id", backController.GetBannerByID)
		back.GET("/banners", backController.GetBanners)
		back.DELETE("/banner/:id", backController.DeleteBannerByID)
		back.GET("/restore-banner/:id", backController.RestoreBannerByID)
		back.DELETE("/delete-banner/:id", backController.DeletePermanentlyBannerByID)

		back.POST("/translation-header", backController.CreateTranslationHeader)
		back.PUT("/translation-header", backController.UpdateTranslationHeaderByID)
		back.GET("/translation-header/:id", backController.GetTranslationHeaderByID)

		back.POST("/translation-footer", backController.CreateTranslationFooter)
		back.PUT("/translation-footer", backController.UpdateTranslationFooterByID)
		back.GET("/translation-footer/:id", backController.GetTranslationFooterByID)

		back.POST("/translation-secure", backController.CreateTranslationSecure)
		back.PUT("/translation-secure", backController.UpdateTranslationSecureByID)
		back.GET("/translation-secure/:id", backController.GetTranslationSecureByID)

		back.POST("/translation-payment", backController.CreateTranslationPayment)
		back.PUT("/translation-payment", backController.UpdateTranslationPaymentByID)
		back.GET("/translation-payment/:id", backController.GetTranslationPaymentByID)

		back.POST("/translation-about", backController.CreateTranslationAbout)
		back.PUT("/translation-about", backController.UpdateTranslationAboutByID)
		back.GET("/translation-about/:id", backController.GetTranslationAboutByID)

		back.POST("/translation-contact", backController.CreateTranslationContact)
		back.PUT("/translation-contact", backController.UpdateTranslationContactByID)
		back.GET("/translation-contact/:id", backController.GetTranslationContactByID) // fully raedy

		back.POST("/translation-my-information-page", backController.CreateTranslationMyInformationPage)
		back.PUT("/translation-my-information-page", backController.UpdateTranslationMyInformationPageByID)
		back.GET("/translation-my-information-page/:id", backController.GetTranslationMyInformationPageByID)

		back.POST("/translation-update-password-page", backController.CreateTranslationUpdatePasswordPage)
		back.PUT("/translation-update-password-page", backController.UpdateTranslationUpdatePasswordPageByID)
		back.GET("/translation-update-password-page/:id", backController.GetTranslationUpdatePasswordPageByID)

		back.POST("/translation-basket-page", backController.CreateTranslationBasketPage)
		back.PUT("/translation-basket-page", backController.UpdateTranslationBasketPageByID)
		back.GET("/translation-basket-page/:id", backController.GetTranslationBasketPageByID)

		back.POST("/translation-order-page", backController.CreateTranslationOrderPage)
		back.PUT("/translation-order-page", backController.UpdateTranslationOrderPageByID)
		back.GET("/translation-order-page/:id", backController.GetTranslationOrderPageByID)

		back.POST("/translation-my-order-page", backController.CreateTranslationMyOrderPage)
		back.PUT("/translation-my-order-page", backController.UpdateTranslationMyOrderPageByID)
		back.GET("/translation-my-order-page/:id", backController.GetTranslationMyOrderPageByID)

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
		back.PUT("/company-phone", backController.UpdateCompanyPhoneByID)
		back.GET("/company-phone/:id", backController.GetCompanyPhoneByID)
		back.DELETE("/company-phone/:id", backController.DeleteCompanyPhoneByID)
		back.GET("/restore-company-phone/:id", backController.RestoreCompanyPhoneByID)
		back.DELETE("/delete-company-phone/:id", backController.DeletePermanentlyCompanyPhoneByID)

		back.POST("/company-address", backController.CreateCompanyAddress)
		back.PUT("/company-address", backController.UpdateCompanyAddressByID)
		back.GET("/company-address/:id", backController.GetCompanyAddressByID)

		back.POST("/payment-type", backController.CreatePaymentType)
		back.PUT("/payment-type", backController.UpdatePaymentTypeByID)
		back.GET("/payment-type/:id", backController.GetPaymentTypeByID)
		back.GET("/payment-types", backController.GetPaymentTypes)

		back.POST("/afisa", backController.CreateAfisa)
		back.PUT("/afisa/:id", backController.UpdateAfisaByID)
		back.GET("/afisa/:id", backController.GetAfisaByID)
		back.GET("/afisas", backController.GetAfisas)
		back.DELETE("/afisa/:id", backController.DeleteAfisaByID)
		back.GET("/restore-afisa/:id", backController.RestoreAfisaByID)
		back.DELETE("/delete-afisa/:id", backController.DeletePermanentlyAfisaByID)

		back.POST("/district", backController.CreateDistrict)

		back.POST("/shop", backController.CreateShop)
		back.PUT("/shop", backController.UpdateShopByID)
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

		// get my order page translation
		front.GET("/translation-my-order-page", backController.GetTranslationMyOrderPageByLangID)

		// get payment ttype by lang id
		front.GET("/payment-types", backController.GetPaymentTypesByLangID)

		// homepage categories
		front.GET("/homepage-categories", frontController.GetHomePageCategories)

		// // get one category with products
		front.GET("/:category_id/:limit/:page", backController.GetOneCategoryWithProducts)

		// get order time
		front.GET("/order-time", backController.GetOrderTime)

		// search
		front.POST("/search/:limit/:page", frontController.Search)

		// get my information page translation
		front.GET("/translation-my-information-page", backController.GetTranslationMyInformationPageByLangID)

		// to order
		front.POST("/to-order", frontController.ToOrder)

		// to order
		front.POST("/send-mail", frontController.SendMail)

		// get like products without customer by product id ->
		// Eger musderi like - a haryt gosup sonam sol haryt bazadan ayrylan bolsa
		// sony bildirmek ucin front - dan mana cookie - daki product_id - leri
		// ugdurkdyryar we men yzyna sol id - leri product - lary ugratyan
		front.POST("/likes-without-customer", frontController.GetLikedProductsWithoutCustomer)

		// get order products without customer by product id ->
		// Eger musderi sebede - e haryt gosup sonam sol haryt bazadan ayrylan bolsa
		// sony bildirmek ucin front - dan mana cookie - daki product_id - leri
		// ugdurkdyryar we men yzyna sol id - leri product - lary ugratyan
		front.POST("/orders-without-customer", frontController.GetOrderedProductsWithoutCustomer)

		securedCustomer := front.Group("/").Use(middlewares.Auth())
		{
			// add like if customer exists
			securedCustomer.POST("/like", frontController.AddLike)

			// remove like if customer exists
			// securedCustomer.POST("/like/:product_id", frontController.RemoveLike)

			// get like products if customer exists
			securedCustomer.GET("/likes", frontController.GetCustomerLikes)

			// add product to cart
			securedCustomer.POST("/add-cart", frontController.AddCart)

			// get product of cart
			securedCustomer.GET("/get-cart", frontController.GetCustomerCartProducts)

			// remove product from cart
			securedCustomer.POST("/remove-cart", frontController.RemoveCart)

			// get customer orders
			securedCustomer.GET("/orders/:limit/:page", frontController.GetCustomerOrders)

			// get customer orders
			securedCustomer.GET("/addresses", frontController.GetCustomerAddresses)

			// get customer informations
			securedCustomer.GET("/my-information", frontController.GetCustomerInformation)

			// get customer informations
			securedCustomer.PUT("/my-information", frontController.UpdateCustomerInformation)

			// update customer address status
			securedCustomer.PUT("/address", frontController.UpdateCustomerAddressStatus)

			// update customer password
			securedCustomer.PUT("/customer-password", frontController.UpdateCustomerPassword)

		}

	}

	return routes

}
