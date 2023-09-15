package controllers

import (
	"context"
	"github/abbgo/isleg/isleg-backend/auth"
	"github/abbgo/isleg/isleg-backend/config"
	"github/abbgo/isleg/isleg-backend/helpers"

	"github/abbgo/isleg/isleg-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gopkg.in/guregu/null.v4"
)

type Login struct {
	PhoneNumber string `json:"phone_number" binding:"required,e164,len=12"`
	Password    string `json:"password" binding:"required,min=3,max=25"`
}

type CustomerInformation struct {
	ID          string      `json:"id"`
	FullName    string      `json:"full_name" binding:"required,min=3"`
	PhoneNumber string      `json:"phone_number" binding:"required,e164,len=12"`
	Birthday    null.String `json:"birthday"`
	Email       string      `json:"email" binding:"email"`
	Gender      null.String `json:"gender" binding:"email"`
	IsRegister  bool        `json:"is_register"`
	Addresses   []Address   `json:"addresses"`
}

type Address struct {
	ID       string `json:"id"`
	Address  string `json:"address"`
	IsActive bool   `json:"is_active"`
}

func RegisterCustomer(c *gin.Context) {
	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	var customer models.Customer
	if err := c.BindJSON(&customer); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	if customer.Password == "" {
		helpers.HandleError(c, 400, "customer password is required")
		return
	}

	if len(customer.Password) < 5 || len(customer.Password) > 25 {
		helpers.HandleError(c, 400, "password length should be between 5 and 25")
		return
	}

	err = models.ValidateCustomerRegister(customer.PhoneNumber, customer.Email)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	hashPassword, err := models.HashPassword(customer.Password)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	// On registrasiya bolman haryt sargyt eden adamlar is_register false bolyar ,
	// solar sayta registrasiya boljak bolsa Olaryn taze at familiyasyny telefon belgisini
	// we parolyny baza yazdyrmak ucin database - den sol customer - leri tapyp update etmeli
	var phone_number string
	db.QueryRow(context.Background(), "SELECT phone_number FROM customers WHERE phone_number = $1 AND is_register = false AND deleted_at IS NULL", customer.PhoneNumber).Scan(&phone_number)

	var customerID string
	if phone_number != "" {
		db.QueryRow(context.Background(), "UPDATE customers SET full_name = $1 , password = $2 , email = $3 , is_register = $4 WHERE phone_number = $5 RETURNING id", customer.FullName, hashPassword, customer.Email, true, customer.PhoneNumber).Scan(&customerID)
	} else {
		db.QueryRow(context.Background(), "INSERT INTO customers (full_name,phone_number,password,email,is_register) VALUES ($1,$2,$3,$4,$5) RETURNING id", customer.FullName, customer.PhoneNumber, hashPassword, customer.Email, true).Scan(&customerID)
	}

	accessTokenString, refreshTokenString, err := auth.GenerateTokenForCustomer(customer.PhoneNumber, customerID)
	if err != nil {
		helpers.HandleError(c, 500, err.Error())
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":        true,
		"message":       "customer successfully added",
		"access_token":  accessTokenString,
		"refresh_token": refreshTokenString,
	})
}

func LoginCustomer(c *gin.Context) {
	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	var customer Login
	if err := c.BindJSON(&customer); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	if customer.Password == "" {
		helpers.HandleError(c, 400, "customer password is required")
		return
	}

	if len(customer.Password) < 5 || len(customer.Password) > 25 {
		helpers.HandleError(c, 400, "password length should be between 5 and 25")
		return
	}

	err = models.ValidateCustomerLogin(customer.PhoneNumber)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}

	// check if email exists and password is correct
	var customerID, oldPassword string
	db.QueryRow(context.Background(), "SELECT id,password FROM customers WHERE phone_number = $1 AND is_register = true AND deleted_at IS NULL", customer.PhoneNumber).Scan(&customerID, &oldPassword)

	if customerID == "" {
		helpers.HandleError(c, 404, "this client does not exist")
		return
	}

	credentialError := models.CheckPassword(customer.Password, oldPassword)
	if credentialError != nil {
		helpers.HandleError(c, 400, "invalid credentials")
		return
	}

	accessTokenString, refreshTokenString, err := auth.GenerateTokenForCustomer(customer.PhoneNumber, customerID)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token":  accessTokenString,
		"refresh_token": refreshTokenString,
	})
}

