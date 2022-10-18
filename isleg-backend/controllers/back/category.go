package controllers

import (
	"github/abbgo/isleg/isleg-backend/config"
	"github/abbgo/isleg/isleg-backend/models"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ResultCategory struct {
	ID            string          `json:"id"`
	Image         string          `json:"image"`
	Name          string          `json:"name"`
	ResultCategor []ResultCategor `json:"child_category"`
}

type ResultCategor struct {
	ID           string         `json:"id"`
	Name         string         `json:"name"`
	ResultCatego []ResultCatego `json:"child_category"`
}

type ResultCatego struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Category struct {
	ID       string    `json:"id"`
	Name     string    `json:"name"`
	Products []Product `json:"products"`
}

type Product struct {
	ID          string           `json:"id"`
	Name        string           `json:"name"`
	Price       float64          `json:"price"`
	OldPrice    float64          `json:"old_price"`
	ProductCode string           `json:"product_code"`
	MainImage   models.MainImage `json:"main_image"`
	Images      []models.Images  `json:"images"`
	Brend       Brend            `json:"brend"`
	LimitAmount uint             `json:"limit_amount"`
	IsNew       bool             `json:"is_new"`
}

type Brend struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type OneCategory struct {
	ID               uuid.UUID             `json:"id"`
	ParentCategoryID uuid.UUID             `json:"parent_category_id"`
	Image            string                `json:"image"`
	IsHomeCategory   bool                  `json:"is_home_category"`
	Translations     []TranslationCategory `json:"translations"`
}

type TranslationCategory struct {
	LangID string `json:"lang_id"`
	Name   string `json:"name"`
}

func CreateCategory(c *gin.Context) {

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

	var categoryID string

	var fileName string

	// GET DATA FROM REQUEST
	isHomeCategoryStr := c.PostForm("is_home_category")
	isHomeCategory, err := strconv.ParseBool(isHomeCategoryStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	parentCategoryID := c.PostForm("parent_category_id")

	if parentCategoryID != "" {
		rowCategory, err := db.Query("SELECT id FROM categories WHERE id = $1 AND deleted_at IS NULL", parentCategoryID)
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

		var parentCategory string

		for rowCategory.Next() {
			if err := rowCategory.Scan(&parentCategory); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  false,
					"message": err.Error(),
				})
				return
			}
		}

		if parentCategory == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": "parent category not found",
			})
			return
		}
	}

	// GET ALL LANGUAGE
	languages, err := GetAllLanguageWithIDAndNameShort()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	// FILE UPLOAD
	file, errFile := c.FormFile("image")
	if errFile != nil {
		fileName = ""
	} else {
		extension := filepath.Ext(file.Filename)
		// VALIDATE IMAGE
		if extension != ".jpg" && extension != ".jpeg" && extension != ".png" && extension != ".gif" {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": "the file must be an image.",
			})
			return
		}

		newFileName := uuid.New().String() + extension
		fileName = "uploads/category/" + newFileName
	}

	// VALIDATE DATA
	for _, v := range languages {
		if c.PostForm("name_"+v.NameShort) == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": "name_" + v.NameShort + " is required",
			})
			return
		}
	}

	if parentCategoryID == "" && fileName == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "parent category image is required",
		})
		return
	}

	if parentCategoryID != "" && fileName != "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "child cannot be an image of the category",
		})
		return
	}

	// CREATE CATEGORY
	if parentCategoryID != "" {
		resultCateor, err := db.Query("INSERT INTO categories (parent_category_id,image,is_home_category) VALUES ($1,$2,$3) RETURNING id", parentCategoryID, fileName, isHomeCategory)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
		defer func() {
			if err := resultCateor.Close(); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  false,
					"message": err.Error(),
				})
				return
			}
		}()

		var category_id string

		for resultCateor.Next() {
			if err := resultCateor.Scan(&category_id); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  false,
					"message": err.Error(),
				})
				return
			}
		}

		categoryID = category_id

	} else {
		result, err := db.Query("INSERT INTO categories (image,is_home_category) VALUES ($1,$2) RETURNING id", fileName, isHomeCategory)
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

		var category_id string

		for result.Next() {
			if err := result.Scan(&category_id); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  false,
					"message": err.Error(),
				})
				return
			}
		}

		categoryID = category_id
	}

	if fileName != "" {
		c.SaveUploadedFile(file, "./"+fileName)
	}

	// CREATE TRANSLATION CATEGORY
	for _, v := range languages {
		result, err := db.Query("INSERT INTO translation_category (lang_id,category_id,name) VALUES ($1,$2,$3)", v.ID, categoryID, c.PostForm("name_"+v.NameShort))
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
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "category successfully added",
	})

}

