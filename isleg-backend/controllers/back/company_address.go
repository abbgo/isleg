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

		rowLang, err := db.Query(context.Background(), "SELECT id FROM languages WHERE id = $1 AND deleted_at IS NULL", v.LangID)
		if err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}

		var langID string
		for rowLang.Next() {
			if err := rowLang.Scan(&langID); err != nil {
				helpers.HandleError(c, 400, err.Error())
				return
			}
		}

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
	rowComAddress, err := db.Query(context.Background(), "SELECT id FROM company_address WHERE id = $1 AND deleted_at IS NULL", companyAddress.ID)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	var id string
	for rowComAddress.Next() {
		if err := rowComAddress.Scan(&id); err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}
	}

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
	rowComAddress, err := db.Query(context.Background(), "SELECT address FROM company_address WHERE id = $1 AND deleted_at IS NULL", ID)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	var adress string
	for rowComAddress.Next() {
		if err := rowComAddress.Scan(&adress); err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}
	}

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

	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	langID, err := CheckLanguage(c)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	// get company address where lang_id equal langID
	addressRow, err := db.Query(context.Background(), "SELECT address FROM company_address WHERE lang_id = $1 AND deleted_at IS NULL", langID)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	var address string
	for addressRow.Next() {
		if err := addressRow.Scan(&address); err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"status":          true,
		"company_address": address,
	})

}
