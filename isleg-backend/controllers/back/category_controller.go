package controllers

import (
	"github/abbgo/isleg/isleg-backend/config"
	"github/abbgo/isleg/isleg-backend/models"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateCategory(c *gin.Context) {

	var languages []models.Language
	var fileName string

	// GET DATA FROM REQUEST
	isHomeCategoryStr := c.PostForm("is_home_category")
	isHomeCategory, err := strconv.ParseBool(isHomeCategoryStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	parentCategoryID := c.PostForm("parent_category_id")
	parentCategoryIDUUID, err := uuid.Parse(parentCategoryID)
	if parentCategoryID != "" {
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	} else {
		parentCategoryIDUUID = uuid.Nil
	}
	if parentCategoryIDUUID != uuid.Nil {
		_, err := config.ConnDB().Query("SELECT id FROM categories WHERE id = $1", parentCategoryID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}

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
		newFileName := "category" + uuid.New().String() + extension
		fileName = "uploads/" + newFileName
	}

	// VALIDATE DATA
	for _, v := range languages {
		if c.PostForm("name_"+v.NameShort) == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": "name_" + v.NameShort + " is required",
			})
			return
		}
	}
	if parentCategoryIDUUID == uuid.Nil && fileName == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "parent category image is required",
		})
		return
	}
	if parentCategoryIDUUID != uuid.Nil && fileName != "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "child cannot be an image of the category",
		})
		return
	}

	// CREATE CATEGORY
	if parentCategoryIDUUID != uuid.Nil {
		_, err = config.ConnDB().Exec("INSERT INTO categories (parent_category_id,image_path,is_home_category) VALUES ($1,$2,$3)", parentCategoryIDUUID, fileName, isHomeCategory)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	} else {
		_, err = config.ConnDB().Exec("INSERT INTO categories (image_path,is_home_category) VALUES ($1,$2)", fileName, isHomeCategory)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}

	if fileName != "" {
		c.SaveUploadedFile(file, "./"+fileName)
	}

	// GET LAST CATEGORY ID
	lastCategoryID, err := config.ConnDB().Query("SELECT id FROM categories ORDER BY created_at DESC LIMIT 1")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	var categoryID string
	for lastCategoryID.Next() {
		if err := lastCategoryID.Scan(&categoryID); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}

	// CREATE TRANSLATION CATEGORY
	for _, v := range languages {
		_, err := config.ConnDB().Exec("INSERT INTO translation_category (lang_id,category_id,name) VALUES ($1,$2,$3)", v.ID, categoryID, c.PostForm("name_"+v.NameShort))
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
		"message": "category successfully added",
	})

}
