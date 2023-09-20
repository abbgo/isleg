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
	"github.com/google/uuid"
	"github.com/gosimple/slug"
	"github.com/jackc/pgx/v5"
	"github.com/lib/pq"
	"gopkg.in/guregu/null.v4"
)

type ResultCategory struct {
	ID                    string          `json:"id,omitempty"`
	Image                 string          `json:"image,omitempty"`
	Name                  string          `json:"name,omitempty"`
	OrderNumber           uint            `json:"order_number,omitempty"`
	OrderNumberInHomePage uint            `json:"order_number_in_home_page,omitempty"`
	ResultCategor         []ResultCategor `json:"child_category,omitempty"`
}

type ResultCategor struct {
	ID                    string         `json:"id,omitempty"`
	Name                  string         `json:"name,omitempty"`
	OrderNumber           uint           `json:"order_number,omitempty"`
	OrderNumberInHomePage uint           `json:"order_number_in_home_page,omitempty"`
	ResultCatego          []ResultCatego `json:"child_category,omitempty"`
}

type ResultCatego struct {
	ID                    string `json:"id,omitempty"`
	Name                  string `json:"name,omitempty"`
	OrderNumber           uint   `json:"order_number,omitempty"`
	OrderNumberInHomePage uint   `json:"order_number_in_home_page,omitempty"`
}

type Category struct {
	ID       string    `json:"id,omitempty"`
	Name     string    `json:"name,omitempty"`
	Image    string    `json:"image,omitempty"`
	Products []Product `json:"products,omitempty"`
}

type Product struct {
	ID           string                                 `json:"id,omitempty"`
	Price        float64                                `json:"price,omitempty"`
	OldPrice     float64                                `json:"old_price,omitempty"`
	Percentage   float64                                `json:"percentage,omitempty"`
	MainImage    string                                 `json:"main_image,omitempty"`
	Brend        Brend                                  `json:"brend,omitempty"`
	LimitAmount  int                                    `json:"limit_amount,omitempty"`
	Amount       int                                    `json:"amount,omitempty"`
	IsNew        bool                                   `json:"is_new,omitempty"`
	Benefit      null.Float                             `json:"-"`
	Translations []map[string]models.TranslationProduct `json:"translations"`
	Code         null.String                            `json:"code,omitempty"`
}

type Brend struct {
	ID   uuid.NullUUID `json:"id,omitempty"`
	Name null.String   `json:"name,omitempty"`
}

func CreateCategory(c *gin.Context) {
	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	var category models.Category
	if err := c.BindJSON(&category); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	var categoryID string
	var parent_category_id interface{}

	// validate other data of category
	if err := models.ValidateCategory("", category.ParentCategoryID.String, category.Image, "create", category.OrderNumber, category.OrderNumberInHomePage, category.IsHomeCategory); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	// CREATE CATEGORY
	if category.ParentCategoryID.String != "" {
		parent_category_id = category.ParentCategoryID.String
	} else {
		parent_category_id = nil
	}

	// add data to categories table
	db.QueryRow(context.Background(), "INSERT INTO categories (parent_category_id,image,is_home_category,order_number,is_visible) VALUES ($1,$2,$3,$4,$5) RETURNING id", parent_category_id, category.Image, category.IsHomeCategory, category.OrderNumber, category.IsVisible).Scan(&categoryID)
	langID, err := GetLangID("tm")
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	var nameTm string

	// CREATE TRANSLATION CATEGORY
	for _, v := range category.TranslationCategory {
		if langID == v.LangID.UUID.String() {
			nameTm = v.Name
		}
		_, err := db.Exec(context.Background(), "INSERT INTO translation_category (lang_id,category_id,name,slug) VALUES ($1,$2,$3,$4)", v.LangID, categoryID, v.Name, slug.MakeLang(v.Name, "en"))
		if err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}
	}

	if parent_category_id != "" {
		c.JSON(http.StatusOK, gin.H{
			"status": true,
			"id":     categoryID,
			"name":   nameTm,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "data successfully added",
	})
}

// func UpdateParentCategoriesOrderNumber(c *gin.Context) {
// 	// initialize database connection
// 	db, err := config.ConnDB()
// 	if err != nil {
// 		helpers.HandleError(c, 400, err.Error())
// 		return
// 	}
// 	defer db.Close()

// 	rowsCategories, err := db.Query(context.Background(), "SELECT id FROM categories WHERE parent_category_id IS NULL")
// 	if err != nil {
// 		helpers.HandleError(c, 400, err.Error())
// 		return
// 	}
// 	defer rowsCategories.Close()

// 	var countOfCategory uint
// 	for rowsCategories.Next() {
// 		countOfCategory++

// 		var categoryID string
// 		if err := rowsCategories.Scan(&categoryID); err != nil {
// 			helpers.HandleError(c, 400, err.Error())
// 			return
// 		}

// 		_, err = db.Exec(context.Background(), "UPDATE categories SET order_number = $1 WHERE id = $2", countOfCategory, categoryID)
// 		if err != nil {
// 			helpers.HandleError(c, 400, err.Error())
// 			return
// 		}
// 	}
// }

