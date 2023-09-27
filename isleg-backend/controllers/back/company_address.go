package controllers

import (
	"context"
	"github/abbgo/isleg/isleg-backend/config"
	"github/abbgo/isleg/isleg-backend/helpers"
	"github/abbgo/isleg/isleg-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateCompanyAddress(c *gin.Context) {
	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	// get data from request
	var companyAddresses []models.CompanyAddress
	if err := c.BindJSON(&companyAddresses); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	// check lans_id
	for _, v := range companyAddresses {
		var langID string
		db.QueryRow(context.Background(), "SELECT id FROM languages WHERE id = $1 AND deleted_at IS NULL", v.LangID).Scan(&langID)
		if langID == "" {
			helpers.HandleError(c, 404, "language not found")
			return
		}
	}

	// create company address
	for _, v := range companyAddresses {
		_, err := db.Exec(context.Background(), "INSERT INTO company_address (lang_id,address) VALUES ($1,$2)", v.LangID, v.Address)
		if err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "data successfully added",
	})
}

func UpdateCompanyAddressByID(c *gin.Context) {
	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	// get id from request parameter
	var companyAddress models.CompanyAddress

	//check id
	var id string
	db.QueryRow(context.Background(), "SELECT id FROM company_address WHERE id = $1 AND deleted_at IS NULL", companyAddress.ID).Scan(&id)
	if id == "" {
		helpers.HandleError(c, 404, "record not found")
		return
	}

	_, err = db.Exec(context.Background(), "UPDATE company_address SET address = $1, lang_id = $2 WHERE id = $3", companyAddress.Address, companyAddress.LangID, companyAddress.ID)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "data successfully updated",
	})
}

func GetCompanyAddressByID(c *gin.Context) {
	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	// get id from request parameter
	ID := c.Param("id")

	// check id and get data from database
	var adress string
	db.QueryRow(context.Background(), "SELECT address FROM company_address WHERE id = $1 AND deleted_at IS NULL", ID).Scan(&adress)
	if adress == "" {
		helpers.HandleError(c, 404, "record not found")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":          true,
		"company_address": adress,
	})
}

// GetCompanyAddress funksiya dil boyunca firmanyn salgysyny getirip beryar
func GetCompanyAddress(c *gin.Context) {
	langID, err := pkg.ValidateMiddlewareData(c, "lang_id")
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	// get company address where lang_id equal langID
	var address string
	db.QueryRow(context.Background(), "SELECT address FROM company_address WHERE lang_id = $1 AND deleted_at IS NULL", langID).Scan(&address)
	c.JSON(http.StatusOK, gin.H{
		"status":          true,
		"company_address": address,
	})

}
