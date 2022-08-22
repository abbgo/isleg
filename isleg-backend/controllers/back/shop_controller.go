package controllers

import (
	"github/abbgo/isleg/isleg-backend/config"
	"github/abbgo/isleg/isleg-backend/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateShop(c *gin.Context) {

	db, err := config.ConnDB()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer db.Close()

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

	resultShops, err := db.Query("INSERT INTO shops (owner_name,address,phone_number,running_time) VALUES ($1,$2,$3,$4)", ownerName, address, phoneNumber, runningTime)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer resultShops.Close()

	// get the id of the added shop
	lastShopID, err := db.Query("SELECT id FROM shops ORDER BY created_at DESC LIMIT 1")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer lastShopID.Close()

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
		resultCategorySHop, err := db.Query("INSERT INTO category_shop (category_id,shop_id) VALUES ($1,$2)", vUUID, shopID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
		defer resultCategorySHop.Close()
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "shop successfully added",
	})

}

func UpdateShopByID(c *gin.Context) {

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

	rowShop, err := db.Query("SELECT id FROM shops WHERE id = $1 AND deleted_at IS NULL", ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer rowShop.Close()

	var shopID string

	for rowShop.Next() {
		if err := rowShop.Scan(&shopID); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}

	if shopID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "record not found",
		})
		return
	}

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

	currentTime := time.Now()

	resultShop, err := db.Query("UPDATE shops SET owner_name = $1 , address = $2 , phone_number = $3 , running_time = $4 , updated_at = $5 WHERE id = $6", ownerName, address, phoneNumber, runningTime, currentTime, ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer resultShop.Close()

	resultCategoryShop, err := db.Query("DELETE FROM category_shop WHERE shop_id = $1", ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer resultCategoryShop.Close()

	for _, v := range categories {
		resultCatShop, err := db.Query("INSERT INTO category_shop (category_id,shop_id,updated_at) VALUES ($1,$2,$3)", v, ID, currentTime)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
		defer resultCatShop.Close()
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "shop successfully updated",
	})

}
