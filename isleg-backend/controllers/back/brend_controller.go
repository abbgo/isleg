package controllers

import (
	"github/abbgo/isleg/isleg-backend/config"
	"github/abbgo/isleg/isleg-backend/pkg"
	"net/http"

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
