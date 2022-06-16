package controllers

import (
	"github/abbgo/isleg/isleg-backend/config"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

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
	file, err := c.FormFile("image_path")
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
	newFileName := "brend" + uuid.New().String() + extension
	c.SaveUploadedFile(file, "./uploads/"+newFileName)

	// CREATE BREND
	_, err = config.ConnDB().Exec("INSERT INTO brends (name,image_path) VALUES ($1,$2)", name, "uploads/"+newFileName)
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