// func UpdateHomePageCategoriesOrderNumber(c *gin.Context) {
// 	// initialize database connection
// 	db, err := config.ConnDB()
// 	if err != nil {
// 		helpers.HandleError(c, 400, err.Error())
// 		return
// 	}
// 	defer db.Close()

// 	rowsCategories, err := db.Query(context.Background(), "SELECT id FROM categories WHERE is_home_category = true")
// 	if err != nil {
// 		helpers.HandleError(c, 400, err.Error())
// 		return
// 	}
// 	defer rowsCategories.Close()

// 	var countOfCategory uint
// 	for rowsCategories.Next() {
// 		countOfCategory++

// 		var categoryID string
// 		if err := rowsCategories.Scan(&categoryID); err != nil {
// 			helpers.HandleError(c, 400, err.Error())
// 			return
// 		}

// 		_, err = db.Exec(context.Background(), "UPDATE categories SET order_number_in_home_page = $1 WHERE id = $2", countOfCategory, categoryID)
// 		if err != nil {
// 			helpers.HandleError(c, 400, err.Error())
// 			return
// 		}
// 	}
// }

func UpdateCategoryByID(c *gin.Context) {
	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	// get id from request parameter
	ID := c.Param("id")

	var category models.Category
	if err := c.BindJSON(&category); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	// var fileName string
	var parent_category_id interface{}
	if err := models.ValidateCategory(ID, category.ParentCategoryID.String, "", "update", category.OrderNumber, category.OrderNumberInHomePage, category.IsHomeCategory); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	if category.ParentCategoryID.String != "" && category.Image != "" {
		if err := os.Remove(pkg.ServerPath + category.Image); err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}
	}

	// UPDATE CATEGORY
	if category.ParentCategoryID.String != "" {
		parent_category_id = category.ParentCategoryID.String
	} else {
		parent_category_id = nil
	}

	if category.Image != "" {
		_, err := db.Exec(context.Background(), "UPDATE categories SET parent_category_id = $1, image = $2, is_home_category = $3 , order_number = $5 , is_visible = $6 WHERE id = $4", parent_category_id, category.Image, category.IsHomeCategory, ID, category.OrderNumber, category.IsVisible)
		if err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}
	} else {
		_, err := db.Exec(context.Background(), "UPDATE categories SET parent_category_id = $1, is_home_category = $2 , order_number = $4 , is_visible = $5 WHERE id = $3", parent_category_id, category.IsHomeCategory, ID, category.OrderNumber, category.IsVisible)
		if err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}
	}

	langID, err := GetLangID("tm")
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	var nameTm string

	// UPDATE TRANSLATION CATEGORY
	for _, v := range category.TranslationCategory {
		if langID == v.LangID.UUID.String() {
			nameTm = v.Name
		}
		_, err := db.Exec(context.Background(), "UPDATE translation_category SET name = $1 , slug = $4 WHERE lang_id = $2 AND category_id = $3", v.Name, v.LangID, ID, slug.MakeLang(v.Name, "en"))
		if err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}
	}

	if parent_category_id != "" {
		c.JSON(http.StatusOK, gin.H{
			"status": true,
			"id":     ID,
			"name":   nameTm,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "data successfully updated",
	})
}

func GetCategoryByID(c *gin.Context) {
	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	// get id from request parameter
	ID := c.Param("id")

	// check id and get data from daabase
	var category models.Category
	db.QueryRow(context.Background(), "SELECT id,parent_category_id,image,is_home_category FROM categories WHERE id = $1 AND deleted_at IS NULL", ID).Scan(&category.ID, &category.ParentCategoryID, &category.Image, &category.IsHomeCategory)
	if category.ID == "" {
		helpers.HandleError(c, 404, "category not found")
		return
	}

	rowsTrCategory, err := db.Query(context.Background(), "SELECT lang_id,name FROM translation_category WHERE category_id = $1 AND deleted_at IS NULL", ID)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer rowsTrCategory.Close()

	var translations []models.TranslationCategory
	for rowsTrCategory.Next() {
		var translation models.TranslationCategory
		if err := rowsTrCategory.Scan(&translation.LangID, &translation.Name); err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}
		translations = append(translations, translation)
	}

	category.TranslationCategory = translations

	c.JSON(http.StatusOK, gin.H{
		"status":   true,
		"category": category,
	})
}

