package controllers

import (
	"context"
	"fmt"
	"github/abbgo/isleg/isleg-backend/config"
	backController "github/abbgo/isleg/isleg-backend/controllers/back"
	"github/abbgo/isleg/isleg-backend/helpers"
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
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	// GET DATA FROM ROUTE PARAMETER
	langShortName := c.Param("lang")

	// GET language id
	langID, err := backController.GetLangID(langShortName)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

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

	var countOfProduct uint

	incomingsSarch := slug.MakeLang(c.Query("search"), "en")
	search := strings.ReplaceAll(incomingsSarch, "-", " | ")
	searchStr := fmt.Sprintf("%%%s%%", search)

	db.QueryRow(context.Background(), "SELECT COUNT(DISTINCT p.id) FROM products p INNER JOIN translation_product tp ON tp.product_id = p.id WHERE p.is_visible = true AND to_tsvector(tp.slug) @@ to_tsquery($1) OR tp.slug LIKE $3 AND tp.lang_id = $2 AND tp.deleted_at IS NULL AND p.amount > 0 AND p.limit_amount > 0 AND p.deleted_at IS NULL", search, langID, searchStr).Scan(&countOfProduct)

	rowsProduct, err := db.Query(context.Background(), "SELECT DISTINCT ON (p.created_at) p.id,p.brend_id,p.price,p.old_price,p.amount,p.limit_amount,p.is_new,p.main_image,p.benefit FROM products p INNER JOIN translation_product tp ON tp.product_id = p.id WHERE p.is_visible = true AND to_tsvector(tp.slug) @@ to_tsquery($1) OR tp.slug LIKE $5 AND tp.lang_id = $2 AND tp.deleted_at IS NULL AND p.amount > 0 AND p.limit_amount > 0 AND p.deleted_at IS NULL ORDER BY p.created_at ASC LIMIT $3 OFFSET $4", search, langID, limit, offset, searchStr)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer rowsProduct.Close()

	var products []LikeProduct
	for rowsProduct.Next() {
		var product LikeProduct
		if err := rowsProduct.Scan(&product.ID, &product.BrendID, &product.Price, &product.OldPrice, &product.Amount, &product.LimitAmount, &product.IsNew, &product.MainImage, &product.Benefit); err != nil {
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
			var trProduct models.TranslationProduct
			translation := make(map[string]models.TranslationProduct)
			db.QueryRow(context.Background(), "SELECT lang_id,name,description FROM translation_product WHERE product_id = $1 AND lang_id = $2 AND deleted_at IS NULL", product.ID, v.ID).Scan(&trProduct.LangID, &trProduct.Name, &trProduct.Description)
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
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	// GET DATA FROM ROUTE PARAMETER
	langShortName := c.Param("lang")

	// GET language id
	langID, err := backController.GetLangID(langShortName)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	//get category_id from param
	categoryID := c.Param("id")
	if categoryID == "" {
		helpers.HandleError(c, 400, " category_id is required")
		return
	}

	var category_id string
	db.QueryRow(context.Background(), "SELECT id FROM categories WHERE id = $1 AND deleted_at IS NULL", categoryID).Scan(&category_id)

	if category_id == "" {
		helpers.HandleError(c, 404, "category not found")
		return
	}

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

	var countOfProduct uint

	priceSort := c.Query("price_sort")
	if priceSort != "" {
		if priceSort != "ASC" && priceSort != "DESC" {
			helpers.HandleError(c, 400, "price_sort is invalid")
			return
		}
	}

	var minPrice float32
	minPriceStr := c.Query("min_price")
	if minPriceStr == "" {
		helpers.HandleError(c, 400, "min_price is required")
		return
	} else {
		min_price, err := strconv.ParseFloat(minPriceStr, 32)
		if err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}
		if min_price < 0 {
			helpers.HandleError(c, 400, "min_price cannot be less than zero")
			return
		}
		minPrice = float32(min_price)
	}

	var maxPrice float32
	maxPriceStr := c.Query("max_price")
	if maxPriceStr == "" {
		helpers.HandleError(c, 400, "max_price is required")
		return
	} else {
		max_price, err := strconv.ParseFloat(maxPriceStr, 32)
		if err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}
		if max_price < 0 {
			helpers.HandleError(c, 400, "max_price cannot be less than zero")
			return
		}
		maxPrice = float32(max_price)
	}

	isDiscountStr := c.Query("is_discount")
	discount := -1
	isDiscount, err := strconv.ParseBool(isDiscountStr)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	if isDiscount {
		discount = 0
	}

	brendIDs := c.QueryArray("brend_ids")
	if len(brendIDs) != 0 {
		for _, v := range brendIDs {
			var brend_id string
			db.QueryRow(context.Background(), "SELECT id FROM brends WHERE id = $1 AND deleted_at IS NULL", v).Scan(&brend_id)

			if brend_id == "" {
				helpers.HandleError(c, 404, "brend not found")
				return
			}
		}
	} else {
		rowsBrend, err := db.Query(context.Background(), "SELECT id FROM brends WHERE deleted_at IS NULL")
		if err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}
		defer rowsBrend.Close()

		for rowsBrend.Next() {
			var brendID string
			if err := rowsBrend.Scan(&brendID); err != nil {
				helpers.HandleError(c, 400, err.Error())
				return
			}
			brendIDs = append(brendIDs, brendID)
		}
	}

	db.QueryRow(context.Background(), "SELECT COUNT(DISTINCT(p.id)) FROM products p LEFT JOIN category_product c ON p.id=c.product_id INNER JOIN translation_product tp ON tp.product_id = p.id WHERE p.is_visible = true AND p.brend_id = ANY($1) AND tp.lang_id = $2 AND c.category_id = $3 AND tp.deleted_at IS NULL AND p.amount > 0 AND p.limit_amount > 0 AND p.deleted_at IS NULL AND p.price >= $4 AND p.price <= $5 AND p.old_price > $6", pq.Array(brendIDs), langID, categoryID, minPrice, maxPrice, discount).Scan(&countOfProduct)

	var rowQuery string
	if priceSort == "" {
		rowQuery = "SELECT p.id,p.brend_id,p.price,p.old_price,p.amount,p.limit_amount,p.is_new,p.main_image,benefit FROM products p LEFT JOIN category_product c ON p.id=c.product_id INNER JOIN translation_product tp ON tp.product_id = p.id WHERE p.is_visible = true AND p.brend_id = ANY($1) AND tp.lang_id = $2 AND c.category_id = $3 AND tp.deleted_at IS NULL AND p.amount > 0 AND p.limit_amount > 0 AND p.deleted_at IS NULL AND p.price >= $4 AND p.price <= $5 AND p.old_price > $6 LIMIT $7 OFFSET $8"
	} else {
		if priceSort == "DESC" {
			rowQuery = "SELECT p.id,p.brend_id,p.price,p.old_price,p.amount,p.limit_amount,p.is_new,p.main_image,benefit FROM products p LEFT JOIN category_product c ON p.id=c.product_id INNER JOIN translation_product tp ON tp.product_id = p.id WHERE p.is_visible = true AND p.brend_id = ANY($1) AND tp.lang_id = $2 AND c.category_id = $3 AND tp.deleted_at IS NULL AND p.amount > 0 AND p.limit_amount > 0 AND p.deleted_at IS NULL AND p.price >= $4 AND p.price <= $5 AND p.old_price > $6 ORDER BY p.price DESC LIMIT $7 OFFSET $8"
		} else {
			rowQuery = "SELECT p.id,p.brend_id,p.price,p.old_price,p.amount,p.limit_amount,p.is_new,p.main_image,benefit FROM products p LEFT JOIN category_product c ON p.id=c.product_id INNER JOIN translation_product tp ON tp.product_id = p.id WHERE p.is_visible = true AND p.brend_id = ANY($1) AND tp.lang_id = $2 AND c.category_id = $3 AND tp.deleted_at IS NULL AND p.amount > 0 AND p.limit_amount > 0 AND p.deleted_at IS NULL AND p.price >= $4 AND p.price <= $5 AND p.old_price > $6 ORDER BY p.price ASC LIMIT $7 OFFSET $8"
		}
	}

	rowsProduct, err := db.Query(context.Background(), rowQuery, pq.Array(brendIDs), langID, categoryID, minPrice, maxPrice, discount, limit, offset)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer rowsProduct.Close()

	var products []LikeProduct
	for rowsProduct.Next() {
		var product LikeProduct
		if err := rowsProduct.Scan(&product.ID, &product.BrendID, &product.Price, &product.OldPrice, &product.Amount, &product.LimitAmount, &product.IsNew, &product.MainImage, &product.Benefit); err != nil {
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
			var trProduct models.TranslationProduct
			translation := make(map[string]models.TranslationProduct)
			db.QueryRow(context.Background(), "SELECT lang_id,name,description FROM translation_product WHERE product_id = $1 AND lang_id = $2 AND deleted_at IS NULL", product.ID, v.ID).Scan(&trProduct.LangID, &trProduct.Name, &trProduct.Description)
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
