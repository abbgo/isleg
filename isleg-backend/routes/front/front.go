package frontApi

import (
	backController "github/abbgo/isleg/isleg-backend/controllers/back"
	frontController "github/abbgo/isleg/isleg-backend/controllers/front"
	"github/abbgo/isleg/isleg-backend/middlewares"

	"github.com/gin-gonic/gin"
)

func FrontRoutes(front *gin.RouterGroup) {

	frontRoutes := front.Group("/:lang")
	{
		// GetHeaderData header - e degisli ahli maglumatlary alyar
		frontRoutes.GET("header", middlewares.CheckLang(), frontController.GetHeaderData)

		// GetFooterData funksiya footer - a degisli maglumnatlary alyar
		frontRoutes.GET("footer", middlewares.CheckLang(), frontController.GetFooterData)

		// GetBrends funksiya ahli brendlerin suratlaryny we id - lerini getiryar
		frontRoutes.GET("brends", frontController.GetBrends)

		// GetCompanyPhones funksiya firmany  ahli telefon belgilerini getirip beryar
		frontRoutes.GET("company-phones", backController.GetCompanyPhones)

		// GetCompanyAddress funksiya dil boyunca firmanyn salgysyny getirip beryar
		frontRoutes.GET("company-address", middlewares.CheckLang(), backController.GetCompanyAddress)

		// GetTranslationSecureByLangID funksiya dil boyunca ulanys duzgunleri we
		// gizlinlik sertleri sahypasynyn terjimesini getirip beryar
		frontRoutes.GET("translation-secure", middlewares.CheckLang(), backController.GetTranslationSecureByLangID)

		// GetTranslationPaymentByLangID funksiya dil boyunca eltip bermek
		// we toleg tertibi sahypasynyn terjimesini getirip beryar
		frontRoutes.GET("translation-payment", middlewares.CheckLang(), backController.GetTranslationPaymentByLangID)

		// GetTranslationAboutByLangID funksiya dil boyunca biz barada sahypanyn
		// terjimesini getirip beryar
		frontRoutes.GET("translation-about", middlewares.CheckLang(), backController.GetTranslationAboutByLangID)

		// GetTranslationContactByLangID funksiya dil boyunca aragatnasyk ( habarlasmak )
		// sahypasynyn terjimesini getirip beryar
		frontRoutes.GET("translation-contact", middlewares.CheckLang(), backController.GetTranslationContactByLangID)

		// GetTranslationUpdatePasswordPageByLangID funksiya dil boyunca
		// musderinin parol uytgetyan sahypasynyn terjimesini getirip beryar
		frontRoutes.GET("translation-update-password-page", middlewares.CheckLang(), backController.GetTranslationUpdatePasswordPageByLangID)

		// GetTranslationBasketPageByLangID funksiya dil boyunca sebet sahypasynyn
		// terjimesini getirip beryar
		frontRoutes.GET("translation-basket-page", middlewares.CheckLang(), backController.GetTranslationBasketPageByLangID)

		// GetTranslationOrderPageByLangID funksiya dil boyunca sargyt sahypanyn
		// terjimesini getirip beryar
		frontRoutes.GET("translation-order-page", middlewares.CheckLang(), backController.GetTranslationOrderPageByLangID)

		// GetTranslationMyOrderPageByLangID funksiya dil boyunca musderinin
		// eden sargytlaryny gorjek sahypasynyn terjimesini getiryar
		frontRoutes.GET("translation-my-order-page", middlewares.CheckLang(), backController.GetTranslationMyOrderPageByLangID)

		// GetPaymentTypesByLangID funksiya dil boyunca toleg gornuslerinin
		// terjimesini getirip beryar
		frontRoutes.GET("payment-types", middlewares.CheckLang(), backController.GetPaymentTypesByLangID)

		// GetNotificationByLangID funksiya dil boyunca ahli bildirislerin ( notification )
		// terjimesini getirip beryar
		frontRoutes.GET("notifications", middlewares.CheckLang(), backController.GetNotificationByLangID)

		// GetHomePageCategories funksiya dil boyunca bas sahypada duryan kategoriyalary
		// 4 sany harydy bilen bilelikde getiryar
		frontRoutes.GET("homepage-categories", middlewares.CheckLang(), frontController.GetHomePageCategories)

		// GetOneCategoryWithDeletedProducts funksiya dil boyunca dine bir kategoriyany
		// ahli pozulan harytlary bilen pagination edip getiryar
		frontRoutes.GET("category-with-deleted-products/:id/:limit/:page", middlewares.CheckLang(), backController.GetOneCategoryWithDeletedProducts)

		// GetOneBrendWithProducts funksiya dil boyunca dine bir brendi
		// ahli harytlary bilen pagination edip getiryar
		frontRoutes.GET("brend/:id/:limit/:page", backController.GetOneBrendWithProducts)

		// GetOrderTime funksiya dil boyunca musderi ucin sargyt edilip bilinjek
		// wagtlary getirip beryar
		frontRoutes.GET("order-time", middlewares.CheckLang(), backController.GetOrderTime)

		// Search funksiya dil boyunca gozlenilen harytlary pagination edip
		// getirip beryar
		frontRoutes.GET("search/:limit/:page", middlewares.CheckLang(), frontController.Search)

		// FilterAndSort funksiya dil boyunca tertiplenen we filter edilen harytlary pagination edip
		// getirip beryar
		// frontRoutes.GET("category/:id/filter-and-sort/:limit/:page", frontController.FilterAndSort)

		// GetOneCategoryWithProducts funksiya dil boyunca dine bir kategoriyany
		// ahli harytlary bilen pagination edip getiryar
		frontRoutes.POST("category/:id/:limit/:page", middlewares.CheckLang(), backController.GetOneCategoryWithProducts)

		// GetTranslationMyInformationPageByLangID funksiya dil boyunca musderinin maglumatlarym
		// sahypasynyn terjimesinin   getirip beryar
		frontRoutes.GET("translation-my-information-page", middlewares.CheckLang(), backController.GetTranslationMyInformationPageByLangID)

		// ToOrder funksiya sargyt sebede gosulan harytlary sargyt etmek ucin ulanylyar
		frontRoutes.POST("to-order", middlewares.CheckLang(), frontController.ToOrder)

		// SendMail funksiya musderi habarlasmak sahypa girip hat yazanda firma hat ugratyar
		frontRoutes.POST("send-mail", frontController.SendMail)

		// get like products without customer by product id ->
		// Eger musderi like - a haryt gosup sonam sol haryt bazadan ayrylan bolsa
		// sony bildirmek ucin front - dan mana cookie - daki product_id - leri
		// ugradyar we men yzyna sol id - leri product - lary ugratyan

		// get order products without customer by product id ->
		// Eger musderi sebede  haryt gosup sonam sol haryt bazadan ayrylan bolsa
		// sony bildirmek ucin front - dan mana cookie - daki product_id - leri
		// ugdurkdyryar we men yzyna sol id - leri product - lary ugratyan

		frontRoutes.POST("likes-or-orders-without-customer", frontController.GetLikedOrOrderedProductsWithoutCustomer)

		// get order products without customer by product id ->
		// Eger musderi sebede - e haryt gosup sonam sol haryt bazadan ayrylan bolsa
		// sony bildirmek ucin front - dan mana cookie - daki product_id - leri
		// ugdurkdyryar we men yzyna sol id - leri product - lary ugratyan
		// frontRoutes.POST("orders-without-customer", frontController.GetOrderedProductsWithoutCustomer)

		frontRoutes.GET("product/:id", backController.GetProductByIDForFront)

		frontRoutes.GET("banners", backController.GetBannersForFront)

		frontRoutes.PUT("update-customer-password", frontController.UpdateCustPassword)

		frontRoutes.POST("check-otp", frontController.CheckOTP)

		frontRoutes.POST("count-of-customers", frontController.SetCountOfCustomers)

		SecuredCustomerRoutes(frontRoutes)
	}

}
