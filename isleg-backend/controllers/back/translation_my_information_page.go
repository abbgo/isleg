package controllers

import (
	"github/abbgo/isleg/isleg-backend/config"
	"github/abbgo/isleg/isleg-backend/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateTranslationMyInformationPage(c *gin.Context) {

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
		resutlTRMyInfPage, err := db.Query("INSERT INTO translation_my_information_page (lang_id,address,birthday,update_password,save) VALUES ($1,$2,$3,$4,$5)", v.ID, c.PostForm("address_"+v.NameShort), c.PostForm("birthday_"+v.NameShort), c.PostForm("update_password_"+v.NameShort), c.PostForm("save_"+v.NameShort))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
		defer func() {
			if err := resutlTRMyInfPage.Close(); err != nil {
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
		"message": "translation my information page successfully added",
	})

}

func UpdateTranslationMyInformationPageByID(c *gin.Context) {

	db, err := config.ConnDB()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer db.Close()

	ID := c.Param("id")

	rowFlag, err := db.Query("SELECT id FROM translation_my_information_page WHERE id = $1 AND deleted_at IS NULL", ID)
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

	dataNames := []string{"address", "birthday", "update_password", "save"}

	// VALIDATE DATA
	err = models.ValidateTranslationMyInformationPageUpdate(dataNames, c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	currentTime := time.Now()

	resultTRMyInfPage, err := db.Query("UPDATE translation_my_information_page SET address = $1, birthday = $2 , update_password = $3, save = $4 , updated_at = $6 WHERE id = $5", c.PostForm("address"), c.PostForm("birthday"), c.PostForm("update_password"), c.PostForm("save"), id, currentTime)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer resultTRMyInfPage.Close()

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "translation_my_information_page successfully updated",
	})

}

func GetTranslationMyInformationPageByID(c *gin.Context) {

	db, err := config.ConnDB()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer db.Close()

	ID := c.Param("id")

	rowFlag, err := db.Query("SELECT address,birthday,update_password,save FROM translation_my_information_page WHERE id = $1 AND deleted_at IS NULL", ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer rowFlag.Close()

	var t TrMyInformationPage

	for rowFlag.Next() {
		if err := rowFlag.Scan(&t.Address, &t.Birthday, &t.UpdatePassword, &t.Save); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}

	if t.Address == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "record not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":                          true,
		"translation_my_information_page": t,
	})

}

func GetTranslationMyInformationPageByLangID(c *gin.Context) {

	db, err := config.ConnDB()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer db.Close()

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
	aboutRow, err := db.Query("SELECT address,birthday,update_password,save FROM translation_my_information_page WHERE lang_id = $1 AND deleted_at IS NULL", langID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer aboutRow.Close()

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

	if trMyInformationPage.Address == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "record not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":                          true,
		"translation_my_information_page": trMyInformationPage,
	})

}
