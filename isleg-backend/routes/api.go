package routes

import (
	backController "github/abbgo/isleg/isleg-backend/controllers/back"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Routes() *gin.Engine {

	routes := gin.Default()
	routes.Use(cors.Default())

	back := routes.Group("/admin")
	{

		back.POST("/language", backController.CreateLanguage)

		back.POST("/company-setting", backController.CreateCompanySetting)

		back.POST("/translation-header", backController.CreateTranslationHeader)

		back.POST("/category", backController.CreateCategory)

		back.POST("/brend", backController.CreateBrend)

	}

	// front := routes.Group("/api/:lang")
	// {
	// 	front.GET("/header", frontController.GetHeaderData)

	// }

	return routes

}
