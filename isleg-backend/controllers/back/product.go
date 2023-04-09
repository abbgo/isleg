package controllers

import (
	"github/abbgo/isleg/isleg-backend/config"
	"github/abbgo/isleg/isleg-backend/models"
	"github/abbgo/isleg/isleg-backend/pkg"
	"math"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gosimple/slug"
	"github.com/lib/pq"
	"gopkg.in/guregu/null.v4"
)

type ProductForFront struct {
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
	Images       []string                               `json:"images,omitempty"`
	Translations []map[string]models.TranslationProduct `json:"translations"`
}

type DeleteImage struct {
	Image string `json:"image"`
}

type ProductForAdmin struct {
	ID                 string                      `json:"id,omitempty"`
	BrendID            null.String                 `json:"brend_id,omitempty"`
	ShopID             null.String                 `json:"shop_id,omitempty"`
	Price              float64                     `json:"price,omitempty" binding:"required"`
	OldPrice           float64                     `json:"old_price"`
	Benefit            null.Float                  `json:"benefit"`
	Amount             uint                        `json:"amount,omitempty" binding:"required"`
	LimitAmount        uint                        `json:"limit_amount,omitempty" binding:"required"`
	IsNew              bool                        `json:"is_new,omitempty"`
	MainImage          string                      `json:"main_image,omitempty"`
	Images             []string                    `json:"images,omitempty"`                                 // one to many
	TranslationProduct []models.TranslationProduct `json:"translation_product,omitempty" binding:"required"` // one to many
	Categories         []string                    `json:"categories,omitempty" binding:"required"`
}

func DeleteProductImages(c *gin.Context) {

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

	var image DeleteImage
	if err := c.Bind(&image); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	if image.Image == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "path of image is required",
		})
		return
	}

	row, err := db.Query("SELECT id FROM helper_images WHERE image = $1 AND deleted_at IS NULL", image.Image)
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
	var helperImageID string
	for row.Next() {
		if err := row.Scan(&helperImageID); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}
	if helperImageID != "" {
		resultHelperImage, err := db.Query("DELETE FROM helper_images WHERE id = $1", helperImageID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
		defer func() {
			if err := resultHelperImage.Close(); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  false,
					"message": err.Error(),
				})
				return
			}
		}()
	}

	if err := os.Remove(pkg.ServerPath + image.Image); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "image successfully deleted",
	})

}

func CreateProductImage(c *gin.Context) {

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

	var path, file_name string
	imageType := c.Query("image")

	oldImage := c.PostForm("old_path")
	if oldImage != "" {
		if err := os.Remove(pkg.ServerPath + oldImage); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}

		resultHelperImage, err := db.Query("DELETE FROM helper_images WHERE image = $1", oldImage)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
		defer func() {
			if err := resultHelperImage.Close(); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  false,
					"message": err.Error(),
				})
				return
			}
		}()
	}

	switch imageType {
	case "product":
		fileName := c.Query("type")
		if fileName != "main_image" && fileName != "image" {
			c.JSON(http.StatusNotFound, gin.H{
				"status":  false,
				"message": "invalid file name",
			})
			return
		}
		path = "product/" + fileName
		file_name = fileName
	case "category":
		path = "category"
		file_name = "image"
	case "brend":
		path = "brend"
		file_name = "image"
	default:
		c.JSON(http.StatusNotFound, gin.H{
			"status":  false,
			"message": "invalid image",
		})
		return
	}

	image, err := pkg.FileUpload(file_name, path, c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	result, err := db.Query("INSERT INTO helper_images (image) VALUES ($1)", image)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
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

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"image":  image,
	})

}

func CreateProduct(c *gin.Context) {

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

	var product ProductForAdmin
	if err := c.BindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if product.MainImage == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "main_image is required",
		})
		return
	}

	benefit, _, price, oldPrice, amount, limitAmount, isNew, err := models.ValidateProductModel("", product.Benefit.Float64, "", product.Price, product.OldPrice, product.Amount, product.LimitAmount, product.IsNew, product.Categories)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	var shopID, brendID interface{}
	if product.BrendID.String == "" {
		brendID = nil
	} else {
		brendID = product.BrendID.String
	}
	if product.ShopID.String == "" {
		shopID = nil
	} else {
		shopID = product.ShopID.String
	}

	// create product
	resultProducts, err := db.Query("INSERT INTO products (brend_id,price,old_price,amount,limit_amount,is_new,shop_id,main_image,benefit) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9) RETURNING id", brendID, price, oldPrice, amount, limitAmount, isNew, shopID, product.MainImage, benefit)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer func() {
		if err := resultProducts.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

	var productID string

	for resultProducts.Next() {
		if err := resultProducts.Scan(&productID); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}

	if len(product.Images) != 0 {
		// create images of product
		resultImages, err := db.Query("INSERT INTO images (product_id,image) VALUES ($1,unnest($2::varchar[]))", productID, pq.Array(product.Images))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
		defer func() {
			if err := resultImages.Close(); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  false,
					"message": err.Error(),
				})
				return
			}
		}()

		resultHelperImages, err := db.Query("DELETE FROM helper_images WHERE image = ANY($1)", pq.Array(product.Images))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
		defer func() {
			if err := resultHelperImages.Close(); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  false,
					"message": err.Error(),
				})
				return
			}
		}()
	}

	for _, v := range product.TranslationProduct {
		resultTrProducts, err := db.Query("INSERT INTO translation_product (lang_id,product_id,name,description,slug) VALUES ($1,$2,$3,$4,$5)", v.LangID, productID, v.Name, v.Description, slug.MakeLang(v.Name, "en"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
		defer func() {
			if err := resultTrProducts.Close(); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  false,
					"message": err.Error(),
				})
				return
			}
		}()
	}

	// create category product
	resultCategoryProduct, err := db.Query("INSERT INTO category_product (category_id,product_id) VALUES (unnest($1::uuid[]),$2)", pq.Array(product.Categories), productID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer func() {
		if err := resultCategoryProduct.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

	resultHelperImage, err := db.Query("DELETE FROM helper_images WHERE image = $1", product.MainImage)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer func() {
		if err := resultHelperImage.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "data successfully added",
	})

}

