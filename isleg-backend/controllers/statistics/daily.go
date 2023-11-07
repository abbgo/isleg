package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetDailyStatistics(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"daily":  "daily statistics",
	})

}
