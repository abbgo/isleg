package controllers

import (
	"github/abbgo/isleg/isleg-backend/config"
	"github/abbgo/isleg/isleg-backend/models"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateAfisa(c *gin.Context) {

	db, err := config.ConnDB()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer db.Close()

	var fileName string

	// GET ALL LANGUAGE
	languages, err := GetAllLanguageWithIDAndNameShort()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	// FILE UPLOAD
	file, errFile := c.FormFile("image")
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

		newFileName := uuid.New().String() + extension
		fileName = "uploads/afisa/" + newFileName
	}

	dataNames := []string{"title", "description"}

	// VALIDATE DATA
	if err = models.ValidateAfisaData(languages, dataNames, c); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	// create afisa
	resultAFisa, err := db.Query("INSERT INTO afisa (image) VALUES ($1)", fileName)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer resultAFisa.Close()

	if fileName != "" {
		c.SaveUploadedFile(file, "./"+fileName)
	}

	// get id of added afisa
	lastAfisaID, err := db.Query("SELECT id FROM afisa WHERE deleted_at IS NULL ORDER BY created_at DESC LIMIT 1")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer lastAfisaID.Close()

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
		resultTRAfisa, err := db.Query("INSERT INTO translation_afisa (afisa_id,lang_id,title,description) VALUES ($1,$2,$3,$4)", afisaID, v.ID, c.PostForm("title_"+v.NameShort), c.PostForm("description"+v.NameShort))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
		defer resultTRAfisa.Close()
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "afisa successfully added",
	})

}

func UpdateAfisaByID(c *gin.Context) {

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
	var fileName string

	rowAfisa, err := db.Query("SELECT id,image FROM afisa WHERE id = $1 AND deleted_at IS NULL", ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer rowAfisa.Close()

	var afisaID, image string

	for rowAfisa.Next() {
		if err := rowAfisa.Scan(&afisaID, &image); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}

	if afisaID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "record not found",
		})
		return
	}

	languages, err := GetAllLanguageWithIDAndNameShort()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	dataNames := []string{"title", "description"}

	// VALIDATE DATA
	if err = models.ValidateAfisaData(languages, dataNames, c); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	// FILE UPLOAD
	file, errFile := c.FormFile("image")
	if errFile != nil {
		fileName = image
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

		newFileName := uuid.New().String() + extension
		fileName = "uploads/afisa/" + newFileName

		if image != "" {
			if err := os.Remove("./" + image); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  false,
					"message": err.Error(),
				})
				return
			}
		}
	}

	currentTime := time.Now()

	resultAfisa, err := db.Query("UPDATE afisa SET image = $1 , updated_at = $2 WHERE id = $3", fileName, currentTime, ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer resultAfisa.Close()

	if fileName != "" {
		c.SaveUploadedFile(file, "./"+fileName)
	}

	for _, v := range languages {
		resultTRAfisa, err := db.Query("UPDATE translation_afisa SET title = $1 , description = $2 , updated_at = $3 WHERE afisa_id = $4 AND lang_id = $5", c.PostForm("title_"+v.NameShort), c.PostForm("description_"+v.NameShort), currentTime, ID, v.ID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
		defer resultTRAfisa.Close()
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "afisa successfully updated",
	})

}
