package controllers

import (
	"github/abbgo/isleg/isleg-backend/config"
	"github/abbgo/isleg/isleg-backend/models"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateAfisa(c *gin.Context) {

	var languages []models.Language
	var fileName string

	// GET ALL LANGUAGE
	languageRows, err := config.ConnDB().Query("SELECT id,name_short FROM languages ORDER BY created_at ASC")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

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

	// FILE UPLOAD
	file, errFile := c.FormFile("image_path")
	if errFile != nil {
		fileName = ""
	} else {
		extension := filepath.Ext(file.Filename)
		// VALIDATE IMAGE
		if extension != ".jpg" && extension != ".jpeg" && extension != ".png" && extension != ".gif" {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": "the file must be an image.",
			})
			return
		}

		newFileName := "afisa" + uuid.New().String() + extension
		fileName = "uploads/" + newFileName
	}

	// VALIDATE DATA
	for _, v := range languages {
		if c.PostForm("title_"+v.NameShort) == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": "title_" + v.NameShort + " is required",
			})
			return
		}
	}

	for _, v := range languages {
		if c.PostForm("description_"+v.NameShort) == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": "description_" + v.NameShort + " is required",
			})
			return
		}
	}

	// create afisa
	_, err = config.ConnDB().Exec("INSERT INTO afisa (image_path) VALUES ($1)", fileName)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	if fileName != "" {
		c.SaveUploadedFile(file, "./"+fileName)
	}

	// get last afisa id
	lastAfisaID, err := config.ConnDB().Query("SELECT id FROM afisa ORDER BY created_at DESC LIMIT 1")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	var afisaID string

	for lastAfisaID.Next() {
		if err := lastAfisaID.Scan(&afisaID); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}

	// create translation afisa
	for _, v := range languages {
		_, err := config.ConnDB().Exec("INSERT INTO translation_afisa (afisa_id,lang_id,title,description) VALUES ($1,$2,$3,$4)", afisaID, v.ID, c.PostForm("title_"+v.NameShort), c.PostForm("description"+v.NameShort))
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
		"message": "afisa successfully added",
	})

}
