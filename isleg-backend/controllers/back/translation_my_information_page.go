package controllers

import (
	"context"
	"github/abbgo/isleg/isleg-backend/config"
	"github/abbgo/isleg/isleg-backend/helpers"
	"github/abbgo/isleg/isleg-backend/models"
	"github/abbgo/isleg/isleg-backend/pkg"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateTranslationMyInformationPage(c *gin.Context) {
	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	//get data from request
	var trMyInforPages []models.TranslationMyInformationPage
	if err := c.BindJSON(&trMyInforPages); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	// check lang_id
	for _, v := range trMyInforPages {
		var langID string
		db.QueryRow(context.Background(), "SELECT id FROM languages WHERE id = $1 AND deleted_at IS NULL", v.LangID).Scan(&langID)

		if langID == "" {
			helpers.HandleError(c, 404, "language not found")
			return
		}
	}

	// create translation_my_information_page
	for _, v := range trMyInforPages {
		_, err := db.Exec(context.Background(), "INSERT INTO translation_my_information_page (lang_id,address,birthday,update_password,save,gender,male,female) VALUES ($1,$2,$3,$4,$5,$6,$7,$8)", v.LangID, v.Address, v.Birthday, v.UpdatePassword, v.Save, v.Gender, v.Male, v.Female)
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

func UpdateTranslationMyInformationPageByID(c *gin.Context) {
	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	// get id of translation my information page from request parameter
	var trMyInforPage models.TranslationMyInformationPage
	if err := c.BindJSON(&trMyInforPage); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	// check id
	var id string
	db.QueryRow(context.Background(), "SELECT id FROM translation_my_information_page WHERE id = $1 AND deleted_at IS NULL", trMyInforPage.ID).Scan(&id)

	if id == "" {
		helpers.HandleError(c, 404, "record not found")
		return
	}

	_, err = db.Exec(context.Background(), "UPDATE translation_my_information_page SET address = $1, birthday = $2 , update_password = $3, save = $4, lang_id = $6, gender = $7, male = $8, female = $9 WHERE id = $5", trMyInforPage.Address, trMyInforPage.Birthday, trMyInforPage.UpdatePassword, trMyInforPage.Save, trMyInforPage.ID, trMyInforPage.LangID, trMyInforPage.Gender, trMyInforPage.Male, trMyInforPage.Female)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "data successfully updated",
	})
}

func GetTranslationMyInformationPageByID(c *gin.Context) {
	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	// get id of translation my information page from request parameter
	ID := c.Param("id")

	// check id and get data
	var t models.TranslationMyInformationPage
	db.QueryRow(context.Background(), "SELECT id,address,birthday,update_password,save,gender,male,female FROM translation_my_information_page WHERE id = $1 AND deleted_at IS NULL", ID).Scan(&t.ID, &t.Address, &t.Birthday, &t.UpdatePassword, &t.Save, &t.Gender, &t.Male, &t.Female)

	if t.ID == "" {
		helpers.HandleError(c, 404, "record not found")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":                          true,
		"translation_my_information_page": t,
	})
}

func GetTranslationMyInformationPageByLangID(c *gin.Context) {
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

	// get translation-my-information-page where lang_id equal langID
	var trMyInformationPage models.TranslationMyInformationPage
	db.QueryRow(context.Background(), "SELECT address,birthday,update_password,save,gender,male,female FROM translation_my_information_page WHERE lang_id = $1 AND deleted_at IS NULL", langID).Scan(&trMyInformationPage.Address, &trMyInformationPage.Birthday, &trMyInformationPage.UpdatePassword, &trMyInformationPage.Save, &trMyInformationPage.Gender, &trMyInformationPage.Male, &trMyInformationPage.Female)

	if trMyInformationPage.Address == "" {
		helpers.HandleError(c, 404, "record not found")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":                          true,
		"translation_my_information_page": trMyInformationPage,
	})
}