func GetCategoryByIDWithChild(c *gin.Context) {
	langID, err := GetLangID("tm")
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	ID := c.Param("id")

	// get all category where parent category id is null
	var result ResultCategory
	rows, err := db.Query(context.Background(), "SELECT c.id,c.image,tc.name FROM categories c LEFT JOIN translation_category tc ON c.id=tc.category_id WHERE tc.lang_id = $1 AND c.id = $2 AND c.parent_category_id IS NULL AND c.deleted_at IS NULL AND tc.deleted_at IS NULL ORDER BY c.created_at DESC", langID, ID)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&result.ID, &result.Image, &result.Name); err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}

		// get all category where parent category id equal category id
		rowss, err := db.Query(context.Background(), "SELECT c.id,tc.name FROM categories c LEFT JOIN translation_category tc ON c.id=tc.category_id WHERE tc.lang_id = $1 AND c.parent_category_id = $2 AND c.deleted_at IS NULL AND tc.deleted_at IS NULL ORDER BY c.created_at DESC", langID, result.ID)
		if err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}
		defer rowss.Close()

		var resuls []ResultCategor
		for rowss.Next() {
			var resul ResultCategor
			if err := rowss.Scan(&resul.ID, &resul.Name); err != nil {
				helpers.HandleError(c, 400, err.Error())
				return
			}

			// get all category where parent category id equal category id
			rowsss, err := db.Query(context.Background(), "SELECT c.id,tc.name FROM categories c LEFT JOIN translation_category tc ON c.id=tc.category_id WHERE tc.lang_id = $1 AND c.parent_category_id =$2 AND c.deleted_at IS NULL AND tc.deleted_at IS NULL ORDER BY c.created_at DESC", langID, resul.ID)
			if err != nil {
				helpers.HandleError(c, 400, err.Error())
				return
			}
			defer rowsss.Close()

			var resus []ResultCatego
			for rowsss.Next() {
				var resu ResultCatego
				if err := rowsss.Scan(&resu.ID, &resu.Name); err != nil {
					helpers.HandleError(c, 400, err.Error())
					return
				}

				resus = append(resus, resu)
			}
			resul.ResultCatego = resus
			resuls = append(resuls, resul)
		}
		result.ResultCategor = resuls
	}

	c.JSON(http.StatusOK, gin.H{
		"status":   true,
		"category": result,
	})
}

func GetAllCategory(c *gin.Context) {
	langID, err := GetLangID("tm")
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	// get all category where parent category id is null
	rows, err := db.Query(context.Background(), "SELECT DISTINCT ON (c.id,c.created_at) c.id,c.image,tc.name FROM categories c LEFT JOIN translation_category tc ON c.id=tc.category_id INNER JOIN category_product cp ON c.id=cp.category_id WHERE tc.lang_id = $1 AND c.parent_category_id IS NULL AND cp.deleted_at IS NOT NULL ORDER BY c.created_at DESC", langID)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer rows.Close()

	var results []ResultCategory
	for rows.Next() {
		var result ResultCategory
		if err := rows.Scan(&result.ID, &result.Image, &result.Name); err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}

		// get all category where parent category id equal category id
		rowss, err := db.Query(context.Background(), "SELECT DISTINCT ON (c.id,c.created_at) c.id,tc.name FROM categories c LEFT JOIN translation_category tc ON c.id=tc.category_id INNER JOIN category_product cp ON cp.category_id=c.id WHERE tc.lang_id = $1 AND c.parent_category_id = $2 AND cp.deleted_at IS NOT NULL ORDER BY c.created_at DESC", langID, result.ID)
		if err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}
		defer rowss.Close()

		var resuls []ResultCategor
		for rowss.Next() {
			var resul ResultCategor
			if err := rowss.Scan(&resul.ID, &resul.Name); err != nil {
				helpers.HandleError(c, 400, err.Error())
				return
			}

			// get all category where parent category id equal category id
			rowsss, err := db.Query(context.Background(), "SELECT categories.id,translation_category.name FROM categories LEFT JOIN translation_category ON categories.id=translation_category.category_id WHERE translation_category.lang_id = $1 AND categories.parent_category_id =$2 ORDER BY categories.created_at DESC", langID, resul.ID)
			if err != nil {
				helpers.HandleError(c, 400, err.Error())
				return
			}
			defer rowsss.Close()

			var resus []ResultCatego
			for rowsss.Next() {
				var resu ResultCatego
				if err := rowsss.Scan(&resu.ID, &resu.Name); err != nil {
					helpers.HandleError(c, 400, err.Error())
					return
				}

				resus = append(resus, resu)
			}
			resul.ResultCatego = resus
			resuls = append(resuls, resul)
		}
		result.ResultCategor = resuls
		results = append(results, result)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":     true,
		"categories": results,
	})
}

func GetDeletedCategories(c *gin.Context) {
	langID, err := GetLangID("tm")
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	// get all category where parent category id is null
	rows, err := db.Query(context.Background(), "SELECT c.id,c.image,tc.name FROM categories c LEFT JOIN translation_category tc ON c.id=tc.category_id WHERE tc.lang_id = $1 AND c.deleted_at IS NOT NULL ORDER BY c.deleted_at DESC", langID)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer rows.Close()

	var results []ResultCategory
	for rows.Next() {
		var result ResultCategory
		if err := rows.Scan(&result.ID, &result.Image, &result.Name); err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}

		// get all category where parent category id equal category id
		rowss, err := db.Query(context.Background(), "SELECT c.id,tc.name FROM categories c LEFT JOIN translation_category tc ON c.id=tc.category_id WHERE tc.lang_id = $1 AND c.parent_category_id = $2 AND c.deleted_at IS NOT NULL ORDER BY c.deleted_at DESC", langID, result.ID)
		if err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}
		defer rowss.Close()

		var resuls []ResultCategor
		for rowss.Next() {
			var resul ResultCategor
			if err := rowss.Scan(&resul.ID, &resul.Name); err != nil {
				helpers.HandleError(c, 400, err.Error())
				return
			}

			// get all category where parent category id equal category id
			rowsss, err := db.Query(context.Background(), "SELECT c.id,tc.name FROM categories c LEFT JOIN translation_category tc ON c.id=tc.category_id WHERE tc.lang_id = $1 AND c.parent_category_id =$2 AND c.deleted_at IS NOT NULL ORDER BY c.deleted_at DESC", langID, resul.ID)
			if err != nil {
				helpers.HandleError(c, 400, err.Error())
				return
			}
			defer rowsss.Close()

			var resus []ResultCatego
			for rowsss.Next() {
				var resu ResultCatego
				if err := rowsss.Scan(&resu.ID, &resu.Name); err != nil {
					helpers.HandleError(c, 400, err.Error())
					return
				}

				resus = append(resus, resu)
			}
			resul.ResultCatego = resus
			resuls = append(resuls, resul)
		}
		result.ResultCategor = resuls
		results = append(results, result)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":     true,
		"categories": results,
	})
}

