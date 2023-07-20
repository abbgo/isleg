package controllers

import (
	"github/abbgo/isleg/isleg-backend/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CustomerAddress struct {
	ID       string `json:"id"`
	Address  string `json:"address"`
	IsActive bool   `json:"is_active"`
}

func GetCustomerAddresses(c *gin.Context) {

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
		return
	}

	// musderinin salgylaryny alyas bazadan
	rowsAddress, err := db.Query("SELECT id,address,is_active FROM customer_address WHERE customer_id = $1 AND deleted_at IS NULL", customerID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer func() {
		if err := rowsAddress.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

	var addresses []CustomerAddress

	for rowsAddress.Next() {

		var address CustomerAddress

		if err := rowsAddress.Scan(&address.ID, &address.Address, &address.IsActive); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}

		addresses = append(addresses, address)

	}

	c.JSON(http.StatusOK, gin.H{
		"status":             true,
		"customer_addresses": addresses,
	})

}

func UpdateCustomerAddressStatus(c *gin.Context) {

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

	addressID := c.Param("id")

	custID, hasCustomer := c.Get("customer_id")
	if !hasCustomer {
		c.JSON(http.StatusBadRequest, "customer_id is required")
		return
	}
	customerID, ok := custID.(string)
	if !ok {
		c.JSON(http.StatusBadRequest, "customer_id must be string")
		return
	}

	// musderide frontdan gelen id - li salgy barmy ya-da yokmy sony barlayas
	rowCusomerAddress, err := db.Query("SELECT id FROM customer_address WHERE id = $1 AND deleted_at IS NULL", addressID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer func() {
		if err := rowCusomerAddress.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

	var address_id string

	for rowCusomerAddress.Next() {
		if err := rowCusomerAddress.Scan(&address_id); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}

	// eger salgy yok bolsa yzyna yalnyslyk goyberyas
	if address_id == "" {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  false,
			"message": "address of customer not found",
		})
		return
	}

	// salgy bar bolsada sol salgynyn statusyny update etyas
	resultCustomerAddress, err := db.Query("UPDATE customer_address SET is_active = true WHERE id = $1", addressID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer func() {
		if err := resultCustomerAddress.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

	// gelen salgydan basga ahli salgylaryn statusyny false etyas
	resultCustAddressIsActive, err := db.Query("UPDATE customer_address SET is_active = false WHERE id != $1 AND customer_id = $2", addressID, customerID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer func() {
		if err := resultCustAddressIsActive.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "address status successfuly updated",
	})

}

func AddAddressToCustomer(c *gin.Context) {

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
		return
	}

	address := c.PostForm("address")
	if address == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "address is required",
		})
		return
	}

	// gosuljak bolyan salgy sol musderide on barmy ya-da yokmy sony barlayas
	rowCusomerAddress, err := db.Query("SELECT id FROM customer_address WHERE address = $1 AND customer_id = $2 AND deleted_at IS NULL", address, customerID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer func() {
		if err := rowCusomerAddress.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

	var address_id string

	for rowCusomerAddress.Next() {
		if err := rowCusomerAddress.Scan(&address_id); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}

	// eger salgy on bar bolsa onda yzyna musdera duydurys goyberyas
	if address_id != "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "address already exists in this customer",
		})
		return
	}

	// eger salgy on sol musderide yok bolsa , onda sol musderinin salgylarynyn arasyna gosyas
	resultCustomerAddress, err := db.Query("INSERT INTO customer_address (customer_id,address,is_active) VALUES ($1,$2,true) RETURNING id", customerID, address)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer func() {
		if err := resultCustomerAddress.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

	var addresID string
	for resultCustomerAddress.Next() {
		if err := resultCustomerAddress.Scan(&addresID); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}

	// gosulan salgyny aktiwe edip musdera degisli beyleki ahli salgylaryn statusynyn false etyas
	resultCustAddressStatus, err := db.Query("UPDATE customer_address SET is_active = false WHERE address != $1 AND customer_id = $2", address, customerID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer func() {
		if err := resultCustAddressStatus.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "customer address successfully created",
		"id":      addresID,
	})

}

func DeleteCustomerAddress(c *gin.Context) {

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
		return
	}

	addressID := c.Param("id")

	rowCusomerAddress, err := db.Query("SELECT id FROM customer_address WHERE id = $1 AND customer_id = $2 AND deleted_at IS NULL", addressID, customerID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer func() {
		if err := rowCusomerAddress.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

	var address_id string

	for rowCusomerAddress.Next() {
		if err := rowCusomerAddress.Scan(&address_id); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}

	if address_id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "address not found in this customer",
		})
		return
	}

	resultCustomerAddress, err := db.Query("DELETE FROM customer_address WHERE id = $1", addressID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer func() {
		if err := resultCustomerAddress.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "customer address successfully deleted",
	})

}
