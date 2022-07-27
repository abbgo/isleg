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

func CreateProduct(c *gin.Context) {

	// validate brend id
	brendID := c.PostForm("brend_id")
	brendIDUUID, err := uuid.Parse(brendID)
	if brendID != "" {
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}
	_, err = config.ConnDB().Query("SELECT id FROM brends WHERE id = $1", brendID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	// validate data from request
	priceStr := c.PostForm("price")
	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	oldPriceStr := c.PostForm("old_price")
	var oldPrice float64
	if oldPriceStr != "" {
		oldPrice, err = strconv.ParseFloat(oldPriceStr, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}

		if oldPrice < price {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": "cannot be less than oldPrice Price",
			})
			return
		}

	} else {
		oldPrice = 0
	}

	amountStr := c.PostForm("amount")
	amount, err := strconv.ParseUint(amountStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	productCode := c.PostForm("product_code")
	if productCode == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "product code is required",
		})
		return
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

	// validate name and description
	for _, v := range languages {
		if c.PostForm("name_"+v.NameShort) == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": "name_" + v.NameShort + " is required",
			})
			return
		}
	}
	for _, v := range languages {
		if c.PostForm("description_"+v.NameShort) == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": "description_" + v.NameShort + " is required",
			})
			return
		}
	}

	// file upload
	if err := c.Request.ParseMultipartForm(2000000); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	// upload main_image
	mainImagePathFile, err := c.FormFile("main_image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	// validate image
	extensionFile := filepath.Ext(mainImagePathFile.Filename)
	if extensionFile != ".jpg" && extensionFile != ".jpeg" && extensionFile != ".png" && extensionFile != ".gif" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "the file must be an image.",
		})
		return
	}

	newFileName := "productMain" + uuid.New().String() + extensionFile
	c.SaveUploadedFile(mainImagePathFile, "./uploads/"+newFileName)

	// upload images
	files := c.Request.MultipartForm.File["images"]
	var imagePaths []string
	for _, v := range files {
		extension := filepath.Ext(v.Filename)
		//validate image
		if extension != ".jpg" && extension != ".jpeg" && extension != ".png" && extension != ".gif" {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": "the file must be an image.",
			})
			return
		}
		fileName := "product" + uuid.New().String() + extension
		c.SaveUploadedFile(v, "./uploads/"+fileName)
		imagePaths = append(imagePaths, "uploads/"+fileName)
	}

	// create product
	_, err = config.ConnDB().Exec("INSERT INTO products (brend_id,price,old_price,amount,product_code,main_image,images) VALUES ($1,$2,$3,$4,$5,$6,$7)", brendIDUUID, price, oldPrice, amount, productCode, "uploads/"+newFileName, pq.StringArray(imagePaths))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	// get the id of the added product
	lastProductID, err := config.ConnDB().Query("SELECT id FROM products ORDER BY created_at DESC LIMIT 1")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	var productID string

	for lastProductID.Next() {
		if err := lastProductID.Scan(&productID); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}

	// create translation product
	for _, v := range languages {
		_, err := config.ConnDB().Exec("INSERT INTO translation_product (lang_id,product_id,name,description) VALUES ($1,$2,$3,$4)", v.ID, productID, c.PostForm("name_"+v.NameShort), c.PostForm("description_"+v.NameShort))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}

	// get all category id from request
	categories, _ := c.GetPostFormArray("category_id")
	if len(categories) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "product category is required",
		})
		return
	}

	for _, v := range categories {
		rawCategory, err := config.ConnDB().Query("SELECT id FROM categories WHERE id = $1", v)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}

		var categoryID string

		for rawCategory.Next() {
			if err := rawCategory.Scan(&categoryID); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  false,
					"message": err.Error(),
				})
				return
			}
		}

		if categoryID == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": "category not found",
			})
			return
		}
	}

	// create category product
	for _, v := range categories {
		vUUID, err := uuid.Parse(v)
		if v != "" {
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  false,
					"message": err.Error(),
				})
				return
			}
		}
		_, err = config.ConnDB().Exec("INSERT INTO category_product (category_id,product_id) VALUES ($1,$2)", vUUID, productID)
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
		"message": "product successfully added",
	})

}