func GetCategories(c *gin.Context) {
	langID, err := GetLangID("tm")
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	searchQuery := c.Query("search")
	var searchStr, search string
	if searchQuery != "" {
		incomingsSarch := slug.MakeLang(searchQuery, "en")
		search = strings.ReplaceAll(incomingsSarch, "-", " | ")
		searchStr = fmt.Sprintf("%%%s%%", search)
	}

	// get limit from param
	limitStr := c.Param("limit")
	var limit uint64
	if limitStr != "" {
		limit, err = strconv.ParseUint(limitStr, 10, 64)
		if err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}
	}

	// get page from param
	pageStr := c.Param("page")
	var page uint64
	if pageStr != "" {
		page, err = strconv.ParseUint(pageStr, 10, 64)
		if err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}
	}

	// get all category from category controller
	categories, countOfCatagories, err := GetAllCategoryForHeader(langID, search, searchStr, uint(limit), uint(page))
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":     true,
		"categories": categories,
		"total":      countOfCatagories,
	})
}

func GetCategoriesForAdmin(c *gin.Context) {
	langID, err := GetLangID("tm")
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	categories, _, err := GetAllCategoryForHeader(langID, "", "", 0, 0)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":     true,
		"categories": categories,
	})
}

func GetAllCategoryForHeader(langID, search, searchStr string, limit, page uint) ([]ResultCategory, uint, error) {
	db, err := config.ConnDB()
	if err != nil {
		return []ResultCategory{}, 0, err
	}
	defer db.Close()

	var countOfCategories uint
	var offset interface{}
	if limit != 0 && page != 0 {
		offset = limit * (page - 1)
	} else {
		offset = nil
	}

	// get all category where parent category id is null
	var rows, rowsCount pgx.Rows
	if search == "" {
		if offset != nil {
			rowsCount, err = db.Query(context.Background(), "SELECT COUNT(*) FROM categories c LEFT JOIN translation_category tc ON c.id=tc.category_id WHERE tc.lang_id = $1 AND c.parent_category_id IS NULL AND tc.deleted_at IS NULL AND c.deleted_at IS NULL", langID)
			if err != nil {
				return []ResultCategory{}, 0, err
			}

			rows, err = db.Query(context.Background(), "SELECT c.id,c.image,tc.name,c.order_number,c.order_number_in_home_page FROM categories c LEFT JOIN translation_category tc ON c.id=tc.category_id WHERE tc.lang_id = $1 AND c.parent_category_id IS NULL AND tc.deleted_at IS NULL AND c.deleted_at IS NULL ORDER BY c.order_number  ASC LIMIT $2 OFFSET $3", langID, limit, offset)
			if err != nil {
				return []ResultCategory{}, 0, err
			}
		} else {
			rows, err = db.Query(context.Background(), "SELECT c.id,c.image,tc.name,c.order_number,c.order_number_in_home_page FROM categories c LEFT JOIN translation_category tc ON c.id=tc.category_id WHERE tc.lang_id = $1 AND c.parent_category_id IS NULL AND tc.deleted_at IS NULL AND c.deleted_at IS NULL ORDER BY c.order_number ASC", langID)
			if err != nil {
				return []ResultCategory{}, 0, err
			}
		}
	} else {
		if offset != nil {
			rowsCount, err = db.Query(context.Background(), "SELECT COUNT(DISTINCT(c.id)) FROM categories c LEFT JOIN translation_category tc ON c.id=tc.category_id WHERE to_tsvector(tc.slug) @@ to_tsquery($2) OR tc.slug LIKE $3 AND tc.lang_id = $1 AND c.parent_category_id IS NULL AND tc.deleted_at IS NULL AND c.deleted_at IS NULL", langID, search, searchStr)
			if err != nil {
				return []ResultCategory{}, 0, err
			}

			rows, err = db.Query(context.Background(), "SELECT DISTINCT ON (c.id,c.created_at) c.id,c.image,tc.name,c.order_number,c.order_number_in_home_page FROM categories c LEFT JOIN translation_category tc ON c.id=tc.category_id WHERE to_tsvector(tc.slug) @@ to_tsquery($2) OR tc.slug LIKE $3 AND tc.lang_id = $1 AND c.parent_category_id IS NULL AND tc.deleted_at IS NULL AND c.deleted_at IS NULL ORDER BY c.order_number ASC LIMIT $4 OFFSET $5", langID, search, searchStr, limit, offset)
			if err != nil {
				return []ResultCategory{}, 0, err
			}
		} else {
			rows, err = db.Query(context.Background(), "SELECT DISTINCT ON (c.id,c.created_at) c.id,c.image,tc.name,c.order_number,c.order_number_in_home_page FROM categories c LEFT JOIN translation_category tc ON c.id=tc.category_id WHERE to_tsvector(tc.slug) @@ to_tsquery($2) OR tc.slug LIKE $3 AND tc.lang_id = $1 AND c.parent_category_id IS NULL AND tc.deleted_at IS NULL AND c.deleted_at IS NULL ORDER BY c.order_number ASC", langID, search, searchStr)
			if err != nil {
				return []ResultCategory{}, 0, err
			}
		}
	}
	defer rows.Close()

	var results []ResultCategory
	for rows.Next() {
		var result ResultCategory
		if err := rows.Scan(&result.ID, &result.Image, &result.Name, &result.OrderNumber, &result.OrderNumberInHomePage); err != nil {
			return []ResultCategory{}, 0, err
		}

		// get all category where parent category id equal category id
		rowss, err := db.Query(context.Background(), "SELECT c.id,tc.name,c.order_number,c.order_number_in_home_page FROM categories c LEFT JOIN translation_category tc ON c.id=tc.category_id WHERE tc.lang_id = $1 AND c.parent_category_id = $2 AND tc.deleted_at IS NULL AND c.deleted_at IS NULL ORDER BY tc.slug ASC", langID, result.ID)
		if err != nil {
			return []ResultCategory{}, 0, err
		}
		defer rowss.Close()

		var resuls []ResultCategor
		for rowss.Next() {
			var resul ResultCategor
			if err := rowss.Scan(&resul.ID, &resul.Name, &resul.OrderNumber, &resul.OrderNumberInHomePage); err != nil {
				return []ResultCategory{}, 0, err
			}

			// get all category where parent category id equal category id
			rowsss, err := db.Query(context.Background(), "SELECT c.id,tc.name,c.order_number,c.order_number_in_home_page FROM categories c LEFT JOIN translation_category tc ON c.id=tc.category_id WHERE tc.lang_id = $1 AND c.parent_category_id =$2 AND tc.deleted_at IS NULL AND c.deleted_at IS NULL ORDER BY tc.slug ASC", langID, resul.ID)
			if err != nil {
				return []ResultCategory{}, 0, err
			}
			defer rowsss.Close()

			var resus []ResultCatego
			for rowsss.Next() {
				var resu ResultCatego
				if err := rowsss.Scan(&resu.ID, &resu.Name, &resu.OrderNumber, &resu.OrderNumberInHomePage); err != nil {
					return []ResultCategory{}, 0, err
				}

				resus = append(resus, resu)
			}
			resul.ResultCatego = resus
			resuls = append(resuls, resul)
		}
		result.ResultCategor = resuls
		results = append(results, result)
	}

	if offset != nil {
		for rowsCount.Next() {
			if err := rowsCount.Scan(&countOfCategories); err != nil {
				if err != nil {
					return []ResultCategory{}, 0, err
				}
			}
		}
	}

	return results, countOfCategories, nil
}

