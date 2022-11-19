package controllers

import (
	"github/abbgo/isleg/isleg-backend/config"
	"github/abbgo/isleg/isleg-backend/models"
	"github/abbgo/isleg/isleg-backend/pkg"
	"math"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

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
	Price       float64          `json:"price"`
	OldPrice    float64          `json:"old_price"`
	Percentage  float64          `json:"percentage"`
	MainImage   models.MainImage `json:"main_image"`
	Images      []models.Images  `json:"images"`
	Brend       Brend            `json:"brend"`
	LimitAmount int              `json:"limit_amount"`
	Amount      int              `json:"amount"`
	IsNew       bool             `json:"is_new"`
	// Translation models.TranslationProduct `json:"translation"`
	Translations []map[string]models.TranslationProduct `json:"translations"`
}

type Brend struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func CreateCategory(c *gin.Context) {

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

	var categoryID, fileName string
	var parent_category_id interface{}

	// GET DATA FROM REQUEST
	isHomeCategoryStr := c.PostForm("is_home_category")
	parentCategoryIDStr := c.PostForm("parent_category_id")

	// GET ALL LANGUAGE
	languages, err := GetAllLanguageWithIDAndNameShort()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	dataNames := []string{"name"}

	// VALIDATE translation category
	if err := pkg.ValidateTranslations(languages, dataNames, c); err != nil {
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

	// validate other data of category
	isHomeCategory, parentCategoryID, err := models.ValidateCategory(isHomeCategoryStr, parentCategoryIDStr, fileName)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	if fileName != "" {
		if err := c.SaveUploadedFile(file, "./"+fileName); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}

	// CREATE CATEGORY
	if parentCategoryID != "" {
		parent_category_id = parentCategoryID
	} else {
		parent_category_id = nil
	}

	// add data to categories table
	resultCateor, err := db.Query("INSERT INTO categories (parent_category_id,image,is_home_category) VALUES ($1,$2,$3) RETURNING id", parent_category_id, fileName, isHomeCategory)
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

	for resultCateor.Next() {
		if err := resultCateor.Scan(&categoryID); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
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
		"message": "data successfully added",
	})

}

func UpdateCategoryByID(c *gin.Context) {

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
	parentCategoryIDStr := c.PostForm("parent_category_id")
	isHomeCategoryStr := c.PostForm("is_home_category")

	var fileName string
	var parent_category_id interface{}

	// GET ALL LANGUAGE
	languages, err := GetAllLanguageWithIDAndNameShort()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	isHomeCategory, parentCategoryID, image, err := models.ValidateCategoryForUpdate(isHomeCategoryStr, ID, parentCategoryIDStr)
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

		if image != "" {
			if err := os.Remove("./" + image); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  false,
					"message": err.Error(),
				})
				return
			}
		}

		if err := c.SaveUploadedFile(file, "./"+fileName); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}

	}

	dataNames := []string{"name"}

	if err := pkg.ValidateTranslations(languages, dataNames, c); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	if parentCategoryID == "" && fileName == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "parent category image is required",
		})
		return
	}

	if parentCategoryID != "" && fileName != "" {
		if err := os.Remove("./" + fileName); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}

	// UPDATE CATEGORY
	if parentCategoryID != "" {
		parent_category_id = parentCategoryID
	} else {
		parent_category_id = nil
	}

	// update data
	result, err := db.Query("UPDATE categories SET parent_category_id = $1, image = $2, is_home_category = $3 WHERE id = $4", parent_category_id, fileName, isHomeCategory, ID)
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

	// UPDATE TRANSLATION CATEGORY
	for _, v := range languages {
		resultTRCate, err := db.Query("UPDATE translation_category SET name = $1 WHERE lang_id = $2 AND category_id = $3", c.PostForm("name_"+v.NameShort), v.ID, ID)
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
		"message": "data successfully updated",
	})

}

