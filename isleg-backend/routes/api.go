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
			// RegisterAdmin admin registrasiya etmek ucin ulanylyar.
			// Admini dine super admin registrasiya edip bilyar. Admin admin registrasiya edip bilenok
			admin.POST("/register", middlewares.IsSuperAdmin(), adminController.RegisterAdmin)

			// Adminlerin maglumatlaryny uytgetmek ucin ulanylyar. Adminlerin maglumatlaryny dine super admin
			// uytgedip bilyar. Admin hic bit adminin maglumatlaryny uytgedip bilenok
			admin.PUT("/information-of-admin", middlewares.IsSuperAdmin(), adminController.UpdateAdminInformation)

			// Adminlerin parollaryny uytgetmek ucin ulanylyar. Islendik adminin parolyny dine
			// super admin uytgedip bilyar. Admin hic bir adminin parolyny uytgedip bilenok
			admin.PUT("/password-of-admin/:id", middlewares.IsSuperAdmin(), adminController.UpdateAdminPassword)

			// LoginAdmin funksiya admin login bolmak ucin ulanylyar.
			admin.POST("/login", adminController.LoginAdmin)

			// Adminlerin access tokenin tazelelap refresh bilen access tokeni bile bermek
			// ucin ulanylyar
			admin.POST("/refresh", auth.RefreshTokenForAdmin)
		}

		// bu group - daki ahli yazylan funksiyalary dine ahli adminler isledip biler
		// CheckAdmin middleware sony kesgitleyar
		securedAdmin := back.Group("/").Use(middlewares.CheckAdmin())
		{

			// GetAdmins funksiya hemme adminlerin spisoygyny almak ucin ulanylyar.
			admin.GET("/admins/:limit/:page", adminController.GetAdmins)

			// GetOrders funksiya musderilerin sargyt edilen harytlaryny almak ucin ulanylyar
			// yagny musderiler tarapyndan sargyt edilen harytlary adminkada gormek ucin
			securedAdmin.GET("/orders/:limit/:page", frontController.GetOrders)

			// OrderConfirmation funksiya musderiler tarapyndan edilen sargydy
			// tassyklamak ucin ulanylyar
			securedAdmin.POST("/order-confirmation/:id", frontController.OrderConfirmation)

			// ReturnOrder funksiya musderi eden sargydyny tassyklamasa edilen sargydy
			// yzyna gaytarmak ucin ulanylyar
			securedAdmin.GET("/return-order/:id", frontController.ReturnOrder)

			// Baslayar --- dil ucin CRUD
			securedAdmin.POST("/language", backController.CreateLanguage)
			securedAdmin.PUT("/language/:id", backController.UpdateLanguageByID)
			securedAdmin.GET("/language/:id", backController.GetLanguageByID)
			securedAdmin.GET("/languages", backController.GetLanguages)
			securedAdmin.DELETE("/language/:id", backController.DeleteLanguageByID)
			securedAdmin.GET("/restore-language/:id", backController.RestoreLanguageByID)
			securedAdmin.DELETE("/delete-language/:id", backController.DeletePermanentlyLanguageByID)
			// Gutaryar --- dil ucin CRUD

			// Baslayar --- firmanyn sazlamalary ucin ( meselem logo , favicon , email addres we s,m ) ucin CRUD
			securedAdmin.POST("/company-setting", backController.CreateCompanySetting)
			securedAdmin.PUT("/company-setting", backController.UpdateCompanySetting)
			securedAdmin.GET("/company-setting", backController.GetCompanySetting)
			// Gutaryar --- firmanyn sazlamalary ucin CRUD

			// Baslayar --- Musderilerin harydy sargyt edip biljek wagtlary ucin CRUD
			securedAdmin.POST("/order-time", backController.CreateOrderTime)
			securedAdmin.PUT("/order-time", backController.UpdateOrderTimeByID)
			securedAdmin.GET("/order-time/:id", backController.GetOrderTimeByID)
			securedAdmin.GET("/order-times", backController.GetOrderTimes)
			securedAdmin.DELETE("/order-time/:id", backController.DeleteOrderTimeByID)
			securedAdmin.GET("/restore-order-time/:id", backController.RestoreOrderTimeByID)
			securedAdmin.DELETE("/delete-order-time/:id", backController.DeletePermanentlyOrderTimeByID)
			// GUtaryar --- Musderilerin harydy sargyt edip biljek wagtlary ucin CRUD

			// Baslayar --- Bas sahypada duryan banner ucin CRUD
			securedAdmin.POST("/banner", backController.CreateBanner)
			securedAdmin.PUT("/banner/:id", backController.UpdateBannerByID)
			securedAdmin.GET("/banner/:id", backController.GetBannerByID)
			securedAdmin.GET("/banners", backController.GetBanners)
			securedAdmin.DELETE("/banner/:id", backController.DeleteBannerByID)
			securedAdmin.GET("/restore-banner/:id", backController.RestoreBannerByID)
			securedAdmin.DELETE("/delete-banner/:id", backController.DeletePermanentlyBannerByID)
			// Gutaryar --- Bas sahypada duryan banner ucin CRUD

			// Baslayar --- header - in terjimesi ucin CRUD
			securedAdmin.POST("/translation-header", backController.CreateTranslationHeader)
			securedAdmin.PUT("/translation-header", backController.UpdateTranslationHeaderByID)
			securedAdmin.GET("/translation-header/:id", backController.GetTranslationHeaderByID)
			// Gutaryar --- header - in terjimesi ucin CRUD

			// Baslayar --- footer - in terjimesi ucin CRUD
			securedAdmin.POST("/translation-footer", backController.CreateTranslationFooter)
			securedAdmin.PUT("/translation-footer", backController.UpdateTranslationFooterByID)
			securedAdmin.GET("/translation-footer/:id", backController.GetTranslationFooterByID)
			// Gutaryar --- footer - in terjimesi ucin CRUD

			// Baslayar --- Ulanys duzgunleri we gizlinlik sertnamasy sahypasynyn terjimesi ucin CRUD
			securedAdmin.POST("/translation-secure", backController.CreateTranslationSecure)
			securedAdmin.PUT("/translation-secure", backController.UpdateTranslationSecureByID)
			securedAdmin.GET("/translation-secure/:id", backController.GetTranslationSecureByID)
			// Gutaryar --- Ulanys duzgunleri we gizlinlik sertnamasy sahypasynyn terjimesi ucin CRUD

			// Baslayar --- eltip bermek we toleg tertibi sahypasynyn terjimesi ucin CRUD
			securedAdmin.POST("/translation-payment", backController.CreateTranslationPayment)
			securedAdmin.PUT("/translation-payment", backController.UpdateTranslationPaymentByID)
			securedAdmin.GET("/translation-payment/:id", backController.GetTranslationPaymentByID)
			// Gutaryar --- eltip bermek we toleg tertibi sahypasynyn terjimesi ucin CRUD

			// Baslayar --- biz barada sahypasynyn terjimesi ucin CRUD
			securedAdmin.POST("/translation-about", backController.CreateTranslationAbout)
			securedAdmin.PUT("/translation-about", backController.UpdateTranslationAboutByID)
			securedAdmin.GET("/translation-about/:id", backController.GetTranslationAboutByID)
			// Gutaryar --- biz barada sahypasynyn terjimesi ucin CRUD

			// Baslayar --- aragatnasyk ( habaralasmak ) sahypasynyn terjimesi ucin CRUD
			securedAdmin.POST("/translation-contact", backController.CreateTranslationContact)
			securedAdmin.PUT("/translation-contact", backController.UpdateTranslationContactByID)
			securedAdmin.GET("/translation-contact/:id", backController.GetTranslationContactByID)
			// Gutaryar --- aragatnasyk ( habaralasmak ) sahypasynyn terjimesi ucin CRUD

			// Baslayar --- musderinin maglumatlarym sahypasynyn terjimesi ucin CRUD
			securedAdmin.POST("/translation-my-information-page", backController.CreateTranslationMyInformationPage)
			securedAdmin.PUT("/translation-my-information-page", backController.UpdateTranslationMyInformationPageByID)
			securedAdmin.GET("/translation-my-information-page/:id", backController.GetTranslationMyInformationPageByID)
			// Gutaryar --- musderinin maglumatlarym sahypasynyn terjimesi ucin CRUD

			// Baslayar --- musderinin acar sozuni uytget sahypasynyn terjimesi ucin CRUD
			securedAdmin.POST("/translation-update-password-page", backController.CreateTranslationUpdatePasswordPage)
			securedAdmin.PUT("/translation-update-password-page", backController.UpdateTranslationUpdatePasswordPageByID)
			securedAdmin.GET("/translation-update-password-page/:id", backController.GetTranslationUpdatePasswordPageByID)
			// Gutaryar --- musderinin acar sozuni uytget sahypasynyn terjimesi ucin CRUD

			// Baslayar --- sebet sahypasynyn terjimesi ucin CRUD
			securedAdmin.POST("/translation-basket-page", backController.CreateTranslationBasketPage)
			securedAdmin.PUT("/translation-basket-page", backController.UpdateTranslationBasketPageByID)
			securedAdmin.GET("/translation-basket-page/:id", backController.GetTranslationBasketPageByID)
			// Gutaryar --- sebet sahypasynyn terjimesi ucin CRUD

			// Baslayar --- haryt sargyt edilyan sahypasynyn terjimesi ucin CRUD
			securedAdmin.POST("/translation-order-page", backController.CreateTranslationOrderPage)
			securedAdmin.PUT("/translation-order-page", backController.UpdateTranslationOrderPageByID)
			securedAdmin.GET("/translation-order-page/:id", backController.GetTranslationOrderPageByID)
			// Gutaryar --- haryt sargyt edilyan sahypasynyn terjimesi ucin CRUD

			// Baslayar --- musderinin sargytlaryny goryan sahypasynyn terjimesi ucin CRUD
			securedAdmin.POST("/translation-my-order-page", backController.CreateTranslationMyOrderPage)
			securedAdmin.PUT("/translation-my-order-page", backController.UpdateTranslationMyOrderPageByID)
			securedAdmin.GET("/translation-my-order-page/:id", backController.GetTranslationMyOrderPageByID)
			// Gutaryar --- musderinin sargytlaryny goryan sahypasynyn terjimesi ucin CRUD

			// Baslayar --- kategoriya ucin CRUD
			securedAdmin.POST("/category", backController.CreateCategory)
			securedAdmin.PUT("/category/:id", backController.UpdateCategoryByID)
			securedAdmin.GET("/category/:id", backController.GetCategoryByID)
			securedAdmin.GET("/categories", backController.GetCategories)
			securedAdmin.DELETE("/category/:id", backController.DeleteCategoryByID)
			securedAdmin.GET("/restore-category/:id", backController.RestoreCategoryByID)
			securedAdmin.DELETE("/delete-category/:id", backController.DeletePermanentlyCategoryByID)
			// Gutaryar --- kategoriya ucin CRUD

			// Baslayar --- brend ucin CRUD
			securedAdmin.POST("/brend", backController.CreateBrend)
			securedAdmin.PUT("/brend/:id", backController.UpdateBrendByID)
			securedAdmin.GET("/brend/:id", backController.GetBrendByID)
			securedAdmin.GET("/brends", backController.GetBrends)
			securedAdmin.DELETE("/brend/:id", backController.DeleteBrendByID)
			securedAdmin.GET("/restore-brend/:id", backController.RestoreBrendByID)
			securedAdmin.DELETE("/delete-brend/:id", backController.DeletePermanentlyBrendByID)
			// Gutaryar --- brend ucin CRUD

			// Baslayar --- haryt ucin CRUD
			securedAdmin.POST("/product", backController.CreateProduct)
			securedAdmin.PUT("/product/:id", backController.UpdateProductByID)
			securedAdmin.GET("/product/:id", backController.GetProductByID)
			securedAdmin.GET("/products", backController.GetProducts)
			securedAdmin.DELETE("/product/:id", backController.DeleteProductByID)
			securedAdmin.GET("/restore-product/:id", backController.RestoreProductByID)
			securedAdmin.DELETE("/delete-product/:id", backController.DeletePermanentlyProductByID)
			// Gutaryar --- haryt ucin CRUD

			// Baslayar --- firmanyn telefon belgisi ucin CRUD
			securedAdmin.POST("/company-phone", backController.CreateCompanyPhone)
			securedAdmin.PUT("/company-phone", backController.UpdateCompanyPhoneByID)
			securedAdmin.GET("/company-phone/:id", backController.GetCompanyPhoneByID)
			securedAdmin.DELETE("/company-phone/:id", backController.DeleteCompanyPhoneByID)
			securedAdmin.GET("/restore-company-phone/:id", backController.RestoreCompanyPhoneByID)
			securedAdmin.DELETE("/delete-company-phone/:id", backController.DeletePermanentlyCompanyPhoneByID)
			// Gutaryar --- firmanyn telefon belgisi ucin CRUD

			// Baslayar --- firmanyn adresi we onun terjimesi ucin CRUD
			securedAdmin.POST("/company-address", backController.CreateCompanyAddress)
			securedAdmin.PUT("/company-address", backController.UpdateCompanyAddressByID)
			securedAdmin.GET("/company-address/:id", backController.GetCompanyAddressByID)
			// GUtaryar --- firmanyn adresi we onun terjimesi ucin CRUD

			// Baslayar --- musderi sargyt edende tolegin gornusleri ucin CRUD
			securedAdmin.POST("/payment-type", backController.CreatePaymentType)
			securedAdmin.PUT("/payment-type", backController.UpdatePaymentTypeByID)
			securedAdmin.GET("/payment-type/:id", backController.GetPaymentTypeByID)
			securedAdmin.GET("/payment-types", backController.GetPaymentTypes)
			// Gutaryar --- musderi sargyt edende tolegin gornusleri ucin CRUD

			// Baslayar --- afisa ucin CRUD
			securedAdmin.POST("/afisa", backController.CreateAfisa)
			securedAdmin.PUT("/afisa/:id", backController.UpdateAfisaByID)
			securedAdmin.GET("/afisa/:id", backController.GetAfisaByID)
			securedAdmin.GET("/afisas", backController.GetAfisas)
			securedAdmin.DELETE("/afisa/:id", backController.DeleteAfisaByID)
			securedAdmin.GET("/restore-afisa/:id", backController.RestoreAfisaByID)
			securedAdmin.DELETE("/delete-afisa/:id", backController.DeletePermanentlyAfisaByID)
			// Gutaryar --- afisa ucin CRUD

			// Baslayar --- duyduryslar ( notification ) we olaryn terjimeleri ucin CRUD
			securedAdmin.POST("/notification", backController.CreateNotification)
			securedAdmin.PUT("/notification", backController.UpdateNotificationByID)
			securedAdmin.GET("/notification/:id", backController.GetNotificationByID)
			securedAdmin.GET("/notifications", backController.GetNotifications)
			securedAdmin.DELETE("/notification/:id", backController.DeleteNotificationByID)
			securedAdmin.GET("/restore-notification/:id", backController.RestoreNotificationByID)
			securedAdmin.DELETE("/delete-notification/:id", backController.DeletePermanentlyNotificationByID)
			// Gutaryar --- duyduryslar ( notification ) we olaryn terjimeleri ucin CRUD

			// eger sargydyn bahasy durli rayonlara gora uytgeyan bolsa yazylmaly funksiya
			// securedAdmin.POST("/district", backController.CreateDistrict)

			// Baslayar --- haryt alynjak magazynlar ucin CRUD
			securedAdmin.POST("/shop", backController.CreateShop)
			securedAdmin.PUT("/shop", backController.UpdateShopByID)
			securedAdmin.GET("/shop/:id", backController.GetShopByID)
			securedAdmin.GET("/shops", backController.GetShops)
			securedAdmin.DELETE("/shop/:id", backController.DeleteShopByID)
			securedAdmin.GET("/restore-shop/:id", backController.RestoreShopByID)
			securedAdmin.DELETE("/delete-shop/:id", backController.DeletePermanentlyShopByID)
			// GUtaryar --- haryt alynjak magazynlar ucin CRUD

		}

	}

	// customer routes
	customer := routes.Group("/api/auth")
	{
		// RegisterCustomer funksiyada musderi ulgama registrasiya bolyar
		customer.POST("/register", frontController.RegisterCustomer)

		// LoginCustomer funksiyada musderi ulgama login bolyar
		customer.POST("/login", frontController.LoginCustomer)

		// Refresh funksiya musderini tokenini tazelemek ucin ulanylyar
		customer.POST("/refresh", auth.Refresh)
	}

	// bu group - a degisli api - lerden maglumat alynanda ( :lang ) paramter boyunca uytgedilip
	// terjime alynyar
	front := routes.Group("/api/:lang")
	{
		// GetHeaderData header - e degisli ahli maglumatlary alyar
		front.GET("/header", frontController.GetHeaderData)

		// GetFooterData funksiya footer - a degisli maglumnatlary alyar
		front.GET("/footer", frontController.GetFooterData)

		// GetBrends funksiya ahli brendlerin suratlaryny we id - lerini getiryar
		front.GET("/brends", frontController.GetBrends)

		// GetCompanyPhones funksiya firmany  ahli telefon belgilerini getirip beryar
		front.GET("/company-phones", backController.GetCompanyPhones)

		// GetCompanyAddress funksiya dil boyunca firmanyn salgysyny getirip beryar
		front.GET("/company-address", backController.GetCompanyAddress)

		// GetTranslationSecureByLangID funksiya dil boyunca ulanys duzgunleri we
		// gizlinlik sertleri sahypasynyn terjimesini getirip beryar
		front.GET("/translation-secure", backController.GetTranslationSecureByLangID)

		// GetTranslationPaymentByLangID funksiya dil boyunca eltip bermek
		// we toleg tertibi sahypasynyn terjimesini getirip beryar
		front.GET("/translation-payment", backController.GetTranslationPaymentByLangID)

		// GetTranslationAboutByLangID funksiya dil boyunca biz barada sahypanyn
		// terjimesini getirip beryar
		front.GET("/translation-about", backController.GetTranslationAboutByLangID)

		// GetTranslationContactByLangID funksiya dil boyunca aragatnasyk ( habarlasmak )
		// sahypasynyn terjimesini getirip beryar
		front.GET("/translation-contact", backController.GetTranslationContactByLangID)

		// GetTranslationUpdatePasswordPageByLangID funksiya dil boyunca
		// musderinin parol uytgetyan sahypasynyn terjimesini getirip beryar
		front.GET("/translation-update-password-page", backController.GetTranslationUpdatePasswordPageByLangID)

		// GetTranslationBasketPageByLangID funksiya dil boyunca sebet sahypasynyn
		// terjimesini getirip beryar
		front.GET("/translation-basket-page", backController.GetTranslationBasketPageByLangID)

		// GetTranslationOrderPageByLangID funksiya dil boyunca sargyt sahypanyn
		// terjimesini getirip beryar
		front.GET("/translation-order-page", backController.GetTranslationOrderPageByLangID)

		// GetTranslationMyOrderPageByLangID funksiya dil boyunca musderinin
		// eden sargytlaryny gorjek sahypasynyn terjimesini getiryar
		front.GET("/translation-my-order-page", backController.GetTranslationMyOrderPageByLangID)

		// GetPaymentTypesByLangID funksiya dil boyunca toleg gornuslerinin
		// terjimesini getirip beryar
		front.GET("/payment-types", backController.GetPaymentTypesByLangID)

		// GetNotificationByLangID funksiya dil boyunca ahli bildirislerin ( notification )
		// terjimesini getirip beryar
		front.GET("/notifications", backController.GetNotificationByLangID)

		// GetHomePageCategories funksiya dil boyunca bas sahypada duryan kategoriyalary
		// 4 sany harydy bilen bilelikde getiryar
		front.GET("/homepage-categories", frontController.GetHomePageCategories)

		// GetOneCategoryWithProducts funksiya dil boyunca dine bir kategoriyany
		// ahli harytlary pagination edip getiryar
		front.GET("/:category_id/:limit/:page", backController.GetOneCategoryWithProducts)

		// GetOrderTime funksiya dil boyunca musderi ucin sargyt edilip bilinjek
		// wagtlary getirip beryar
		front.GET("/order-time", backController.GetOrderTime)

		// Search funksiya dil boyunca gozlenilen harytlary pagination edip
		// getirip beryar
		front.POST("/search/:limit/:page", frontController.Search)

		// GetTranslationMyInformationPageByLangID funksiya dil boyunca musderinin maglumatlarym
		// sahypasynyn terjimesinin   getirip beryar
		front.GET("/translation-my-information-page", backController.GetTranslationMyInformationPageByLangID)

		// ToOrder funksiya sargyt sebede gosulan harytlary sargyt etmek ucin ulanylyar
		front.POST("/to-order", frontController.ToOrder)

		// SendMail funksiya musderi habarlasmak sahypa girip hat yazanda firma hat ugratyar
		front.POST("/send-mail", frontController.SendMail)

		// get like products without customer by product id ->
		// Eger musderi like - a haryt gosup sonam sol haryt bazadan ayrylan bolsa
		// sony bildirmek ucin front - dan mana cookie - daki product_id - leri
		// ugradyar we men yzyna sol id - leri product - lary ugratyan

		// get order products without customer by product id ->
		// Eger musderi sebede - e haryt gosup sonam sol haryt bazadan ayrylan bolsa
		// sony bildirmek ucin front - dan mana cookie - daki product_id - leri
		// ugdurkdyryar we men yzyna sol id - leri product - lary ugratyan

		front.POST("/likes-or-orders-without-customer", frontController.GetLikedOrOrderedProductsWithoutCustomer)

		// get order products without customer by product id ->
		// Eger musderi sebede - e haryt gosup sonam sol haryt bazadan ayrylan bolsa
		// sony bildirmek ucin front - dan mana cookie - daki product_id - leri
		// ugdurkdyryar we men yzyna sol id - leri product - lary ugratyan
		// front.POST("/orders-without-customer", frontController.GetOrderedProductsWithoutCustomer)

		securedCustomer := front.Group("/").Use(middlewares.Auth())
		{
			// AddLike funksiya musderinin tokeni bar bolan yagdayynda
			// halanlarym sahypa haryt gosmak ucin ya-da halanlarym sahypadan
			// haryt pozmak ucin ulanylyar
			securedCustomer.POST("/like", frontController.AddOrRemoveLike)

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

			// update customer informations
			securedCustomer.PUT("/my-information", frontController.UpdateCustomerInformation)

			// update customer address status
			securedCustomer.PUT("/address", frontController.UpdateCustomerAddressStatus)

			// add address to customer
			securedCustomer.POST("/address", frontController.AddAddressToCustomer)

			// update customer password
			securedCustomer.PUT("/customer-password", frontController.UpdateCustomerPassword)

			// add address to customer
			securedCustomer.POST("/customer-password", frontController.UpdateCustomerPassword)

		}

	}

	return routes

}
