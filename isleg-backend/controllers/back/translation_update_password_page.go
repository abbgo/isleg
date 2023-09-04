package controllers

import (
	"context"
	"github/abbgo/isleg/isleg-backend/config"
	"github/abbgo/isleg/isleg-backend/helpers"
	"github/abbgo/isleg/isleg-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateTranslationUpdatePasswordPage(c *gin.Context) {

	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	// get data from request
	var trUpdPassPages []models.TranslationUpdatePasswordPage

	if err := c.BindJSON(&trUpdPassPages); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	// check lang_id
	for _, v := range trUpdPassPages {

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

	// add data in database
	for _, v := range trUpdPassPages {
		_, err := db.Exec(context.Background(), "INSERT INTO translation_update_password_page (lang_id,title,password,verify_password,explanation,save) VALUES ($1,$2,$3,$4,$5,$6)", v.LangID, v.Title, v.Password, v.VerifyPassword, v.Explanation, v.Save)
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

func UpdateTranslationUpdatePasswordPageByID(c *gin.Context) {

	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	// get id of translation update password page from request data
	var trUpdPassPage models.TranslationUpdatePasswordPage

	if err := c.BindJSON(&trUpdPassPage); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	// check id
	rowFlag, err := db.Query(context.Background(), "SELECT id FROM translation_update_password_page WHERE id = $1 AND deleted_at IS NULL", trUpdPassPage.ID)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	var id string
	for rowFlag.Next() {
		if err := rowFlag.Scan(&id); err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}
	}

	if id == "" {
		helpers.HandleError(c, 404, "record not found")
		return
	}

	_, err = db.Exec(context.Background(), "UPDATE translation_update_password_page SET title = $1, verify_password = $2 , explanation = $3 , save = $4 , password = $5, lang_id = $7 WHERE id = $6", trUpdPassPage.Title, trUpdPassPage.VerifyPassword, trUpdPassPage.Explanation, trUpdPassPage.Save, trUpdPassPage.Password, trUpdPassPage.ID, trUpdPassPage.LangID)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "data successfully updated",
	})

}

func GetTranslationUpdatePasswordPageByID(c *gin.Context) {

	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	// get id from request parameter
	ID := c.Param("id")

	// check id and get data
	rowFlag, err := db.Query(context.Background(), "SELECT id,title,verify_password,explanation,save,password FROM translation_update_password_page WHERE id = $1 AND deleted_at IS NULL", ID)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	var t models.TranslationUpdatePasswordPage
	for rowFlag.Next() {
		if err := rowFlag.Scan(&t.ID, &t.Title, &t.VerifyPassword, &t.Explanation, &t.Save, &t.Password); err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}
	}

	if t.ID == "" {
		helpers.HandleError(c, 404, "record not found")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":                           true,
		"translation_update_password_page": t,
	})

}

func GetTranslationUpdatePasswordPageByLangID(c *gin.Context) {

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

	// get translation-update-password-page where lang_id equal langID
	aboutRow, err := db.Query(context.Background(), "SELECT title,password,verify_password,explanation,save FROM translation_update_password_page WHERE lang_id = $1 AND deleted_at IS NULL", langID)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	var trUpdatePasswordPage models.TranslationUpdatePasswordPage
	for aboutRow.Next() {
		if err := aboutRow.Scan(&trUpdatePasswordPage.Title, &trUpdatePasswordPage.Password, &trUpdatePasswordPage.VerifyPassword, &trUpdatePasswordPage.Explanation, &trUpdatePasswordPage.Save); err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}
	}

	if trUpdatePasswordPage.Title == "" {
		helpers.HandleError(c, 404, "record not found")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":                           true,
		"translation_update_password_page": trUpdatePasswordPage,
	})

}
