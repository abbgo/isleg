package controllers

import (
	"context"
	"fmt"
	"github/abbgo/isleg/isleg-backend/config"
	"github/abbgo/isleg/isleg-backend/helpers"
	"github/abbgo/isleg/isleg-backend/models"
	"github/abbgo/isleg/isleg-backend/pkg"
	"math"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gosimple/slug"
	"github.com/jackc/pgx/v5"
)

func CreateBrend(c *gin.Context) {
	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	var brend models.Brend
	if err := c.BindJSON(&brend); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	// CREATE BREND
	_, err = db.Exec(context.Background(), "INSERT INTO brends (name,image,slug) VALUES ($1,$2,$3)", brend.Name, brend.Image, slug.MakeLang(brend.Name, "en"))
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "data successfully added",
	})
}

func UpdateBrendByID(c *gin.Context) {
	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	// get id from reequest parameter
	ID := c.Param("id")

	var brend models.Brend
	if err := c.BindJSON(&brend); err != nil {
		helpers.HandleError(c, http.StatusBadRequest, err.Error())
		return
	}

	// check id and get image of brend
	var image, brendID string
	db.QueryRow(context.Background(), "SELECT id,image FROM brends WHERE id = $1 AND deleted_at IS NULL", ID).Scan(&brendID, &image)
	if brendID == "" {
		helpers.HandleError(c, 404, "record not found")
		return
	}

	var fileName string
	if brend.Image == "" {
		fileName = image
	} else {
		fileName = brend.Image
	}

	// update data
	_, err = db.Exec(context.Background(), "UPDATE brends SET name = $1 , image = $2 , slug = $4 WHERE id = $3", brend.Name, fileName, ID, slug.MakeLang(brend.Name, "en"))
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "data successfully updated",
	})
}

func GetBrendByID(c *gin.Context) {
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
	var brend models.Brend
	db.QueryRow(context.Background(), "SELECT id,name,image FROM brends WHERE id = $1 AND deleted_at IS NULL", ID).Scan(&brend.ID, &brend.Name, &brend.Image)
	if brend.ID == "" {
		helpers.HandleError(c, 404, "record not found")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"brend":  brend,
	})
}

func GetBrends(c *gin.Context) {
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
	var countOfBrends uint

	searchQuery := c.Query("search")
	var searchStr, search string
	if searchQuery != "" {
		incomingsSarch := slug.MakeLang(searchQuery, "en")
		search = strings.ReplaceAll(incomingsSarch, "-", " | ")
		searchStr = fmt.Sprintf("%%%s%%", search)
	}

	statusQuery := c.DefaultQuery("status", "false")
	status, err := strconv.ParseBool(statusQuery)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	if !status {
		if search == "" {
			db.QueryRow(context.Background(), "SELECT COUNT(id) FROM brends WHERE deleted_at IS NULL").Scan(&countOfBrends)
		} else {
			db.QueryRow(context.Background(), "SELECT COUNT(id) FROM brends WHERE deleted_at IS NULL AND (to_tsvector(slug) @@ to_tsquery($1) OR slug LIKE $2)", search, searchStr).Scan(&countOfBrends)
		}
	} else {
		if search == "" {
			db.QueryRow(context.Background(), "SELECT COUNT(id) FROM brends WHERE deleted_at IS NOT NULL").Scan(&countOfBrends)
		} else {
			db.QueryRow(context.Background(), "SELECT COUNT(id) FROM brends WHERE deleted_at IS NOT NULL AND (to_tsvector(slug) @@ to_tsquery($1) OR slug LIKE $2)", search, searchStr).Scan(&countOfBrends)
		}
	}

	var rowBrends pgx.Rows
	if !status {
		if search == "" {
			rowBrends, err = db.Query(context.Background(), "SELECT id,name,image FROM brends WHERE deleted_at IS NULL ORDER BY created_at DESC LIMIT $1 OFFSET $2", limit, offset)
			if err != nil {
				helpers.HandleError(c, 400, err.Error())
				return
			}
		} else {
			rowBrends, err = db.Query(context.Background(), "SELECT id,name,image FROM brends WHERE deleted_at IS NULL AND (to_tsvector(slug) @@ to_tsquery($3) OR slug LIKE $4) ORDER BY created_at DESC LIMIT $1 OFFSET $2", limit, offset, search, searchStr)
			if err != nil {
				helpers.HandleError(c, 400, err.Error())
				return
			}
		}
	} else {
		if search == "" {
			rowBrends, err = db.Query(context.Background(), "SELECT id,name,image FROM brends WHERE deleted_at IS NOT NULL ORDER BY created_at DESC LIMIT $1 OFFSET $2", limit, offset)
			if err != nil {
				helpers.HandleError(c, 400, err.Error())
				return
			}
		} else {
			rowBrends, err = db.Query(context.Background(), "SELECT id,name,image FROM brends WHERE deleted_at IS NOT NULL AND (to_tsvector(slug) @@ to_tsquery($3) OR slug LIKE $4) ORDER BY created_at DESC LIMIT $1 OFFSET $2", limit, offset, search, searchStr)
			if err != nil {
				helpers.HandleError(c, 400, err.Error())
				return
			}
		}
	}
	defer rowBrends.Close()

	// get data from database

	var brends []models.Brend
	for rowBrends.Next() {
		var brend models.Brend
		if err := rowBrends.Scan(&brend.ID, &brend.Name, &brend.Image); err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}
		brends = append(brends, brend)
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"brends": brends,
		"total":  countOfBrends,
	})
}

