package controllers

import (
	"github/abbgo/isleg/isleg-backend/config"
	"net/http"
	"path/filepath"
	"strconv"

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
	ID            string         `json:"id"`
	Name          string         `json:"name"`
	Price         float64        `json:"price"`
	OldPrice      float64        `json:"old_price"`
	MainImagePath string         `json:"main_image_path"`
	ProductCode   string         `json:"product_code"`
	ImagePaths    pq.StringArray `json:"image_paths"`
	Brend         Brend          `json:"brend"`
}

type Brend struct {
	ID   string `json:"id"`
	Name string `json:"name"`
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
	parentCategoryIDUUID, err := uuid.Parse(parentCategoryID)
	if parentCategoryID != "" {
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	} else {
		parentCategoryIDUUID = uuid.Nil
	}

	if parentCategoryIDUUID != uuid.Nil {
		_, err := config.ConnDB().Query("SELECT id FROM categories WHERE id = $1", parentCategoryID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
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
	file, errFile := c.FormFile("image_path")
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

	if parentCategoryIDUUID == uuid.Nil && fileName == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "parent category image is required",
		})
		return
	}

	if parentCategoryIDUUID != uuid.Nil && fileName != "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "child cannot be an image of the category",
		})
		return
	}

	// CREATE CATEGORY
	if parentCategoryIDUUID != uuid.Nil {
		_, err = config.ConnDB().Exec("INSERT INTO categories (parent_category_id,image_path,is_home_category) VALUES ($1,$2,$3)", parentCategoryIDUUID, fileName, isHomeCategory)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	} else {
		_, err = config.ConnDB().Exec("INSERT INTO categories (image_path,is_home_category) VALUES ($1,$2)", fileName, isHomeCategory)
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

	// GET LAST CATEGORY ID
	lastCategoryID, err := config.ConnDB().Query("SELECT id FROM categories ORDER BY created_at DESC LIMIT 1")
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

func GetAllCategoryForHeader(langID string) ([]ResultCategory, error) {

	// get all category where parent category id is null
	rows, err := config.ConnDB().Query("SELECT categories.id,categories.image_path,translation_category.name FROM categories LEFT JOIN translation_category ON categories.id=translation_category.category_id WHERE translation_category.lang_id = $1 AND categories.parent_category_id IS NULL", langID)
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
		rowss, err := config.ConnDB().Query("SELECT categories.id,translation_category.name FROM categories LEFT JOIN translation_category ON categories.id=translation_category.category_id WHERE translation_category.lang_id = $1 AND categories.parent_category_id = $2", langID, result.ID)
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
			rowsss, err := config.ConnDB().Query("SELECT categories.id,translation_category.name FROM categories LEFT JOIN translation_category ON categories.id=translation_category.category_id WHERE translation_category.lang_id = $1 AND categories.parent_category_id =$2", langID, resul.ID)
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
	categoryRow, err := config.ConnDB().Query("SELECT c.id,t.name FROM categories c LEFT JOIN translation_category t ON c.id=t.category_id WHERE t.lang_id = $1 AND c.id = $2", langID, categoryID)
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
		productCount, err := config.ConnDB().Query("SELECT COUNT(p.id) FROM products p LEFT JOIN category_product c ON p.id=c.product_id WHERE c.category_id = $1", categoryID)
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
		productRows, err := config.ConnDB().Query("SELECT p.id,t.name,p.price,p.old_price,p.main_image_path,p.product_code,p.image_paths FROM products p LEFT JOIN category_product c ON p.id=c.product_id LEFT JOIN translation_product t ON p.id=t.product_id WHERE t.lang_id = $1 AND c.category_id = $2 ORDER BY p.created_at ASC LIMIT $3 OFFSET $4", langID, categoryID, limit, offset)
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
			if err := productRows.Scan(&product.ID, &product.Name, &product.Price, &product.OldPrice, &product.MainImagePath, &product.ProductCode, &product.ImagePaths); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  false,
					"message": err.Error(),
				})
				return
			}

			// get brend where id equal brend_id of product
			brendRows, err := config.ConnDB().Query("SELECT b.id,b.name FROM products p LEFT JOIN brends b ON p.brend_id=b.id WHERE p.id = $1", product.ID)
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
