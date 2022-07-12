package controllers

import (
	"github/abbgo/isleg/isleg-backend/config"
	"github/abbgo/isleg/isleg-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TrMyInformationPage struct {
	Birthday       string `json:"birthday"`
	Address        string `json:"address"`
	UpdatePassword string `json:"update_password"`
	Save           string `json:"save"`
}

func CreateTranslationMyInformationPage(c *gin.Context) {

	// GET ALL LANGUAGE
	languages, err := GetAllLanguageWithIDAndNameShort()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	dataNames := []string{"address", "birthday", "update_password", "save"}

	// VALIDATE DATA
	if err = models.ValidateTranslationMyInformationPageData(languages, dataNames, c); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	// create translation_my_information_page
	for _, v := range languages {
		_, err := config.ConnDB().Exec("INSERT INTO translation_my_information_page (lang_id,address,birthday,update_password,save) VALUES ($1,$2,$3,$4,$5)", v.ID, c.PostForm("address_"+v.NameShort), c.PostForm("birthday_"+v.NameShort), c.PostForm("update_password_"+v.NameShort), c.PostForm("save_"+v.NameShort))
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
		"message": "translation my information page successfully added",
	})

}

func GetTranslationMyInformationPage(c *gin.Context) {

	// GET DATA FROM ROUTE PARAMETER
	langShortName := c.Param("lang")

	// GET language id
	langID, err := GetLangID(langShortName)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	// get translation-my-information-page where lang_id equal langID
	aboutRow, err := config.ConnDB().Query("SELECT address,birthday,update_password,save FROM translation_my_information_page WHERE lang_id = $1", langID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	var trMyInformationPage TrMyInformationPage

	for aboutRow.Next() {
		if err := aboutRow.Scan(&trMyInformationPage.Address, &trMyInformationPage.Birthday, &trMyInformationPage.UpdatePassword, &trMyInformationPage.Save); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"status":                          true,
		"translation_my_information_page": trMyInformationPage,
	})

}
