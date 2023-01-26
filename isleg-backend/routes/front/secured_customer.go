package frontApi

import (
	frontController "github/abbgo/isleg/isleg-backend/controllers/front"
	"github/abbgo/isleg/isleg-backend/middlewares"

	"github.com/gin-gonic/gin"
)

func SecuredCustomerRoutes(front *gin.RouterGroup) {

	securedCustomer := front.Group("").Use(middlewares.Auth())
	{
		// AddOrRemoveLike funksiya musderinin tokeni bar bolan yagdayynda
		// halanlarym sahypa haryt gosmak ucin ya-da halanlarym sahypadan
		// haryt pozmak ucin ulanylyar
		securedCustomer.POST("like", frontController.AddOrRemoveLike)

		// remove like if customer exists
		// securedCustomer.POST("like/:product_id", frontController.RemoveLike)

		// GetCustomerLikes funksiya frontdan token bar bolan yagdayynda
		// musderinin halanlarym sahypasyna gosan harytlaryny getiryar
		securedCustomer.GET("likes", frontController.GetCustomerLikes)

		// AddCart funksiya sebede haryt gosmak ucin ulanylyar
		// musderinin tokeni gelen yagdayynda
		securedCustomer.POST("add-cart", frontController.AddCart)

		// GetCustomerCartProducts funksiya musderinin sebedindaki harytlary fronda bermek ucin ulanylyar
		// token bar bolan yagdayynda
		securedCustomer.GET("get-cart", frontController.GetCustomerCartProducts)

		// RemoveCart funksiya musderinin sebedinden haryt pozmak ucin ulanylyar
		// token bar bolan yagdayynda
		securedCustomer.POST("remove-cart", frontController.RemoveCart)

		// GetCustomerOrders funkisya musderinin bazadaky onki sargytlaryny
		// getirip beryar. token bar bolan yagdayynda
		securedCustomer.GET("orders/:limit/:page", frontController.GetCustomerOrders)

		// GetCustomerAddresses funksiya musderinin ahli salgylaryny alyar
		// token bar bolan yagdayynda
		securedCustomer.GET("addresses", frontController.GetCustomerAddresses)

		// GetCustomerInformation funksiya musderinin maglumatllaryny alyar
		// token bar bolan yagdayynda
		securedCustomer.GET("my-information", frontController.GetCustomerInformation)

		// UpdateCustomerInformation funksiya musderinin maglumatlary uytgetmek
		// ucin ulanylyar. token bar bolan yagdayynda
		securedCustomer.PUT("my-information", frontController.UpdateCustomerInformation)

		// UpdateCustomerAddressStatus funksiya musderinin salgysynyn aktiwligini uytgetyar
		// musderi haryt sargyt edende haysy salgysyny ulanjak bolsa sony aktiwe edip goyar yaly
		securedCustomer.PUT("address", frontController.UpdateCustomerAddressStatus)

		// AddAddressToCustomer musderi ozune taze salgy gosup biler yaly yazylan funksiya
		// token bar bolan yagdayynda
		securedCustomer.POST("address", frontController.AddAddressToCustomer)

		// DeleteCustomerAddress musderi onki gosan salgysyny pozup biler yaly yazylan funksiya
		// token bar bolan yagdayynda
		securedCustomer.DELETE("address/:id", frontController.DeleteCustomerAddress)

		// UpdateCustomerPassword funksiya musderinin parolyny uytgetmek ucin
		// token bar bolan yagdayynda
		securedCustomer.PUT("customer-password", frontController.UpdateCustomerPassword)

	}

}
