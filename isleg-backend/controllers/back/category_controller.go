package controllers

import (
	"github/abbgo/isleg/isleg-backend/config"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/lib/pq"
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
	ID          string         `json:"id"`
	Name        string         `json:"name"`
	Price       float64        `json:"price"`
	OldPrice    float64        `json:"old_price"`
	MainImage   string         `json:"main_image"`
	ProductCode string         `json:"product_code"`
	Images      pq.StringArray `json:"images"`
	Brend       Brend          `json:"brend"`
}

type Brend struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type OneCategory struct {
	ID               string    `json:"id"`
	ParentCategoryID uuid.UUID `json:"parent_category_id"`
	Image            string    `json:"image"`
	IsHomeCategory   bool      `json:"is_home_category"`
}

func CreateCategory(c *gin.Context) {

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
		rowCategory, err := config.ConnDB().Query("SELECT id FROM categories WHERE id = $1 AND deleted_at IS NULL", parentCategoryID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}

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

		newFileName := "category" + uuid.New().String() + extension
		fileName = "uploads/" + newFileName
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
		_, err = config.ConnDB().Exec("INSERT INTO categories (parent_category_id,image,is_home_category) VALUES ($1,$2,$3)", parentCategoryID, fileName, isHomeCategory)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	} else {
		_, err = config.ConnDB().Exec("INSERT INTO categories (image,is_home_category) VALUES ($1,$2)", fileName, isHomeCategory)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}

	if fileName != "" {
		c.SaveUploadedFile(file, "./"+fileName)
	}

	// GET ID OFF ADDED CATEGORY
	lastCategoryID, err := config.ConnDB().Query("SELECT id FROM categories WHERE deleted_at IS NULL ORDER BY created_at DESC LIMIT 1")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	var categoryID string

	for lastCategoryID.Next() {
		if err := lastCategoryID.Scan(&categoryID); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}

	// CREATE TRANSLATION CATEGORY
	for _, v := range languages {
		_, err := config.ConnDB().Exec("INSERT INTO translation_category (lang_id,category_id,name) VALUES ($1,$2,$3)", v.ID, categoryID, c.PostForm("name_"+v.NameShort))
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
		"message": "category successfully added",
	})

}

func UpdateCategory(c *gin.Context) {

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

	rowCategor, err := config.ConnDB().Query("SELECT id,image FROM categories WHERE id = $1 AND deleted_at IS NULL", ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

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
		rowCategory, err := config.ConnDB().Query("SELECT id FROM categories WHERE id = $1 AND deleted_at IS NULL", parentCategoryID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}

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

		newFileName := "category" + uuid.New().String() + extension
		fileName = "uploads/" + newFileName

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

	// UPDATE CATEGORY
	if parentCategoryID != "" {
		_, err = config.ConnDB().Exec("UPDATE categories SET parent_category_id = $1, image = $2, is_home_category = $3 WHERE id = $4", parentCategoryID, fileName, isHomeCategory, ID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	} else {
		_, err = config.ConnDB().Exec("UPDATE categories SET image = $1, is_home_category = $2 WHERE id = $3", fileName, isHomeCategory, ID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}

	if fileName != "" {
		c.SaveUploadedFile(file, "./"+fileName)
	}

	// UPDATE TRANSLATION CATEGORY
	for _, v := range languages {
		_, err := config.ConnDB().Exec("UPDATE translation_category SET name = $1 WHERE lang_id = $2 AND category_id = $3", c.PostForm("name_"+v.NameShort), v.ID, ID)
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
		"message": "category successfully updated",
	})

}

func GetOneCategory(c *gin.Context) {

	ID := c.Param("id")

	rowCategor, err := config.ConnDB().Query("SELECT id,parent_category_id,image,is_home_category FROM categories WHERE id = $1 AND deleted_at IS NULL", ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	var category OneCategory

	for rowCategor.Next() {
		if err := rowCategor.Scan(&category.ID, &category.ParentCategoryID, &category.Image, &category.Image); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}

	if category.ID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "category not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":   true,
		"category": category,
	})

}

func GetAllCategory(c *gin.Context) {

	rowCategor, err := config.ConnDB().Query("SELECT id,parent_category_id,image,is_home_category FROM categories WHERE deleted_at IS NULL")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	var categories []OneCategory

	for rowCategor.Next() {
		var category OneCategory

		if err := rowCategor.Scan(&category.ID, &category.ParentCategoryID, &category.Image, &category.Image); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}

		categories = append(categories, category)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":     true,
		"categories": categories,
	})

}

