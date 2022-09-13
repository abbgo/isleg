package controllers

import (
	"github/abbgo/isleg/isleg-backend/config"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

type Cart struct {
	CustomerID string        `json:"customer_id"`
	Products   []CartProduct `json:"products"`
}

type CartProduct struct {
	ProductID         string `json:"product_id"`
	QuantityOfProduct int    `json:"quantity_of_product"`
}

func AddCart(c *gin.Context) {

	db, err := config.ConnDB()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer db.Close()

	var cart Cart
	var ids []string
	var counts []int

	if err := c.BindJSON(&cart); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	rowCustomer, err := db.Query("SELECT id FROM customers WHERE id = $1 AND deleted_at IS NULL", cart.CustomerID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer rowCustomer.Close()

	var customer_id string

	for rowCustomer.Next() {
		if err := rowCustomer.Scan(&customer_id); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}

	if customer_id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "Customer not found",
		})
		return
	}

	for k, v := range cart.Products {

		if v.QuantityOfProduct < 1 {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": "quantity of product cannot be less than 1",
			})
			return
		}

		for _, x := range cart.Products[(k + 1):] {
			if v.ProductID == x.ProductID {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  false,
					"message": "the same product cannot be repeated twice",
				})
				return
			}
		}

		rowProduct, err := db.Query("SELECT id FROM products WHERE id = $1 AND deleted_at IS NULL", v.ProductID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
		defer rowProduct.Close()

		var product_id string

		for rowProduct.Next() {
			if err := rowProduct.Scan(&product_id); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  false,
					"message": err.Error(),
				})
				return
			}
		}

		if product_id == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": "product not found",
			})
			return
		}

		ids = append(ids, v.ProductID)
		counts = append(counts, v.QuantityOfProduct)

	}

	resultCart, err := db.Query("INSERT INTO cart (customer_id,product_id,quantity_of_product) VALUES ($1,unnest($2::uuid[]),unnest($3::int[]))", cart.CustomerID, pq.Array(ids), pq.Array(counts))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			// "message": err.Error(),
			"message": "yalnys",
		})
		return
	}
	defer resultCart.Close()

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "product added successfull to cart",
	})

}

// func AddProductToBasket(c *gin.Context) {

// 	db, err := config.ConnDB()
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"status":  false,
// 			"message": err.Error(),
// 		})
// 		return
// 	}
// 	defer db.Close()

// 	customerID := c.PostForm("customer_id")
// 	productID := c.PostForm("product_id")
// 	quantityOfProduct := c.PostForm("quantity_of_product")

// 	quantityInt, err := models.ValidateCustomerBasket(customerID, productID, quantityOfProduct)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"status":  false,
// 			"message": err.Error(),
// 		})
// 		return
// 	}

// 	rowBasket, err := db.Query("SELECT COUNT(*) FROM basket WHERE product_id = $1 AND customer_id = $2 AND deleted_at IS NULL", productID, customerID)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"status":  false,
// 			"message": err.Error(),
// 		})
// 		return
// 	}

// 	var quantity int

// 	for rowBasket.Next() {
// 		if err := rowBasket.Scan(&quantity); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{
// 				"status":  false,
// 				"message": err.Error(),
// 			})
// 			return
// 		}
// 	}

// 	if quantity == 0 {
// 		resultBasket, err := db.Query("INSERT INTO basket (product_id,customer_id,quantity_of_product) VALUES ($1,$2,$3)", productID, customerID, quantityInt)
// 		if err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{
// 				"status":  false,
// 				"message": err.Error(),
// 			})
// 			return
// 		}
// 		defer resultBasket.Close()
// 	} else {
// 		resultBasket, err := db.Query("UPDATE basket SET quantity_of_product = $1 WHERE product_id = $2 AND customer_id = $3", quantityInt, productID, customerID)
// 		if err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{
// 				"status":  false,
// 				"message": err.Error(),
// 			})
// 			return
// 		}
// 		defer resultBasket.Close()
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"status":  true,
// 		"message": "product successfully added to card",
// 	})

// }
