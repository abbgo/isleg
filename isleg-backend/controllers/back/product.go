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

// struct used in function GetProductByID
type OneProduct struct {
	ID           uuid.UUID            `json:"id"`
	BrendID      uuid.UUID            `json:"brend_id"`
	Price        float64              `json:"price"`
	OldPrice     float64              `json:"old_price"`
	Amount       uint                 `json:"amount"`
	ProductCode  string               `json:"product_code"`
	MainImage    models.MainImage     `json:"main_image"`
	Images       []models.Images      `json:"images"`
	Categories   []string             `json:"categories"`
	Translations []TranslationProduct `json:"translations"`
	LimitAmount  uint                 `json:"limit_amount"`
	IsNew        bool                 `json:"is_new"`
}
type TranslationProduct struct {
	LanguageID  string `json:"lang_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func CreateProduct(c *gin.Context) {

	db, err := config.ConnDB()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer db.Close()

	// validate brend id
	brendID := c.PostForm("brend_id")
	rowBrend, err := db.Query("SELECT id FROM brends WHERE id = $1 AND deleted_at IS NULL", brendID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer rowBrend.Close()

	var brend_id string

	for rowBrend.Next() {
		if err := rowBrend.Scan(&brend_id); err != nil {
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

	limitAmountStr := c.PostForm("limit_amount")
	limitAmount, err := strconv.ParseUint(limitAmountStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	if limitAmount > amount {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "cannot be less than limit_amount amount",
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

	isNewStr := c.PostForm("is_new")
	isNew, err := strconv.ParseBool(isNewStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
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

	mainSamllFile, err := c.FormFile("main_small")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	extMainSamll := filepath.Ext(mainSamllFile.Filename)
	if extMainSamll != ".jpg" && extMainSamll != ".jpeg" && extMainSamll != ".png" && extMainSamll != ".gif" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "the file must be an image.",
		})
		return
	}

	mainSamll := uuid.New().String() + extMainSamll
	c.SaveUploadedFile(mainSamllFile, "./uploads/product/"+mainSamll)

	mainMediumFile, err := c.FormFile("main_medium")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	extMainMedium := filepath.Ext(mainMediumFile.Filename)
	if extMainMedium != ".jpg" && extMainMedium != ".jpeg" && extMainMedium != ".png" && extMainMedium != ".gif" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "the file must be an image.",
		})
		return
	}

	mainMedium := uuid.New().String() + extMainMedium
	c.SaveUploadedFile(mainMediumFile, "./uploads/product/"+mainMedium)

	mainLargeFile, err := c.FormFile("main_large")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	extMainLarge := filepath.Ext(mainLargeFile.Filename)
	if extMainLarge != ".jpg" && extMainLarge != ".jpeg" && extMainLarge != ".png" && extMainLarge != ".gif" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "the file must be an image.",
		})
		return
	}

	mainLarge := uuid.New().String() + extMainLarge
	c.SaveUploadedFile(mainLargeFile, "./uploads/product/"+mainLarge)

	smallFiles := c.Request.MultipartForm.File["small"]
	largeFiles := c.Request.MultipartForm.File["large"]

	if len(smallFiles) != len(largeFiles) {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "images are incomplete",
		})
		return
	}

	var smalls []string

	for _, v := range smallFiles {
		extension := filepath.Ext(v.Filename)
		//validate image
		if extension != ".jpg" && extension != ".jpeg" && extension != ".png" && extension != ".gif" {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": "the file must be an image.",
			})
			return
		}
		fileName := uuid.New().String() + extension
		c.SaveUploadedFile(v, "./uploads/product/"+fileName)
		smalls = append(smalls, "uploads/product/"+fileName)
	}

	var larges []string

	for _, v := range largeFiles {
		extension := filepath.Ext(v.Filename)
		//validate image
		if extension != ".jpg" && extension != ".jpeg" && extension != ".png" && extension != ".gif" {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": "the file must be an image.",
			})
			return
		}
		fileName := uuid.New().String() + extension
		c.SaveUploadedFile(v, "./uploads/product/"+fileName)
		larges = append(larges, "uploads/product/"+fileName)
	}

	// create product
	resultProducts, err := db.Query("INSERT INTO products (brend_id,price,old_price,amount,product_code,limit_amount,is_new) VALUES ($1,$2,$3,$4,$5,$6,$7)", brendID, price, oldPrice, amount, productCode, limitAmount, isNew)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer resultProducts.Close()

	// get the id of the added product
	lastProductID, err := db.Query("SELECT id FROM products WHERE deleted_at IS NULL ORDER BY created_at DESC LIMIT 1")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer lastProductID.Close()

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

	resultMainImage, err := db.Query("INSERT INTO main_image (product_id,small,medium,large) VALUES ($1,$2,$3,$4)", productID, "uploads/product/"+mainSamll, "uploads/product/"+mainMedium, "uploads/product/"+mainLarge)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer resultMainImage.Close()

	for i := 0; i < len(smalls); i++ {

		resultImages, err := db.Query("INSERT INTO images (product_id,small,large) VALUES ($1,$2,$3)", productID, smalls[i], larges[i])
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
		defer resultImages.Close()

	}

	// create translation product
	for _, v := range languages {
		resultTrProducts, err := db.Query("INSERT INTO translation_product (lang_id,product_id,name,description) VALUES ($1,$2,$3,$4)", v.ID, productID, c.PostForm("name_"+v.NameShort), c.PostForm("description_"+v.NameShort))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
		defer resultTrProducts.Close()
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
		rawCategory, err := db.Query("SELECT id FROM categories WHERE id = $1 AND deleted_at IS NULL", v)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
		defer rawCategory.Close()

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
		resultCategoryProduct, err := db.Query("INSERT INTO category_product (category_id,product_id) VALUES ($1,$2)", v, productID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
		defer resultCategoryProduct.Close()
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "product successfully added",
	})

}

func UpdateProductByID(c *gin.Context) {

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
	var mainSmall, mainMedium, mainLarge string

	rowProduct, err := db.Query("SELECT id FROM products WHERE id = $1 AND deleted_at IS NULL", ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer rowProduct.Close()

	// var product ProductImage
	var id string

	for rowProduct.Next() {
		if err := rowProduct.Scan(&id); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}

	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "record not found",
		})
		return
	}

	// validate brend id
	brendID := c.PostForm("brend_id")
	rowBrend, err := db.Query("SELECT id FROM brends WHERE id = $1 AND deleted_at IS NULL", brendID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer rowBrend.Close()

	var brend_id string

	for rowBrend.Next() {
		if err := rowBrend.Scan(&brend_id); err != nil {
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

	limitAmountStr := c.PostForm("limit_amount")
	limitAmount, err := strconv.ParseUint(limitAmountStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	if limitAmount > amount {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "cannot be less than limit_amount amount",
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

	isNewStr := c.PostForm("is_new")
	isNew, err := strconv.ParseBool(isNewStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
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

	rowMainImage, err := db.Query("SELECT small,medium,large FROM main_image WHERE deleted_at IS NULL AND product_id = $1", ID)
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

	if mainImage.Small == "" || mainImage.Medium == "" || mainImage.Large == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "main image of product not found",
		})
		return
	}

	mainSamllFile, err := c.FormFile("main_small")
	if err != nil {
		mainSmall = mainImage.Small
	} else {
		extMainSamll := filepath.Ext(mainSamllFile.Filename)
		if extMainSamll != ".jpg" && extMainSamll != ".jpeg" && extMainSamll != ".png" && extMainSamll != ".gif" {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": "the file must be an image.",
			})
			return
		}

		newFileName := uuid.New().String() + extMainSamll
		c.SaveUploadedFile(mainSamllFile, "./uploads/product/"+newFileName)

		if err := os.Remove("./" + mainImage.Small); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}

		mainSmall = "uploads/product/" + newFileName
	}

	mainMediumFile, err := c.FormFile("main_medium")
	if err != nil {
		mainMedium = mainImage.Medium
	} else {
		extMainMedium := filepath.Ext(mainMediumFile.Filename)
		if extMainMedium != ".jpg" && extMainMedium != ".jpeg" && extMainMedium != ".png" && extMainMedium != ".gif" {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": "the file must be an image.",
			})
			return
		}

		newFileName := uuid.New().String() + extMainMedium
		c.SaveUploadedFile(mainMediumFile, "./uploads/product/"+newFileName)

		if err := os.Remove("./" + mainImage.Medium); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}

		mainMedium = "uploads/product/" + newFileName
	}

	mainLargeFile, err := c.FormFile("main_large")
	if err != nil {
		mainLarge = mainImage.Large
	} else {
		extMainLarge := filepath.Ext(mainLargeFile.Filename)
		if extMainLarge != ".jpg" && extMainLarge != ".jpeg" && extMainLarge != ".png" && extMainLarge != ".gif" {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": "the file must be an image.",
			})
			return
		}

		newFileName := uuid.New().String() + extMainLarge
		c.SaveUploadedFile(mainLargeFile, "./uploads/product/"+newFileName)

		if err := os.Remove("./" + mainImage.Large); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}

		mainLarge = "uploads/product/" + newFileName
	}

	rowImages, err := db.Query("SELECT small,large FROM images WHERE deleted_at IS NULL AND product_id = $1", ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer rowImages.Close()

	var images []models.Images

	for rowImages.Next() {
		var image models.Images

		if err := rowImages.Scan(&image.Small, &image.Large); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}

		images = append(images, image)
	}

	smallFiles := c.Request.MultipartForm.File["small"]
	var smalls []string

	largeFiles := c.Request.MultipartForm.File["large"]
	var larges []string

	if len(smallFiles) != len(largeFiles) {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "images are incomplete",
		})
		return
	}

	if len(smallFiles) != 0 {

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

		resultImages, err := db.Query("DELETE FROM images WHERE product_id = $1", ID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
		defer resultImages.Close()

		for _, v := range smallFiles {
			extension := filepath.Ext(v.Filename)
			//validate image
			if extension != ".jpg" && extension != ".jpeg" && extension != ".png" && extension != ".gif" {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  false,
					"message": "the file must be an image.",
				})
				return
			}
			fileName := uuid.New().String() + extension
			c.SaveUploadedFile(v, "./uploads/product/"+fileName)
			smalls = append(smalls, "uploads/product/"+fileName)
		}

		for _, v := range largeFiles {
			extension := filepath.Ext(v.Filename)
			//validate image
			if extension != ".jpg" && extension != ".jpeg" && extension != ".png" && extension != ".gif" {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  false,
					"message": "the file must be an image.",
				})
				return
			}
			fileName := uuid.New().String() + extension
			c.SaveUploadedFile(v, "./uploads/product/"+fileName)
			larges = append(larges, "uploads/product/"+fileName)
		}

		for i := 0; i < len(smalls); i++ {

			resultImages, err := db.Query("INSERT INTO images (product_id,small,large) VALUES ($1,$2,$3)", ID, smalls[i], larges[i])
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  false,
					"message": err.Error(),
				})
				return
			}
			defer resultImages.Close()

		}

	}

	currentTime := time.Now()

	resultProducts, err := db.Query("UPDATE products SET brend_id = $1 , price = $2 , old_price = $3, amount = $4, product_code = $5, limit_amount = $6 , is_new = $7 , updated_at = $8 WHERE id = $9", brendID, price, oldPrice, amount, productCode, limitAmount, isNew, currentTime, ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer resultProducts.Close()

	resultMainImage, err := db.Query("UPDATE main_image SET small = $1 , medium = $2 , large = $3 , updated_at = $4 WHERE product_id = $5", mainSmall, mainMedium, mainLarge, currentTime, ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer resultMainImage.Close()

	// update translation product
	for _, v := range languages {
		resultTrProduct, err := db.Query("UPDATE translation_product SET name = $1, description = $2, updated_at = $3 WHERE product_id = $4 AND lang_id = $5", c.PostForm("name_"+v.NameShort), c.PostForm("description_"+v.NameShort), currentTime, ID, v.ID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
		defer resultTrProduct.Close()
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
		rawCategory, err := db.Query("SELECT id FROM categories WHERE id = $1 AND deleted_at IS NULL", v)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
		defer rawCategory.Close()

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

	// update category product
	resultCategoryProduct, err := db.Query("DELETE FROM category_product WHERE product_id = $1", ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer resultCategoryProduct.Close()

	for _, v := range categories {
		resultCategProduct, err := db.Query("INSERT INTO category_product (category_id,product_id,updated_at) VALUES ($1,$2,$3)", v, ID, currentTime)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
		defer resultCategProduct.Close()
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "product successfully updated",
	})

}

func GetProductByID(c *gin.Context) {

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

	rowProduct, err := db.Query("SELECT id,brend_id,price,old_price,amount,product_code,limit_amount,is_new FROM products WHERE id = $1 AND deleted_at IS NULL", ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer rowProduct.Close()

	var product OneProduct

	for rowProduct.Next() {
		if err := rowProduct.Scan(&product.ID, &product.BrendID, &product.Price, &product.OldPrice, &product.Amount, &product.ProductCode, &product.LimitAmount, &product.IsNew); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}

	if product.ID.String() == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "record not found",
		})
		return
	}

	rowMainImage, err := db.Query("SELECT small,medium,large FROM main_image WHERE deleted_at IS NULL AND product_id = $1", ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer rowMainImage.Close()

	for rowMainImage.Next() {
		if err := rowMainImage.Scan(&product.MainImage.Small, &product.MainImage.Medium, &product.MainImage.Large); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}

	rowsImages, err := db.Query("SELECT small,large FROM images WHERE deleted_at IS NULL AND product_id = $1", ID)
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

	rowsCategoryProduct, err := db.Query("SELECT category_id FROM category_product WHERE product_id = $1 AND deleted_at IS NULL", ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer rowsCategoryProduct.Close()

	var categories []string

	for rowsCategoryProduct.Next() {
		var category string

		if err := rowsCategoryProduct.Scan(&category); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}

		categories = append(categories, category)
	}

	if len(categories) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "record not found",
		})
		return
	}

	product.Categories = categories

	rowTranslationProduct, err := db.Query("SELECT lang_id,name,description FROM translation_product WHERE product_id = $1 AND deleted_at IS NULL", ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer rowTranslationProduct.Close()

	var translations []TranslationProduct

	for rowTranslationProduct.Next() {
		var translation TranslationProduct
		if err := rowTranslationProduct.Scan(&translation.LanguageID, &translation.Name, &translation.Description); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
		if translation.Name == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": "record not found",
			})
			return
		}
		translations = append(translations, translation)
	}

	product.Translations = translations

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"product": product,
	})

}

func GetProducts(c *gin.Context) {

	db, err := config.ConnDB()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer db.Close()

	rowsProduct, err := db.Query("SELECT id,brend_id,price,old_price,amount,product_code,limit_amount,is_new FROM products WHERE deleted_at IS NULL")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer rowsProduct.Close()

	var products []OneProduct
	// var ids []string

	for rowsProduct.Next() {
		var product OneProduct

		if err := rowsProduct.Scan(&product.ID, &product.BrendID, &product.Price, &product.OldPrice, &product.Amount, &product.ProductCode, &product.LimitAmount, &product.IsNew); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}

		rowMainImage, err := db.Query("SELECT small,medium,large FROM main_image WHERE deleted_at IS NULL AND product_id = $1", product.ID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
		defer rowMainImage.Close()

		for rowMainImage.Next() {
			if err := rowMainImage.Scan(&product.MainImage.Small, &product.MainImage.Medium, &product.MainImage.Large); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  false,
					"message": err.Error(),
				})
				return
			}
		}

		rowsImages, err := db.Query("SELECT small,large FROM images WHERE deleted_at IS NULL AND product_id = $1", product.ID)
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

		rowsCategoryProduct, err := db.Query("SELECT category_id FROM category_product WHERE product_id = $1 AND deleted_at IS NULL", product.ID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
		defer rowsCategoryProduct.Close()

		var categories []string

		for rowsCategoryProduct.Next() {
			var category string
			if err := rowsCategoryProduct.Scan(&category); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  false,
					"message": err.Error(),
				})
				return
			}

			categories = append(categories, category)
		}

		product.Categories = categories

		rowTranslationProduct, err := db.Query("SELECT lang_id,name,description FROM translation_product WHERE product_id = $1 AND deleted_at IS NULL", product.ID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
		defer rowTranslationProduct.Close()

		var translations []TranslationProduct

		for rowTranslationProduct.Next() {
			var translation TranslationProduct
			if err := rowTranslationProduct.Scan(&translation.LanguageID, &translation.Name, &translation.Description); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  false,
					"message": err.Error(),
				})
				return
			}
			translations = append(translations, translation)
		}

		product.Translations = translations

		products = append(products, product)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":   true,
		"products": products,
	})

}

func DeleteProductByID(c *gin.Context) {

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

	rowProduct, err := db.Query("SELECT id FROM products WHERE id = $1 AND deleted_at IS NULL", ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer rowProduct.Close()

	var productID string

	for rowProduct.Next() {
		if err := rowProduct.Scan(&productID); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}

	if productID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "record not found",
		})
		return
	}

	currentTime := time.Now()

	resultProduct, err := db.Query("UPDATE products SET deleted_at = $1 WHERE id = $2", currentTime, ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer resultProduct.Close()

	resultCategoryProduct, err := db.Query("UPDATE category_product SET deleted_at = $1 WHERE product_id = $2", currentTime, ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer resultCategoryProduct.Close()

	resultTRProduct, err := db.Query("UPDATE translation_product SET deleted_at = $1 WHERE product_id = $2", currentTime, ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer resultTRProduct.Close()

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "product successfully deleted",
	})

}

func RestoreProductByID(c *gin.Context) {

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

	rowProduct, err := db.Query("SELECT id FROM products WHERE id = $1 AND deleted_at IS NOT NULL", ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer rowProduct.Close()

	var productID string

	for rowProduct.Next() {
		if err := rowProduct.Scan(&productID); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}

	if productID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "record not found",
		})
		return
	}

	resultProduct, err := db.Query("UPDATE products SET deleted_at = NULL WHERE id = $1", ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer resultProduct.Close()

	resultCategoryProduct, err := db.Query("UPDATE category_product SET deleted_at = NULL WHERE product_id = $1", ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer resultCategoryProduct.Close()

	resultTrProduct, err := db.Query("UPDATE translation_product SET deleted_at = NULL WHERE product_id = $1", ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer resultTrProduct.Close()

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "product successfully restored",
	})

}

func DeletePermanentlyProductByID(c *gin.Context) {

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

	rowProduct, err := db.Query("SELECT id FROM products WHERE id = $1 AND deleted_at IS NOT NULL", ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer rowProduct.Close()

	var productID string

	for rowProduct.Next() {
		if err := rowProduct.Scan(&productID); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}

	if productID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "product not found",
		})
		return
	}

	rowMainImage, err := db.Query("SELECT small,medium,large FROM main_image WHERE product_id = $1", productID)
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

	if err := os.Remove("./" + mainImage.Small); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	if err := os.Remove("./" + mainImage.Medium); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	if err := os.Remove("./" + mainImage.Large); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	rowsImages, err := db.Query("SELECT small,large FROM images WHERE product_id = $1", productID)
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

	if len(images) != 0 {

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

	}

	resultProduct, err := db.Query("DELETE FROM products WHERE id = $1", ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer resultProduct.Close()

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "product successfully deleted",
	})

}
