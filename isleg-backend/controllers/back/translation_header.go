package controllers

import (
	"github/abbgo/isleg/isleg-backend/config"
	"github/abbgo/isleg/isleg-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateTranslationHeader(c *gin.Context) {

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

	var trHeaders []models.TranslationHeader

	if err := c.BindJSON(&trHeaders); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	// check lang_id
	for _, v := range trHeaders {

		rowLang, err := db.Query("SELECT id FROM languages WHERE id = $1 AND deleted_at IS NULL", v.LangID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
		defer func() {
			if err := rowLang.Close(); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  false,
					"message": err.Error(),
				})
				return
			}
		}()

		var langID string

		for rowLang.Next() {
			if err := rowLang.Scan(&langID); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  false,
					"message": err.Error(),
				})
				return
			}
		}

		if langID == "" {
			c.JSON(http.StatusNotFound, gin.H{
				"status":  false,
				"message": "language not found",
			})
			return
		}

	}

	// add data to translation_header table
	for _, v := range trHeaders {

		resultTrHedaer, err := db.Query("INSERT INTO translation_header (lang_id,research,phone,password,forgot_password,sign_in,sign_up,name,password_verification,verify_secure,my_information,my_favorites,my_orders,log_out,basket,email,add_to_basket) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17)", v.LangID, v.Research, v.Phone, v.Password, v.ForgotPassword, v.SignIn, v.SignUp, v.Name, v.PasswordVerification, v.VerifySecure, v.MyInformation, v.MyFavorites, v.MyOrders, v.LogOut, v.Basket, v.Email, v.AddToBasket)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
		defer func() {
			if err := resultTrHedaer.Close(); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  false,
					"message": err.Error(),
				})
				return
			}
		}()

	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "data added successfully",
	})

}

func UpdateTranslationHeaderByID(c *gin.Context) {

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

	// get id of translation_header table from request parameter
	var trHeader models.TranslationHeader

	if err := c.BindJSON(&trHeader); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	// trHead := c.Param("id")

	// check id
	rowFlag, err := db.Query("SELECT id FROM translation_header WHERE id = $1 AND deleted_at IS NULL", trHeader.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer func() {
		if err := rowFlag.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

	var id string

	for rowFlag.Next() {
		if err := rowFlag.Scan(&id); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}

	if id == "" {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  false,
			"message": "record not found",
		})
		return
	}

	// update translation_header table data
	resultTrHedaer, err := db.Query("UPDATE translation_header SET research = $1 , phone = $2, password = $3, forgot_password = $4, sign_in = $5, sign_up = $6, name = $7, password_verification = $8, verify_secure = $9, my_information = $10, my_favorites = $11, my_orders = $12, log_out = $13, basket = $14 , email = $16, add_to_basket = $17, lang_id = $18 WHERE id = $15", trHeader.Research, trHeader.Phone, trHeader.Password, trHeader.ForgotPassword, trHeader.SignIn, trHeader.SignUp, trHeader.Name, trHeader.PasswordVerification, trHeader.VerifySecure, trHeader.MyInformation, trHeader.MyFavorites, trHeader.MyOrders, trHeader.LogOut, trHeader.Basket, trHeader.ID, trHeader.Email, trHeader.AddToBasket, trHeader.LangID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer func() {
		if err := resultTrHedaer.Close(); err != nil {
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

func GetTranslationHeaderByID(c *gin.Context) {

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

	// get id of translation header from request parameter
	trHead := c.Param("id")

	// check id and get data
	rowFlag, err := db.Query("SELECT id,research,phone,password,forgot_password,sign_in,sign_up,name,password_verification,verify_secure,my_information,my_favorites,my_orders,log_out,basket,email,add_to_basket FROM translation_header WHERE id = $1 AND deleted_at IS NULL", trHead)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer func() {
		if err := rowFlag.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

	var t models.TranslationHeader

	for rowFlag.Next() {
		if err := rowFlag.Scan(&t.ID, &t.Research, &t.Phone, &t.Password, &t.ForgotPassword, &t.SignIn, &t.SignUp, &t.Name, &t.PasswordVerification, &t.VerifySecure, &t.MyInformation, &t.MyFavorites, &t.MyOrders, &t.LogOut, &t.Basket, &t.Email, &t.AddToBasket); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}

	if t.ID == "" {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  false,
			"message": "record not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":             true,
		"translation_header": t,
	})

}

func GetTranslationHeaderForHeader(langID string) (models.TranslationHeader, error) {

	db, err := config.ConnDB()
	if err != nil {
		return models.TranslationHeader{}, err
	}
	defer func() (models.TranslationHeader, error) {
		if err := db.Close(); err != nil {
			return models.TranslationHeader{}, err
		}
		return models.TranslationHeader{}, nil
	}()

	var t models.TranslationHeader

	// GET TranslationHeader For Header
	row, err := db.Query("SELECT research,phone,password,forgot_password,sign_in,sign_up,name,password_verification,verify_secure,my_information,my_favorites,my_orders,log_out,basket,email,add_to_basket FROM translation_header WHERE lang_id = $1 AND deleted_at IS NULL", langID)
	if err != nil {
		return models.TranslationHeader{}, err
	}
	defer func() (models.TranslationHeader, error) {
		if err := row.Close(); err != nil {
			return models.TranslationHeader{}, err
		}
		return models.TranslationHeader{}, nil
	}()

	for row.Next() {
		if err := row.Scan(&t.Research, &t.Phone, &t.Password, &t.ForgotPassword, &t.SignIn, &t.SignUp, &t.Name, &t.PasswordVerification, &t.VerifySecure, &t.MyInformation, &t.MyFavorites, &t.MyOrders, &t.LogOut, &t.Basket, &t.Email, &t.AddToBasket); err != nil {
			return models.TranslationHeader{}, err
		}
	}

	return t, nil

}
