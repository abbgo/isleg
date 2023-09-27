package controllers

import (
	"context"
	"github/abbgo/isleg/isleg-backend/config"
	"github/abbgo/isleg/isleg-backend/helpers"
	"github/abbgo/isleg/isleg-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetSearchsOfCustomers(c *gin.Context) {
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	rowsSearchs, err := db.Query(context.Background(), "SELECT search FROM searchs_of_customers WHRE deleted_at IS NULL")
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer rowsSearchs.Close()

	var searchs []models.SearchsOfCustomers
	for rowsSearchs.Next() {
		var search models.SearchsOfCustomers
		if err := rowsSearchs.Scan(&search.Search); err != nil {
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
