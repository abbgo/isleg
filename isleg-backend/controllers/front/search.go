package controllers

import (
	"fmt"
	"github/abbgo/isleg/isleg-backend/config"
	backController "github/abbgo/isleg/isleg-backend/controllers/back"
	"github/abbgo/isleg/isleg-backend/models"
	"math"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gosimple/slug"
	"github.com/lib/pq"
)

func Search(c *gin.Context) {

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

	// GET DATA FROM ROUTE PARAMETER
	langShortName := c.Param("lang")

	// GET language id
	langID, err := backController.GetLangID(langShortName)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

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

	var countOfProduct uint

	incomingsSarch := slug.MakeLang(c.Query("search"), "en")
	search := strings.ReplaceAll(incomingsSarch, "-", " | ")
	searchStr := fmt.Sprintf("%%%s%%", search)

	countProduct, err := db.Query("SELECT COUNT(DISTINCT p.id) FROM products p inner join translation_product tp on tp.product_id = p.id WHERE to_tsvector(tp.slug) @@ to_tsquery($1) OR tp.slug LIKE $3 AND tp.lang_id = $2 AND tp.deleted_at IS NULL AND p.amount > 0 AND p.limit_amount > 0 AND p.deleted_at IS NULL", search, langID, searchStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer func() {
		if err := countProduct.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

	for countProduct.Next() {
		if err := countProduct.Scan(&countOfProduct); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}

	rowsProduct, err := db.Query("SELECT DISTINCT ON (p.created_at) p.id,p.brend_id,p.price,p.old_price,p.amount,p.limit_amount,p.is_new,p.main_image,p.benefit FROM products p inner join translation_product tp on tp.product_id = p.id WHERE to_tsvector(tp.slug) @@ to_tsquery($1) OR tp.slug LIKE $5 AND tp.lang_id = $2 AND tp.deleted_at IS NULL AND p.amount > 0 AND p.limit_amount > 0 AND p.deleted_at IS NULL ORDER BY p.created_at ASC LIMIT $3 OFFSET $4", search, langID, limit, offset, searchStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
			"yalnys":  "yalnys",
		})
		return
	}
	defer func() {
		if err := rowsProduct.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

	var products []LikeProduct

	for rowsProduct.Next() {
		var product LikeProduct

		if err := rowsProduct.Scan(&product.ID, &product.BrendID, &product.Price, &product.OldPrice, &product.Amount, &product.LimitAmount, &product.IsNew, &product.MainImage, &product.Benefit); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
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

		rowsLang, err := db.Query("SELECT id,name_short FROM languages WHERE deleted_at IS NULL")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
		defer func() {
			if err := rowsLang.Close(); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  false,
					"message": err.Error(),
				})
				return
			}
		}()

		var languages []models.Language

		for rowsLang.Next() {
			var language models.Language

			if err := rowsLang.Scan(&language.ID, &language.NameShort); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  false,
					"message": err.Error(),
				})
				return
			}

			languages = append(languages, language)
		}

		for _, v := range languages {

			rowTranslationProduct, err := db.Query("SELECT lang_id,name,description FROM translation_product WHERE product_id = $1 AND lang_id = $2 AND deleted_at IS NULL", product.ID, v.ID)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  false,
					"message": err.Error(),
				})
				return
			}
			defer func() {
				if err := rowTranslationProduct.Close(); err != nil {
					c.JSON(http.StatusBadRequest, gin.H{
						"status":  false,
						"message": err.Error(),
					})
					return
				}
			}()

			var trProduct models.TranslationProduct

			translation := make(map[string]models.TranslationProduct)

			for rowTranslationProduct.Next() {
				if err := rowTranslationProduct.Scan(&trProduct.LangID, &trProduct.Name, &trProduct.Description); err != nil {
					c.JSON(http.StatusBadRequest, gin.H{
						"status":  false,
						"message": err.Error(),
					})
					return
				}
			}

			translation[v.NameShort] = trProduct

			product.Translations = append(product.Translations, translation)

		}

		products = append(products, product)

	}

	c.JSON(http.StatusOK, gin.H{
		"status":            true,
		"products":          products,
		"count_of_products": countOfProduct,
	})

}