func GetCategoryByID(c *gin.Context) {

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

	// check id and get data from daabase
	rowCategor, err := db.Query("SELECT id,parent_category_id,image,is_home_category FROM categories WHERE id = $1 AND deleted_at IS NULL", ID)
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

	var category models.Category

	for rowCategor.Next() {
		if err := rowCategor.Scan(&category.ID, &category.ParentCategoryID, &category.Image, &category.IsHomeCategory); err != nil {
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

	rowsTrCategory, err := db.Query("SELECT lang_id,name FROM translation_category WHERE category_id = $1 AND deleted_at IS NULL", ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer func() {
		if err := rowsTrCategory.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

	var translations []models.TranslationCategory

	for rowsTrCategory.Next() {
		var translation models.TranslationCategory
		if err := rowsTrCategory.Scan(&translation.LangID, &translation.Name); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
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

func GetCategories(c *gin.Context) {

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

	// get data from deatabase
	rowCategor, err := db.Query("SELECT id,parent_category_id,image,is_home_category FROM categories WHERE deleted_at IS NULL")
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

	var categories []models.Category

	for rowCategor.Next() {
		var category models.Category

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
		defer func() {
			if err := rowsTrCategory.Close(); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  false,
					"message": err.Error(),
				})
				return
			}
		}()

		var translations []models.TranslationCategory

		for rowsTrCategory.Next() {
			var translation models.TranslationCategory
			if err := rowsTrCategory.Scan(&translation.LangID, &translation.Name); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  false,
					"message": err.Error(),
				})
				return
			}
			translations = append(translations, translation)
		}

		category.TranslationCategory = translations

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
		return []ResultCategory{}, err
	}
	defer func() ([]ResultCategory, error) {
		if err := db.Close(); err != nil {
			return []ResultCategory{}, err
		}
		return []ResultCategory{}, nil
	}()

	// get all category where parent category id is null
	rows, err := db.Query("SELECT categories.id,categories.image,translation_category.name FROM categories LEFT JOIN translation_category ON categories.id=translation_category.category_id WHERE translation_category.lang_id = $1 AND categories.parent_category_id IS NULL AND translation_category.deleted_at IS NULL AND categories.deleted_at IS NULL", langID)
	if err != nil {
		return []ResultCategory{}, err
	}
	defer func() ([]ResultCategory, error) {
		if err := rows.Close(); err != nil {
			return []ResultCategory{}, err
		}
		return []ResultCategory{}, nil
	}()

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
		defer func() ([]ResultCategory, error) {
			if err := rowss.Close(); err != nil {
				return []ResultCategory{}, err
			}
			return []ResultCategory{}, nil
		}()

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
			defer func() ([]ResultCategory, error) {
				if err := rowsss.Close(); err != nil {
					return []ResultCategory{}, err
				}
				return []ResultCategory{}, nil
			}()

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

	// check id
	rowCategor, err := db.Query("SELECT id FROM categories WHERE id = $1 AND deleted_at IS NULL", ID)
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

	resultProc, err := db.Query("CALL delete_category($1)", ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer func() {
		if err := resultProc.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

	resultPROC, err := db.Query("CALL after_delete_category($1)", ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer func() {
		if err := resultPROC.Close(); err != nil {
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

func RestoreCategoryByID(c *gin.Context) {

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

	// check ids
	rowCategor, err := db.Query("SELECT id FROM categories WHERE id = $1 AND deleted_at IS NOT NULL", ID)
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

	resultProc, err := db.Query("CALL restore_category($1)", ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer func() {
		if err := resultProc.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

	resultPROC, err := db.Query("CALL after_restore_category($1)", ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer func() {
		if err := resultPROC.Close(); err != nil {
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

func DeletePermanentlyCategoryByID(c *gin.Context) {

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

	// check id and get image of categories
	rowCategor, err := db.Query("SELECT id,image FROM categories WHERE id = $1 AND deleted_at IS NOT NULL", ID)
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

	// kategoriyanyn suraty uploads papkadan pozulyar
	if image != "" {
		if err := os.Remove("./" + image); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}

	// kategoriyanyn produktalarynyn main suratlaryn direktlary bazadan alynyar
	rowsMainImageProduct, err := db.Query("SELECT m.product_id,m.small,m.medium,m.large FROM products p INNER JOIN category_product c ON c.product_id=p.id INNER JOIN main_image m ON m.product_id = p.id WHERE c.category_id = $1", ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer func() {
		if err := rowsMainImageProduct.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

	var mainImages []models.MainImage

	for rowsMainImageProduct.Next() {
		var mainImage models.MainImage

		if err := rowsMainImageProduct.Scan(&mainImage.ProductID, &mainImage.Small, &mainImage.Medium, &mainImage.Large); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}

		mainImages = append(mainImages, mainImage)
	}

	// kategoriyanyn produktalarynyn main suratlary uploads papkadan pozulyar
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

	// kategoriyanyn produktalarynyn suratlarynyn direktlary bazadan alynyar
	rowsImagesProduct, err := db.Query("SELECT i.small,i.large FROM products p INNER JOIN category_product c ON c.product_id=p.id INNER JOIN images i ON i.product_id = p.id WHERE c.category_id = $1", ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer func() {
		if err := rowsImagesProduct.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

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

	// kategoriyanyn produktalarynyn suratlaryny uploads papkadan pozyaryn
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

	// kategoriya degisli harytlar bazadan pozulyar
	resultProduct, err := db.Query("DELETE FROM products USING category_product WHERE category_product.product_id = products.id AND category_product.category_id = $1", category_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer func() {
		if err := resultProduct.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

	// haryt pozulanda sol haryda degisli likelar hem pozulmaly
	for _, v := range mainImages {
		resultLike, err := db.Query("DELETE FROM likes WHERE product_id = $1", v.ProductID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
		defer func() {
			if err := resultLike.Close(); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  false,
					"message": err.Error(),
				})
				return
			}
		}()
	}

	// haryt pozulanda sol haryda degisli sebetdaki harytlar hem pozulmaly
	for _, v := range mainImages {
		resultCart, err := db.Query("DELETE FROM cart WHERE product_id = $1", v.ProductID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
		defer func() {
			if err := resultCart.Close(); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  false,
					"message": err.Error(),
				})
				return
			}
		}()
	}

	// haryt pozulanda sol haryda degisli sargyt edilen harytlar hem pozulmaly
	for _, v := range mainImages {
		resultOrderedProduct, err := db.Query("DELETE FROM ordered_products WHERE product_id = $1", v.ProductID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
		defer func() {
			if err := resultOrderedProduct.Close(); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  false,
					"message": err.Error(),
				})
				return
			}
		}()
	}

	// kategoriya pozulanda category_product tablisadaky maglumatlar hem pozulmaly
	resultCatProc, err := db.Query("DELETE FROM category_product WHERE category_id = $1", category_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer func() {
		if err := resultCatProc.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

	// kategoriyanyn child kategoriyalarynyn id - leri alynyar database - den
	rowChildCategory, err := db.Query("SELECT id FROM categories WHERE parent_category_id = $1", ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer func() {
		if err := rowChildCategory.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

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

	// child kategoriya degisli harytlaryn suratlarynyn direktleri bazadan alynyar we uploads papkadam pozulyar
	for _, v := range childCategoryIDS {
		rowPrdcs, err := db.Query("SELECT m.product_id,m.small,m.medium,m.large FROM products p INNER JOIN category_product c ON c.product_id=p.id INNER JOIN main_image m ON m.product_id = p.id WHERE c.category_id = $1", v)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
		defer func() {
			if err := rowPrdcs.Close(); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  false,
					"message": err.Error(),
				})
				return
			}
		}()

		var childMainImages []models.MainImage

		for rowsMainImageProduct.Next() {
			var childMainImage models.MainImage

			if err := rowsMainImageProduct.Scan(&childMainImage.ProductID, &childMainImage.Small, &childMainImage.Medium, &childMainImage.Large); err != nil {
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

		rowsChildImagesProduct, err := db.Query("SELECT i.small,i.large FROM products p INNER JOIN category_product c ON c.product_id=p.id INNER JOIN images i ON i.product_id = p.id WHERE c.category_id = $1", v)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
		defer func() {
			if err := rowsChildImagesProduct.Close(); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  false,
					"message": err.Error(),
				})
				return
			}
		}()

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

		// child kategoriya degisli harytlar bazadan pozulyar
		childresultProduct, err := db.Query("DELETE FROM products USING category_product WHERE category_product.product_id = products.id AND category_product.category_id = $1", v)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
		defer func() {
			if err := childresultProduct.Close(); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  false,
					"message": err.Error(),
				})
				return
			}
		}()

		// haryt pozulanda sol haryda degisli likelar hem pozulmaly
		for _, v := range childMainImages {
			resultLike, err := db.Query("DELETE FROM likes WHERE product_id = $1", v.ProductID)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  false,
					"message": err.Error(),
				})
				return
			}
			defer func() {
				if err := resultLike.Close(); err != nil {
					c.JSON(http.StatusBadRequest, gin.H{
						"status":  false,
						"message": err.Error(),
					})
					return
				}
			}()
		}

		// haryt pozulanda sol haryda degisli sebetdaki harytlar hem pozulmaly
		for _, v := range childMainImages {
			resultCart, err := db.Query("DELETE FROM cart WHERE product_id = $1", v.ProductID)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  false,
					"message": err.Error(),
				})
				return
			}
			defer func() {
				if err := resultCart.Close(); err != nil {
					c.JSON(http.StatusBadRequest, gin.H{
						"status":  false,
						"message": err.Error(),
					})
					return
				}
			}()
		}

		// haryt pozulanda sol haryda degisli sargyt edilen harytlar hem pozulmaly
		for _, v := range childMainImages {
			resultOrderedProduct, err := db.Query("DELETE FROM ordered_products WHERE product_id = $1", v.ProductID)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  false,
					"message": err.Error(),
				})
				return
			}
			defer func() {
				if err := resultOrderedProduct.Close(); err != nil {
					c.JSON(http.StatusBadRequest, gin.H{
						"status":  false,
						"message": err.Error(),
					})
					return
				}
			}()
		}

		// kategoriya pozulanda category_product tablisadaky maglumatlar hem pozulmaly
		resultCatProc, err := db.Query("DELETE FROM category_product WHERE category_id = $1", v)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
		defer func() {
			if err := resultCatProc.Close(); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  false,
					"message": err.Error(),
				})
				return
			}
		}()

	}

	// in sonunda kategpriyanyn ozi pozulyar
	resutCateg, err := db.Query("DELETE FROM categories WHERE id = $1", ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer func() {
		if err := resutCateg.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

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
	defer func() {
		if err := db.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

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
	defer func() {
		if err := categoryRow.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

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
		defer func() {
			if err := productCount.Close(); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  false,
					"message": err.Error(),
				})
				return
			}
		}()

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
		productRows, err := db.Query("SELECT p.id,p.price,p.old_price,p.limit_amount,p.is_new,p.amount FROM products p LEFT JOIN category_product c ON p.id=c.product_id WHERE c.category_id = $1 AND p.deleted_at IS NULL AND c.deleted_at IS NULL ORDER BY p.created_at ASC LIMIT $2 OFFSET $3", categoryID, limit, offset)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
		defer func() {
			if err := productRows.Close(); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  false,
					"message": err.Error(),
				})
				return
			}
		}()

		var products []Product

		for productRows.Next() {
			var product Product

			if err := productRows.Scan(&product.ID, &product.Price, &product.OldPrice, &product.LimitAmount, &product.IsNew, &product.Amount); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  false,
					"message": err.Error(),
				})
				return
			}

			if product.OldPrice != 0 {
				product.Percentage = -math.Round(((product.OldPrice - product.Price) * 100) / product.OldPrice)
			} else {
				product.Percentage = 0
			}

			if product.Amount != 0 {

				rowMainImage, err := db.Query("SELECT small,medium,large FROM main_image WHERE product_id = $1", product.ID)
				if err != nil {
					c.JSON(http.StatusBadRequest, gin.H{
						"status":  false,
						"message": err.Error(),
					})
					return
				}
				defer func() {
					if err := rowMainImage.Close(); err != nil {
						c.JSON(http.StatusBadRequest, gin.H{
							"status":  false,
							"message": err.Error(),
						})
						return
					}
				}()

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

					rowTrProduct, err := db.Query("SELECT name,description FROM translation_product WHERE lang_id = $1 AND product_id = $2 AND deleted_at IS NULL", langID, product.ID)
					if err != nil {
						c.JSON(http.StatusBadRequest, gin.H{
							"status":  false,
							"message": err.Error(),
						})
						return
					}
					defer func() {
						if err := rowTrProduct.Close(); err != nil {
							c.JSON(http.StatusBadRequest, gin.H{
								"status":  false,
								"message": err.Error(),
							})
							return
						}
					}()

					var trProduct models.TranslationProduct

					translation := make(map[string]models.TranslationProduct)

					for rowTrProduct.Next() {
						if err := rowTrProduct.Scan(&trProduct.Name, &trProduct.Description); err != nil {
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

				product.MainImage = mainImage

				rowsImages, err := db.Query("SELECT small,large FROM images WHERE product_id = $1", product.ID)
				if err != nil {
					c.JSON(http.StatusBadRequest, gin.H{
						"status":  false,
						"message": err.Error(),
					})
					return
				}
				defer func() {
					if err := rowsImages.Close(); err != nil {
						c.JSON(http.StatusBadRequest, gin.H{
							"status":  false,
							"message": err.Error(),
						})
						return
					}
				}()

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
				defer func() {
					if err := brendRows.Close(); err != nil {
						c.JSON(http.StatusBadRequest, gin.H{
							"status":  false,
							"message": err.Error(),
						})
						return
					}
				}()

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
		}
		category.Products = products
	}
	c.JSON(http.StatusOK, gin.H{
		"status":            true,
		"category":          category,
		"count_of_products": countOfProducts,
	})

}
