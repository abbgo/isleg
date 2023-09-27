package middlewares

import (
	"github/abbgo/isleg/isleg-backend/config"
	controllers "github/abbgo/isleg/isleg-backend/controllers/back"

	"github.com/gin-gonic/gin"
)

func CheckLang() gin.HandlerFunc {
	return func(context *gin.Context) {
		db, err := config.ConnDB()
		if err != nil {
			context.AbortWithStatusJSON(400, gin.H{"message": err.Error()})
			return
		}
		defer db.Close()

		// GET DATA FROM ROUTE PARAMETER
		langShortName := context.Param("lang")

		// GET ID OFF LANGUAGE
		langID, err := controllers.GetLangID(langShortName)
		if err != nil {
			context.AbortWithStatusJSON(404, gin.H{"message": err.Error()})
			return
		}

		context.Set("lang_id", langID)
		context.Next()
	}
}
