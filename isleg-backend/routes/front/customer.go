package frontApi

import (
	"github/abbgo/isleg/isleg-backend/auth"
	frontController "github/abbgo/isleg/isleg-backend/controllers/front"

	"github.com/gin-gonic/gin"
)

func CustomerRoutes(front *gin.RouterGroup) {

	customer := front.Group("/auth")
	{
		// RegisterCustomer funksiyada musderi ulgama registrasiya bolyar
		customer.POST("register", frontController.RegisterCustomer)

		// LoginCustomer funksiyada musderi ulgama login bolyar
		customer.POST("login", frontController.LoginCustomer)

		// Refresh funksiya musderini tokenini tazelemek ucin ulanylyar
		customer.POST("refresh", auth.Refresh)
	}

}