func DeleteCategoryByID(c *gin.Context) {
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
	var category_id string
	db.QueryRow(context.Background(), "SELECT id FROM categories WHERE id = $1 AND deleted_at IS NULL", ID).Scan(&category_id)
	if category_id == "" {
		helpers.HandleError(c, 404, "category not found")
		return
	}

	_, err = db.Exec(context.Background(), "CALL delete_category($1)", ID)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	_, err = db.Exec(context.Background(), "CALL after_delete_category($1)", ID)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "data successfully deleted",
	})
}

func RestoreCategoryByID(c *gin.Context) {
	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	// get id from request parameter
	ID := c.Param("id")

	// check ids
	var category_id string
	db.QueryRow(context.Background(), "SELECT id FROM categories WHERE id = $1 AND deleted_at IS NOT NULL", ID).Scan(&category_id)
	if category_id == "" {
		helpers.HandleError(c, 404, "category not found")
		return
	}

	_, err = db.Exec(context.Background(), "CALL restore_category($1)", ID)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	_, err = db.Exec(context.Background(), "CALL after_restore_category($1)", ID)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "data successfully restored",
	})
}

func DeletePermanentlyCategoryByID(c *gin.Context) {
	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	// get id from request parameter
	ID := c.Param("id")

	// check id and get image of categories
	var category_id, image string
	db.QueryRow(context.Background(), "SELECT id,image FROM categories WHERE id = $1 AND deleted_at IS NOT NULL", ID).Scan(&category_id, &image)
	if category_id == "" {
		helpers.HandleError(c, 404, "category not found")
		return
	}

	// kategoriyanyn suraty uploads papkadan pozulyar
	if image != "" {
		if err := os.Remove(pkg.ServerPath + image); err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}
	}

	// kategoriyanyn produktalarynyn main suratlaryn direktlary bazadan alynyar
	rowsMainImageProduct, err := db.Query(context.Background(), "SELECT p.id,p.main_image FROM products p INNER JOIN category_product c ON c.product_id=p.id WHERE c.category_id = $1", ID)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer rowsMainImageProduct.Close()

	var mainImages []models.Product
	for rowsMainImageProduct.Next() {
		var mainImage models.Product
		if err := rowsMainImageProduct.Scan(&mainImage.ID, &mainImage.MainImage); err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}
		mainImages = append(mainImages, mainImage)
	}

	// kategoriyanyn produktalarynyn main suratlary uploads papkadan pozulyar
	for _, v := range mainImages {
		if err := os.Remove(pkg.ServerPath + v.MainImage); err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}
	}

	// kategoriyanyn produktalarynyn suratlarynyn direktlary bazadan alynyar
	rowsImagesProduct, err := db.Query(context.Background(), "SELECT i.image FROM products p INNER JOIN category_product c ON c.product_id=p.id INNER JOIN images i ON i.product_id = p.id WHERE c.category_id = $1", ID)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer rowsImagesProduct.Close()

	var images []models.Images
	for rowsImagesProduct.Next() {
		var image models.Images
		if err := rowsImagesProduct.Scan(&image.Image); err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}
		images = append(images, image)
	}

	// kategoriyanyn produktalarynyn suratlaryny uploads papkadan pozyaryn
	for _, v := range images {
		if err := os.Remove(pkg.ServerPath + v.Image); err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}
	}

	// kategoriya degisli harytlar bazadan pozulyar
	_, err = db.Exec(context.Background(), "DELETE FROM products USING category_product WHERE category_product.product_id = products.id AND category_product.category_id = $1", category_id)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	// haryt pozulanda sol haryda degisli likelar hem pozulmaly
	for _, v := range mainImages {
		_, err := db.Exec(context.Background(), "DELETE FROM likes WHERE product_id = $1", v.ID)
		if err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}
	}

	// haryt pozulanda sol haryda degisli sebetdaki harytlar hem pozulmaly
	for _, v := range mainImages {
		_, err := db.Exec(context.Background(), "DELETE FROM cart WHERE product_id = $1", v.ID)
		if err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}
	}

	// haryt pozulanda sol haryda degisli sargyt edilen harytlar hem pozulmaly
	for _, v := range mainImages {
		_, err := db.Exec(context.Background(), "DELETE FROM ordered_products WHERE product_id = $1", v.ID)
		if err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}
	}

	// kategoriya pozulanda category_product tablisadaky maglumatlar hem pozulmaly
	_, err = db.Exec(context.Background(), "DELETE FROM category_product WHERE category_id = $1", category_id)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	// kategoriyanyn child kategoriyalarynyn id - leri alynyar database - den
	rowChildCategory, err := db.Query(context.Background(), "SELECT id FROM categories WHERE parent_category_id = $1", ID)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer rowChildCategory.Close()

	var childCategoryIDS []string
	for rowChildCategory.Next() {
		var childCategoryID string
		if err := rowChildCategory.Scan(&childCategoryID); err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}
		childCategoryIDS = append(childCategoryIDS, childCategoryID)
	}

	// child kategoriya degisli harytlaryn suratlarynyn direktleri bazadan alynyar we uploads papkadam pozulyar
	for _, v := range childCategoryIDS {
		rowPrdcs, err := db.Query(context.Background(), "SELECT p.id,p.main_image FROM products p INNER JOIN category_product c ON c.product_id=p.id WHERE c.category_id = $1", v)
		if err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}
		defer rowPrdcs.Close()

		var childMainImages []models.Product
		for rowPrdcs.Next() {
			var childMainImage models.Product
			if err := rowPrdcs.Scan(&childMainImage.ID, &childMainImage.MainImage); err != nil {
				helpers.HandleError(c, 400, err.Error())
				return
			}
			childMainImages = append(childMainImages, childMainImage)
		}

		for _, v := range childMainImages {
			if err := os.Remove(pkg.ServerPath + v.MainImage); err != nil {
				helpers.HandleError(c, 400, err.Error())
				return
			}
		}

		rowsChildImagesProduct, err := db.Query(context.Background(), "SELECT i.image FROM products p INNER JOIN category_product c ON c.product_id=p.id INNER JOIN images i ON i.product_id = p.id WHERE c.category_id = $1", v)
		if err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}
		defer rowsChildImagesProduct.Close()

		var childImages []models.Images
		for rowsChildImagesProduct.Next() {
			var childImage models.Images
			if err := rowsChildImagesProduct.Scan(&childImage.Image); err != nil {
				helpers.HandleError(c, 400, err.Error())
				return
			}
			childImages = append(childImages, childImage)
		}

		for _, v := range childImages {
			if err := os.Remove(pkg.ServerPath + v.Image); err != nil {
				helpers.HandleError(c, 400, err.Error())
				return
			}
		}

		// child kategoriya degisli harytlar bazadan pozulyar
		_, err = db.Exec(context.Background(), "DELETE FROM products USING category_product WHERE category_product.product_id = products.id AND category_product.category_id = $1", v)
		if err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}

		// haryt pozulanda sol haryda degisli likelar hem pozulmaly
		for _, v := range childMainImages {
			_, err := db.Exec(context.Background(), "DELETE FROM likes WHERE product_id = $1", v.ID)
			if err != nil {
				helpers.HandleError(c, 400, err.Error())
				return
			}
		}

		// haryt pozulanda sol haryda degisli sebetdaki harytlar hem pozulmaly
		for _, v := range childMainImages {
			_, err := db.Exec(context.Background(), "DELETE FROM cart WHERE product_id = $1", v.ID)
			if err != nil {
				helpers.HandleError(c, 400, err.Error())
				return
			}
		}

		// haryt pozulanda sol haryda degisli sargyt edilen harytlar hem pozulmaly
		for _, v := range childMainImages {
			_, err := db.Exec(context.Background(), "DELETE FROM ordered_products WHERE product_id = $1", v.ID)
			if err != nil {
				helpers.HandleError(c, 400, err.Error())
				return
			}
		}

		// kategoriya pozulanda category_product tablisadaky maglumatlar hem pozulmaly
		_, err = db.Exec(context.Background(), "DELETE FROM category_product WHERE category_id = $1", v)
		if err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}

	}

	// in sonunda kategpriyanyn ozi pozulyar
	_, err = db.Exec(context.Background(), "DELETE FROM categories WHERE id = $1", ID)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "category successfully deleted",
	})
}

