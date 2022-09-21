package controllers

import (
	"github/abbgo/isleg/isleg-backend/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Order struct {
	FullName     string        `json:"full_name" binding:"required,min=3"`
	PhoneNumber  string        `json:"phone_number" binding:"required,e164,len=12"`
	Address      string        `json:"address" binding:"required,min=3"`
	CustomerMark string        `json:"customer_mark"`
	OrderTime    string        `json:"order_time" binding:"required"`
	PaymentType  string        `json:"payment_type" binding:"required"`
	TotalPrice   float64       `json:"total_price" binding:"required"`
	Products     []CartProduct `json:"products" binding:"required"`
}

func ToOrder(c *gin.Context) {

	db, err := config.ConnDB()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer db.Close()

	var order Order

	if err := c.BindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	rowCustomer, err := db.Query("SELECT id,phone_number FROM customers WHERE phone_number = $1 AND deleted_at IS NULL", order.PhoneNumber)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer rowCustomer.Close()

	var customerPhoneNumber string
	var customerID string

	for rowCustomer.Next() {
		if err := rowCustomer.Scan(&customerID, &customerPhoneNumber); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}

	if customerPhoneNumber != "" {

		resultCustomerAddres, err := db.Query("INSERT INTO customer_address (customer_id,address,is_active) VALUES ($1,$2,$3)", customerID, order.Address, false)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
		defer resultCustomerAddres.Close()

	} else {

		resultCustomer, err := db.Query("INSERT INTO customers (full_name,phone_number,is_register) VALUES ($1,$2,$3)", order.FullName, order.PhoneNumber, false)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
		defer resultCustomer.Close()

		lastCustomerID, err := db.Query("SELECT id FROM customers ORDER BY created_at DESC LIMIT 1")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
		defer lastCustomerID.Close()

		var customer_id string

		for lastCustomerID.Next() {
			if err := lastCustomerID.Scan(&customer_id); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  false,
					"message": err.Error(),
				})
				return
			}
		}

		resultCustomerAddress, err := db.Query("INSERT INTO customer_address (customer_id,address,is_active) VALUES ($1,$2,$3)", customer_id, order.Address, false)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
		defer resultCustomerAddress.Close()

		customerID = customer_id

	}

	resultOrders, err := db.Query("INSERT INTO orders (customer_id,customer_mark,order_time,payment_type,total_price) VALUES ($1,$2,$3,$4,$5)", customerID, order.CustomerMark, order.OrderTime, order.PaymentType, order.TotalPrice)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer resultOrders.Close()

	lastOrderID, err := db.Query("SELECT id FROM orders ORDER BY created_at DESC LIMIT 1")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer lastOrderID.Close()

	var order_id string

	for lastOrderID.Next() {
		if err := lastOrderID.Scan(&order_id); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}

	for _, v := range order.Products {

		if v.QuantityOfProduct < 1 {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": "quantity of product cannot be less than 1",
			})
			return
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

		resultOrderedProduct, err := db.Query("INSERT INTO ordered_products (product_id,quantity_of_product,order_id) VALUES ($1,$2,$3)", v.ProductID, v.QuantityOfProduct, order_id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
		defer resultOrderedProduct.Close()

	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "success",
	})

}
