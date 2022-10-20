package controllers

import (
	"github/abbgo/isleg/isleg-backend/config"
	"github/abbgo/isleg/isleg-backend/models"
	"github/abbgo/isleg/isleg-backend/pkg"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TranslationHeaderForHeader struct {
	Research             string `json:"research"`
	Phone                string `json:"phone"`
	Password             string `json:"password"`
	ForgotPassword       string `json:"forgot_password"`
	SignIn               string `json:"sign_in"`
	SignUp               string `json:"sign_up"`
	Name                 string `json:"name"`
	PasswordVerification string `json:"password_verification"`
	VerifySecure         string `json:"verify_secure"`
	MyInformation        string `json:"my_information"`
	MyFavorites          string `json:"my_favorites"`
	MyOrders             string `json:"my_orders"`
	LogOut               string `json:"log_out"`
	Basket               string `json:"basket"`
	Email                string `json:"email"`
	AddToBasket          string `json:"add_to_basket"`
}

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

	// GET ALL LANGUAGE
	languages, err := GetAllLanguageWithIDAndNameShort()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	dataNames := []string{"research", "phone", "password", "forgot_password", "sign_in", "sign_up", "name", "password_verification", "verify_secure", "my_information", "my_favorites", "my_orders", "log_out", "basket", "email", "add_to_basket"}

	// VALIDATE DATA
	err = pkg.ValidateTranslations(languages, dataNames, c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	// add data to translation_header table
	for _, v := range languages {
		resultTrHedaer, err := db.Query("INSERT INTO translation_header (lang_id,research,phone,password,forgot_password,sign_in,sign_up,name,password_verification,verify_secure,my_information,my_favorites,my_orders,log_out,basket,email,add_to_basket) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17)", v.ID, c.PostForm("research_"+v.NameShort), c.PostForm("phone_"+v.NameShort), c.PostForm("password_"+v.NameShort), c.PostForm("forgot_password_"+v.NameShort), c.PostForm("sign_in_"+v.NameShort), c.PostForm("sign_up_"+v.NameShort), c.PostForm("name_"+v.NameShort), c.PostForm("password_verification_"+v.NameShort), c.PostForm("verify_secure_"+v.NameShort), c.PostForm("my_information_"+v.NameShort), c.PostForm("my_favorites_"+v.NameShort), c.PostForm("my_orders_"+v.NameShort), c.PostForm("log_out_"+v.NameShort), c.PostForm("basket_"+v.NameShort), c.PostForm("email_"+v.NameShort), c.PostForm("add_to_basket_"+v.NameShort))
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
	trHead := c.Param("id")

	// check id
	rowFlag, err := db.Query("SELECT id FROM translation_header WHERE id = $1 AND deleted_at IS NULL", trHead)
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
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "record not found",
		})
		return
	}

	dataNames := []string{"research", "phone", "password", "forgot_password", "sign_in", "sign_up", "name", "password_verification", "verify_secure", "my_information", "my_favorites", "my_orders", "log_out", "basket", "email", "add_to_basket"}

	// VALIDATE DATA
	err = pkg.ValidateTranslationsForUpdate(dataNames, c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	// update translation_header table data
	resultTrHedaer, err := db.Query("UPDATE translation_header SET research = $1 , phone = $2, password = $3,forgot_password = $4,sign_in = $5,sign_up = $6,name = $7,password_verification = $8,verify_secure = $9, my_information = $10, my_favorites = $11, my_orders = $12, log_out = $13, basket = $14 , email = $16, add_to_basket = $17 WHERE id = $15", c.PostForm("research"), c.PostForm("phone"), c.PostForm("password"), c.PostForm("forgot_password"), c.PostForm("sign_in"), c.PostForm("sign_up"), c.PostForm("name"), c.PostForm("password_verification"), c.PostForm("verify_secure"), c.PostForm("my_information"), c.PostForm("my_favorites"), c.PostForm("my_orders"), c.PostForm("log_out"), c.PostForm("basket"), id, c.PostForm("email"), c.PostForm("add_to_basket"))
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

	trHead := c.Param("id")

	rowFlag, err := db.Query("SELECT research,phone,password,forgot_password,sign_in,sign_up,name,password_verification,verify_secure,my_information,my_favorites,my_orders,log_out,basket,email,add_to_basket FROM translation_header WHERE id = $1 AND deleted_at IS NULL", trHead)
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
		if err := rowFlag.Scan(&t.Research, &t.Phone, &t.Password, &t.ForgotPassword, &t.SignIn, &t.SignUp, &t.Name, &t.PasswordVerification, &t.VerifySecure, &t.MyInformation, &t.MyFavorites, &t.MyOrders, &t.LogOut, &t.Basket, &t.Email, &t.AddToBasket); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}

	if t.Research == "" {
		c.JSON(http.StatusBadRequest, gin.H{
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
