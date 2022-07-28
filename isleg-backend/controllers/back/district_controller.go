package controllers

import (
	"github/abbgo/isleg/isleg-backend/config"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateDistrict(c *gin.Context) {

	// GET ALL LANGUAGE
	languages, err := GetAllLanguageWithIDAndNameShort()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	// validate data from request
	priceStr := c.PostForm("price")
	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	for _, v := range languages {
		if c.PostForm("name_"+v.NameShort) == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": "name_" + v.NameShort + " is required",
			})
			return
		}
	}

	// create district
	_, err = config.ConnDB().Exec("INSERT INTO district (price) VALUES ($1)", price)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	// get id off added district
	lastDistrictID, err := config.ConnDB().Query("SELECT id FROM district ORDER BY created_at DESC LIMIT 1")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	var districtID string

	for lastDistrictID.Next() {
		if err := lastDistrictID.Scan(&districtID); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}

	// create translation afisa
	for _, v := range languages {
		_, err := config.ConnDB().Exec("INSERT INTO translation_district (lang_id,district_id,name) VALUES ($1,$2,$3)", v.ID, districtID, c.PostForm("name_"+v.NameShort))
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
		"message": "district successfully added",
	})

}