func UpdateCategoryByID(c *gin.Context) {

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

	ID := c.Param("id")
	var fileName string

	isHomeCategoryStr := c.PostForm("is_home_category")
	isHomeCategory, err := strconv.ParseBool(isHomeCategoryStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	rowCategor, err := db.Query("SELECT id,image FROM categories WHERE id = $1 AND deleted_at IS NULL", ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer func() {
		if err := rowCategor.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

	var category_id, image string

	for rowCategor.Next() {
		if err := rowCategor.Scan(&category_id, &image); err != nil {
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

	parentCategoryID := c.PostForm("parent_category_id")

	if parentCategoryID != "" {
		rowCategory, err := db.Query("SELECT id FROM categories WHERE id = $1 AND deleted_at IS NULL", parentCategoryID)
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

		var parentCategory string

		for rowCategory.Next() {
			if err := rowCategory.Scan(&parentCategory); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  false,
					"message": err.Error(),
				})
				return
			}
		}

		if parentCategory == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": "parent category not found",
			})
			return
		}
	}

	// GET ALL LANGUAGE
	languages, err := GetAllLanguageWithIDAndNameShort()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	// FILE UPLOAD
	file, errFile := c.FormFile("image")
	if errFile != nil {
		fileName = image
	} else {
		extension := filepath.Ext(file.Filename)
		// VALIDATE IMAGE
		if extension != ".jpg" && extension != ".jpeg" && extension != ".png" && extension != ".gif" {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": "the file must be an image.",
			})
			return
		}

		newFileName := uuid.New().String() + extension
		fileName = "uploads/category/" + newFileName

		if err := os.Remove("./" + image); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}

	// VALIDATE DATA
	for _, v := range languages {
		if c.PostForm("name_"+v.NameShort) == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": "name_" + v.NameShort + " is required",
			})
			return
		}
	}

	if parentCategoryID == "" && fileName == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "parent category image is required",
		})
		return
	}

	if parentCategoryID != "" && fileName != "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "child cannot be an image of the category",
		})
		return
	}

	currentTime := time.Now()

	// UPDATE CATEGORY
	if parentCategoryID != "" {
		result, err := db.Query("UPDATE categories SET parent_category_id = $1, image = $2, is_home_category = $3 , updated_at = $5 WHERE id = $4", parentCategoryID, fileName, isHomeCategory, ID, currentTime)
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
	} else {
		resultCat, err := db.Query("UPDATE categories SET image = $1, is_home_category = $2 , updated_at = $4 WHERE id = $3", fileName, isHomeCategory, ID, currentTime)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
		defer func() {
			if err := resultCat.Close(); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  false,
					"message": err.Error(),
				})
				return
			}
		}()
	}

	if fileName != "" {
		c.SaveUploadedFile(file, "./"+fileName)
	}

	// UPDATE TRANSLATION CATEGORY
	for _, v := range languages {
		resultTRCate, err := db.Query("UPDATE translation_category SET name = $1 , updated_at = $4 WHERE lang_id = $2 AND category_id = $3", c.PostForm("name_"+v.NameShort), v.ID, ID, currentTime)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
		defer func() {
			if err := resultTRCate.Close(); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  false,
					"message": err.Error(),
				})
				return
			}
		}()
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "category successfully updated",
	})

}

