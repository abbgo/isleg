package controllers

import (
	"context"
	"github/abbgo/isleg/isleg-backend/config"
	"github/abbgo/isleg/isleg-backend/helpers"
	"github/abbgo/isleg/isleg-backend/models"
	"github/abbgo/isleg/isleg-backend/pkg"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// create new language
func CreateLanguage(c *gin.Context) {

	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	var language models.Language
	if err := c.BindJSON(&language); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	_, err = models.ValidateLanguage(language.NameShort, "create", "")
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	// add language to database , used after_insert_language trigger
	_, err = db.Exec(context.Background(), "INSERT INTO languages (name_short,flag) VALUES ($1,$2)", strings.ToLower(language.NameShort), language.Flag)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "Language added successfully",
	})

}

func UpdateLanguageByID(c *gin.Context) {

	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	// get language id from paramter
	langID := c.Param("id")

	var language models.Language
	if err := c.BindJSON(&language); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	image, err := models.ValidateLanguage(language.NameShort, "update", langID)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	var fileName string
	if language.Flag == "" {
		fileName = image
	} else {
		fileName = language.Flag
	}

	// update language in database
	_, err = db.Exec(context.Background(), "UPDATE languages SET name_short = $1 , flag = $2  WHERE id = $3", strings.ToLower(language.NameShort), fileName, langID)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "language successfully updated",
	})

}

func GetLanguageByID(c *gin.Context) {

	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	// get id of language from parameter
	langID := c.Param("id")

	// get  name_short and flag of language from database
	rowLanguage, err := db.Query(context.Background(), "SELECT id,name_short,flag FROM languages WHERE id = $1 AND deleted_at IS NULL", langID)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	var lang models.Language
	for rowLanguage.Next() {
		if err := rowLanguage.Scan(&lang.ID, &lang.NameShort, &lang.Flag); err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}
	}

	if lang.ID == "" {
		helpers.HandleError(c, 404, "language not found")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":   true,
		"language": lang,
	})

}

func GetLanguages(c *gin.Context) {

	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		if err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}
	}
	defer db.Close()

	var ls []models.Language

	statusQuery := c.DefaultQuery("status", "false")
	status, err := strconv.ParseBool(statusQuery)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	var rowsQuery string
	if !status {
		rowsQuery = `SELECT id,name_short,flag FROM languages WHERE deleted_at IS NULL`
	} else {
		rowsQuery = `SELECT id,name_short,flag FROM languages WHERE deleted_at IS NOT NULL`
	}

	// get name_short,flag of all languages from database
	rows, err := db.Query(context.Background(), rowsQuery)
	if err != nil {
		if err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}
	}

	for rows.Next() {
		var l models.Language
		if err := rows.Scan(&l.ID, &l.NameShort, &l.Flag); err != nil {
			if err != nil {
				helpers.HandleError(c, 400, err.Error())
				return
			}
		}
		ls = append(ls, l)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":    true,
		"languages": ls,
	})

}

func DeleteLanguageByID(c *gin.Context) {

	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	// get id of language from request parameter
	langID := c.Param("id")

	// Check if there is a language, id equal to langID
	rowFlag, err := db.Query(context.Background(), "SELECT id,name_short FROM languages WHERE id = $1 AND deleted_at IS NULL", langID)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	var id, name_short string
	for rowFlag.Next() {
		if err := rowFlag.Scan(&id, &name_short); err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}
	}

	if id == "" {
		helpers.HandleError(c, 404, "language not found")
		return
	}

	if name_short == "tm" {
		helpers.HandleError(c, 400, "You cannot delete it, as it is default language")
		return
	}

	// set current time to deleted_at row of language, used delete_language procedure
	_, err = db.Exec(context.Background(), "CALL delete_language($1)", langID)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "language successfully deleted",
	})

}

func RestoreLanguageByID(c *gin.Context) {

	//initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	// get id of language from request parameter
	langID := c.Param("id")

	// Check if there is a language, id equal to langID
	rowFlag, err := db.Query(context.Background(), "SELECT id FROM languages WHERE id = $1 AND deleted_at IS NOT NULL", langID)
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
		helpers.HandleError(c, 404, "language not found")
		return
	}

	// set null to deleted_at row of language, used restore_language procedure
	_, err = db.Exec(context.Background(), "CALL restore_language($1)", langID)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "language successfully restored",
	})

}

func DeletePermanentlyLanguageByID(c *gin.Context) {

	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	// get id of language from request parameter
	langID := c.Param("id")

	// Check if there is a language, id equal to langID and get image of language from database
	rowFlag, err := db.Query(context.Background(), "SELECT flag,name_short FROM languages WHERE id = $1 AND deleted_at IS NOT NULL", langID)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	var flag, name_short string
	for rowFlag.Next() {
		if err := rowFlag.Scan(&flag, &name_short); err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}
	}

	if flag == "" {
		helpers.HandleError(c, 404, "language not found")
		return
	}

	if name_short == "tm" {
		helpers.HandleError(c, 400, "You cannot delete it, as it is default language")
		return
	}

	// remove image of language
	if err := os.Remove(pkg.ServerPath + flag); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	_, err = db.Exec(context.Background(), "DELETE FROM languages WHERE id = $1", langID)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "language successfully deleted",
	})

}

func GetAllLanguageForHeader() ([]models.Language, error) {

	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		return nil, nil
	}
	defer db.Close()

	var ls []models.Language

	// get name_short,flag of all languages from database
	rows, err := db.Query(context.Background(), "SELECT name_short,flag FROM languages WHERE deleted_at IS NULL")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var l models.Language
		if err := rows.Scan(&l.NameShort, &l.Flag); err != nil {
			return nil, err
		}
		ls = append(ls, l)
	}

	return ls, nil

}

func GetAllLanguageWithIDAndNameShort() ([]models.Language, error) {

	db, err := config.ConnDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	languageRows, err := db.Query(context.Background(), "SELECT id,name_short FROM languages WHERE deleted_at IS NULL ORDER BY created_at ASC")
	if err != nil {
		return []models.Language{}, err
	}

	var languages []models.Language

	for languageRows.Next() {
		var language models.Language
		if err := languageRows.Scan(&language.ID, &language.NameShort); err != nil {
			return []models.Language{}, err
		}
		languages = append(languages, language)
	}

	return languages, nil

}

// GetLangID funksiya berilen langShortName parameter boyunca dilin id - sini getirip beryar
// bu yerde langShortName - dilin gysga ady
func GetLangID(langShortName string) (string, error) {

	db, err := config.ConnDB()
	if err != nil {
		return "", nil
	}
	defer db.Close()

	var langID string

	row, err := db.Query(context.Background(), "SELECT id FROM languages WHERE name_short = $1 AND deleted_at IS NULL", langShortName)
	if err != nil {
		return "", err
	}

	for row.Next() {
		if err := row.Scan(&langID); err != nil {
			return "", err
		}
	}

	return langID, nil

}

// router - daki lang parameter boyunca dilin id - sini getirip beryar
func CheckLanguage(c *gin.Context) (string, error) {

	// GET DATA FROM ROUTE PARAMETER
	langShortName := c.Param("lang")

	// GET ID OFF LANGUAGE
	langID, err := GetLangID(langShortName)
	if err != nil {
		return "", err
	}

	return langID, nil

}