func GetOneCategoryWithProducts(c *gin.Context) {
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	var countOfProducts uint64

	langID, err := CheckLanguage(c)
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

	categoryID := c.Param("id")

	// get category where id equal categiryID
	categoryRow, err := db.Query(context.Background(), "SELECT c.id,c.image,t.name FROM categories c LEFT JOIN translation_category t ON c.id=t.category_id WHERE t.lang_id = $1 AND c.id = $2 AND c.deleted_at IS NULL AND t.deleted_at IS NULL", langID, categoryID)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer categoryRow.Close()

	var category Category
	var productIDS []string
	for categoryRow.Next() {
		if err := categoryRow.Scan(&category.ID, &category.Image, &category.Name); err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}

		if category.ID == "" {
			helpers.HandleError(c, 404, "category not found")
			return
		}

		// get count product where product_id equal categoryID
		productCount, err := db.Query(context.Background(), "SELECT COUNT(DISTINCT p.id) FROM products p LEFT JOIN category_product c ON p.id=c.product_id WHERE c.category_id = $1 AND p.amount > 0 AND p.limit_amount > 0 AND p.deleted_at IS NULL AND c.deleted_at IS NULL", categoryID)
		if err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}
		defer productCount.Close()

		for productCount.Next() {
			if err := productCount.Scan(&countOfProducts); err != nil {
				if err != nil {
					helpers.HandleError(c, 400, err.Error())
					return
				}
			}
		}

		// get all product where category id equal categoryID
		productRows, err := db.Query(context.Background(), "SELECT DISTINCT ON (p.created_at) p.id,p.price,p.old_price,p.limit_amount,p.is_new,p.amount,p.main_image,p.benefit,p.code FROM products p LEFT JOIN category_product c ON p.id=c.product_id WHERE c.category_id = $1 AND p.amount > 0 AND p.limit_amount > 0 AND p.deleted_at IS NULL AND c.deleted_at IS NULL ORDER BY p.created_at ASC LIMIT $2 OFFSET $3", categoryID, limit, offset)
		if err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}
		defer productRows.Close()

		var products []Product
		for productRows.Next() {
			var product Product
			if err := productRows.Scan(&product.ID, &product.Price, &product.OldPrice, &product.LimitAmount, &product.IsNew, &product.Amount, &product.MainImage, &product.Benefit, &product.Code); err != nil {
				helpers.HandleError(c, 400, err.Error())
				return
			}

			productIDS = append(productIDS, product.ID)

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
			brendRows, err := db.Query(context.Background(), "SELECT b.id,b.name FROM products p LEFT JOIN brends b ON p.brend_id=b.id WHERE p.id = $1 AND p.deleted_at IS NULL AND b.deleted_at IS NULL", product.ID)
			if err != nil {
				helpers.HandleError(c, 400, err.Error())
				return
			}
			defer brendRows.Close()

			var brend Brend
			for brendRows.Next() {
				if err := brendRows.Scan(&brend.ID, &brend.Name); err != nil {
					helpers.HandleError(c, 400, err.Error())
					return
				}
			}
			product.Brend = brend
			products = append(products, product)

		}
		category.Products = products
	}

	rowsBrend, err := db.Query(context.Background(), "SELECT DISTINCT(b.id),b.name FROM brends b INNER JOIN products p ON p.brend_id = b.id WHERE p.id = ANY($1) AND b.deleted_at IS NULL AND p.deleted_at IS NULL", pq.Array(productIDS))
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer rowsBrend.Close()

	var brends []Brend
	for rowsBrend.Next() {
		var brend Brend
		if err := rowsBrend.Scan(&brend.ID, &brend.Name); err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}
		brends = append(brends, brend)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":            true,
		"category":          category,
		"count_of_products": countOfProducts,
		"brends":            brends,
	})
}

