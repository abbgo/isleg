package controllers

import (
	"github/abbgo/isleg/isleg-backend/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetFooterData(c *gin.Context) {

	// GET DATA FROM ROUTE PARAMETER
	langShortName := c.Param("lang")

	// GET language id
	var langID string
	row, err := config.ConnDB().Query("SELECT id FROM languages WHERE name_short = $1", langShortName)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	for row.Next() {
		if err := row.Scan(&langID); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}

}