func GetCategoryByID(c *gin.Context) {

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

	rowCategor, err := db.Query("SELECT id,parent_category_id,image,is_home_category FROM categories WHERE id = $1 AND deleted_at IS NULL", ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer rowCategor.Close()

	var category OneCategory

	for rowCategor.Next() {
		if err := rowCategor.Scan(&category.ID, &category.ParentCategoryID, &category.Image, &category.IsHomeCategory); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}

	if category.ID == uuid.Nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "category not found",
		})
		return
	}

	rowsTrCategory, err := db.Query("SELECT lang_id,name FROM translation_category WHERE category_id = $1 AND deleted_at IS NULL", ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer rowsTrCategory.Close()

	var translations []TranslationCategory

	for rowsTrCategory.Next() {
		var translation TranslationCategory
		if err := rowsTrCategory.Scan(&translation.LangID, &translation.Name); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
		translations = append(translations, translation)
	}

	category.Translations = translations

	c.JSON(http.StatusOK, gin.H{
		"status":   true,
		"category": category,
	})

}

func GetCategories(c *gin.Context) {

	db, err := config.ConnDB()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer db.Close()

	rowCategor, err := db.Query("SELECT id,parent_category_id,image,is_home_category FROM categories WHERE deleted_at IS NULL")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer rowCategor.Close()

	var categories []OneCategory

	for rowCategor.Next() {
		var category OneCategory

		if err := rowCategor.Scan(&category.ID, &category.ParentCategoryID, &category.Image, &category.IsHomeCategory); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}

		rowsTrCategory, err := db.Query("SELECT lang_id,name FROM translation_category WHERE deleted_at IS NULL AND category_id = $1", category.ID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
		defer rowsTrCategory.Close()

		var translations []TranslationCategory

		for rowsTrCategory.Next() {
			var translation TranslationCategory
			if err := rowsTrCategory.Scan(&translation.LangID, &translation.Name); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  false,
					"message": err.Error(),
				})
				return
			}
			translations = append(translations, translation)
		}

		category.Translations = translations

		categories = append(categories, category)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":     true,
		"categories": categories,
	})

}

func GetAllCategoryForHeader(langID string) ([]ResultCategory, error) {

	db, err := config.ConnDB()
	if err != nil {
		return []ResultCategory{}, nil
	}
	defer db.Close()

	// get all category where parent category id is null
	rows, err := db.Query("SELECT categories.id,categories.image,translation_category.name FROM categories LEFT JOIN translation_category ON categories.id=translation_category.category_id WHERE translation_category.lang_id = $1 AND categories.parent_category_id IS NULL AND translation_category.deleted_at IS NULL AND categories.deleted_at IS NULL", langID)
	if err != nil {
		return []ResultCategory{}, err
	}
	defer rows.Close()

	var results []ResultCategory

	for rows.Next() {
		var result ResultCategory
		if err := rows.Scan(&result.ID, &result.Image, &result.Name); err != nil {
			return []ResultCategory{}, err
		}

		// get all category where parent category id equal category id
		rowss, err := db.Query("SELECT categories.id,translation_category.name FROM categories LEFT JOIN translation_category ON categories.id=translation_category.category_id WHERE translation_category.lang_id = $1 AND categories.parent_category_id = $2 AND translation_category.deleted_at IS NULL AND categories.deleted_at IS NULL", langID, result.ID)
		if err != nil {
			return []ResultCategory{}, err
		}
		defer rowss.Close()

		var resuls []ResultCategor

		for rowss.Next() {
			var resul ResultCategor
			if err := rowss.Scan(&resul.ID, &resul.Name); err != nil {
				return []ResultCategory{}, err
			}

			// get all category where parent category id equal category id
			rowsss, err := db.Query("SELECT categories.id,translation_category.name FROM categories LEFT JOIN translation_category ON categories.id=translation_category.category_id WHERE translation_category.lang_id = $1 AND categories.parent_category_id =$2 AND translation_category.deleted_at IS NULL AND categories.deleted_at IS NULL", langID, resul.ID)
			if err != nil {
				return []ResultCategory{}, err
			}
			defer rowsss.Close()

			var resus []ResultCatego

			for rowsss.Next() {
				var resu ResultCatego
				if err := rowsss.Scan(&resu.ID, &resu.Name); err != nil {
					return []ResultCategory{}, err
				}

				resus = append(resus, resu)
			}
			resul.ResultCatego = resus

			resuls = append(resuls, resul)
		}
		result.ResultCategor = resuls

		results = append(results, result)
	}
	return results, nil

}