func GetCustomerInformation(c *gin.Context) {
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	custID, hasCustomer := c.Get("customer_id")
	if !hasCustomer {
		helpers.HandleError(c, 400, "customer_id is required")
		return
	}
	customerID, ok := custID.(string)
	if !ok {
		helpers.HandleError(c, 400, "customer_id must be string")
		return
	}

	// bazadan musderinin maglumatlary alynyar
	var customer CustomerInformation
	db.QueryRow(context.Background(), "SELECT id , full_name , phone_number , birthday , email , gender FROM customers WHERE id = $1 AND is_register = true AND deleted_at IS NULL", customerID).Scan(&customer.ID, &customer.FullName, &customer.PhoneNumber, &customer.Birthday, &customer.Email, &customer.Gender)

	if customer.ID == "" {
		helpers.HandleError(c, 404, "customer not found")
		return
	}

	// bazadan musderinin salgylary alynyar
	rowsCustomerAddress, err := db.Query(context.Background(), "SELECT id , address , is_active FROM customer_address WHERE deleted_at IS NULL AND customer_id = $1", customer.ID)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	var addresses []Address
	for rowsCustomerAddress.Next() {
		var address Address
		if err := rowsCustomerAddress.Scan(&address.ID, &address.Address, &address.IsActive); err != nil {
			helpers.HandleError(c, 400, err.Error())
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

func UpdateCustomerInformation(c *gin.Context) {
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	custID, hasCustomer := c.Get("customer_id")
	if !hasCustomer {
		helpers.HandleError(c, 400, "customer_id is required")
		return
	}
	customerID, ok := custID.(string)
	if !ok {
		helpers.HandleError(c, 400, "customer_id must be string")
		return
	}

	var customer models.Customer
	if err := c.BindJSON(&customer); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	// musderinin  maglumatlaryny uytgetyar
	_, err = db.Exec(context.Background(), "UPDATE customers SET full_name = $1, phone_number = $2, email = $3, birthday = $4 , gender = $6 WHERE id = $5", customer.FullName, customer.PhoneNumber, customer.Email, customer.Birthday, customerID, customer.Gender)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "data successfully updated",
	})
}

func UpdateCustomerPassword(c *gin.Context) {
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	custID, hasCustomer := c.Get("customer_id")
	if !hasCustomer {
		helpers.HandleError(c, 400, "customer_id is required")
		return
	}
	customerID, ok := custID.(string)
	if !ok {
		helpers.HandleError(c, 400, "customer_id must be string")
		return
	}

	password := c.PostForm("password")

	// parol update edilmanka paroly kotlayas
	hashPassword, err := models.HashPassword(password)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	// sonrada musderinin parolyny uytgetyas
	_, err = db.Exec(context.Background(), "UPDATE customers SET password = $1 WHERE id = $2", hashPassword, customerID)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "password of customer successfuly updated",
	})
}

func UpdateCustPassword(c *gin.Context) {
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	phoneNumber := c.PostForm("phone_number")

	// sonrada musderinin parolyny uytgetyas
	var customerID string
	db.QueryRow(context.Background(), "SELECT id FROM customers WHERE phone_number =  $1 AND deleted_at IS NULL AND is_register = true", phoneNumber).Scan(&customerID)

	if customerID == "" {
		helpers.HandleError(c, 404, "customer not found")
		return
	}

	password := c.PostForm("password")
	// parol update edilmanka paroly kotlayas
	hashPassword, err := models.HashPassword(password)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	// sonrada musderinin parolyny uytgetyas
	_, err = db.Exec(context.Background(), "UPDATE customers SET password = $1 WHERE phone_number = $2", hashPassword, phoneNumber)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "password of customer successfuly updated",
	})
}
