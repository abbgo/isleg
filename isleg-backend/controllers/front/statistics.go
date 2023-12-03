package controllers

import (
	"context"
	"github/abbgo/isleg/isleg-backend/config"
	"github/abbgo/isleg/isleg-backend/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetCountOfCustomers(c *gin.Context) {
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	var count uint
	db.QueryRow(context.Background(), "SELECT count FROM count_of_customers WHERE day = CURRENT_DATE").Scan(&count)
	if count == 0 {
		_, err := db.Exec(context.Background(), "INSERT INTO count_of_customers (count,day) VALUES (1,CURRENT_DATE)")
		if err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}
	} else {
		_, err := db.Exec(context.Background(), "UPDATE count_of_customers SET count = count + 1 WHERE day = CURRENT_DATE")
		if err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "success",
	})

}