func DeleteCategoryByID(c *gin.Context) {

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

	rowCategor, err := db.Query("SELECT id FROM categories WHERE id = $1 AND deleted_at IS NULL", ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer rowCategor.Close()

	var category_id string

	for rowCategor.Next() {
		if err := rowCategor.Scan(&category_id); err != nil {
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

	currentTime := time.Now()

	resultCate, err := db.Query("UPDATE categories SET deleted_at = $1 WHERE id = $2", currentTime, ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer resultCate.Close()

	resultTRCate, err := db.Query("UPDATE translation_category SET deleted_at = $1 WHERE category_id = $2", currentTime, ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer resultTRCate.Close()

	resultCateg, err := db.Query("UPDATE categories SET deleted_at = $1 WHERE parent_category_id = $2", currentTime, ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer resultCateg.Close()

	rowChildCategory, err := db.Query("SELECT id FROM categories WHERE parent_category_id = $1", ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer rowChildCategory.Close()

	var childCategoryIDS []string

	for rowChildCategory.Next() {
		var childCategoryID string
		if err := rowChildCategory.Scan(&childCategoryID); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}

		childCategoryIDS = append(childCategoryIDS, childCategoryID)
	}

	for _, v := range childCategoryIDS {
		resultTRCate, err := db.Query("UPDATE translation_category SET deleted_at = $1 WHERE category_id = $2", currentTime, v)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
		defer resultTRCate.Close()

		resultCaetProd, err := db.Query("UPDATE category_product SET deleted_at = $1 WHERE category_id = $2", currentTime, v)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
		defer resultCaetProd.Close()

		resultProd, err := db.Query("UPDATE products SET deleted_at = $1 FROM category_product WHERE category_product.product_id = products.id AND category_product.category_id = $2", currentTime, v)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
		defer resultProd.Close()

		resultTRPr, err := db.Query("UPDATE translation_product SET deleted_at = $1 FROM products,category_product WHERE translation_product.product_id = products.id AND category_product.product_id = products.id  AND category_product.category_id = $2", currentTime, v)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
		defer resultTRPr.Close()

		resultCateSho, err := db.Query("UPDATE category_shop SET deleted_at = $1 WHERE category_id = $2", currentTime, v)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
		defer resultCateSho.Close()

		resultSHop, err := db.Query("UPDATE shops SET deleted_at = $1 FROM category_shop WHERE category_shop.shop_id = shops.id AND category_shop.category_id = $2", currentTime, v)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
		defer resultSHop.Close()
	}

	resultCategProd, err := db.Query("UPDATE category_product SET deleted_at = $1 WHERE category_id = $2", currentTime, ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer resultCategProd.Close()

	resultProduct, err := db.Query("UPDATE products SET deleted_at = $1 FROM category_product WHERE category_product.product_id = products.id AND category_product.category_id = $2", currentTime, ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer resultProduct.Close()

	resultTRPRD, err := db.Query("UPDATE translation_product SET deleted_at = $1 FROM products,category_product WHERE translation_product.product_id = products.id AND category_product.product_id = products.id  AND category_product.category_id = $2", currentTime, ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer resultTRPRD.Close()

	resultCateShop, err := db.Query("UPDATE category_shop SET deleted_at = $1 WHERE category_id = $2", currentTime, ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer resultCateShop.Close()

	resultsHI, err := db.Query("UPDATE shops SET deleted_at = $1 FROM category_shop WHERE category_shop.shop_id = shops.id AND category_shop.category_id = $2", currentTime, ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer resultsHI.Close()

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "category successfully deleted",
	})

}

func RestoreCategoryByID(c *gin.Context) {

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

	rowCategor, err := db.Query("SELECT id FROM categories WHERE id = $1 AND deleted_at IS NOT NULL", ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer rowCategor.Close()

	var category_id string

	for rowCategor.Next() {
		if err := rowCategor.Scan(&category_id); err != nil {
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

	rESUTCate, err := db.Query("UPDATE categories SET deleted_at = NULL WHERE id = $1", ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer rESUTCate.Close()

	resultTrCateg, err := db.Query("UPDATE translation_category SET deleted_at = NULL WHERE category_id = $1", ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer resultTrCateg.Close()

	resultCt, err := db.Query("UPDATE categories SET deleted_at = NULL WHERE parent_category_id = $1", ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer resultCt.Close()

	rowChildCategory, err := db.Query("SELECT id FROM categories WHERE parent_category_id = $1", ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer rowChildCategory.Close()

	var childCategoryIDS []string

	for rowChildCategory.Next() {
		var childCategoryID string
		if err := rowChildCategory.Scan(&childCategoryID); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}

		childCategoryIDS = append(childCategoryIDS, childCategoryID)
	}

	for _, v := range childCategoryIDS {
		resultTRCategory, err := db.Query("UPDATE translation_category SET deleted_at = NULL WHERE category_id = $1", v)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
		defer resultTRCategory.Close()

		resultCateProd, err := db.Query("UPDATE category_product SET deleted_at = NULL WHERE category_id = $1", v)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
		defer resultCateProd.Close()

		resultProd, err := db.Query("UPDATE products SET deleted_at = NULL FROM category_product WHERE category_product.product_id = products.id AND category_product.category_id = $1", v)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
		defer resultProd.Close()

		resultTRProduct, err := db.Query("UPDATE translation_product SET deleted_at = NULL FROM products,category_product WHERE translation_product.product_id = products.id AND category_product.product_id = products.id  AND category_product.category_id = $1", v)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
		defer resultTRProduct.Close()

		resultCateShop, err := db.Query("UPDATE category_shop SET deleted_at = NULL WHERE category_id = $1", v)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
		defer resultCateShop.Close()

		resultSHops, err := db.Query("UPDATE shops SET deleted_at = NULL FROM category_shop WHERE category_shop.shop_id = shops.id AND category_shop.category_id = $1", v)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
		defer resultSHops.Close()
	}

	resutlCategPro, err := db.Query("UPDATE category_product SET deleted_at = NULL WHERE category_id = $1", ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer resutlCategPro.Close()

	resultProd, err := db.Query("UPDATE products SET deleted_at = NULL FROM category_product WHERE category_product.product_id = products.id AND category_product.category_id = $1", ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer resultProd.Close()

	resultTRProd, err := db.Query("UPDATE translation_product SET deleted_at = NULL FROM products,category_product WHERE translation_product.product_id = products.id AND category_product.product_id = products.id  AND category_product.category_id = $1", ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer resultTRProd.Close()

	resultCategShop, err := db.Query("UPDATE category_shop SET deleted_at = NULL WHERE category_id = $1", ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer resultCategShop.Close()

	resutShops, err := db.Query("UPDATE shops SET deleted_at = NULL FROM category_shop WHERE category_shop.shop_id = shops.id AND category_shop.category_id = $1", ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer resutShops.Close()

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "category successfully restored",
	})

}

func DeletePermanentlyCategoryByID(c *gin.Context) {

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

	rowCategor, err := db.Query("SELECT id,image FROM categories WHERE id = $1 AND deleted_at IS NOT NULL", ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer rowCategor.Close()

	var category_id, image string

	for rowCategor.Next() {
		if err := rowCategor.Scan(&category_id, &image); err != nil {
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

	if image != "" {
		if err := os.Remove("./" + image); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}

	rowsMainImageProduct, err := db.Query("SELECT m.small,m.medium,m.large FROM products p INNER JOIN category_product c ON c.product_id=p.id INNER JOIN main_image m ON m.product_id = p.id WHERE c.category_id = $1", ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer rowsMainImageProduct.Close()

	var mainImages []models.MainImage

	for rowsMainImageProduct.Next() {
		var mainImage models.MainImage

		if err := rowsMainImageProduct.Scan(&mainImage.Small, &mainImage.Medium, &mainImage.Large); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}

		mainImages = append(mainImages, mainImage)
	}

	for _, v := range mainImages {
		if err := os.Remove("./" + v.Small); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}

		if err := os.Remove("./" + v.Medium); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}

		if err := os.Remove("./" + v.Large); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}

	}

	rowsImagesProduct, err := db.Query("SELECT i.small,i.large FROM products p INNER JOIN category_product c ON c.product_id=p.id INNER JOIN images i ON i.product_id = p.id WHERE c.category_id = $1", ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer rowsImagesProduct.Close()

	var images []models.Images

	for rowsImagesProduct.Next() {
		var image models.Images

		if err := rowsImagesProduct.Scan(&image.Small, &image.Large); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}

		images = append(images, image)
	}

	for _, v := range images {
		if err := os.Remove("./" + v.Small); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}

		if err := os.Remove("./" + v.Large); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}

	}

	resultProduct, err := db.Query("DELETE FROM products USING category_product WHERE category_product.product_id = products.id AND category_product.category_id = $1", category_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer resultProduct.Close()

	rowChildCategory, err := db.Query("SELECT id FROM categories WHERE parent_category_id = $1", ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer rowChildCategory.Close()

	var childCategoryIDS []string

	for rowChildCategory.Next() {
		var childCategoryID string
		if err := rowChildCategory.Scan(&childCategoryID); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}

		childCategoryIDS = append(childCategoryIDS, childCategoryID)
	}

	for _, v := range childCategoryIDS {
		rowPrdcs, err := db.Query("SELECT m.small,m.medium,m.large FROM products p INNER JOIN category_product c ON c.product_id=p.id INNER JOIN main_image m ON m.product_id = p.id WHERE c.category_id = $1", v)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
		defer rowPrdcs.Close()

		var childMainImages []models.MainImage

		for rowsMainImageProduct.Next() {
			var childMainImage models.MainImage

			if err := rowsMainImageProduct.Scan(&childMainImage.Small, &childMainImage.Medium, &childMainImage.Large); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  false,
					"message": err.Error(),
				})
				return
			}

			childMainImages = append(childMainImages, childMainImage)
		}

		for _, v := range childMainImages {
			if err := os.Remove("./" + v.Small); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  false,
					"message": err.Error(),
				})
				return
			}

			if err := os.Remove("./" + v.Medium); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  false,
					"message": err.Error(),
				})
				return
			}

			if err := os.Remove("./" + v.Large); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  false,
					"message": err.Error(),
				})
				return
			}

		}

		rowsChildImagesProduct, err := db.Query("SELECT i.small,i.large FROM products p INNER JOIN category_product c ON c.product_id=p.id INNER JOIN images i ON i.product_id = p.id WHERE c.category_id = $1", ID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
		defer rowsChildImagesProduct.Close()

		var childImages []models.Images

		for rowsChildImagesProduct.Next() {
			var childImage models.Images

			if err := rowsChildImagesProduct.Scan(&childImage.Small, &childImage.Large); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  false,
					"message": err.Error(),
				})
				return
			}

			childImages = append(childImages, childImage)
		}

		for _, v := range childImages {
			if err := os.Remove("./" + v.Small); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  false,
					"message": err.Error(),
				})
				return
			}

			if err := os.Remove("./" + v.Large); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  false,
					"message": err.Error(),
				})
				return
			}

		}

		childresultProduct, err := db.Query("DELETE FROM products USING category_product WHERE category_product.product_id = products.id AND category_product.category_id = $1", category_id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
		defer childresultProduct.Close()

	}

	resutCateg, err := db.Query("DELETE FROM categories WHERE id = $1", ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer resutCateg.Close()

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "category successfully deleted",
	})

}

