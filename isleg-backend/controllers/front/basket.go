package controllers

import (
	"github/abbgo/isleg/isleg-backend/config"
	"github/abbgo/isleg/isleg-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddProductToBasket(c *gin.Context) {

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
	quantityOfProduct := c.PostForm("quantity_of_product")

	quantityInt, err := models.ValidateCustomerBasket(customerID, productID, quantityOfProduct)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	rowBasket, err := db.Query("SELECT COUNT(*) FROM basket WHERE product_id = $1 AND customer_id = $2 AND deleted_at IS NULL", productID, customerID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	var quantity int

	for rowBasket.Next() {
		if err := rowBasket.Scan(&quantity); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}

	if quantity == 0 {
		resultBasket, err := db.Query("INSERT INTO basket (product_id,customer_id,quantity_of_product) VALUES ($1,$2,$3)", productID, customerID, quantityInt)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
		defer resultBasket.Close()
	} else {
		resultBasket, err := db.Query("UPDATE basket SET quantity_of_product = $1 WHERE product_id = $2 AND customer_id = $3", quantityInt, productID, customerID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
		defer resultBasket.Close()
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "product successfully added to card",
	})

}