func UpdateProductByID(c *gin.Context) {

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

	var product ProductForAdmin
	if err := c.BindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	// get id from request parameter
	ID := c.Param("id")

	// validate data
	benefit, mainImage, price, oldPrice, amount, limitAmount, isNew, err := models.ValidateProductModel(product.MainImage, product.Benefit.Float64, ID, product.Price, product.OldPrice, product.Amount, product.LimitAmount, product.IsNew, product.Categories)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	var shopID, brendID interface{}
	if product.BrendID.String == "" {
		brendID = nil
	} else {
		brendID = product.BrendID.String
	}
	if product.ShopID.String == "" {
		shopID = nil
	} else {
		shopID = product.ShopID.String
	}

	resultProducts, err := db.Query("UPDATE products SET brend_id = $1 , price = $2 , old_price = $3, amount = $4, limit_amount = $5 , is_new = $6, shop_id = $8 , main_image = $9 , benefit = $10 WHERE id = $7", brendID, price, oldPrice, amount, limitAmount, isNew, ID, shopID, mainImage, benefit)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer func() {
		if err := resultProducts.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

	// update translation product
	for _, v := range product.TranslationProduct {
		resultTrProduct, err := db.Query("UPDATE translation_product SET name = $1, description = $2, slug = $3 WHERE product_id = $4 AND lang_id = $5", v.Name, v.Description, slug.MakeLang(v.Name, "en"), ID, v.LangID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
		defer func() {
			if err := resultTrProduct.Close(); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  false,
					"message": err.Error(),
				})
				return
			}
		}()
	}

	resultImages, err := db.Query("DELETE FROM images WHERE product_id = $1", ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer func() {
		if err := resultImages.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

	if len(product.Images) != 0 {
		resultImage, err := db.Query("INSERT INTO images (image,product_id) VALUES (unnest($1::varchar[]),$2)", pq.Array(product.Images), ID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
		defer func() {
			if err := resultImage.Close(); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  false,
					"message": err.Error(),
				})
				return
			}
		}()

		resultHelperImages, err := db.Query("DELETE FROM helper_images WHERE image = ANY($1)", pq.Array(product.Images))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
		defer func() {
			if err := resultHelperImages.Close(); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  false,
					"message": err.Error(),
				})
				return
			}
		}()
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
	defer func() {
		if err := resultCategoryProduct.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

	resultCategProduct, err := db.Query("INSERT INTO category_product (category_id,product_id) VALUES (unnest($1::uuid[]),$2)", pq.Array(product.Categories), ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer func() {
		if err := resultCategProduct.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "data successfully updated",
	})

}