func GetDeletedBrends(c *gin.Context) {
	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	// get data from database
	rowBrends, err := db.Query(context.Background(), "SELECT id,name,image FROM brends WHERE deleted_at IS NOT NULL ORDER BY created_at DESC")
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer rowBrends.Close()

	var brends []models.Brend
	for rowBrends.Next() {
		var brend models.Brend
		if err := rowBrends.Scan(&brend.ID, &brend.Name, &brend.Image); err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}
		brends = append(brends, brend)
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"brends": brends,
	})
}

func DeleteBrendByID(c *gin.Context) {
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
	db.QueryRow(context.Background(), "SELECT id FROM brends WHERE id = $1 AND deleted_at IS NULL", ID).Scan(&id)
	if id == "" {
		helpers.HandleError(c, 404, "record not found")
		return
	}

	_, err = db.Exec(context.Background(), "CALL delete_brend($1)", ID)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "data successfully deleted",
	})
}

func RestoreBrendByID(c *gin.Context) {
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
	var id string
	db.QueryRow(context.Background(), "SELECT id FROM brends WHERE id = $1 AND deleted_at IS NOT NULL", ID).Scan(&id)
	if id == "" {
		helpers.HandleError(c, 404, "record not found")
		return
	}

	_, err = db.Exec(context.Background(), "CALL restore_brend($1)", ID)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "data successfully restored",
	})
}

func DeletePermanentlyBrendByID(c *gin.Context) {
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
	db.QueryRow(context.Background(), "SELECT image FROM brends WHERE id = $1 AND deleted_at IS NOT NULL", ID).Scan(&image)
	if image == "" {
		helpers.HandleError(c, 404, "record not found")
		return
	}

	if err := os.Remove(pkg.ServerPath + image); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	rowsMainImage, err := db.Query(context.Background(), "SELECT main_image FROM products WHERE brend_id = $1", ID)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer rowsMainImage.Close()

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

	rowsImages, err := db.Query(context.Background(), "SELECT i.image FROM images i INNER JOIN products p ON p.id = i.product_id WHERE p.brend_id = $1", ID)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer rowsImages.Close()

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

	_, err = db.Exec(context.Background(), "DELETE FROM brends WHERE id = $1", ID)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "data successfully deleted",
	})
}

