package controllers

import (
	"github/abbgo/isleg/isleg-backend/config"
	"github/abbgo/isleg/isleg-backend/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type TrUpdatePasswordPage struct {
	Title          string `json:"title"`
	Password       string `json:"password"`
	VerifyPassword string `json:"verify_password"`
	Explanation    string `json:"explanation"`
	Save           string `json:"save"`
}

func CreateTranslationUpdatePasswordPage(c *gin.Context) {

	db, err := config.ConnDB()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer db.Close()

	// GET ALL LANGUAGE
	languages, err := GetAllLanguageWithIDAndNameShort()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	dataNames := []string{"title", "password", "verify_password", "explanation", "save"}

	// VALIDATE DATA
	if err = models.ValidateTranslationUpdatePasswordPageData(languages, dataNames, c); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	for _, v := range languages {
		result, err := db.Query("INSERT INTO translation_update_password_page (lang_id,title,password,verify_password,explanation,save) VALUES ($1,$2,$3,$4,$5,$6)", v.ID, c.PostForm("title_"+v.NameShort), c.PostForm("password_"+v.NameShort), c.PostForm("verify_password_"+v.NameShort), c.PostForm("explanation_"+v.NameShort), c.PostForm("save_"+v.NameShort))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
		defer result.Close()
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "translation update password page successfully added",
	})

}

func UpdateTranslationUpdatePasswordPageByID(c *gin.Context) {

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

	rowFlag, err := db.Query("SELECT id FROM translation_update_password_page WHERE id = $1 AND deleted_at IS NULL", ID)
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

	dataNames := []string{"title", "password", "verify_password", "explanation", "save"}

	// VALIDATE DATA
	err = models.ValidateTranslationUpdatePasswordPageUpdate(dataNames, c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	currentTime := time.Now()

	result, err := db.Query("UPDATE translation_update_password_page SET title = $1, verify_password = $2 , explanation = $3 , save = $4 , password = $5 , updated_at = $7 WHERE id = $6", c.PostForm("title"), c.PostForm("verify_password"), c.PostForm("explanation"), c.PostForm("save"), c.PostForm("password"), id, currentTime)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer result.Close()

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "translation_update_password_page successfully updated",
	})

}

func GetTranslationUpdatePasswordPageByID(c *gin.Context) {

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

	rowFlag, err := db.Query("SELECT title,verify_password,explanation,save,password FROM translation_update_password_page WHERE id = $1 AND deleted_at IS NULL", ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer rowFlag.Close()

	var t TrUpdatePasswordPage

	for rowFlag.Next() {
		if err := rowFlag.Scan(&t.Title, &t.VerifyPassword, &t.Explanation, &t.Save, &t.Password); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}

	if t.Title == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "record not found",
		})
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
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer db.Close()

	langID, err := CheckLanguage(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	// get translation-update-password-page where lang_id equal langID
	aboutRow, err := db.Query("SELECT title,password,verify_password,explanation,save FROM translation_update_password_page WHERE lang_id = $1 AND deleted_at IS NULL", langID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer aboutRow.Close()

	var trUpdatePasswordPage TrUpdatePasswordPage

	for aboutRow.Next() {
		if err := aboutRow.Scan(&trUpdatePasswordPage.Title, &trUpdatePasswordPage.Password, &trUpdatePasswordPage.VerifyPassword, &trUpdatePasswordPage.Explanation, &trUpdatePasswordPage.Save); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}

	if trUpdatePasswordPage.Title == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "record not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":                           true,
		"translation_update_password_page": trUpdatePasswordPage,
	})

}
