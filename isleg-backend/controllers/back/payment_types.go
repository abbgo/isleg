package controllers

import (
	"context"
	"github/abbgo/isleg/isleg-backend/config"
	"github/abbgo/isleg/isleg-backend/helpers"
	"github/abbgo/isleg/isleg-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreatePaymentType(c *gin.Context) {
	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	// get data from request
	var paymentTypes []models.PaymentTypes
	if err := c.BindJSON(&paymentTypes); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	// check lang_id
	for _, v := range paymentTypes {
		var langID string
		if err := db.QueryRow(context.Background(), "SELECT id FROM languages WHERE id = $1 AND deleted_at IS NULL", v.LangID).Scan(&langID); err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}

		if langID == "" {
			helpers.HandleError(c, 404, "language not found")
			return
		}
	}

	// create company address
	for _, v := range paymentTypes {
		_, err := db.Exec(context.Background(), "INSERT INTO payment_types (lang_id,type) VALUES ($1,$2)", v.LangID, v.Type)
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

func UpdatePaymentTypeByID(c *gin.Context) {
	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	// get data from request
	var paymentType models.PaymentTypes
	if err := c.BindJSON(&paymentType); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	// check id
	var id string
	if err := db.QueryRow(context.Background(), "SELECT id FROM payment_types WHERE id = $1 AND deleted_at IS NULL", paymentType.ID).Scan(&id); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	if id == "" {
		helpers.HandleError(c, 404, "record not found")
		return
	}

	_, err = db.Exec(context.Background(), "UPDATE payment_types SET type = $1, lang_id = $3 WHERE id = $2", paymentType.Type, paymentType.ID, paymentType.LangID)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "data successfully updated",
	})
}

func GetPaymentTypeByID(c *gin.Context) {
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
	var paymentType string
	if err := db.QueryRow(context.Background(), "SELECT type FROM payment_types WHERE id = $1 AND deleted_at IS NULL", ID).Scan(&paymentType); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	if paymentType == "" {
		helpers.HandleError(c, 404, "record not found")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":       true,
		"payment_type": paymentType,
	})
}

func GetPaymentTypes(c *gin.Context) {
	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	// get data from database
	rowsPaymentType, err := db.Query(context.Background(), "SELECT lang_id,type FROM payment_types WHERE deleted_at IS NULL")
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	var paymentTypes []models.PaymentTypes
	for rowsPaymentType.Next() {
		var paymentType models.PaymentTypes
		if err := rowsPaymentType.Scan(&paymentType.LangID, &paymentType.Type); err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}
		paymentTypes = append(paymentTypes, paymentType)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":        true,
		"payment_types": paymentTypes,
	})
}

func GetPaymentTypesByLangID(c *gin.Context) {
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	// GET DATA FROM ROUTE PARAMETER
	langShortName := c.Param("lang")

	// GET language id
	langID, err := GetLangID(langShortName)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	rowsPaymentType, err := db.Query(context.Background(), "SELECT type FROM payment_types WHERE lang_id = $1 AND deleted_at IS NULL", langID)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	var paymentTypes []string
	for rowsPaymentType.Next() {
		var paymentType string
		if err := rowsPaymentType.Scan(&paymentType); err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}
		paymentTypes = append(paymentTypes, paymentType)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":        true,
		"payment_types": paymentTypes,
	})
}