func FilterAndSort(c *gin.Context) {

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

	// GET DATA FROM ROUTE PARAMETER
	langShortName := c.Param("lang")

	// GET language id
	langID, err := backController.GetLangID(langShortName)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	//get category_id from param
	categoryID := c.Param("id")
	if categoryID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": " category_id is required",
		})
		return
	}

	rowCategory, err := db.Query("SELECT id FROM categories WHERE id = $1 AND deleted_at IS NULL", categoryID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer func() {
		if err := rowCategory.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()
	var category_id string
	for rowCategory.Next() {
		if err := rowCategory.Scan(&category_id); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}
	if category_id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "category not found",
		})
		return
	}

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

	var countOfProduct uint

	priceSort := c.Query("price_sort")
	if priceSort != "" {
		if priceSort != "ASC" && priceSort != "DESC" {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": "price_sort is invalid",
			})
			return
		}
	}

	var minPrice float32
	minPriceStr := c.Query("min_price")
	if minPriceStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "min_price is required",
		})
		return
	} else {
		min_price, err := strconv.ParseFloat(minPriceStr, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
		if min_price < 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": "min_price cannot be less than zero",
			})
			return
		}
		minPrice = float32(min_price)
	}

	var maxPrice float32
	maxPriceStr := c.Query("max_price")
	if maxPriceStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "max_price is required",
		})
		return
	} else {
		max_price, err := strconv.ParseFloat(maxPriceStr, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
		if max_price < 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": "max_price cannot be less than zero",
			})
			return
		}
		maxPrice = float32(max_price)
	}

	isDiscountStr := c.Query("is_discount")
	discount := -1
	isDiscount, err := strconv.ParseBool(isDiscountStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	if isDiscount {
		discount = 0
	}

	brendIDs := c.QueryArray("brend_ids")
	if len(brendIDs) != 0 {
		for _, v := range brendIDs {
			row, err := db.Query("SELECT id FROM brends WHERE id = $1 AND deleted_at IS NULL", v)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  false,
					"message": err.Error(),
				})
				return
			}
			defer func() {
				if err := row.Close(); err != nil {
					c.JSON(http.StatusBadRequest, gin.H{
						"status":  false,
						"message": err.Error(),
					})
					return
				}
			}()

			var brend_id string

			for row.Next() {
				if err := row.Scan(&brend_id); err != nil {
					c.JSON(http.StatusBadRequest, gin.H{
						"status":  false,
						"message": err.Error(),
					})
					return
				}
			}

			if brend_id == "" {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  false,
					"message": "brend not found",
				})
				return
			}
		}
	}

	countProduct, err := db.Query("SELECT COUNT(*) FROM products p LEFT JOIN category_product c ON p.id=c.product_id INNER JOIN translation_product tp ON tp.product_id = p.id WHERE p.brend_id = ANY($1) AND tp.lang_id = $2 AND c.category_id = $3 AND tp.deleted_at IS NULL AND p.amount > 0 AND p.limit_amount > 0 AND p.deleted_at IS NULL AND p.price >= $4 AND p.price <= $5 AND p.old_price > $6", pq.Array(brendIDs), langID, categoryID, minPrice, maxPrice, discount)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer func() {
		if err := countProduct.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

	for countProduct.Next() {
		if err := countProduct.Scan(&countOfProduct); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}

	rowQuery := "SELECT p.id,p.brend_id,p.price,p.old_price,p.amount,p.limit_amount,p.is_new,p.main_image,benefit FROM products p LEFT JOIN category_product c ON p.id=c.product_id INNER JOIN translation_product tp ON tp.product_id = p.id WHERE p.brend_id = ANY($1) AND tp.lang_id = $2 AND c.category_id = $3 AND tp.deleted_at IS NULL AND p.amount > 0 AND p.limit_amount > 0 AND p.deleted_at IS NULL AND p.price >= $4 AND p.price <= $5 AND p.old_price > $6 ORDER BY p.price ASC LIMIT $7 OFFSET $8"
	if priceSort == "DESC" {
		rowQuery = "SELECT p.id,p.brend_id,p.price,p.old_price,p.amount,p.limit_amount,p.is_new,p.main_image,benefit FROM products p LEFT JOIN category_product c ON p.id=c.product_id INNER JOIN translation_product tp ON tp.product_id = p.id WHERE p.brend_id = ANY($1) AND tp.lang_id = $2 AND c.category_id = $3 AND tp.deleted_at IS NULL AND p.amount > 0 AND p.limit_amount > 0 AND p.deleted_at IS NULL AND p.price >= $4 AND p.price <= $5 AND p.old_price > $6 ORDER BY p.price DESC LIMIT $7 OFFSET $8"
	}
	rowsProduct, err := db.Query(rowQuery, pq.Array(brendIDs), langID, categoryID, minPrice, maxPrice, discount, limit, offset)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer func() {
		if err := rowsProduct.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

	var products []LikeProduct

	for rowsProduct.Next() {
		var product LikeProduct

		if err := rowsProduct.Scan(&product.ID, &product.BrendID, &product.Price, &product.OldPrice, &product.Amount, &product.LimitAmount, &product.IsNew, &product.MainImage, &product.Benefit); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
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

		rowsLang, err := db.Query("SELECT id,name_short FROM languages WHERE deleted_at IS NULL")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
		defer func() {
			if err := rowsLang.Close(); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  false,
					"message": err.Error(),
				})
				return
			}
		}()

		var languages []models.Language

		for rowsLang.Next() {
			var language models.Language

			if err := rowsLang.Scan(&language.ID, &language.NameShort); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  false,
					"message": err.Error(),
				})
				return
			}

			languages = append(languages, language)
		}

		for _, v := range languages {

			rowTranslationProduct, err := db.Query("SELECT lang_id,name,description FROM translation_product WHERE product_id = $1 AND lang_id = $2 AND deleted_at IS NULL", product.ID, v.ID)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  false,
					"message": err.Error(),
				})
				return
			}
			defer func() {
				if err := rowTranslationProduct.Close(); err != nil {
					c.JSON(http.StatusBadRequest, gin.H{
						"status":  false,
						"message": err.Error(),
					})
					return
				}
			}()

			var trProduct models.TranslationProduct

			translation := make(map[string]models.TranslationProduct)

			for rowTranslationProduct.Next() {
				if err := rowTranslationProduct.Scan(&trProduct.LangID, &trProduct.Name, &trProduct.Description); err != nil {
					c.JSON(http.StatusBadRequest, gin.H{
						"status":  false,
						"message": err.Error(),
					})
					return
				}
			}

			translation[v.NameShort] = trProduct

			product.Translations = append(product.Translations, translation)

		}

		products = append(products, product)

	}

	c.JSON(http.StatusOK, gin.H{
		"status":            true,
		"products":          products,
		"count_of_products": countOfProduct,
	})

}