func GetOneCategoryWithDeletedProducts(c *gin.Context) {
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	var countOfProducts uint64

	langID, err := CheckLanguage(c)
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
	categoryID := c.Param("id")

	// get category where id equal categiryID
	categoryRow, err := db.Query(context.Background(), "SELECT c.id,c.image,t.name FROM categories c LEFT JOIN translation_category t ON c.id=t.category_id WHERE t.lang_id = $1 AND c.id = $2 AND c.deleted_at IS NULL AND t.deleted_at IS NULL", langID, categoryID)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer categoryRow.Close()

	var category Category
	for categoryRow.Next() {
		if err := categoryRow.Scan(&category.ID, &category.Image, &category.Name); err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}

		if category.ID == "" {
			helpers.HandleError(c, 404, "category not found")
			return
		}

		// get count product where product_id equal categoryID
		productCount, err := db.Query(context.Background(), "SELECT COUNT(DISTINCT p.id) FROM products p LEFT JOIN category_product c ON p.id=c.product_id WHERE c.category_id = $1 AND p.amount > 0 AND p.limit_amount > 0 AND p.deleted_at IS NOT NULL", categoryID)
		if err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}
		defer productCount.Close()

		for productCount.Next() {
			if err := productCount.Scan(&countOfProducts); err != nil {
				if err != nil {
					helpers.HandleError(c, 400, err.Error())
					return
				}
			}
		}

		// get all product where category id equal categoryID
		productRows, err := db.Query(context.Background(), "SELECT DISTINCT ON (p.created_at) p.id,p.price,p.old_price,p.limit_amount,p.is_new,p.amount,p.main_image,p.benefit,p.code FROM products p LEFT JOIN category_product c ON p.id=c.product_id WHERE c.category_id = $1 AND p.amount > 0 AND p.limit_amount > 0 AND p.deleted_at IS NOT NULL ORDER BY p.created_at ASC LIMIT $2 OFFSET $3", categoryID, limit, offset)
		if err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}
		defer productRows.Close()

		var products []Product
		for productRows.Next() {
			var product Product
			if err := productRows.Scan(&product.ID, &product.Price, &product.OldPrice, &product.LimitAmount, &product.IsNew, &product.Amount, &product.MainImage, &product.Benefit, &product.Code); err != nil {
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
		category.Products = products
	}
	c.JSON(http.StatusOK, gin.H{
		"status":            true,
		"category":          category,
		"count_of_products": countOfProducts,
	})
}

