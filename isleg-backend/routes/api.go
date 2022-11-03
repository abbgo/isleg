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

		securedAdmin := back.Group("/").Use(middlewares.CheckAdmin())
		{
			securedAdmin.GET("/orders/:limit/:page", frontController.GetOrders)

			securedAdmin.POST("/language", backController.CreateLanguage)
			securedAdmin.PUT("/language/:id", backController.UpdateLanguageByID)
			securedAdmin.GET("/language/:id", backController.GetLanguageByID)
			securedAdmin.GET("/languages", backController.GetLanguages)
			securedAdmin.DELETE("/language/:id", backController.DeleteLanguageByID)
			securedAdmin.GET("/restore-language/:id", backController.RestoreLanguageByID)
			securedAdmin.DELETE("/delete-language/:id", backController.DeletePermanentlyLanguageByID)

			securedAdmin.POST("/company-setting", backController.CreateCompanySetting)
			securedAdmin.PUT("/company-setting", backController.UpdateCompanySetting)
			securedAdmin.GET("/company-setting", backController.GetCompanySetting)

			securedAdmin.POST("/order-time", backController.CreateOrderTime)
			securedAdmin.PUT("/order-time", backController.UpdateOrderTimeByID)
			securedAdmin.GET("/order-time/:id", backController.GetOrderTimeByID)
			securedAdmin.GET("/order-times", backController.GetOrderTimes)
			securedAdmin.DELETE("/order-time/:id", backController.DeleteOrderTimeByID)
			securedAdmin.GET("/restore-order-time/:id", backController.RestoreOrderTimeByID)
			securedAdmin.DELETE("/delete-order-time/:id", backController.DeletePermanentlyOrderTimeByID)

			securedAdmin.POST("/banner", backController.CreateBanner)
			securedAdmin.PUT("/banner/:id", backController.UpdateBannerByID)
			securedAdmin.GET("/banner/:id", backController.GetBannerByID)
			securedAdmin.GET("/banners", backController.GetBanners)
			securedAdmin.DELETE("/banner/:id", backController.DeleteBannerByID)
			securedAdmin.GET("/restore-banner/:id", backController.RestoreBannerByID)
			securedAdmin.DELETE("/delete-banner/:id", backController.DeletePermanentlyBannerByID)

			securedAdmin.POST("/translation-header", backController.CreateTranslationHeader)
			securedAdmin.PUT("/translation-header", backController.UpdateTranslationHeaderByID)
			securedAdmin.GET("/translation-header/:id", backController.GetTranslationHeaderByID)

			securedAdmin.POST("/translation-footer", backController.CreateTranslationFooter)
			securedAdmin.PUT("/translation-footer", backController.UpdateTranslationFooterByID)
			securedAdmin.GET("/translation-footer/:id", backController.GetTranslationFooterByID)

			securedAdmin.POST("/translation-secure", backController.CreateTranslationSecure)
			securedAdmin.PUT("/translation-secure", backController.UpdateTranslationSecureByID)
			securedAdmin.GET("/translation-secure/:id", backController.GetTranslationSecureByID)

			securedAdmin.POST("/translation-payment", backController.CreateTranslationPayment)
			securedAdmin.PUT("/translation-payment", backController.UpdateTranslationPaymentByID)
			securedAdmin.GET("/translation-payment/:id", backController.GetTranslationPaymentByID)

			securedAdmin.POST("/translation-about", backController.CreateTranslationAbout)
			securedAdmin.PUT("/translation-about", backController.UpdateTranslationAboutByID)
			securedAdmin.GET("/translation-about/:id", backController.GetTranslationAboutByID)

			securedAdmin.POST("/translation-contact", backController.CreateTranslationContact)
			securedAdmin.PUT("/translation-contact", backController.UpdateTranslationContactByID)
			securedAdmin.GET("/translation-contact/:id", backController.GetTranslationContactByID) // fully raedy

			securedAdmin.POST("/translation-my-information-page", backController.CreateTranslationMyInformationPage)
			securedAdmin.PUT("/translation-my-information-page", backController.UpdateTranslationMyInformationPageByID)
			securedAdmin.GET("/translation-my-information-page/:id", backController.GetTranslationMyInformationPageByID)

			securedAdmin.POST("/translation-update-password-page", backController.CreateTranslationUpdatePasswordPage)
			securedAdmin.PUT("/translation-update-password-page", backController.UpdateTranslationUpdatePasswordPageByID)
			securedAdmin.GET("/translation-update-password-page/:id", backController.GetTranslationUpdatePasswordPageByID)

			securedAdmin.POST("/translation-basket-page", backController.CreateTranslationBasketPage)
			securedAdmin.PUT("/translation-basket-page", backController.UpdateTranslationBasketPageByID)
			securedAdmin.GET("/translation-basket-page/:id", backController.GetTranslationBasketPageByID)

			securedAdmin.POST("/translation-order-page", backController.CreateTranslationOrderPage)
			securedAdmin.PUT("/translation-order-page", backController.UpdateTranslationOrderPageByID)
			securedAdmin.GET("/translation-order-page/:id", backController.GetTranslationOrderPageByID)

			securedAdmin.POST("/translation-my-order-page", backController.CreateTranslationMyOrderPage)
			securedAdmin.PUT("/translation-my-order-page", backController.UpdateTranslationMyOrderPageByID)
			securedAdmin.GET("/translation-my-order-page/:id", backController.GetTranslationMyOrderPageByID)

			securedAdmin.POST("/category", backController.CreateCategory)
			securedAdmin.PUT("/category/:id", backController.UpdateCategoryByID)
			securedAdmin.GET("/category/:id", backController.GetCategoryByID)
			securedAdmin.GET("/categories", backController.GetCategories)
			securedAdmin.DELETE("/category/:id", backController.DeleteCategoryByID)
			securedAdmin.GET("/restore-category/:id", backController.RestoreCategoryByID)
			securedAdmin.DELETE("/delete-category/:id", backController.DeletePermanentlyCategoryByID)

			securedAdmin.POST("/brend", backController.CreateBrend)
			securedAdmin.PUT("/brend/:id", backController.UpdateBrendByID)
			securedAdmin.GET("/brend/:id", backController.GetBrendByID)
			securedAdmin.GET("/brends", backController.GetBrends)
			securedAdmin.DELETE("/brend/:id", backController.DeleteBrendByID)
			securedAdmin.GET("/restore-brend/:id", backController.RestoreBrendByID)
			securedAdmin.DELETE("/delete-brend/:id", backController.DeletePermanentlyBrendByID)

			securedAdmin.POST("/product", backController.CreateProduct)
			securedAdmin.PUT("/product/:id", backController.UpdateProductByID)
			securedAdmin.GET("/product/:id", backController.GetProductByID)
			securedAdmin.GET("/products", backController.GetProducts)
			securedAdmin.DELETE("/product/:id", backController.DeleteProductByID)
			securedAdmin.GET("/restore-product/:id", backController.RestoreProductByID)
			securedAdmin.DELETE("/delete-product/:id", backController.DeletePermanentlyProductByID)

			securedAdmin.POST("/company-phone", backController.CreateCompanyPhone)
			securedAdmin.PUT("/company-phone", backController.UpdateCompanyPhoneByID)
			securedAdmin.GET("/company-phone/:id", backController.GetCompanyPhoneByID)
			securedAdmin.DELETE("/company-phone/:id", backController.DeleteCompanyPhoneByID)
			securedAdmin.GET("/restore-company-phone/:id", backController.RestoreCompanyPhoneByID)
			securedAdmin.DELETE("/delete-company-phone/:id", backController.DeletePermanentlyCompanyPhoneByID)

			securedAdmin.POST("/company-address", backController.CreateCompanyAddress)
			securedAdmin.PUT("/company-address", backController.UpdateCompanyAddressByID)
			securedAdmin.GET("/company-address/:id", backController.GetCompanyAddressByID)

			securedAdmin.POST("/payment-type", backController.CreatePaymentType)
			securedAdmin.PUT("/payment-type", backController.UpdatePaymentTypeByID)
			securedAdmin.GET("/payment-type/:id", backController.GetPaymentTypeByID)
			securedAdmin.GET("/payment-types", backController.GetPaymentTypes)

			securedAdmin.POST("/afisa", backController.CreateAfisa)
			securedAdmin.PUT("/afisa/:id", backController.UpdateAfisaByID)
			securedAdmin.GET("/afisa/:id", backController.GetAfisaByID)
			securedAdmin.GET("/afisas", backController.GetAfisas)
			securedAdmin.DELETE("/afisa/:id", backController.DeleteAfisaByID)
			securedAdmin.GET("/restore-afisa/:id", backController.RestoreAfisaByID)
			securedAdmin.DELETE("/delete-afisa/:id", backController.DeletePermanentlyAfisaByID)

			securedAdmin.POST("/district", backController.CreateDistrict)

			securedAdmin.POST("/shop", backController.CreateShop)
			securedAdmin.PUT("/shop", backController.UpdateShopByID)
			securedAdmin.GET("/shop/:id", backController.GetShopByID)
			securedAdmin.GET("/shops", backController.GetShops)
			securedAdmin.DELETE("/shop/:id", backController.DeleteShopByID)
			securedAdmin.GET("/restore-shop/:id", backController.RestoreShopByID)
			securedAdmin.DELETE("/delete-shop/:id", backController.DeletePermanentlyShopByID)

		}

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
