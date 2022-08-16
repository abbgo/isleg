package controllers

import (
	"github/abbgo/isleg/isleg-backend/config"
	"github/abbgo/isleg/isleg-backend/models"
	"net/http"
	"time"

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
}

func CreateTranslationHeader(c *gin.Context) {

	// GET ALL LANGUAGE
	languages, err := GetAllLanguageWithIDAndNameShort()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	dataNames := []string{"research", "phone", "password", "forgot_password", "sign_in", "sign_up", "name", "password_verification", "verify_secure", "my_information", "my_favorites", "my_orders", "log_out", "basket"}

	// VALIDATE DATA
	err = models.ValidateTranslationHeaderCreate(languages, dataNames, c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	// CREATE TRANSLATION HEADER
	for _, v := range languages {
		resultTrHedaer, err := config.ConnDB().Query("INSERT INTO translation_header (lang_id,research,phone,password,forgot_password,sign_in,sign_up,name,password_verification,verify_secure,my_information,my_favorites,my_orders,log_out,basket) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15)", v.ID, c.PostForm("research_"+v.NameShort), c.PostForm("phone_"+v.NameShort), c.PostForm("password_"+v.NameShort), c.PostForm("forgot_password_"+v.NameShort), c.PostForm("sign_in_"+v.NameShort), c.PostForm("sign_up_"+v.NameShort), c.PostForm("name_"+v.NameShort), c.PostForm("password_verification_"+v.NameShort), c.PostForm("verify_secure_"+v.NameShort), c.PostForm("my_information_"+v.NameShort), c.PostForm("my_favorites_"+v.NameShort), c.PostForm("my_orders_"+v.NameShort), c.PostForm("log_out_"+v.NameShort), c.PostForm("basket_"+v.NameShort))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
		defer resultTrHedaer.Close()
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "translation header successfully added",
	})

}

func UpdateTranslationHeaderByID(c *gin.Context) {

	trHead := c.Param("id")

	rowFlag, err := config.ConnDB().Query("SELECT id FROM translation_header WHERE id = $1 AND deleted_at IS NULL", trHead)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer rowFlag.Close()

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

	dataNames := []string{"research", "phone", "password", "forgot_password", "sign_in", "sign_up", "name", "password_verification", "verify_secure", "my_information", "my_favorites", "my_orders", "log_out", "basket"}

	// VALIDATE DATA
	err = models.ValidateTranslationHeaderUpdate(dataNames, c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	currentTime := time.Now()

	resultTrHedaer, err := config.ConnDB().Query("UPDATE translation_header SET research = $1 , phone = $2, password = $3,forgot_password = $4,sign_in = $5,sign_up = $6,name = $7,password_verification = $8,verify_secure = $9, my_information = $10, my_favorites = $11, my_orders = $12, log_out = $13, basket = $14 , updated_at = $16 WHERE id = $15", c.PostForm("research"), c.PostForm("phone"), c.PostForm("password"), c.PostForm("forgot_password"), c.PostForm("sign_in"), c.PostForm("sign_up"), c.PostForm("name"), c.PostForm("password_verification"), c.PostForm("verify_secure"), c.PostForm("my_information"), c.PostForm("my_favorites"), c.PostForm("my_orders"), c.PostForm("log_out"), c.PostForm("basket"), id, currentTime)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer resultTrHedaer.Close()

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "translation header successfully updated",
	})

}

func GetTranslationHeaderByID(c *gin.Context) {

	trHead := c.Param("id")

	rowFlag, err := config.ConnDB().Query("SELECT research,phone,password,forgot_password,sign_in,sign_up,name,password_verification,verify_secure,my_information,my_favorites,my_orders,log_out,basket FROM translation_header WHERE id = $1 AND deleted_at IS NULL", trHead)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer rowFlag.Close()

	var t TranslationHeaderForHeader

	for rowFlag.Next() {
		if err := rowFlag.Scan(&t.Research, &t.Phone, &t.Password, &t.ForgotPassword, &t.SignIn, &t.SignUp, &t.Name, &t.PasswordVerification, &t.VerifySecure, &t.MyInformation, &t.MyFavorites, &t.MyOrders, &t.LogOut, &t.Basket); err != nil {
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

func GetTranslationHeaderForHeader(langID string) (TranslationHeaderForHeader, error) {

	var t TranslationHeaderForHeader

	// GET TranslationHeader For Header
	row, err := config.ConnDB().Query("SELECT research,phone,password,forgot_password,sign_in,sign_up,name,password_verification,verify_secure,my_information,my_favorites,my_orders,log_out,basket FROM translation_header WHERE lang_id = $1 AND deleted_at IS NULL", langID)
	if err != nil {
		return TranslationHeaderForHeader{}, err
	}
	defer row.Close()

	for row.Next() {
		if err := row.Scan(&t.Research, &t.Phone, &t.Password, &t.ForgotPassword, &t.SignIn, &t.SignUp, &t.Name, &t.PasswordVerification, &t.VerifySecure, &t.MyInformation, &t.MyFavorites, &t.MyOrders, &t.LogOut, &t.Basket); err != nil {
			return TranslationHeaderForHeader{}, err
		}
	}

	return t, nil

}
