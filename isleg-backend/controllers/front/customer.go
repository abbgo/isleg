package controllers

import (
	"github/abbgo/isleg/isleg-backend/auth"
	"github/abbgo/isleg/isleg-backend/config"
	"time"

	"github/abbgo/isleg/isleg-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

type Login struct {
	PhoneNumber string `json:"phone_number" binding:"required,e164,len=12"`
	Password    string `json:"password" binding:"required,min=5,max=25"`
}

type CustomerInformation struct {
	ID          string      `json:"id"`
	FullName    string      `json:"full_name" binding:"required,min=3"`
	PhoneNumber string      `json:"phone_number" binding:"required,e164,len=12"`
	Birthday    pq.NullTime `json:"birthday"`
	Email       string      `json:"email" binding:"email"`
	IsRegister  bool        `json:"is_register"`
	Addresses   []Address   `json:"addresses"`
}

type Address struct {
	ID       string `json:"id"`
	Address  string `json:"address"`
	IsActive bool   `json:"is_active"`
}

func RegisterCustomer(c *gin.Context) {

	db, err := config.ConnDB()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer db.Close()

	var customer models.Customer

	if err := c.BindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err = models.ValidateCustomerRegister(customer.PhoneNumber, customer.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	hashPassword, err := models.HashPassword(customer.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	currentTime := time.Now()

	rowCustomer, err := db.Query("SELECT phone_number FROM customers WHERE phone_number = $1 AND is_register = false AND deleted_at IS NULL", customer.PhoneNumber)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer rowCustomer.Close()

	var phone_number string

	for rowCustomer.Next() {
		if err := rowCustomer.Scan(&phone_number); err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}

	if phone_number != "" {

		resultCustomer, err := db.Query("UPDATE customers SET full_name = $1 , password = $2 , email = $3 , is_register = $4 , created_at = $5 , updated_at = $6 WHERE phone_number = $7", customer.FullName, hashPassword, customer.Email, true, currentTime, currentTime, customer.PhoneNumber)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
		defer resultCustomer.Close()

	} else {

		resultCustomers, err := db.Query("INSERT INTO customers (full_name,phone_number,password,email,is_register) VALUES ($1,$2,$3,$4,$5)", customer.FullName, customer.PhoneNumber, hashPassword, customer.Email, true)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
		defer resultCustomers.Close()

	}

	row, err := db.Query("SELECT id FROM customers WHERE deleted_at IS NULL ORDER BY created_at ASC LIMIT 1")
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer row.Close()

	var customerID string

	for row.Next() {
		if err := row.Scan(&customerID); err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}

	if customerID == "" {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  false,
			"message": "record not found",
		})
		return
	}

	accessTokenString, err := auth.GenerateAccessToken(customer.PhoneNumber)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		c.Abort()
		return
	}

	refreshTokenString, err := auth.GenerateRefreshToken(customer.PhoneNumber)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":        true,
		"message":       "customer successfully added",
		"access_token":  accessTokenString,
		"refresh_token": refreshTokenString,
		"customer_id":   customerID,
	})

}

func LoginCustomer(c *gin.Context) {

	db, err := config.ConnDB()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer db.Close()

	var customer Login

	if err := c.BindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err = models.ValidateCustomerLogin(customer.PhoneNumber)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}

	// check if email exists and password is correct
	row, err := db.Query("SELECT id,password FROM customers WHERE phone_number = $1 AND is_register = true AND deleted_at IS NULL", customer.PhoneNumber)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}
	defer row.Close()

	var customerID, oldPassword string

	for row.Next() {
		if err := row.Scan(&customerID, &oldPassword); err != nil {
			c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
			return
		}
	}

	if oldPassword == "" || customerID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "this client does not exist"})
		return
	}

	credentialError := models.CheckPassword(customer.Password, oldPassword)
	if credentialError != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "invalid credentials"})
		return
	}

	accessTokenString, err := auth.GenerateAccessToken(customer.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	refreshTokenString, err := auth.GenerateRefreshToken(customer.PhoneNumber)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token":  accessTokenString,
		"refresh_token": refreshTokenString,
		"customer_id":   customerID,
	})

}

func GetCustomerInformation(c *gin.Context) {

	db, err := config.ConnDB()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer db.Close()

	customerID := c.Param("customer_id")

	rowCustomer, err := db.Query("SELECT id , full_name , phone_number , birthday , email FROM customers WHERE id = $1 AND is_register = true AND deleted_at IS NULL", customerID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer rowCustomer.Close()

	var customer CustomerInformation

	for rowCustomer.Next() {
		if err := rowCustomer.Scan(&customer.ID, &customer.FullName, &customer.PhoneNumber, &customer.Birthday, &customer.Email); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}

	if customer.ID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "customer not found",
		})
		return
	}

	rowsCustomerAddress, err := db.Query("SELECT id , address , is_active FROM customer_address WHERE deleted_at IS NULL AND customer_id = $1", customer.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer rowsCustomerAddress.Close()

	var addresses []Address

	for rowsCustomerAddress.Next() {
		var address Address

		if err := rowsCustomerAddress.Scan(&address.ID, &address.Address, &address.IsActive); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}

		addresses = append(addresses, address)
	}

	customer.Addresses = addresses

	c.JSON(http.StatusOK, gin.H{
		"status":                true,
		"customer_informations": customer,
	})

}

func UpdateCustomerPassword(c *gin.Context) {

	db, err := config.ConnDB()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer db.Close()

	customerID := c.Param("customer_id")
	password := c.PostForm("password")

	rowCustomer, err := db.Query("SELECT id FROM customers WHERE id = $1 AND deleted_at IS NULL", customerID)
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
			"message": "customer not found",
		})
		return
	}

	hashPassword, err := models.HashPassword(password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	resultCustomer, err := db.Query("UPDATE customers SET password = $1 WHERE id = $2", hashPassword, customerID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer resultCustomer.Close()

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "password of customer successfuly updated",
	})

}
