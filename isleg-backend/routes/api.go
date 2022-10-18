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
		back.POST("/language", backController.CreateLanguage)                             //+
		back.PUT("/language/:id", backController.UpdateLanguageByID)                      //+
		back.GET("/language/:id", backController.GetLanguageByID)                         //+
		back.GET("/languages", backController.GetLanguages)                               //+
		back.DELETE("/language/:id", backController.DeleteLanguageByID)                   //+
		back.GET("/restore-language/:id", backController.RestoreLanguageByID)             //+
		back.DELETE("/delete-language/:id", backController.DeletePermanentlyLanguageByID) //+

		back.POST("/company-setting", backController.CreateCompanySetting) //+
		back.PUT("/company-setting", backController.UpdateCompanySetting)  //+
		back.GET("/company-setting", backController.GetCompanySetting)     //+

		back.POST("/order-time", backController.CreateOrderTime) // funksiyany optimize etmeli

		back.POST("/translation-header", backController.CreateTranslationHeader)        //+
		back.PUT("/translation-header/:id", backController.UpdateTranslationHeaderByID) //+
		back.GET("/translation-header/:id", backController.GetTranslationHeaderByID)    //+

		back.POST("/translation-footer", backController.CreateTranslationFooter)        //+
		back.PUT("/translation-footer/:id", backController.UpdateTranslationFooterByID) //+
		back.GET("/translation-footer/:id", backController.GetTranslationFooterByID)    //+

		back.POST("/translation-secure", backController.CreateTranslationSecure)        //+
		back.PUT("/translation-secure/:id", backController.UpdateTranslationSecureByID) //+
		back.GET("/translation-secure/:id", backController.GetTranslationSecureByID)    //+

		back.POST("/translation-payment", backController.CreateTranslationPayment)        //+
		back.PUT("/translation-payment/:id", backController.UpdateTranslationPaymentByID) //+
		back.GET("/translation-payment/:id", backController.GetTranslationPaymentByID)    //+

		back.POST("/translation-about", backController.CreateTranslationAbout)        //+
		back.PUT("/translation-about/:id", backController.UpdateTranslationAboutByID) //+
		back.GET("/translation-about/:id", backController.GetTranslationAboutByID)    //+

		back.POST("/translation-contact", backController.CreateTranslationContact)        //+
		back.PUT("/translation-contact/:id", backController.UpdateTranslationContactByID) //+
		back.GET("/translation-contact/:id", backController.GetTranslationContactByID)    //+

		back.POST("/translation-my-information-page", backController.CreateTranslationMyInformationPage)        //+
		back.PUT("/translation-my-information-page/:id", backController.UpdateTranslationMyInformationPageByID) //+
		back.GET("/translation-my-information-page/:id", backController.GetTranslationMyInformationPageByID)    //+

		back.POST("/translation-update-password-page", backController.CreateTranslationUpdatePasswordPage)        //+
		back.PUT("/translation-update-password-page/:id", backController.UpdateTranslationUpdatePasswordPageByID) //+
		back.GET("/translation-update-password-page/:id", backController.GetTranslationUpdatePasswordPageByID)    //+

		back.POST("/translation-basket-page", backController.CreateTranslationBasketPage)        //+
		back.PUT("/translation-basket-page/:id", backController.UpdateTranslationBasketPageByID) //+
		back.GET("/translation-basket-page/:id", backController.GetTranslationBasketPageByID)    //+

		back.POST("/translation-order-page", backController.CreateTranslationOrderPage)        //+
		back.PUT("/translation-order-page/:id", backController.UpdateTranslationOrderPageByID) //+
		back.GET("/translation-order-page/:id", backController.GetTranslationOrderPageByID)    //+

		back.POST("/translation-my-order-page", backController.CreateTranslationMyOrderPage)        //+
		back.PUT("/translation-my-order-page/:id", backController.UpdateTranslationMyOrderPageByID) //+
		back.GET("/translation-my-order-page/:id", backController.GetTranslationMyOrderPageByID)    //+

		back.POST("/category", backController.CreateCategory)                             //+
		back.PUT("/category/:id", backController.UpdateCategoryByID)                      //+
		back.GET("/category/:id", backController.GetCategoryByID)                         //+
		back.GET("/categories", backController.GetCategories)                             //+
		back.DELETE("/category/:id", backController.DeleteCategoryByID)                   // funksiyany gowy optimize etmeli
		back.GET("/restore-category/:id", backController.RestoreCategoryByID)             // funksiyany gowy optimize etmeli
		back.DELETE("/delete-category/:id", backController.DeletePermanentlyCategoryByID) // funksiyany gowy optimize etmeli

		back.POST("/brend", backController.CreateBrend)                             //+
		back.PUT("/brend/:id", backController.UpdateBrendByID)                      //+
		back.GET("/brend/:id", backController.GetBrendByID)                         //+
		back.GET("/brends", backController.GetBrends)                               //+
		back.DELETE("/brend/:id", backController.DeleteBrendByID)                   // funksiyany optimize etmeli
		back.GET("/restore-brend/:id", backController.RestoreBrendByID)             // funksiyany optimize etmeli
		back.DELETE("/delete-brend/:id", backController.DeletePermanentlyBrendByID) // funksiyany optimize etmeli

		back.POST("/product", backController.CreateProduct)                             // funksiyany optimize etmeli
		back.PUT("/product/:id", backController.UpdateProductByID)                      // funksiyany optimize etmeli
		back.GET("/product/:id", backController.GetProductByID)                         // funksiyany optimize etmeli
		back.GET("/products", backController.GetProducts)                               // funksiyany optimize etmeli
		back.DELETE("/product/:id", backController.DeleteProductByID)                   // funksiyany optimize etmeli
		back.GET("/restore-product/:id", backController.RestoreProductByID)             // funksiyany optimize etmeli
		back.DELETE("/delete-product/:id", backController.DeletePermanentlyProductByID) // funksiyany optimie temeli

		back.POST("/company-phone", backController.CreateCompanyPhone)                             //+
		back.PUT("/company-phone/:id", backController.UpdateCompanyPhoneByID)                      //+
		back.GET("/company-phone/:id", backController.GetCompanyPhoneByID)                         //+
		back.DELETE("/company-phone/:id", backController.DeleteCompanyPhoneByID)                   //+
		back.GET("/restore-company-phone/:id", backController.RestoreCompanyPhoneByID)             //+
		back.DELETE("/delete-company-phone/:id", backController.DeletePermanentlyCompanyPhoneByID) //+

		back.POST("/company-address", backController.CreateCompanyAddress)        //+
		back.PUT("/company-address/:id", backController.UpdateCompanyAddressByID) //+
		back.GET("/company-address/:id", backController.GetCompanyAddressByID)    //+

		back.POST("/payment-type", backController.CreatePaymentType)        //+
		back.PUT("/payment-type/:id", backController.UpdatePaymentTypeByID) //+
		back.GET("/payment-type/:id", backController.GetPaymentTypeByID)    //+
		back.GET("/payment-types", backController.GetPaymentTypes)          //+

		back.POST("/afisa", backController.CreateAfisa)                             // funksiyany optimize etmeli
		back.PUT("/afisa/:id", backController.UpdateAfisaByID)                      // funksiyany optimize etmeli
		back.GET("/afisa/:id", backController.GetAfisaByID)                         // funksiyany optimize etmeli
		back.GET("/afisas", backController.GetAfisas)                               // funksiyany optimize etmeli
		back.DELETE("/afisa/:id", backController.DeleteAfisaByID)                   // funksiyany optimize etmeli
		back.GET("/restore-afisa/:id", backController.RestoreAfisaByID)             // funksiyany optimize etmeli
		back.DELETE("/delete-afisa/:id", backController.DeletePermanentlyAfisaByID) // funksiyany optimize etmeli

		back.POST("/district", backController.CreateDistrict) // funksiyany optimize etmeli

		back.POST("/shop", backController.CreateShop)                             // funksiyany optimize etmeli
		back.PUT("/shop/:id", backController.UpdateShopByID)                      // funksiyany optimize etmeli
		back.GET("/shop/:id", backController.GetShopByID)                         // funksiyany optimize etmeli
		back.GET("/shops", backController.GetShops)                               // funksiyany optimize etmeli
		back.DELETE("/shop/:id", backController.DeleteShopByID)                   // funksiyany optimize etmeli
		back.GET("/restore-shop/:id", backController.RestoreShopByID)             // funksiyany optimize etmeli
		back.DELETE("/delete-shop/:id", backController.DeletePermanentlyShopByID) // funksiyany optimize etmeli

	}

	// customer routes
	customer := routes.Group("/api/auth")
	{
		customer.POST("/register", frontController.RegisterCustomer) // funksiyany optimize etmeli
		customer.POST("/login", frontController.LoginCustomer)       // funksiyany optimize etmeli
		customer.POST("/refresh", auth.Refresh)                      //+
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
		front.GET("/homepage-categories", frontController.GetHomePageCategories) // funksiyany optimize etmeli

		// // get one category with products
		front.GET("/:category_id/:limit/:page", backController.GetOneCategoryWithProducts) // funksiyany gowy optimize etmeli

		// get order time
		front.GET("/order-time", backController.GetOrderTime) // funksiyany optimize etmeli

		// search
		front.POST("/search", frontController.Search) // funksiyany optimize etmeli

		// get my information page translation
		front.GET("/translation-my-information-page", backController.GetTranslationMyInformationPageByLangID)

		securedCustomer := front.Group("/").Use(middlewares.Auth())
		{
			// add like if customer exists
			securedCustomer.POST("/like", frontController.AddLike) // funksiyany optimize etmeli

			// remove like if customer exists
			securedCustomer.DELETE("/like/:customer_id/:product_id", frontController.RemoveLike) // funksiyany optimize etmeli

			// get like products if customer exists
			// securedCustomer.GET("/likes/:customer_id", frontController.GetLikes)

			// get like products without customer by product id
			securedCustomer.GET("/likes-without-customer", frontController.GetLikedProductsWithoutCustomer) // funksiyany optimize etmeli

			// add product to cart
			securedCustomer.POST("/add-cart", frontController.AddCart)

			// get product of cart
			// securedCustomer.GET("/get-cart/:customer_id", frontController.GetCartProducts)

			// remove product from cart
			securedCustomer.DELETE("/remove-cart", frontController.RemoveCart)

			// to order
			securedCustomer.POST("/to-order", frontController.ToOrder)

			// get customer orders
			securedCustomer.GET("/orders/:customer_id", frontController.GetCustomerOrders)

			// get customer orders
			securedCustomer.GET("/addresses/:customer_id", frontController.GetCustomerAddresses)

			// get customer informations
			securedCustomer.GET("/my-information/:customer_id", frontController.GetCustomerInformation)

			// update customer address status
			securedCustomer.PUT("/address", frontController.UpdateCustomerAddressStatus)

			// update customer password
			securedCustomer.PUT("/customer-password/:customer_id", frontController.UpdateCustomerPassword)

		}

	}

	return routes

}
