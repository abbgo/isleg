package controllers

import (
	"github/abbgo/isleg/isleg-backend/config"
	"github/abbgo/isleg/isleg-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddLike(c *gin.Context) {

	db, err := config.ConnDB()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer db.Close()

	customerID := c.PostForm("customer_id")
	productID := c.PostForm("product_id")

	err = models.ValidateCustomerLike(customerID, productID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	resultLike, err := db.Query("INSERT INTO likes (product_id,customer_id) VALUES ($1,$2)", productID, customerID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer resultLike.Close()

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "like successfully added",
	})

}
