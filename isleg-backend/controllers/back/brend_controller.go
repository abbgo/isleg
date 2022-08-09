package controllers

import (
	"github/abbgo/isleg/isleg-backend/config"
	"github/abbgo/isleg/isleg-backend/pkg"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type BrendForHomePage struct {
	ID    uuid.UUID `json:"id"`
	Image string    `json:"image"`
}

func CreateBrend(c *gin.Context) {

	// GET DATA FROM REQUEST
	name := c.PostForm("name")

	// VALIDATE DATA
	if name == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "brend name is required",
		})
		return
	}

	// FILE UPLOAD
	newFileName, err := pkg.FileUpload("image", "brend", c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	// CREATE BREND
	_, err = config.ConnDB().Exec("INSERT INTO brends (name,image) VALUES ($1,$2)", name, "uploads/brend/"+newFileName)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "brend successfully added",
	})

}

func UpdateBrend(c *gin.Context) {

	ID := c.Param("id")
	name := c.PostForm("name")
	var fileName string

	rowBrend, err := config.ConnDB().Query("SELECT image FROM brends WHERE id = $1 AND deleted_at IS NULL", ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	var image string

	for rowBrend.Next() {
		if err := rowBrend.Scan(&image); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}

	if image == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "record not found",
		})
		return
	}

	if name == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "name of brend is required",
		})
		return
	}

	file, err := c.FormFile("image")
	if err != nil {
		fileName = image
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
		c.SaveUploadedFile(file, "./uploads/brend/"+newFileName)

		if err := os.Remove("./" + image); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}

		fileName = "uploads/brend/" + newFileName
	}

	_, err = config.ConnDB().Exec("UPDATE brends SET name = $1 , image = $2 WHERE id = $3", name, fileName, ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "brend successfully updated",
	})

}

func GetAllBrendForHomePage() ([]BrendForHomePage, error) {

	var brends []BrendForHomePage

	// get all brends
	rows, err := config.ConnDB().Query("SELECT id,image FROM brends WHERE deleted_at IS NULL")
	if err != nil {
		return []BrendForHomePage{}, err
	}

	for rows.Next() {
		var brend BrendForHomePage
		if err := rows.Scan(&brend.ID, &brend.Image); err != nil {
			return []BrendForHomePage{}, err
		}

		brends = append(brends, brend)
	}

	return brends, nil

}