func GetProductByID(c *gin.Context) {

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

	rowProduct, err := db.Query("SELECT id,brend_id,price,old_price,amount,limit_amount,is_new,main_image,benefit FROM products WHERE id = $1 AND deleted_at IS NULL", ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer func() {
		if err := rowProduct.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

	var product ProductForAdmin

	for rowProduct.Next() {
		if err := rowProduct.Scan(&product.ID, &product.BrendID, &product.Price, &product.OldPrice, &product.Amount, &product.LimitAmount, &product.IsNew, &product.MainImage, &product.Benefit); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}

	if product.ID == "" {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  false,
			"message": "record not found",
		})
		return
	}

	rowsImages, err := db.Query("SELECT image FROM images WHERE deleted_at IS NULL AND product_id = $1", ID)
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

	var images []string

	for rowsImages.Next() {
		var image string

		if err := rowsImages.Scan(&image); err != nil {
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
	defer func() {
		if err := rowsCategoryProduct.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

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
		c.JSON(http.StatusNotFound, gin.H{
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
	defer func() {
		if err := rowTranslationProduct.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

	var translations []models.TranslationProduct

	for rowTranslationProduct.Next() {
		var translation models.TranslationProduct
		if err := rowTranslationProduct.Scan(&translation.LangID, &translation.Name, &translation.Description); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
		if translation.Name == "" {
			c.JSON(http.StatusNotFound, gin.H{
				"status":  false,
				"message": "record not found",
			})
			return
		}
		translations = append(translations, translation)
	}

	product.TranslationProduct = translations

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"product": product,
	})

}

func GetProducts(c *gin.Context) {

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

	rowsProduct, err := db.Query("SELECT id,brend_id,price,old_price,amount,limit_amount,is_new,main_image,benefit FROM products WHERE deleted_at IS NULL")
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

	var products []models.Product
	// var ids []string

	for rowsProduct.Next() {
		var product models.Product

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

		rowsImages, err := db.Query("SELECT image FROM images WHERE deleted_at IS NULL AND product_id = $1", product.ID)
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

		var images []string

		for rowsImages.Next() {
			var image string

			if err := rowsImages.Scan(&image); err != nil {
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
		defer func() {
			if err := rowsCategoryProduct.Close(); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  false,
					"message": err.Error(),
				})
				return
			}
		}()

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
		defer func() {
			if err := rowTranslationProduct.Close(); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  false,
					"message": err.Error(),
				})
				return
			}
		}()

		var translations []models.TranslationProduct

		for rowTranslationProduct.Next() {
			var translation models.TranslationProduct
			if err := rowTranslationProduct.Scan(&translation.LangID, &translation.Name, &translation.Description); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  false,
					"message": err.Error(),
				})
				return
			}
			translations = append(translations, translation)
		}

		product.TranslationProduct = translations

		products = append(products, product)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":   true,
		"products": products,
	})

}

func DeleteProductByID(c *gin.Context) {

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
	rowProduct, err := db.Query("SELECT id FROM products WHERE id = $1 AND deleted_at IS NULL", ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer func() {
		if err := rowProduct.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

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
		c.JSON(http.StatusNotFound, gin.H{
			"status":  false,
			"message": "record not found",
		})
		return
	}

	resultProc, err := db.Query("CALL delete_product($1)", ID)
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

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "data successfully deleted",
	})

}

func RestoreProductByID(c *gin.Context) {

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
	rowProduct, err := db.Query("SELECT id FROM products WHERE id = $1 AND deleted_at IS NOT NULL", ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer func() {
		if err := rowProduct.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

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
		c.JSON(http.StatusNotFound, gin.H{
			"status":  false,
			"message": "record not found",
		})
		return
	}

	resultProc, err := db.Query("CALL restore_product($1)", ID)
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

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "data successfully restored",
	})

}

func DeletePermanentlyProductByID(c *gin.Context) {

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
	rowProduct, err := db.Query("SELECT id,main_image FROM products WHERE id = $1 AND deleted_at IS NOT NULL", ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer func() {
		if err := rowProduct.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

	var productID, mainImage string

	for rowProduct.Next() {
		if err := rowProduct.Scan(&productID, &mainImage); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}

	if productID == "" {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  false,
			"message": "product not found",
		})
		return
	}

	// remove main image of product
	if err := os.Remove(pkg.ServerPath + mainImage); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	// get images of product
	rowsImages, err := db.Query("SELECT image FROM images WHERE product_id = $1", productID)
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

		if err := rowsImages.Scan(&image.Image); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}

		images = append(images, image)
	}

	// remove images of product
	if len(images) != 0 {

		for _, v := range images {

			if err := os.Remove(pkg.ServerPath + v.Image); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  false,
					"message": err.Error(),
				})
				return
			}

		}

	}

	resultCart, err := db.Query("DELETE FROM cart WHERE product_id = $1", ID)
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

	resultLike, err := db.Query("DELETE FROM likes WHERE product_id = $1", ID)
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

	resultOrderedProc, err := db.Query("DELETE FROM ordered_products WHERE product_id = $1", ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer func() {
		if err := resultOrderedProc.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

	resultProduct, err := db.Query("DELETE FROM products WHERE id = $1", ID)
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

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "data successfully deleted",
	})

}

func GetProductByIDForFront(c *gin.Context) {

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

	rowProduct, err := db.Query("SELECT id,price,old_price,amount,limit_amount,is_new,main_image,benefit FROM products WHERE id = $1 AND deleted_at IS NULL", ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer func() {
		if err := rowProduct.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

	var product ProductForFront

	for rowProduct.Next() {
		if err := rowProduct.Scan(&product.ID, &product.Price, &product.OldPrice, &product.Amount, &product.LimitAmount, &product.IsNew, &product.MainImage, &product.Benefit); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
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

	if product.ID == "" {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  false,
			"message": "record not found",
		})
		return
	}

	rowsImages, err := db.Query("SELECT image FROM images WHERE deleted_at IS NULL AND product_id = $1", ID)
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

	var images []string

	for rowsImages.Next() {
		var image string

		if err := rowsImages.Scan(&image); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}

		images = append(images, image)
	}

	product.Images = images

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

		rowTrProduct, err := db.Query("SELECT name,description FROM translation_product WHERE lang_id = $1 AND product_id = $2 AND deleted_at IS NULL", v.ID, product.ID)
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

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"product": product,
	})

}