func GetOneCategoryWithProducts(c *gin.Context) {

	db, err := config.ConnDB()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer db.Close()

	var countOfProducts uint64

	langID, err := CheckLanguage(c)
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

	categoryID := c.Param("category_id")

	// get category where id equal categiryID
	categoryRow, err := db.Query("SELECT c.id,t.name FROM categories c LEFT JOIN translation_category t ON c.id=t.category_id WHERE t.lang_id = $1 AND c.id = $2 AND c.deleted_at IS NULL AND t.deleted_at IS NULL", langID, categoryID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer categoryRow.Close()

	var category Category

	for categoryRow.Next() {
		if err := categoryRow.Scan(&category.ID, &category.Name); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}

		if category.ID == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": "category not found",
			})
			return
		}

		// get count product where product_id equal categoryID
		productCount, err := db.Query("SELECT COUNT(p.id) FROM products p LEFT JOIN category_product c ON p.id=c.product_id WHERE c.category_id = $1 AND p.deleted_at IS NULL AND c.deleted_at IS NULL", categoryID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
		defer productCount.Close()

		for productCount.Next() {
			if err := productCount.Scan(&countOfProducts); err != nil {
				if err != nil {
					c.JSON(http.StatusBadRequest, gin.H{
						"status":  false,
						"message": err.Error(),
					})
					return
				}
			}
		}

		// get all product where category id equal categoryID
		productRows, err := db.Query("SELECT p.id,t.name,p.price,p.old_price,p.product_code,p.limit_amount,p.is_new FROM products p LEFT JOIN category_product c ON p.id=c.product_id LEFT JOIN translation_product t ON p.id=t.product_id WHERE t.lang_id = $1 AND c.category_id = $2 AND p.deleted_at IS NULL AND c.deleted_at IS NULL AND t.deleted_at IS NULL ORDER BY p.created_at ASC LIMIT $3 OFFSET $4", langID, categoryID, limit, offset)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
		defer productRows.Close()

		var products []Product

		for productRows.Next() {
			var product Product

			if err := productRows.Scan(&product.ID, &product.Name, &product.Price, &product.OldPrice, &product.ProductCode, &product.LimitAmount, &product.IsNew); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  false,
					"message": err.Error(),
				})
				return
			}

			rowMainImage, err := db.Query("SELECT small,medium,large FROM main_image WHERE product_id = $1", product.ID)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  false,
					"message": err.Error(),
				})
				return
			}
			defer rowMainImage.Close()

			var mainImage models.MainImage

			for rowMainImage.Next() {
				if err := rowMainImage.Scan(&mainImage.Small, &mainImage.Medium, &mainImage.Large); err != nil {
					c.JSON(http.StatusBadRequest, gin.H{
						"status":  false,
						"message": err.Error(),
					})
					return
				}
			}

			product.MainImage = mainImage

			rowsImages, err := db.Query("SELECT small,large FROM images WHERE product_id = $1", product.ID)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  false,
					"message": err.Error(),
				})
				return
			}
			defer rowsImages.Close()

			var images []models.Images

			for rowsImages.Next() {
				var image models.Images

				if err := rowsImages.Scan(&image.Small, &image.Large); err != nil {
					c.JSON(http.StatusBadRequest, gin.H{
						"status":  false,
						"message": err.Error(),
					})
					return
				}

				images = append(images, image)
			}

			product.Images = images

			// get brend where id equal brend_id of product
			brendRows, err := db.Query("SELECT b.id,b.name FROM products p LEFT JOIN brends b ON p.brend_id=b.id WHERE p.id = $1 AND p.deleted_at IS NULL AND b.deleted_at IS NULL", product.ID)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  false,
					"message": err.Error(),
				})
				return
			}
			defer brendRows.Close()

			var brend Brend

			for brendRows.Next() {
				if err := brendRows.Scan(&brend.ID, &brend.Name); err != nil {
					c.JSON(http.StatusBadRequest, gin.H{
						"status":  false,
						"message": err.Error(),
					})
					return
				}
			}
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
