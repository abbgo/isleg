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
	languageRows, err := config.ConnDB().Query("SELECT id,name_short FROM languages ORDER BY created_at ASC")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	var languages []models.Language

	for languageRows.Next() {
		var language models.Language
		if err := languageRows.Scan(&language.ID, &language.NameShort); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
		languages = append(languages, language)
	}

	// VALIDATE DATA
	for _, v := range languages {
		if c.PostForm("address_"+v.NameShort) == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": "address_" + v.NameShort + " is required",
			})
			return
		}
	}

	for _, v := range languages {
		if c.PostForm("birthday_"+v.NameShort) == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": "birthday_" + v.NameShort + " is required",
			})
			return
		}
	}

	for _, v := range languages {
		if c.PostForm("update_password_"+v.NameShort) == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": "update_password_" + v.NameShort + " is required",
			})
			return
		}
	}

	for _, v := range languages {
		if c.PostForm("save_"+v.NameShort) == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": "save_" + v.NameShort + " is required",
			})
			return
		}
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
