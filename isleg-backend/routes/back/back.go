package backApi

import (
	backController "github/abbgo/isleg/isleg-backend/controllers/back"
	frontController "github/abbgo/isleg/isleg-backend/controllers/front"
	"github/abbgo/isleg/isleg-backend/middlewares"

	"github.com/gin-gonic/gin"
)

func BackRoutes(back *gin.RouterGroup) {

	securedAdmin := back.Group("/").Use(middlewares.CheckAdmin())
	{

		// // GetAdmins funksiya hemme adminlerin spisoygyny almak ucin ulanylyar.
		// admin.GET("/admins/:limit/:page", adminController.GetAdmins)

		// GetOrders funksiya musderilerin sargyt edilen harytlaryny almak ucin ulanylyar
		// yagny musderiler tarapyndan sargyt edilen harytlary adminkada gormek ucin
		securedAdmin.GET("orders/:limit/:page", frontController.GetOrders)

		// OrderConfirmation funksiya musderiler tarapyndan edilen sargydy
		// tassyklamak ucin ulanylyar
		securedAdmin.GET("order-confirmation/:id", frontController.OrderConfirmation)

		// ReturnOrder funksiya musderi eden sargydyny tassyklamasa edilen sargydy
		// yzyna gaytarmak ucin ulanylyar
		securedAdmin.GET("return-order/:id", frontController.ReturnOrder)

		// Baslayar --- dil ucin CRUD
		securedAdmin.POST("language", backController.CreateLanguage)
		securedAdmin.PUT("language/:id", backController.UpdateLanguageByID)
		securedAdmin.GET("language/:id", backController.GetLanguageByID)
		securedAdmin.GET("languages", backController.GetLanguages)
		securedAdmin.DELETE("language/:id", backController.DeleteLanguageByID)
		securedAdmin.GET("restore-language/:id", backController.RestoreLanguageByID)
		securedAdmin.DELETE("delete-language/:id", backController.DeletePermanentlyLanguageByID)
		// Gutaryar --- dil ucin CRUD

		// Baslayar --- firmanyn sazlamalary ucin ( meselem logo , favicon , email addres we s,m ) ucin CRUD
		securedAdmin.POST("company-setting", backController.CreateCompanySetting)
		securedAdmin.PUT("company-setting", backController.UpdateCompanySetting)
		securedAdmin.GET("company-setting", backController.GetCompanySetting)
		// Gutaryar --- firmanyn sazlamalary ucin CRUD

		// Baslayar --- Musderilerin harydy sargyt edip biljek wagtlary ucin CRUD
		securedAdmin.POST("order-date", backController.CreateOrderDate)
		// securedAdmin.PUT("order-date", backController.UpdateOrderDateByID)
		// securedAdmin.GET("order-time/:id", backController.GetOrderTimeByID)
		// securedAdmin.GET("order-times", backController.GetOrderTimes)
		// securedAdmin.DELETE("order-time/:id", backController.DeleteOrderTimeByID)
		// securedAdmin.GET("restore-order-time/:id", backController.RestoreOrderTimeByID)
		// securedAdmin.DELETE("delete-order-time/:id", backController.DeletePermanentlyOrderTimeByID)
		// GUtaryar --- Musderilerin harydy sargyt edip biljek wagtlary ucin CRUD

		securedAdmin.POST("order-time", backController.CreateOrderTime)

		securedAdmin.POST("order-date-hour", backController.CreateOrderDateHour)

		// Baslayar --- Bas sahypada duryan banner ucin CRUD
		securedAdmin.POST("banner", backController.CreateBanner)
		securedAdmin.PUT("banner/:id", backController.UpdateBannerByID)
		securedAdmin.GET("banner/:id", backController.GetBannerByID)
		securedAdmin.GET("banners/:limit/:page", backController.GetBanners)
		securedAdmin.DELETE("banner/:id", backController.DeleteBannerByID)
		securedAdmin.GET("restore-banner/:id", backController.RestoreBannerByID)
		securedAdmin.DELETE("delete-banner/:id", backController.DeletePermanentlyBannerByID)
		// Gutaryar --- Bas sahypada duryan banner ucin CRUD

		// Baslayar --- header - in terjimesi ucin CRUD
		securedAdmin.POST("translation-header", backController.CreateTranslationHeader)
		securedAdmin.PUT("translation-header", backController.UpdateTranslationHeaderByID)
		securedAdmin.GET("translation-header/:id", backController.GetTranslationHeaderByID)
		// Gutaryar --- header - in terjimesi ucin CRUD

		// Baslayar --- footer - in terjimesi ucin CRUD
		securedAdmin.POST("translation-footer", backController.CreateTranslationFooter)
		securedAdmin.PUT("translation-footer", backController.UpdateTranslationFooterByID)
		securedAdmin.GET("translation-footer/:id", backController.GetTranslationFooterByID)
		// Gutaryar --- footer - in terjimesi ucin CRUD

		// Baslayar --- Ulanys duzgunleri we gizlinlik sertnamasy sahypasynyn terjimesi ucin CRUD
		securedAdmin.POST("translation-secure", backController.CreateTranslationSecure)
		securedAdmin.PUT("translation-secure", backController.UpdateTranslationSecureByID)
		securedAdmin.GET("translation-secure/:id", backController.GetTranslationSecureByID)
		// Gutaryar --- Ulanys duzgunleri we gizlinlik sertnamasy sahypasynyn terjimesi ucin CRUD

		// Baslayar --- eltip bermek we toleg tertibi sahypasynyn terjimesi ucin CRUD
		securedAdmin.POST("translation-payment", backController.CreateTranslationPayment)
		securedAdmin.PUT("translation-payment", backController.UpdateTranslationPaymentByID)
		securedAdmin.GET("translation-payment/:id", backController.GetTranslationPaymentByID)
		// Gutaryar --- eltip bermek we toleg tertibi sahypasynyn terjimesi ucin CRUD

		// Baslayar --- biz barada sahypasynyn terjimesi ucin CRUD
		securedAdmin.POST("translation-about", backController.CreateTranslationAbout)
		securedAdmin.PUT("translation-about", backController.UpdateTranslationAboutByID)
		securedAdmin.GET("translation-about/:id", backController.GetTranslationAboutByID)
		// Gutaryar --- biz barada sahypasynyn terjimesi ucin CRUD

		// Baslayar --- aragatnasyk ( habaralasmak ) sahypasynyn terjimesi ucin CRUD
		securedAdmin.POST("translation-contact", backController.CreateTranslationContact)
		securedAdmin.PUT("translation-contact", backController.UpdateTranslationContactByID)
		securedAdmin.GET("translation-contact/:id", backController.GetTranslationContactByID)
		// Gutaryar --- aragatnasyk ( habaralasmak ) sahypasynyn terjimesi ucin CRUD

		// Baslayar --- musderinin maglumatlarym sahypasynyn terjimesi ucin CRUD
		securedAdmin.POST("translation-my-information-page", backController.CreateTranslationMyInformationPage)
		securedAdmin.PUT("translation-my-information-page", backController.UpdateTranslationMyInformationPageByID)
		securedAdmin.GET("translation-my-information-page/:id", backController.GetTranslationMyInformationPageByID)
		// Gutaryar --- musderinin maglumatlarym sahypasynyn terjimesi ucin CRUD

		// Baslayar --- musderinin acar sozuni uytget sahypasynyn terjimesi ucin CRUD
		securedAdmin.POST("translation-update-password-page", backController.CreateTranslationUpdatePasswordPage)
		securedAdmin.PUT("translation-update-password-page", backController.UpdateTranslationUpdatePasswordPageByID)
		securedAdmin.GET("translation-update-password-page/:id", backController.GetTranslationUpdatePasswordPageByID)
		// Gutaryar --- musderinin acar sozuni uytget sahypasynyn terjimesi ucin CRUD

		// Baslayar --- sebet sahypasynyn terjimesi ucin CRUD
		securedAdmin.POST("translation-basket-page", backController.CreateTranslationBasketPage)
		securedAdmin.PUT("translation-basket-page", backController.UpdateTranslationBasketPageByID)
		securedAdmin.GET("translation-basket-page/:id", backController.GetTranslationBasketPageByID)
		// Gutaryar --- sebet sahypasynyn terjimesi ucin CRUD

		// Baslayar --- haryt sargyt edilyan sahypasynyn terjimesi ucin CRUD
		securedAdmin.POST("translation-order-page", backController.CreateTranslationOrderPage)
		securedAdmin.PUT("translation-order-page", backController.UpdateTranslationOrderPageByID)
		securedAdmin.GET("translation-order-page/:id", backController.GetTranslationOrderPageByID)
		// Gutaryar --- haryt sargyt edilyan sahypasynyn terjimesi ucin CRUD

		// Baslayar --- musderinin sargytlaryny goryan sahypasynyn terjimesi ucin CRUD
		securedAdmin.POST("translation-my-order-page", backController.CreateTranslationMyOrderPage)
		securedAdmin.PUT("translation-my-order-page", backController.UpdateTranslationMyOrderPageByID)
		securedAdmin.GET("translation-my-order-page/:id", backController.GetTranslationMyOrderPageByID)
		// Gutaryar --- musderinin sargytlaryny goryan sahypasynyn terjimesi ucin CRUD

		// Baslayar --- kategoriya ucin CRUD
		securedAdmin.POST("category", backController.CreateCategory)
		securedAdmin.PUT("category/:id", backController.UpdateCategoryByID)
		securedAdmin.GET("category/:id", backController.GetCategoryByID)
		// GetCategoryByIDWithChild funksiya bir kategoriyany cagalary bilen bile almak ucin ulanylyar
		securedAdmin.GET("category-with-child/:id", backController.GetCategoryByIDWithChild)
		securedAdmin.GET("categories/:limit/:page", backController.GetCategories)
		securedAdmin.GET("categories", backController.GetCategoriesForAdmin)
		// GetAllCategory korzina salynan harytlaryn kategoriyalaryny cekmek ucin yazylan funksiya
		securedAdmin.GET("all-category", backController.GetAllCategory)
		securedAdmin.GET("deleted-categories", backController.GetDeletedCategories)
		securedAdmin.DELETE("category/:id", backController.DeleteCategoryByID)
		securedAdmin.GET("restore-category/:id", backController.RestoreCategoryByID)
		securedAdmin.DELETE("delete-category/:id", backController.DeletePermanentlyCategoryByID)
		// Gutaryar --- kategoriya ucin CRUD

		// Baslayar --- brend ucin CRUD
		securedAdmin.POST("brend", backController.CreateBrend)
		securedAdmin.PUT("brend/:id", backController.UpdateBrendByID)
		securedAdmin.GET("brend/:id", backController.GetBrendByID)
		securedAdmin.GET("brends/:limit/:page", backController.GetBrends)
		securedAdmin.DELETE("brend/:id", backController.DeleteBrendByID)
		securedAdmin.GET("restore-brend/:id", backController.RestoreBrendByID)
		securedAdmin.DELETE("delete-brend/:id", backController.DeletePermanentlyBrendByID)
		// Gutaryar --- brend ucin CRUD

		// Baslayar --- haryt ucin CRUD

		securedAdmin.POST("product", backController.CreateProduct)
		// securedAdmin.POST("product-in-excel", backController.CreateProductsByExcelFile)
		securedAdmin.PUT("product/:id", backController.UpdateProductByID)
		securedAdmin.GET("product/:id", backController.GetProductByID)
		securedAdmin.GET("products", backController.GetProducts)
		securedAdmin.DELETE("product/:id", backController.DeleteProductByID)
		securedAdmin.GET("restore-product/:id", backController.RestoreProductByID)
		securedAdmin.DELETE("delete-product/:id", backController.DeletePermanentlyProductByID)
		// Gutaryar --- haryt ucin CRUD

		// Baslayar --- firmanyn telefon belgisi ucin CRUD
		securedAdmin.POST("company-phone", backController.CreateCompanyPhone)
		securedAdmin.PUT("company-phone", backController.UpdateCompanyPhoneByID)
		securedAdmin.GET("company-phone/:id", backController.GetCompanyPhoneByID)
		securedAdmin.DELETE("company-phone/:id", backController.DeleteCompanyPhoneByID)
		securedAdmin.GET("restore-company-phone/:id", backController.RestoreCompanyPhoneByID)
		securedAdmin.DELETE("delete-company-phone/:id", backController.DeletePermanentlyCompanyPhoneByID)
		// Gutaryar --- firmanyn telefon belgisi ucin CRUD

		// Baslayar --- firmanyn adresi we onun terjimesi ucin CRUD
		securedAdmin.POST("company-address", backController.CreateCompanyAddress)
		securedAdmin.PUT("company-address", backController.UpdateCompanyAddressByID)
		securedAdmin.GET("company-address/:id", backController.GetCompanyAddressByID)
		// GUtaryar --- firmanyn adresi we onun terjimesi ucin CRUD

		// Baslayar --- musderi sargyt edende tolegin gornusleri ucin CRUD
		securedAdmin.POST("payment-type", backController.CreatePaymentType)
		securedAdmin.PUT("payment-type", backController.UpdatePaymentTypeByID)
		securedAdmin.GET("payment-type/:id", backController.GetPaymentTypeByID)
		securedAdmin.GET("payment-types", backController.GetPaymentTypes)
		// Gutaryar --- musderi sargyt edende tolegin gornusleri ucin CRUD

		// Baslayar --- afisa ucin CRUD
		securedAdmin.POST("afisa", backController.CreateAfisa)
		securedAdmin.PUT("afisa/:id", backController.UpdateAfisaByID)
		securedAdmin.GET("afisa/:id", backController.GetAfisaByID)
		securedAdmin.GET("afisas/:limit/:page", backController.GetAfisas)
		securedAdmin.DELETE("afisa/:id", backController.DeleteAfisaByID)
		securedAdmin.GET("restore-afisa/:id", backController.RestoreAfisaByID)
		securedAdmin.DELETE("delete-afisa/:id", backController.DeletePermanentlyAfisaByID)
		// Gutaryar --- afisa ucin CRUD

		// Baslayar --- duyduryslar ( notification ) we olaryn terjimeleri ucin CRUD
		securedAdmin.POST("notification", backController.CreateNotification)
		securedAdmin.PUT("notification", backController.UpdateNotificationByID)
		securedAdmin.GET("notification/:id", backController.GetNotificationByID)
		securedAdmin.GET("notifications", backController.GetNotifications)
		securedAdmin.DELETE("notification/:id", backController.DeleteNotificationByID)
		securedAdmin.GET("restore-notification/:id", backController.RestoreNotificationByID)
		securedAdmin.DELETE("delete-notification/:id", backController.DeletePermanentlyNotificationByID)
		// Gutaryar --- duyduryslar ( notification ) we olaryn terjimeleri ucin CRUD

		// eger sargydyn bahasy durli rayonlara gora uytgeyan bolsa yazylmaly funksiya
		// securedAdmin.POST("district", backController.CreateDistrict)

		// Baslayar --- haryt alynjak magazynlar ucin CRUD
		securedAdmin.POST("shop", backController.CreateShop)
		securedAdmin.PUT("shop/:id", backController.UpdateShopByID)
		securedAdmin.GET("shop/:id", backController.GetShopByID)
		securedAdmin.GET("shops/:limit/:page", backController.GetShops)
		securedAdmin.DELETE("shop/:id", backController.DeleteShopByID)
		securedAdmin.GET("restore-shop/:id", backController.RestoreShopByID)
		securedAdmin.DELETE("delete-shop/:id", backController.DeletePermanentlyShopByID)
		// GUtaryar --- haryt alynjak magazynlar ucin CRUD

		securedAdmin.POST("image", backController.CreateProductImage)
		securedAdmin.DELETE("image", backController.DeleteProductImages)
		securedAdmin.POST("excel", backController.UploadExcelFile)
		securedAdmin.DELETE("excel", backController.RemoveExcelFile)

	}

}
