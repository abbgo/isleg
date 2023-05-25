package controllers

import (
	"github/abbgo/isleg/isleg-backend/config"
	"github/abbgo/isleg/isleg-backend/models"
	"github/abbgo/isleg/isleg-backend/pkg"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

// create new language
func CreateLanguage(c *gin.Context) {

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

	var language models.Language
	if err := c.BindJSON(&language); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	_, err = models.ValidateLanguage(language.NameShort, "create", "")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	// add language to database , used after_insert_language trigger
	resultLang, err := db.Query("INSERT INTO languages (name_short,flag) VALUES ($1,$2)", strings.ToLower(language.NameShort), language.Flag)
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
		"message": "Language added successfully",
	})

}

func UpdateLanguageByID(c *gin.Context) {

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

	// get language id from paramter
	langID := c.Param("id")

	var language models.Language
	if err := c.BindJSON(&language); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	image, err := models.ValidateLanguage(language.NameShort, "update", langID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	var fileName string
	if language.Flag == "" {
		fileName = image
	} else {
		fileName = language.Flag
	}

	// update language in database
	resultLang, err := db.Query("UPDATE languages SET name_short = $1 , flag = $2  WHERE id = $3", strings.ToLower(language.NameShort), fileName, langID)
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

	// get id of language from parameter
	langID := c.Param("id")

	// get  name_short and flag of language from database
	rowLanguage, err := db.Query("SELECT id,name_short,flag FROM languages WHERE id = $1 AND deleted_at IS NULL", langID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer func() {
		if err := rowLanguage.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

	var lang models.Language

	for rowLanguage.Next() {
		if err := rowLanguage.Scan(&lang.ID, &lang.NameShort, &lang.Flag); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}

	if lang.ID == "" {
		c.JSON(http.StatusNotFound, gin.H{
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

	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}
	defer func() {
		if err := db.Close(); err != nil {
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  false,
					"message": err.Error(),
				})
				return
			}
		}
	}()

	var ls []models.Language

	// get name_short,flag of all languages from database
	rows, err := db.Query("SELECT id,name_short,flag FROM languages WHERE deleted_at IS NULL")
	if err != nil {
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}
	defer func() {
		if err := rows.Close(); err != nil {
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  false,
					"message": err.Error(),
				})
				return
			}
		}
	}()

	for rows.Next() {
		var l models.Language
		if err := rows.Scan(&l.ID, &l.NameShort, &l.Flag); err != nil {
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  false,
					"message": err.Error(),
				})
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

	// get id of language from request parameter
	langID := c.Param("id")

	// Check if there is a language, id equal to langID
	rowFlag, err := db.Query("SELECT id,name_short FROM languages WHERE id = $1 AND deleted_at IS NULL", langID)
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

	var id, name_short string

	for rowFlag.Next() {
		if err := rowFlag.Scan(&id, &name_short); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}

	if id == "" {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  false,
			"message": "language not found",
		})
		return
	}

	if name_short == "tm" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "You cannot delete it, as it is default language",
		})
		return
	}

	// set current time to deleted_at row of language, used delete_language procedure
	resultPROC, err := db.Query("CALL delete_language($1)", langID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer func() {
		if err := resultPROC.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "language successfully deleted",
	})

}

func RestoreLanguageByID(c *gin.Context) {

	//initialize database connection
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

	// get id of language from request parameter
	langID := c.Param("id")

	// Check if there is a language, id equal to langID
	rowFlag, err := db.Query("SELECT id FROM languages WHERE id = $1 AND deleted_at IS NOT NULL", langID)
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
		c.JSON(http.StatusNotFound, gin.H{
			"status":  false,
			"message": "language not found",
		})
		return
	}

	// set null to deleted_at row of language, used restore_language procedure
	resultPROC, err := db.Query("CALL restore_language($1)", langID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer func() {
		if err := resultPROC.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "language successfully restored",
	})

}

func DeletePermanentlyLanguageByID(c *gin.Context) {

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

	// get id of language from request parameter
	langID := c.Param("id")

	// Check if there is a language, id equal to langID and get image of language from database
	rowFlag, err := db.Query("SELECT flag,name_short FROM languages WHERE id = $1 AND deleted_at IS NOT NULL", langID)
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

	var flag, name_short string

	for rowFlag.Next() {
		if err := rowFlag.Scan(&flag, &name_short); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}

	if flag == "" {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  false,
			"message": "language not found",
		})
		return
	}

	if name_short == "tm" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "You cannot delete it, as it is default language",
		})
		return
	}

	// remove image of language
	if err := os.Remove(pkg.ServerPath + flag); err != nil {
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
		"message": "language successfully deleted",
	})

}

func GetAllLanguageForHeader() ([]models.Language, error) {

	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		return nil, nil
	}
	defer func() ([]models.Language, error) {
		if err := db.Close(); err != nil {
			return nil, err
		}
		return nil, err
	}()

	var ls []models.Language

	// get name_short,flag of all languages from database
	rows, err := db.Query("SELECT name_short,flag FROM languages WHERE deleted_at IS NULL")
	if err != nil {
		return nil, err
	}
	defer func() ([]models.Language, error) {
		if err := rows.Close(); err != nil {
			return nil, err
		}
		return nil, err
	}()

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
	defer func() ([]models.Language, error) {
		if err := db.Close(); err != nil {
			return nil, err
		}
		return nil, err
	}()

	languageRows, err := db.Query("SELECT id,name_short FROM languages WHERE deleted_at IS NULL ORDER BY created_at ASC")
	if err != nil {
		return []models.Language{}, err
	}
	defer func() ([]models.Language, error) {
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

// GetLangID funksiya berilen langShortName parameter boyunca dilin id - sini getirip beryar
// bu yerde langShortName - dilin gysga ady
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
