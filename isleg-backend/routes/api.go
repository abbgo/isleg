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

	}

	// front := routes.Group("/api/:lang")
	// {
	// 	front.GET("/header", frontController.GetHeaderData)

	// }

	return routes

}
