package controllers

import (
	"context"
	"github/abbgo/isleg/isleg-backend/auth"
	"github/abbgo/isleg/isleg-backend/config"
	"github/abbgo/isleg/isleg-backend/helpers"
	"github/abbgo/isleg/isleg-backend/pkg"
	"strconv"

	"github/abbgo/isleg/isleg-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gopkg.in/guregu/null.v4"
)

type Login struct {
	PhoneNumber string `json:"phone_number" binding:"required,e164,len=12"`
	// Password    string `json:"password" binding:"required,min=3,max=25"`
	Password string `json:"password"`
}

type OTP struct {
	PhoneNumber string `json:"phone_number" binding:"required,e164,len=12"`
	Code        string `json:"code" binding:"required,len=6"`
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

func SendSmsToCustomer(c *gin.Context) {
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

	forRegister := c.DefaultQuery("for_register", "true")
	for_register, err := strconv.ParseBool(forRegister)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	if for_register {
		if customer.Password == "" || len(customer.Password) < 3 || len(customer.Password) > 25 {
			helpers.HandleError(c, 400, "customer password is required and 3 <= length <= 25")
			return
		}

		if err := models.ValidateCustomer(customer.PhoneNumber, "create", ""); err != nil {
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
		if phone_number == "" {
			_, err = db.Exec(context.Background(), "INSERT INTO customers (phone_number,password,is_register) VALUES ($1,$2,$3)", customer.PhoneNumber, hashPassword, false)
			if err != nil {
				helpers.HandleError(c, 400, err.Error())
				return
			}
		}

		seckretKey, err := pkg.SendOTPSmsCode(customer.PhoneNumber)
		if err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}

		_, err = db.Exec(context.Background(), "UPDATE customers SET otp_secret_key = $1 , password = $3 WHERE phone_number = $2", seckretKey, customer.PhoneNumber, hashPassword)
		if err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}
	} else {
		seckretKey, err := pkg.SendOTPSmsCode(customer.PhoneNumber)
		if err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}

		_, err = db.Exec(context.Background(), "UPDATE customers SET otp_secret_key = $1 WHERE phone_number = $2", seckretKey, customer.PhoneNumber)
		if err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"status":       true,
		"phone_number": customer.PhoneNumber,
	})
}

func RegisterCustomer(c *gin.Context) {
	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	var otp OTP
	if err := c.BindJSON(&otp); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	if err := models.ValidateCustomer(otp.PhoneNumber, "create", ""); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	var otpSeckretKey null.String
	db.QueryRow(context.Background(), "SELECT otp_secret_key FROM customers WHERE is_register = false AND phone_number = $1 AND deleted_at IS NULL", otp.PhoneNumber).Scan(&otpSeckretKey)

	if otpSeckretKey.String == "" {
		helpers.HandleError(c, 404, "customer not found")
		return
	}

	if !pkg.ValidateOTPCode(otp.Code, otpSeckretKey.String) {
		helpers.HandleError(c, 400, "invalid credentials")
		return
	}

	var customerID string
	db.QueryRow(context.Background(), "UPDATE customers SET is_register = $1 WHERE phone_number = $2 RETURNING id", true, otp.PhoneNumber).Scan(&customerID)

	accessTokenString, refreshTokenString, err := auth.GenerateTokenForCustomer(otp.PhoneNumber, customerID)
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

	err = models.ValidateCustomer(customer.PhoneNumber, "update", "")
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

	customerID, err := pkg.ValidateMiddlewareData(c, "customer_id")
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
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
	defer rowsCustomerAddress.Close()

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

	customerID, err := pkg.ValidateMiddlewareData(c, "customer_id")
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
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

	customerID, err := pkg.ValidateMiddlewareData(c, "customer_id")
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
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

func CheckOTP(c *gin.Context) {
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	var otp OTP
	if err := c.BindJSON(&otp); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	// sonrada musderinin parolyny uytgetyas
	var customerID string
	var otpSeckretKey null.String
	db.QueryRow(context.Background(), "SELECT id,otp_secret_key FROM customers WHERE phone_number =  $1 AND deleted_at IS NULL AND is_register = true", otp.PhoneNumber).Scan(&customerID, &otpSeckretKey)
	if customerID == "" {
		helpers.HandleError(c, 404, "customer not found")
		return
	}

	if !pkg.ValidateOTPCode(otp.Code, otpSeckretKey.String) {
		helpers.HandleError(c, 400, "invalid credentials")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "success",
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
