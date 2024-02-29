package controllers

import (
	"context"
	"github/abbgo/isleg/isleg-backend/config"
	"github/abbgo/isleg/isleg-backend/helpers"
	"github/abbgo/isleg/isleg-backend/models"
	"github/abbgo/isleg/isleg-backend/pkg"
	"net/http"
	"net/url"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateBanner(c *gin.Context) {
	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	var banner models.Banner
	if err := c.BindJSON(&banner); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	// VALIDATE DATA
	// if banner.Url == "" {
	// 	helpers.HandleError(c, 400, "url is required")
	// 	return
	// } else {
	// 	_, err := url.ParseRequestURI(banner.Url)
	// 	if err != nil {
	// 		helpers.HandleError(c, 400, err.Error())
	// 		return
	// 	}
	// }

	// CREATE BREND
	_, err = db.Exec(context.Background(), "INSERT INTO banners (image,url) VALUES ($1,$2)", banner.Image, banner.Url)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "data successfully added",
	})
}

func UpdateBannerByID(c *gin.Context) {
	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	// get id from request parameter
	ID := c.Param("id")

	var banner models.Banner
	if err := c.BindJSON(&banner); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	// check id and get image of banner
	var image, bannerID string
	db.QueryRow(context.Background(), "SELECT id,image FROM banners WHERE id = $1 AND deleted_at IS NULL", ID).Scan(&bannerID, &image)
	if bannerID == "" {
		helpers.HandleError(c, 404, "record not found")
		return
	}

	// VALIDATE DATA
	if banner.Url == "" {
		helpers.HandleError(c, 400, "banner url is required")
		return
	} else {
		_, err := url.ParseRequestURI(banner.Url)
		if err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}
	}

	var fileName string
	if banner.Image == "" {
		fileName = image
	} else {
		fileName = banner.Image
	}

	// update data
	_, err = db.Exec(context.Background(), "UPDATE banners SET url = $1 , image = $2 WHERE id = $3", banner.Url, fileName, ID)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "data successfully updated",
	})
}

func GetBannerByID(c *gin.Context) {
	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	// get id from request paramter
	ID := c.Param("id")

	// check id and get data from database
	var banner models.Banner
	db.QueryRow(context.Background(), "SELECT id,url,image FROM banners WHERE id = $1 AND deleted_at IS NULL", ID).Scan(&banner.ID, &banner.Url, &banner.Image)
	if banner.ID == "" {
		helpers.HandleError(c, 404, "record not found")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"banner": banner,
	})
}

func GetBanners(c *gin.Context) {
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
	var countOfBanners uint

	statusQuery := c.DefaultQuery("status", "false")
	status, err := strconv.ParseBool(statusQuery)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	var countBannerssQuery, rowBannersQuery string
	if !status {
		countBannerssQuery = `SELECT COUNT(id) FROM banners WHERE deleted_at IS NULL`
	} else {
		countBannerssQuery = `SELECT COUNT(id) FROM banners WHERE deleted_at IS NOT NULL`
	}

	// get data from database
	db.QueryRow(context.Background(), countBannerssQuery).Scan(&countOfBanners)
	if !status {
		rowBannersQuery = `SELECT id,url,image FROM banners WHERE deleted_at IS NULL ORDER BY created_at DESC LIMIT $1 OFFSET $2`
	} else {
		rowBannersQuery = `SELECT id,url,image FROM banners WHERE deleted_at IS NOT NULL ORDER BY created_at DESC LIMIT $1 OFFSET $2`
	}

	// get data from database
	rowBanners, err := db.Query(context.Background(), rowBannersQuery, limit, offset)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer rowBanners.Close()

	var banners []models.Banner
	for rowBanners.Next() {
		var banner models.Banner
		if err := rowBanners.Scan(&banner.ID, &banner.Url, &banner.Image); err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}
		banners = append(banners, banner)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"banners": banners,
		"total":   countOfBanners,
	})
}

func GetBannersForFront(c *gin.Context) {
	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	// get data from database
	rowBanners, err := db.Query(context.Background(), "SELECT id,url,image FROM banners WHERE deleted_at IS NULL")
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer rowBanners.Close()

	var banners []models.Banner
	for rowBanners.Next() {
		var banner models.Banner
		if err := rowBanners.Scan(&banner.ID, &banner.Url, &banner.Image); err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}
		banners = append(banners, banner)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"banners": banners,
	})
}

func DeleteBannerByID(c *gin.Context) {
	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	// get id from request parameter
	ID := c.Param("id")

	// check id and get image of brend
	var id string
	db.QueryRow(context.Background(), "SELECT id FROM banners WHERE id = $1 AND deleted_at IS NULL", ID).Scan(&id)
	if id == "" {
		helpers.HandleError(c, 404, "record not found")
		return
	}

	_, err = db.Exec(context.Background(), "UPDATE banners SET deleted_at = now() WHERE id = $1", ID)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "data successfully deleted",
	})
}

func RestoreBannerByID(c *gin.Context) {
	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	// get id from request parameter
	ID := c.Param("id")

	// check id and get image of brend
	var id string
	db.QueryRow(context.Background(), "SELECT id FROM banners WHERE id = $1 AND deleted_at IS NOT NULL", ID).Scan(&id)
	if id == "" {
		helpers.HandleError(c, 404, "record not found")
		return
	}

	_, err = db.Exec(context.Background(), "UPDATE banners SET deleted_at = NULL WHERE id = $1", ID)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "data successfully restored",
	})
}

func DeletePermanentlyBannerByID(c *gin.Context) {
	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	// get id from request parameter
	ID := c.Param("id")

	// check id and get image of brend
	var image string
	db.QueryRow(context.Background(), "SELECT image FROM banners WHERE id = $1 AND deleted_at IS NOT NULL", ID).Scan(&image)
	if image == "" {
		helpers.HandleError(c, 404, "record not found")
		return
	}

	if err := os.Remove(pkg.ServerPath + image); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	_, err = db.Exec(context.Background(), "DELETE FROM banners WHERE id = $1", ID)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "data successfully deleted",
	})
}
