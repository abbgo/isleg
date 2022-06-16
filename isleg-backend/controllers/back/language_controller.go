package controllers

import (
	"github/abbgo/isleg/isleg-backend/config"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateLanguage(c *gin.Context) {

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
	file, err := c.FormFile("flag_path")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	extension := filepath.Ext(file.Filename)
	// VALIDATE IMAGE
	if extension != ".jpg" && extension != ".jpeg" && extension != ".png" && extension != ".gif" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "the file must be an image.",
		})
		return
	}
	newFileName := "language" + uuid.New().String() + extension
	c.SaveUploadedFile(file, "./uploads/"+newFileName)

	// CREATE LANGUAGE
	_, err = config.ConnDB().Exec("INSERT INTO languages (name_short,flag_path) VALUES ($1,$2)", strings.ToLower(nameShort), "uploads/"+newFileName)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	// GET LAST LANGUAGE ID
	lastLandID, err := config.ConnDB().Query("SELECT id FROM languages ORDER BY created_at DESC LIMIT 1")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	var langID string
	for lastLandID.Next() {
		if err := lastLandID.Scan(&langID); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}

	// CREATE TRANSLATION HEADER
	_, err = config.ConnDB().Exec("INSERT INTO translation_header (lang_id) VALUES ($1)", langID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "language successfully added",
	})

}
