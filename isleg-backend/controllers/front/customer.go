package controllers

import (
	"github/abbgo/isleg/isleg-backend/auth"
	"github/abbgo/isleg/isleg-backend/config"

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
	ID          string    `json:"id"`
	FullName    string    `json:"full_name" binding:"required,min=3"`
	PhoneNumber string    `json:"phone_number" binding:"required,e164,len=12"`
	Birthday    null.Time `json:"birthday"`
	Email       string    `json:"email" binding:"email"`
	IsRegister  bool      `json:"is_register"`
	Addresses   []Address `json:"addresses"`
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
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer func() {
		if err := db.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

	var customer models.Customer

	if err := c.BindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if customer.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "customer password is required",
		})
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

	// On registrasiya bolman haryt sargyt eden adamlar is_register false bolyar ,
	// solar sayta registrasiya boljak bolsa Olaryn taze at familiyasyny telefon belgisini
	// we parolyny baza yazdyrmak ucin database - den sol customer - leri tapyp update etmeli
	rowCustomer, err := db.Query("SELECT phone_number FROM customers WHERE phone_number = $1 AND is_register = false AND deleted_at IS NULL", customer.PhoneNumber)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer func() {
		if err := rowCustomer.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

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

	var customerID string

	if phone_number != "" {

		resultCustomer, err := db.Query("UPDATE customers SET full_name = $1 , password = $2 , email = $3 , is_register = $4 WHERE phone_number = $5 RETURNING id", customer.FullName, hashPassword, customer.Email, true, customer.PhoneNumber)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
		defer func() {
			if err := resultCustomer.Close(); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  false,
					"message": err.Error(),
				})
				return
			}
		}()

		for resultCustomer.Next() {
			if err := resultCustomer.Scan(&customerID); err != nil {
				c.JSON(http.StatusNotFound, gin.H{
					"status":  false,
					"message": err.Error(),
				})
				return
			}
		}

	} else {

		resultCustomers, err := db.Query("INSERT INTO customers (full_name,phone_number,password,email,is_register) VALUES ($1,$2,$3,$4,$5) RETURNING id", customer.FullName, customer.PhoneNumber, hashPassword, customer.Email, true)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
		defer func() {
			if err := resultCustomers.Close(); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  false,
					"message": err.Error(),
				})
				return
			}
		}()

		for resultCustomers.Next() {
			if err := resultCustomers.Scan(&customerID); err != nil {
				c.JSON(http.StatusNotFound, gin.H{
					"status":  false,
					"message": err.Error(),
				})
				return
			}
		}

	}

	accessTokenString, err := auth.GenerateAccessToken(customer.PhoneNumber, customerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		c.Abort()
		return
	}

	refreshTokenString, err := auth.GenerateRefreshToken(customer.PhoneNumber, customerID)
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
	})

}

func LoginCustomer(c *gin.Context) {

	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer func() {
		if err := db.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

	var customer Login

	if err := c.BindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if customer.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "customer password is required",
		})
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
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer func() {
		if err := row.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

	var customerID, oldPassword string

	for row.Next() {
		if err := row.Scan(&customerID, &oldPassword); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}

	if customerID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "this client does not exist",
		})
		return
	}

	credentialError := models.CheckPassword(customer.Password, oldPassword)
	if credentialError != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "invalid credentials",
		})
		return
	}

	accessTokenString, err := auth.GenerateAccessToken(customer.Password, customerID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	refreshTokenString, err := auth.GenerateRefreshToken(customer.PhoneNumber, customerID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
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
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer func() {
		if err := db.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

	custID, hasCustomer := c.Get("customer_id")
	if !hasCustomer {
		c.JSON(http.StatusBadRequest, "customer_id is required")
		return
	}
	customerID, ok := custID.(string)
	if !ok {
		c.JSON(http.StatusBadRequest, "customer_id must be string")
	}

	// bazadan musderinin maglumatlary alynyar
	rowCustomer, err := db.Query("SELECT id , full_name , phone_number , birthday , email FROM customers WHERE id = $1 AND is_register = true AND deleted_at IS NULL", customerID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer func() {
		if err := rowCustomer.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

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

	// bazadan musderinin salgylary alynyar
	rowsCustomerAddress, err := db.Query("SELECT id , address , is_active FROM customer_address WHERE deleted_at IS NULL AND customer_id = $1", customer.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer func() {
		if err := rowsCustomerAddress.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

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

func UpdateCustomerInformation(c *gin.Context) {

	db, err := config.ConnDB()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer func() {
		if err := db.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

	custID, hasCustomer := c.Get("customer_id")
	if !hasCustomer {
		c.JSON(http.StatusBadRequest, "customer_id is required")
		return
	}
	customerID, ok := custID.(string)
	if !ok {
		c.JSON(http.StatusBadRequest, "customer_id must be string")
	}

	var customer models.Customer

	if err := c.BindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	// musderinin  maglumatlaryny uytgetyar
	resultCustomer, err := db.Query("UPDATE customers SET full_name = $1, phone_number = $2, email = $3, birthday = $4 WHERE id = $5", customer.FullName, customer.PhoneNumber, customer.Email, customer.Birthday, customerID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer func() {
		if err := resultCustomer.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "data successfully updated",
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
	defer func() {
		if err := db.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

	custID, hasCustomer := c.Get("customer_id")
	if !hasCustomer {
		c.JSON(http.StatusBadRequest, "customer_id is required")
		return
	}
	customerID, ok := custID.(string)
	if !ok {
		c.JSON(http.StatusBadRequest, "customer_id must be string")
	}

	password := c.PostForm("password")

	// parol update edilmanka paroly kotlayas
	hashPassword, err := models.HashPassword(password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	// sonrada musderinin parolyny uytgetyas
	resultCustomer, err := db.Query("UPDATE customers SET password = $1 WHERE id = $2", hashPassword, customerID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer func() {
		if err := resultCustomer.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "password of customer successfuly updated",
	})

}
