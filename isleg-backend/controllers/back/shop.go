package controllers

import (
	"context"
	"fmt"
	"github/abbgo/isleg/isleg-backend/config"
	"github/abbgo/isleg/isleg-backend/helpers"
	"github/abbgo/isleg/isleg-backend/models"
	"github/abbgo/isleg/isleg-backend/pkg"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func CreateShop(c *gin.Context) {
	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	// get data from request
	var shop models.Shop
	if err := c.BindJSON(&shop); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	var phoneNumber string
	if err := db.QueryRow(context.Background(), "SELECT id FROM shops WHERE phone_number = $1 AND deleted_at IS NULL", shop.PhoneNumber).Scan(&phoneNumber); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	if phoneNumber != "" {
		helpers.HandleError(c, 400, "this shop already exists")
		return
	}

	_, err = db.Exec(context.Background(), "INSERT INTO shops (owner_name,address,phone_number,running_time) VALUES ($1,$2,$3,$4)", shop.OwnerName, shop.Address, shop.PhoneNumber, shop.RunningTime)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "data successfully added",
	})
}

func UpdateShopByID(c *gin.Context) {
	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	// gat data from request
	var shop models.Shop
	if err := c.BindJSON(&shop); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	shop_id := c.Param("id")

	// check id
	var shopID string
	if err := db.QueryRow(context.Background(), "SELECT id FROM shops WHERE id = $1 AND deleted_at IS NULL", shop_id).Scan(&shopID); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	if shopID == "" {
		helpers.HandleError(c, 404, "record not found")
		return
	}

	_, err = db.Exec(context.Background(), "UPDATE shops SET owner_name = $1 , address = $2 , phone_number = $3 , running_time = $4 WHERE id = $5", shop.OwnerName, shop.Address, shop.PhoneNumber, shop.RunningTime, shop_id)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "data successfully updated",
	})
}

func GetShopByID(c *gin.Context) {
	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	// get id from requets parameter
	ID := c.Param("id")

	// check id and get data from database
	var shop models.Shop
	if err := db.QueryRow(context.Background(), "SELECT id,owner_name,address,phone_number,running_time FROM shops WHERE id = $1 AND deleted_at IS NULL", ID).Scan(&shop.ID, &shop.OwnerName, &shop.Address, &shop.PhoneNumber, &shop.RunningTime); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	if shop.ID == "" {
		helpers.HandleError(c, 404, "record not found")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"shop":   shop,
	})
}

func GetShops(c *gin.Context) {
	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	// get limit from param
	limitStr := c.Param("limit")
	if limitStr == "" {
		helpers.HandleError(c, 400, "limit is required")
		return
	}
	limit, err := strconv.ParseUint(limitStr, 10, 64)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	// get page from param
	pageStr := c.Param("page")
	if pageStr == "" {
		helpers.HandleError(c, 400, "page is required")
		return
	}
	page, err := strconv.ParseUint(pageStr, 10, 64)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	offset := limit * (page - 1)
	var countOfShops uint

	searchQuery := c.Query("search")
	search := fmt.Sprintf("%%%s%%", searchQuery)

	statusQuery := c.DefaultQuery("status", "false")
	status, err := strconv.ParseBool(statusQuery)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	var rowsShop pgx.Rows
	if !status {
		if searchQuery == "" {
			if err = db.QueryRow(context.Background(), "SELECT COUNT(id) FROM shops WHERE deleted_at IS NULL").Scan(&countOfShops); err != nil {
				helpers.HandleError(c, 400, err.Error())
				return
			}
		} else {
			if err = db.QueryRow(context.Background(), "SELECT COUNT(id) FROM shops WHERE deleted_at IS NULL AND phone_number LIKE $1", search).Scan(&countOfShops); err != nil {
				helpers.HandleError(c, 400, err.Error())
				return
			}
		}
	} else {
		if searchQuery == "" {
			if err = db.QueryRow(context.Background(), "SELECT COUNT(id) FROM shops WHERE deleted_at IS NOT NULL").Scan(&countOfShops); err != nil {
				helpers.HandleError(c, 400, err.Error())
				return
			}
		} else {
			if err = db.QueryRow(context.Background(), "SELECT COUNT(id) FROM shops WHERE deleted_at IS NOT NULL AND phone_number LIKE $1", search).Scan(&countOfShops); err != nil {
				helpers.HandleError(c, 400, err.Error())
				return
			}
		}
	}

	if !status {
		if searchQuery == "" {
			rowsShop, err = db.Query(context.Background(), "SELECT id,owner_name,address,phone_number,running_time FROM shops WHERE deleted_at IS NULL ORDER BY created_at DESC LIMIT $1 OFFSET $2", limit, offset)
			if err != nil {
				helpers.HandleError(c, 400, err.Error())
				return
			}
		} else {
			rowsShop, err = db.Query(context.Background(), "SELECT id,owner_name,address,phone_number,running_time FROM shops WHERE deleted_at IS NULL AND phone_number LIKE $3 ORDER BY created_at DESC LIMIT $1 OFFSET $2", limit, offset, search)
			if err != nil {
				helpers.HandleError(c, 400, err.Error())
				return
			}
		}
	} else {
		if searchQuery == "" {
			rowsShop, err = db.Query(context.Background(), "SELECT id,owner_name,address,phone_number,running_time FROM shops WHERE deleted_at IS NOT NULL ORDER BY created_at DESC LIMIT $1 OFFSET $2", limit, offset)
			if err != nil {
				helpers.HandleError(c, 400, err.Error())
				return
			}
		} else {
			rowsShop, err = db.Query(context.Background(), "SELECT id,owner_name,address,phone_number,running_time FROM shops WHERE deleted_at IS NOT NULL AND phone_number LIKE $3 ORDER BY created_at DESC LIMIT $1 OFFSET $2", limit, offset, search)
			if err != nil {
				helpers.HandleError(c, 400, err.Error())
				return
			}
		}
	}

	var shops []models.Shop
	for rowsShop.Next() {
		var shop models.Shop
		if err := rowsShop.Scan(&shop.ID, &shop.OwnerName, &shop.Address, &shop.PhoneNumber, &shop.RunningTime); err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}
		shops = append(shops, shop)
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"shops":  shops,
		"total":  countOfShops,
	})
}

