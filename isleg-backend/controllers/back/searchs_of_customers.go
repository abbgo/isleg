package controllers

import (
	"context"
	"github/abbgo/isleg/isleg-backend/config"
	"github/abbgo/isleg/isleg-backend/helpers"
	"github/abbgo/isleg/isleg-backend/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetSearchsOfCustomers(c *gin.Context) {
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	hasProducts := c.DefaultQuery("has_products", "true")
	has_products, err := strconv.ParseBool(hasProducts)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	rowsSearchs, err := db.Query(context.Background(), "SELECT search,count FROM searchs_of_customers WHRE deleted_at IS NULL AND has_products = $1", has_products)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer rowsSearchs.Close()

	var searchs []models.SearchsOfCustomers
	for rowsSearchs.Next() {
		var search models.SearchsOfCustomers
		if err := rowsSearchs.Scan(&search.Search, &search.Count); err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}
		searchs = append(searchs, search)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"searchs": searchs,
	})
}
