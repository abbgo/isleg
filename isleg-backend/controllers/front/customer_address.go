package controllers

import (
	"context"
	"github/abbgo/isleg/isleg-backend/config"
	"github/abbgo/isleg/isleg-backend/helpers"
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

	// musderinin salgylaryny alyas bazadan
	rowsAddress, err := db.Query(context.Background(), "SELECT id,address,is_active FROM customer_address WHERE customer_id = $1 AND deleted_at IS NULL", customerID)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	var addresses []CustomerAddress
	for rowsAddress.Next() {
		var address CustomerAddress
		if err := rowsAddress.Scan(&address.ID, &address.Address, &address.IsActive); err != nil {
			helpers.HandleError(c, 400, err.Error())
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
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	addressID := c.Param("id")

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

	// musderide frontdan gelen id - li salgy barmy ya-da yokmy sony barlayas
	var address_id string
	db.QueryRow(context.Background(), "SELECT id FROM customer_address WHERE id = $1 AND deleted_at IS NULL", addressID).Scan(&address_id)

	// eger salgy yok bolsa yzyna yalnyslyk goyberyas
	if address_id == "" {
		helpers.HandleError(c, 404, "address of customer not found")
		return
	}

	// salgy bar bolsada sol salgynyn statusyny update etyas
	_, err = db.Exec(context.Background(), "UPDATE customer_address SET is_active = true WHERE id = $1", addressID)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	// gelen salgydan basga ahli salgylaryn statusyny false etyas
	_, err = db.Exec(context.Background(), "UPDATE customer_address SET is_active = false WHERE id != $1 AND customer_id = $2", addressID, customerID)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "address status successfuly updated",
	})
}

func AddAddressToCustomer(c *gin.Context) {
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

	address := c.PostForm("address")
	if address == "" {
		helpers.HandleError(c, 400, "address is required")
		return
	}

	// gosuljak bolyan salgy sol musderide on barmy ya-da yokmy sony barlayas
	var address_id string
	db.QueryRow(context.Background(), "SELECT id FROM customer_address WHERE address = $1 AND customer_id = $2 AND deleted_at IS NULL", address, customerID).Scan(&address_id)

	// eger salgy on bar bolsa onda yzyna musdera duydurys goyberyas
	if address_id != "" {
		helpers.HandleError(c, 400, "address already exists in this customer")
		return
	}

	// eger salgy on sol musderide yok bolsa , onda sol musderinin salgylarynyn arasyna gosyas
	var addresID string
	db.QueryRow(context.Background(), "INSERT INTO customer_address (customer_id,address,is_active) VALUES ($1,$2,true) RETURNING id", customerID, address).Scan(&addresID)

	// gosulan salgyny aktiwe edip musdera degisli beyleki ahli salgylaryn statusynyn false etyas
	_, err = db.Exec(context.Background(), "UPDATE customer_address SET is_active = false WHERE address != $1 AND customer_id = $2", address, customerID)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "customer address successfully created",
		"id":      addresID,
	})
}

func DeleteCustomerAddress(c *gin.Context) {
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

	addressID := c.Param("id")

	var address_id string
	db.QueryRow(context.Background(), "SELECT id FROM customer_address WHERE id = $1 AND customer_id = $2 AND deleted_at IS NULL", addressID, customerID).Scan(&address_id)

	if address_id == "" {
		helpers.HandleError(c, 404, "address not found in this customer")
		return
	}

	_, err = db.Exec(context.Background(), "DELETE FROM customer_address WHERE id = $1", addressID)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "customer address successfully deleted",
	})
}
