package controllers

import (
	"github/abbgo/isleg/isleg-backend/config"
	"github/abbgo/isleg/isleg-backend/models"
	"github/abbgo/isleg/isleg-backend/pkg"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type LanguageForHeader struct {
	NameShort string `json:"name_short"`
	Flag      string `json:"flag"`
}

func CreateLanguage(c *gin.Context) {

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

	// GET DATA FROM REQUEST
	nameShort := c.PostForm("name_short")

	// VALIDATE DATA
	if nameShort == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "language name_short is required",
		})
		return
	}

	// FILE UPLOAD
	newFileName, err := pkg.FileUpload("flag", "language", c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	// CREATE LANGUAGE
	resultLang, err := db.Query("INSERT INTO languages (name_short,flag) VALUES ($1,$2)", strings.ToLower(nameShort), "uploads/language/"+newFileName)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer func() {
		if err := resultLang.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "language successfully added",
	})

}

func UpdateLanguageByID(c *gin.Context) {

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

	langID := c.Param("id")
	nameShort := c.PostForm("name_short")
	var fileName string

	rowFlag, err := db.Query("SELECT flag FROM languages WHERE id = $1 AND deleted_at IS NULL", langID)
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

	var flag string

	for rowFlag.Next() {
		if err := rowFlag.Scan(&flag); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}

	if flag == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "language not found",
		})
		return
	}

	if nameShort == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "language name_short is required",
		})
		return
	}

	file, err := c.FormFile("flag")
	if err != nil {
		fileName = flag
	} else {
		extensionFile := filepath.Ext(file.Filename)

		if extensionFile != ".jpg" && extensionFile != ".jpeg" && extensionFile != ".png" && extensionFile != ".gif" {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": "the file must be an image",
			})
			return
		}

		newFileName := uuid.New().String() + extensionFile
		c.SaveUploadedFile(file, "./uploads/language/"+newFileName)

		if err := os.Remove("./" + flag); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}

		fileName = "uploads/language/" + newFileName
	}

	currentTime := time.Now()

	resultLang, err := db.Query("UPDATE languages SET name_short = $1 , flag = $2 , updated_at = $4 WHERE id = $3", nameShort, fileName, langID, currentTime)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer func() {
		if err := resultLang.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "language successfully updated",
	})

}

func GetLanguageByID(c *gin.Context) {

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

	langID := c.Param("id")

	rowLanguage, err := db.Query("SELECT name_short,flag FROM languages WHERE id = $1 AND deleted_at IS NULL", langID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer rowLanguage.Close()

	var lang LanguageForHeader

	for rowLanguage.Next() {
		if err := rowLanguage.Scan(&lang.NameShort, &lang.Flag); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}

	if lang.NameShort == "" || lang.Flag == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "language not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":   true,
		"language": lang,
	})

}

func GetLanguages(c *gin.Context) {
	languages, err := GetAllLanguageForHeader()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":   true,
		"language": languages,
	})

}

func DeleteLanguageByID(c *gin.Context) {

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

	langID := c.Param("id")

	rowFlag, err := db.Query("SELECT flag FROM languages WHERE id = $1 AND deleted_at IS NULL", langID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer rowFlag.Close()

	var flag string

	for rowFlag.Next() {
		if err := rowFlag.Scan(&flag); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}

	if flag == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "language not found",
		})
		return
	}

	currentTime := time.Now()

	resutlTRLand, err := db.Query("UPDATE languages SET deleted_at = $1 WHERE id = $2", currentTime, langID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer resutlTRLand.Close()

	resultPROC, err := db.Query("CALL after_delete_language($1)", langID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer resultPROC.Close()

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "language successfully deleted",
	})

}

func RestoreLanguageByID(c *gin.Context) {

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

	langID := c.Param("id")

	rowFlag, err := db.Query("SELECT flag FROM languages WHERE id = $1 AND deleted_at IS NOT NULL", langID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer rowFlag.Close()

	var flag string

	for rowFlag.Next() {
		if err := rowFlag.Scan(&flag); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}

	if flag == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "language not found",
		})
		return
	}

	resultLang, err := db.Query("UPDATE languages SET deleted_at = NULL WHERE id = $1", langID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer resultLang.Close()

	resultPROC, err := db.Query("CALL after_restore_language($1)", langID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer resultPROC.Close()

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "language successfully restored",
	})

}

func DeletePermanentlyLanguageByID(c *gin.Context) {

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

	langID := c.Param("id")

	rowFlag, err := db.Query("SELECT flag FROM languages WHERE id = $1 AND deleted_at IS NOT NULL", langID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer rowFlag.Close()

	var flag string

	for rowFlag.Next() {
		if err := rowFlag.Scan(&flag); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}

	if flag == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "language not found",
		})
		return
	}

	if err := os.Remove("./" + flag); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	resultLang, err := db.Query("DELETE FROM languages WHERE id = $1", langID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer resultLang.Close()

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "language successfully deleted",
	})

}

func GetAllLanguageForHeader() ([]LanguageForHeader, error) {

	db, err := config.ConnDB()
	if err != nil {
		return nil, nil
	}
	defer func() ([]LanguageForHeader, error) {
		if err := db.Close(); err != nil {
			return nil, err
		}
		return nil, err
	}()

	var ls []LanguageForHeader

	// GET Language For Header
	rows, err := db.Query("SELECT name_short,flag FROM languages WHERE deleted_at IS NULL")
	if err != nil {
		return nil, err
	}
	defer func() ([]LanguageForHeader, error) {
		if err := rows.Close(); err != nil {
			return nil, err
		}
		return nil, err
	}()

	for rows.Next() {
		var l LanguageForHeader
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
	defer func() ([]LanguageForHeader, error) {
		if err := db.Close(); err != nil {
			return nil, err
		}
		return nil, err
	}()

	languageRows, err := db.Query("SELECT id,name_short FROM languages WHERE deleted_at IS NULL ORDER BY created_at ASC")
	if err != nil {
		return []models.Language{}, err
	}
	defer func() ([]LanguageForHeader, error) {
		if err := languageRows.Close(); err != nil {
			return nil, err
		}
		return nil, err
	}()

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

func GetLangID(langShortName string) (string, error) {

	db, err := config.ConnDB()
	if err != nil {
		return "", nil
	}
	defer func() (string, error) {
		if err := db.Close(); err != nil {
			return "", err
		}
		return "", nil
	}()

	var langID string

	row, err := db.Query("SELECT id FROM languages WHERE name_short = $1 AND deleted_at IS NULL", langShortName)
	if err != nil {
		return "", err
	}
	defer func() (string, error) {
		if err := row.Close(); err != nil {
			return "", err
		}
		return "", nil
	}()

	for row.Next() {
		if err := row.Scan(&langID); err != nil {
			return "", err
		}
	}

	return langID, nil

}

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
