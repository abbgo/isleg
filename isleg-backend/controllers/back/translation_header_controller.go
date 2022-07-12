package controllers

import (
	"github/abbgo/isleg/isleg-backend/config"
	"github/abbgo/isleg/isleg-backend/models"
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

	dataNames := []string{"research", "phone", "password", "forgot_password", "sign_in", "sign_up", "name", "password_verification", "verify_secure", "my_information", "my_favorites", "my_orders", "log_out"}

	// VALIDATE DATA
	err = models.ValidateTranslationHeaderData(languages, dataNames, c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	// CREATE TRANSLATION HEADER
	for _, v := range languages {
		_, err := config.ConnDB().Exec("INSERT INTO translation_header (lang_id,research,phone,password,forgot_password,sign_in,sign_up,name,password_verification,verify_secure,my_information,my_favorites,my_orders,log_out) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14)", v.ID, c.PostForm("research_"+v.NameShort), c.PostForm("phone_"+v.NameShort), c.PostForm("password_"+v.NameShort), c.PostForm("forgot_password_"+v.NameShort), c.PostForm("sign_in_"+v.NameShort), c.PostForm("sign_up_"+v.NameShort), c.PostForm("name_"+v.NameShort), c.PostForm("password_verification_"+v.NameShort), c.PostForm("verify_secure_"+v.NameShort), c.PostForm("my_information_"+v.NameShort), c.PostForm("my_favorites_"+v.NameShort), c.PostForm("my_orders_"+v.NameShort), c.PostForm("log_out_"+v.NameShort))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "translation header successfully added",
	})

}

func GetTranslationHeaderForHeader(langID string) (TranslationHeaderForHeader, error) {

	var t TranslationHeaderForHeader

	// GET TranslationHeader For Header
	row, err := config.ConnDB().Query("SELECT research,phone,password,forgot_password,sign_in,sign_up,name,password_verification,verify_secure,my_information,my_favorites,my_orders,log_out FROM translation_header WHERE lang_id = $1", langID)
	if err != nil {
		return TranslationHeaderForHeader{}, err
	}
	for row.Next() {
		if err := row.Scan(&t.Research, &t.Phone, &t.Password, &t.ForgotPassword, &t.SignIn, &t.SignUp, &t.Name, &t.PasswordVerification, &t.VerifySecure, &t.MyInformation, &t.MyFavorites, &t.MyOrders, &t.LogOut); err != nil {
			return TranslationHeaderForHeader{}, err
		}
	}

	return t, nil

}
