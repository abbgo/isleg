package controllers

import (
	"github/abbgo/isleg/isleg-backend/config"
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
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer func() {
		if err := db.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

	var banner models.Banner
	if err := c.BindJSON(&banner); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	// VALIDATE DATA
	if banner.Url == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "url is required",
		})
		return
	} else {
		_, err := url.ParseRequestURI(banner.Url)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}

	// CREATE BREND
	result, err := db.Query("INSERT INTO banner (image,url) VALUES ($1,$2)", banner.Image, banner.Url)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer func() {
		if err := result.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "data successfully added",
	})

}

func UpdateBannerByID(c *gin.Context) {

	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer func() {
		if err := db.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

	// get id from request parameter
	ID := c.Param("id")

	// get data from request
	bannerUrl := c.PostForm("url")
	var fileName string

	// check id and get image of banner
	rowBrend, err := db.Query("SELECT id,image FROM banner WHERE id = $1 AND deleted_at IS NULL", ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer func() {
		if err := rowBrend.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

	var image, bannerID string

	for rowBrend.Next() {
		if err := rowBrend.Scan(&bannerID, &image); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}

	if bannerID == "" {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  false,
			"message": "record not found",
		})
		return
	}

	// VALIDATE DATA
	if bannerUrl == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "banner url is required",
		})
		return
	} else {
		_, err := url.ParseRequestURI(bannerUrl)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}

	fileName, err = pkg.FileUploadForUpdate("image", "banner", image, c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	// update data
	resultBrend, err := db.Query("UPDATE banner SET url = $1 , image = $2 WHERE id = $3", bannerUrl, fileName, ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer func() {
		if err := resultBrend.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "data successfully updated",
	})

}

func GetBannerByID(c *gin.Context) {

	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer func() {
		if err := db.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

	// get id from request paramter
	ID := c.Param("id")

	// check id and get data from database
	rowBanner, err := db.Query("SELECT id,url,image FROM banner WHERE id = $1 AND deleted_at IS NULL", ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer func() {
		if err := rowBanner.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

	var banner models.Banner

	for rowBanner.Next() {
		if err := rowBanner.Scan(&banner.ID, &banner.Url, &banner.Image); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}

	if banner.ID == "" {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  false,
			"message": "record not found",
		})
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
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer func() {
		if err := db.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

	// get limit from param
	limitStr := c.Param("limit")
	if limitStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "limit is required",
		})
		return
	}
	limit, err := strconv.ParseUint(limitStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	// get page from param
	pageStr := c.Param("page")
	if pageStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "page is required",
		})
		return
	}
	page, err := strconv.ParseUint(pageStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	offset := limit * (page - 1)
	var countOfBanners uint

	statusQuery := c.DefaultQuery("status", "false")
	status, err := strconv.ParseBool(statusQuery)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	var countBannerssQuery, rowBannersQuery string
	if !status {
		countBannerssQuery = `SELECT COUNT(id) FROM banner WHERE deleted_at IS NULL`
	} else {
		countBannerssQuery = `SELECT COUNT(id) FROM banner WHERE deleted_at IS NOT NULL`
	}

	// get data from database
	countBanners, err := db.Query(countBannerssQuery)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer func() {
		if err := countBanners.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()
	for countBanners.Next() {
		if err := countBanners.Scan(&countOfBanners); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}

	if !status {
		rowBannersQuery = `SELECT id,url,image FROM banner WHERE deleted_at IS NULL ORDER BY created_at DESC LIMIT $1 OFFSET $2`
	} else {
		rowBannersQuery = `SELECT id,url,image FROM banner WHERE deleted_at IS NOT NULL ORDER BY created_at DESC LIMIT $1 OFFSET $2`
	}

	// get data from database
	rowBanners, err := db.Query(rowBannersQuery, limit, offset)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer func() {
		if err := rowBanners.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

	var banners []models.Banner

	for rowBanners.Next() {
		var banner models.Banner

		if err := rowBanners.Scan(&banner.ID, &banner.Url, &banner.Image); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
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
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer func() {
		if err := db.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

	// get data from database
	rowBanners, err := db.Query("SELECT id,url,image FROM banner WHERE deleted_at IS NULL")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer func() {
		if err := rowBanners.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

	var banners []models.Banner

	for rowBanners.Next() {
		var banner models.Banner

		if err := rowBanners.Scan(&banner.ID, &banner.Url, &banner.Image); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
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
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer func() {
		if err := db.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

	// get id from request parameter
	ID := c.Param("id")

	// check id and get image of brend
	rowBanner, err := db.Query("SELECT id FROM banner WHERE id = $1 AND deleted_at IS NULL", ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer func() {
		if err := rowBanner.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

	var id string

	for rowBanner.Next() {
		if err := rowBanner.Scan(&id); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}

	if id == "" {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  false,
			"message": "record not found",
		})
		return
	}

	resultBanner, err := db.Query("UPDATE banner SET deleted_at = now() WHERE id = $1", ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer func() {
		if err := resultBanner.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "data successfully deleted",
	})

}

func RestoreBannerByID(c *gin.Context) {

	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer func() {
		if err := db.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

	// get id from request parameter
	ID := c.Param("id")

	// check id and get image of brend
	rowBanner, err := db.Query("SELECT id FROM banner WHERE id = $1 AND deleted_at IS NULL", ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer func() {
		if err := rowBanner.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

	var id string

	for rowBanner.Next() {
		if err := rowBanner.Scan(&id); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}

	if id == "" {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  false,
			"message": "record not found",
		})
		return
	}

	resultBanner, err := db.Query("UPDATE banner SET deleted_at = NULL WHERE id = $1", ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer func() {
		if err := resultBanner.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "data successfully restored",
	})

}

func DeletePermanentlyBannerByID(c *gin.Context) {

	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer func() {
		if err := db.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

	// get id from request parameter
	ID := c.Param("id")

	// check id and get image of brend
	rowBrend, err := db.Query("SELECT image FROM banner WHERE id = $1 AND deleted_at IS NOT NULL", ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer func() {
		if err := rowBrend.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

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
		c.JSON(http.StatusNotFound, gin.H{
			"status":  false,
			"message": "record not found",
		})
		return
	}

	if err := os.Remove(pkg.ServerPath + image); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	resultBanners, err := db.Query("DELETE FROM banner WHERE id = $1", ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer func() {
		if err := resultBanners.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "data successfully deleted",
	})

}