func DeleteShopByID(c *gin.Context) {
	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	// get id from request parameter
	ID := c.Param("id")

	// check id
	var shopID string
	if err := db.QueryRow(context.Background(), "SELECT id FROM shops WHERE id = $1 AND deleted_at IS NULL", ID).Scan(&shopID); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	if shopID == "" {
		helpers.HandleError(c, 404, "record not found")
		return
	}

	_, err = db.Exec(context.Background(), "CALL delete_shop($1)", ID)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "data successfully deleted",
	})
}

func RestoreShopByID(c *gin.Context) {
	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	// get id from request
	ID := c.Param("id")

	// check id
	var shopID string
	if err := db.QueryRow(context.Background(), "SELECT id FROM shops WHERE id = $1 AND deleted_at IS NOT NULL", ID).Scan(&shopID); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	if shopID == "" {
		helpers.HandleError(c, 404, "record not found")
		return
	}

	_, err = db.Exec(context.Background(), "CALL restore_shop($1)", ID)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "data successfully restored",
	})
}

func DeletePermanentlyShopByID(c *gin.Context) {
	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	// get id from request parameter
	ID := c.Param("id")

	// check id
	var shopID string
	if err := db.QueryRow(context.Background(), "SELECT id FROM shops WHERE id = $1 AND deleted_at IS NOT NULL", ID).Scan(&shopID); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	if shopID == "" {
		helpers.HandleError(c, 404, "record not found")
		return
	}

	rowsMainImage, err := db.Query(context.Background(), "SELECT main_image FROM products WHERE shop_id = $1", ID)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	var mainImages []string
	for rowsMainImage.Next() {
		var mainImage string
		if err := rowsMainImage.Scan(&mainImage); err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}
		mainImages = append(mainImages, mainImage)
	}

	for _, v := range mainImages {
		if err := os.Remove(pkg.ServerPath + v); err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}
	}

	rowsImages, err := db.Query(context.Background(), "SELECT i.image FROM images i INNER JOIN products p ON p.id = i.product_id WHERE p.shop_id = $1", ID)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	var images []models.Images
	for rowsImages.Next() {
		var image models.Images
		if err := rowsImages.Scan(&image.Image); err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}
		images = append(images, image)
	}

	for _, v := range images {
		if err := os.Remove(pkg.ServerPath + v.Image); err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}
	}

	_, err = db.Exec(context.Background(), "DELETE FROM shops WHERE id = $1", ID)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "data successfully deleted",
	})
}
