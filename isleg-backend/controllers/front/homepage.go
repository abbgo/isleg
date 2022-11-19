package controllers

import (
	"net/http"

	"github/abbgo/isleg/isleg-backend/config"
	backController "github/abbgo/isleg/isleg-backend/controllers/back"
	"github/abbgo/isleg/isleg-backend/models"

	"github.com/gin-gonic/gin"
)

type HomePageCategory struct {
	ID       string    `json:"id"`
	Name     string    `json:"name"`
	Products []Product `json:"products"`
}

type Product struct {
	ID          string           `json:"id"`
	Price       float64          `json:"price"`
	OldPrice    float64          `json:"old_price"`
	MainImage   models.MainImage `json:"main_image"`
	Images      []models.Images  `json:"images"`
	Brend       Brend            `json:"brend"`
	LimitAmount int              `json:"limit_amount"`
	IsNew       bool             `json:"is_new"`
	Amount      int              `json:"amount"`
	// Translation models.TranslationProduct `json:"translation"`
	Translations []map[string]models.TranslationProduct `json:"translations"`
}

type Brend struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// ahli brendlerin suratlaryny we id - lerini getiryar
func GetBrends(c *gin.Context) {

	// get all brend from brend controller
	brends, err := backController.GetAllBrendForHomePage()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"brends": brends,
	})

}

func GetHomePageCategories(c *gin.Context) {

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

	langID, err := backController.CheckLanguage(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	// get all homepage category where translation category id equal langID
	categoryRows, err := db.Query("SELECT c.id,t.name FROM categories c LEFT JOIN translation_category t ON c.id=t.category_id WHERE t.lang_id = $1 AND c.is_home_category = true AND t.deleted_at IS NULL AND c.deleted_at IS NULL", langID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer func() {
		if err := categoryRows.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

	var homePageCategories []HomePageCategory

	for categoryRows.Next() {
		var homePageCategory HomePageCategory
		if err := categoryRows.Scan(&homePageCategory.ID, &homePageCategory.Name); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}

		// get all product where category id equal homePageCategory.ID and lang_id equal langID
		productRows, err := db.Query("SELECT p.id,p.price,p.old_price,p.limit_amount,p.is_new,p.amount FROM products p LEFT JOIN category_product c ON p.id=c.product_id WHERE c.category_id = $1 AND p.deleted_at IS NULL AND c.deleted_at IS NULL ORDER BY p.created_at DESC LIMIT 4", homePageCategory.ID)
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

				product.MainImage = mainImage

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

				// get brend where id of product brend_id
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
		homePageCategory.Products = products
		homePageCategories = append(homePageCategories, homePageCategory)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":              true,
		"homepage_categories": homePageCategories,
	})

}
