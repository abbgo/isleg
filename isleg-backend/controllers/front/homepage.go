package controllers

import (
	"net/http"

	backController "github/abbgo/isleg/isleg-backend/controllers/back"

	"github.com/gin-gonic/gin"
)

func GetBrends(c *gin.Context) {

	brends, err := backController.GetAllBrendForHomePage()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"brends": brends,
	})

}