func SearchCategory(c *gin.Context) {
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	// GET language id
	langID, err := GetLangID("tm")
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	incomingsSarch := slug.MakeLang(c.Query("search"), "en")
	search := strings.ReplaceAll(incomingsSarch, "-", " | ")
	searchStr := fmt.Sprintf("%%%s%%", search)

	rowsCategory, err := db.Query(context.Background(), "SELECT DISTINCT ON (c.created_at) c.id,c.image,tc.name FROM categories c inner join translation_category tc on tc.category_id = c.id WHERE to_tsvector(tc.slug) @@ to_tsquery($1) OR tc.slug LIKE $3 AND tc.lang_id = $2 AND tc.deleted_at IS NULL AND c.deleted_at IS NULL ORDER BY c.created_at DESC", search, langID, searchStr)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer rowsCategory.Close()

	var categories []ResultCategory
	for rowsCategory.Next() {
		var category ResultCategory
		if err := rowsCategory.Scan(&category.ID, &category.Name, &category.Name); err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}

		rowss, err := db.Query(context.Background(), "SELECT c.id,tc.name FROM categories c LEFT JOIN translation_category tc ON c.id=tc.category_id WHERE tc.lang_id = $1 AND c.parent_category_id = $2 AND c.deleted_at IS NULL ORDER BY c.created_at DESC", langID, category.ID)
		if err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}
		defer rowss.Close()

		var resuls []ResultCategor
		for rowss.Next() {
			var resul ResultCategor
			if err := rowss.Scan(&resul.ID, &resul.Name); err != nil {
				helpers.HandleError(c, 400, err.Error())
				return
			}

			rowsss, err := db.Query(context.Background(), "SELECT c.id,tc.name FROM categories c LEFT JOIN translation_category tc ON c.id=tc.category_id WHERE tc.lang_id = $1 AND c.parent_category_id =$2 AND c.deleted_at IS NOT NULL ORDER BY c.created_at DESC", langID, resul.ID)
			if err != nil {
				helpers.HandleError(c, 400, err.Error())
				return
			}
			defer rowsss.Close()

			var resus []ResultCatego
			for rowsss.Next() {
				var resu ResultCatego
				if err := rowsss.Scan(&resu.ID, &resu.Name); err != nil {
					helpers.HandleError(c, 400, err.Error())
					return
				}

				resus = append(resus, resu)
			}
			resul.ResultCatego = resus
			resuls = append(resuls, resul)
		}
		category.ResultCategor = resuls
		categories = append(categories, category)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":     true,
		"categories": categories,
	})
}