// GetAllBrendForHomePage ahli brendlerin id - lerini we suratlaryny alyar
func GetAllBrendForHomePage() ([]models.Brend, error) {
	db, err := config.ConnDB()
	if err != nil {
		return []models.Brend{}, err
	}
	defer db.Close()

	var brends []models.Brend

	// get all brends
	rows, err := db.Query(context.Background(), "SELECT id,name,image FROM brends WHERE deleted_at IS NULL")
	if err != nil {
		return []models.Brend{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var brend models.Brend
		if err := rows.Scan(&brend.ID, &brend.Name, &brend.Image); err != nil {
			return []models.Brend{}, err
		}
		brends = append(brends, brend)
	}

	return brends, nil
}

func GetOneBrendWithProducts(c *gin.Context) {
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	var countOfProducts uint64

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

	brendID := c.Param("id")

	// get brend where id equal brendID
	var brend Category
	brendRow, err := db.Query(context.Background(), "SELECT id,image,name FROM brends where id = $1", brendID)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer brendRow.Close()

	for brendRow.Next() {
		if err := brendRow.Scan(&brend.ID, &brend.Image, &brend.Name); err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}

		if brend.ID == "" {
			helpers.HandleError(c, 404, "brend not found")
			return
		}

		// get count product where product_id equal brendID
		brendCount, err := db.Query(context.Background(), "SELECT COUNT(id) FROM products WHERE brend_id = $1 AND amount > 0 AND limit_amount > 0 AND deleted_at IS NULL", brendID)
		if err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}
		defer brendCount.Close()

		for brendCount.Next() {
			if err := brendCount.Scan(&countOfProducts); err != nil {
				if err != nil {
					helpers.HandleError(c, 400, err.Error())
					return
				}
			}
		}

		// get all product where brend id equal brendID
		productRows, err := db.Query(context.Background(), "SELECT id,price,old_price,limit_amount,is_new,amount,main_image,benefit FROM products WHERE brend_id = $1 AND amount > 0 AND limit_amount > 0 AND deleted_at IS NULL ORDER BY created_at ASC LIMIT $2 OFFSET $3", brendID, limit, offset)
		if err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}
		defer productRows.Close()

		var products []Product
		for productRows.Next() {
			var product Product
			if err := productRows.Scan(&product.ID, &product.Price, &product.OldPrice, &product.LimitAmount, &product.IsNew, &product.Amount, &product.MainImage, &product.Benefit); err != nil {
				helpers.HandleError(c, 400, err.Error())
				return
			}

			if product.Benefit.Float64 != 0 {
				product.Price = product.Price + (product.Price*product.Benefit.Float64)/100
				product.OldPrice = product.OldPrice + (product.OldPrice*product.Benefit.Float64)/100
			}

			if product.OldPrice != 0 {
				product.Percentage = -math.Round(((product.OldPrice - product.Price) * 100) / product.OldPrice)
			} else {
				product.Percentage = 0
			}

			rowsLang, err := db.Query(context.Background(), "SELECT id,name_short FROM languages WHERE deleted_at IS NULL")
			if err != nil {
				helpers.HandleError(c, 400, err.Error())
				return
			}
			defer rowsLang.Close()

			var languages []models.Language
			for rowsLang.Next() {
				var language models.Language
				if err := rowsLang.Scan(&language.ID, &language.NameShort); err != nil {
					helpers.HandleError(c, 400, err.Error())
					return
				}
				languages = append(languages, language)
			}

			for _, v := range languages {
				rowTrProduct, err := db.Query(context.Background(), "SELECT name,description FROM translation_product WHERE lang_id = $1 AND product_id = $2 AND deleted_at IS NULL", v.ID, product.ID)
				if err != nil {
					helpers.HandleError(c, 400, err.Error())
					return
				}
				defer rowTrProduct.Close()

				var trProduct models.TranslationProduct
				translation := make(map[string]models.TranslationProduct)
				for rowTrProduct.Next() {
					if err := rowTrProduct.Scan(&trProduct.Name, &trProduct.Description); err != nil {
						helpers.HandleError(c, 400, err.Error())
						return
					}
				}

				translation[v.NameShort] = trProduct
				product.Translations = append(product.Translations, translation)
			}

			// get brend where id equal brend_id of product
			var brend Brend
			db.QueryRow(context.Background(), "SELECT b.id,b.name FROM products p LEFT JOIN brends b ON p.brend_id=b.id WHERE p.id = $1 AND p.deleted_at IS NULL AND b.deleted_at IS NULL", product.ID).Scan(&brend.ID, &brend.Name)
			product.Brend = brend
			products = append(products, product)
		}
		brend.Products = products
	}

	c.JSON(http.StatusOK, gin.H{
		"status":            true,
		"brend":             brend,
		"count_of_products": countOfProducts,
	})
}