func GetAllCategoryForHeader(langID string) ([]ResultCategory, error) {

	// get all category where parent category id is null
	rows, err := config.ConnDB().Query("SELECT categories.id,categories.image,translation_category.name FROM categories LEFT JOIN translation_category ON categories.id=translation_category.category_id WHERE translation_category.lang_id = $1 AND categories.parent_category_id IS NULL AND translation_category.deleted_at IS NULL AND categories.deleted_at IS NULL", langID)
	if err != nil {
		return []ResultCategory{}, err
	}

	var results []ResultCategory

	for rows.Next() {
		var result ResultCategory
		if err := rows.Scan(&result.ID, &result.Image, &result.Name); err != nil {
			return []ResultCategory{}, err
		}

		// get all category where parent category id equal category id
		rowss, err := config.ConnDB().Query("SELECT categories.id,translation_category.name FROM categories LEFT JOIN translation_category ON categories.id=translation_category.category_id WHERE translation_category.lang_id = $1 AND categories.parent_category_id = $2 AND translation_category.deleted_at IS NULL AND categories.deleted_at IS NULL", langID, result.ID)
		if err != nil {
			return []ResultCategory{}, err
		}

		var resuls []ResultCategor

		for rowss.Next() {
			var resul ResultCategor
			if err := rowss.Scan(&resul.ID, &resul.Name); err != nil {
				return []ResultCategory{}, err
			}

			// get all category where parent category id equal category id
			rowsss, err := config.ConnDB().Query("SELECT categories.id,translation_category.name FROM categories LEFT JOIN translation_category ON categories.id=translation_category.category_id WHERE translation_category.lang_id = $1 AND categories.parent_category_id =$2 AND translation_category.deleted_at IS NULL AND categories.deleted_at IS NULL", langID, resul.ID)
			if err != nil {
				return []ResultCategory{}, err
			}

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

func DeleteCategory(c *gin.Context) {

	ID := c.Param("id")

	rowCategor, err := config.ConnDB().Query("SELECT id FROM categories WHERE id = $1 AND deleted_at IS NULL", ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

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

	_, err = config.ConnDB().Exec("UPDATE categories SET deleted_at = $1 WHERE id = $2", currentTime, ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	_, err = config.ConnDB().Exec("UPDATE categories SET deleted_at = $1 WHERE parent_category_id = $2", currentTime, ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	_, err = config.ConnDB().Exec("UPDATE translation_category SET deleted_at = $1 WHERE category_id = $2", currentTime, ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	_, err = config.ConnDB().Exec("UPDATE category_product SET deleted_at = $1 WHERE category_id = $2", currentTime, ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	_, err = config.ConnDB().Exec("UPDATE products SET deleted_at = $1 FROM category_product WHERE category_product.product_id = products.id AND category_product.category_id = $2", currentTime, ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	_, err = config.ConnDB().Exec("UPDATE category_shop SET deleted_at = $1 WHERE category_id = $2", currentTime, ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	_, err = config.ConnDB().Query("UPDATE shops SET deleted_at = $1 FROM category_shop WHERE category_shop.shop_id = shops.id AND category_shop.category_id = $2", currentTime, ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "category successfully deleted",
	})

}

func GetOneCategoryWithProducts(c *gin.Context) {

	var countOfProducts uint64

	// GET DATA FROM ROUTE PARAMETER
	langShortName := c.Param("lang")

	// GET language id
	var langID string
	langID, err := GetLangID(langShortName)
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
	categoryRow, err := config.ConnDB().Query("SELECT c.id,t.name FROM categories c LEFT JOIN translation_category t ON c.id=t.category_id WHERE t.lang_id = $1 AND c.id = $2 AND categories.deleted_at IS NULL AND translation_category.deleted_at IS NULL", langID, categoryID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	var category Category

	for categoryRow.Next() {
		if err := categoryRow.Scan(&category.ID, &category.Name); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}

		// get count product where product_id equal categoryID
		productCount, err := config.ConnDB().Query("SELECT COUNT(p.id) FROM products p LEFT JOIN category_product c ON p.id=c.product_id WHERE c.category_id = $1 AND products.deleted_at IS NULL AND category_product.deleted_at IS NULL", categoryID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}

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
		productRows, err := config.ConnDB().Query("SELECT p.id,t.name,p.price,p.old_price,p.main_image,p.product_code,p.images FROM products p LEFT JOIN category_product c ON p.id=c.product_id LEFT JOIN translation_product t ON p.id=t.product_id WHERE t.lang_id = $1 AND c.category_id = $2 AND products.deleted_at IS NULL AND category_product.deleted_at IS NULL AND translation_product.deleted_at IS NULL ORDER BY p.created_at ASC LIMIT $3 OFFSET $4", langID, categoryID, limit, offset)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}

		var products []Product

		for productRows.Next() {
			var product Product
			if err := productRows.Scan(&product.ID, &product.Name, &product.Price, &product.OldPrice, &product.MainImage, &product.ProductCode, &product.Images); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  false,
					"message": err.Error(),
				})
				return
			}

			// get brend where id equal brend_id of product
			brendRows, err := config.ConnDB().Query("SELECT b.id,b.name FROM products p LEFT JOIN brends b ON p.brend_id=b.id WHERE p.id = $1 AND products.deleted_at IS NULL AND brends.deleted_at IS NULL", product.ID)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  false,
					"message": err.Error(),
				})
				return
			}

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
