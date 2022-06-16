package controllers

import (
	"github/abbgo/isleg/isleg-backend/config"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type LanguageForHeader struct {
	NameShort string `json:"name_short"`
	Flag      string `json:"flag"`
}

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

	// CREATE TRANSLATION FOOTER
	_, err = config.ConnDB().Exec("INSERT INTO translation_footer (lang_id) VALUES ($1)", langID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	// GET ALL CATEGORY id
	var categoryIDs []string
	categoryRows, err := config.ConnDB().Query("SELECT id FROM categories ORDER BY created_at ASC")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	for categoryRows.Next() {
		var categoryID string
		if err := categoryRows.Scan(&categoryID); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
		categoryIDs = append(categoryIDs, categoryID)
	}

	// CREATE TRANSLATION CATEGORY
	for _, v := range categoryIDs {
		_, err = config.ConnDB().Exec("INSERT INTO translation_category (lang_id,category_id) VALUES ($1,$2)", langID, v)
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
		"message": "language successfully added",
	})

}

func GetAllLanguageForHeader() ([]LanguageForHeader, error) {

	var ls []LanguageForHeader

	// GET Language For Header
	rows, err := config.ConnDB().Query("SELECT name_short,flag_path FROM languages")
	if err != nil {
		return []LanguageForHeader{}, err
	}
	for rows.Next() {
		var l LanguageForHeader
		if err := rows.Scan(&l.NameShort, &l.Flag); err != nil {
			return []LanguageForHeader{}, err
		}
		ls = append(ls, l)
	}

	return ls, nil

}
