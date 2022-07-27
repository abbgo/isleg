package controllers

import (
	"github/abbgo/isleg/isleg-backend/config"
	"github/abbgo/isleg/isleg-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateShop(c *gin.Context) {

	ownerName := c.PostForm("owner_name")
	address := c.PostForm("address")
	phoneNumber := c.PostForm("phone_number")
	runningTime := c.PostForm("running_time")
	categories, _ := c.GetPostFormArray("category_id")

	if err := models.ValidateShopData(ownerName, address, phoneNumber, runningTime, categories); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	_, err := config.ConnDB().Exec("INSERT INTO shops (owner_name,address,phone_number,running_time) VALUES ($1,$2,$3,$4)", ownerName, address, phoneNumber, runningTime)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	// get the id of the added shop
	lastShopID, err := config.ConnDB().Query("SELECT id FROM shops ORDER BY created_at DESC LIMIT 1")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	var shopID string

	for lastShopID.Next() {
		if err := lastShopID.Scan(&shopID); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}

	// create category shop
	for _, v := range categories {
		vUUID, err := uuid.Parse(v)
		if v != "" {
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  false,
					"message": err.Error(),
				})
				return
			}
		}
		_, err = config.ConnDB().Exec("INSERT INTO category_shop (category_id,shop_id) VALUES ($1,$2)", vUUID, shopID)
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
		"message": "shop successfully added",
	})

}
