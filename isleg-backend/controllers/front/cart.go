package controllers

import (
	"errors"
	"github/abbgo/isleg/isleg-backend/config"
	backController "github/abbgo/isleg/isleg-backend/controllers/back"
	"github/abbgo/isleg/isleg-backend/models"
	"math"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// type DataForAddCart struct {
// 	Products []CartProduct `json:"products"`
// }

type CartProduct struct {
	ProductID         string `json:"product_id" binding:"required"`
	QuantityOfProduct int    `json:"quantity_of_product"`
}

type ProductOfCart struct {
	ID                 uuid.UUID                 `json:"id"`
	BrendID            uuid.UUID                 `json:"brend_id"`
	Price              float64                   `json:"price"`
	OldPrice           float64                   `json:"old_price"`
	Percentage         float64                   `json:"percentage"`
	Amount             uint                      `json:"amount"`
	LimitAmount        uint                      `json:"limit_amount"`
	IsNew              bool                      `json:"is_new"`
	QuantityOfProduct  int                       `json:"quantity_of_product"`
	MainImage          models.MainImage          `json:"main_image"`
	Images             []models.Images           `json:"images"`
	TranslationProduct models.TranslationProduct `json:"translation_product"`
}

func AddCart(c *gin.Context) {

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

	custID, hasCustomer := c.Get("customer_id")
	if !hasCustomer {
		c.JSON(http.StatusBadRequest, "customer_id is required")
		return
	}
	customerID, ok := custID.(string)
	if !ok {
		c.JSON(http.StatusBadRequest, "customer_id must be string")
	}

	langShortName := c.Param("lang")

	var cart []CartProduct

	if err := c.BindJSON(&cart); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if len(cart) != 0 {

		for k, v := range cart {

			if v.QuantityOfProduct < 1 {
				if err := DeleteCart(customerID, v.ProductID); err != nil {
					c.JSON(http.StatusBadRequest, gin.H{
						"status":  false,
						"message": err.Error(),
					})
					return
				}
				break
			}

			for _, x := range cart[(k + 1):] {
				if v.ProductID == x.ProductID {
					c.JSON(http.StatusBadRequest, gin.H{
						"status":  false,
						"message": "the same product cannot be repeated twice",
					})
					return
				}
			}

			rowProduct, err := db.Query("SELECT id FROM products WHERE id = $1 AND deleted_at IS NULL", v.ProductID)
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

			var product_id string

			for rowProduct.Next() {
				if err := rowProduct.Scan(&product_id); err != nil {
					c.JSON(http.StatusBadRequest, gin.H{
						"status":  false,
						"message": err.Error(),
					})
					return
				}
			}

			if product_id != "" {

				rowCart, err := db.Query("SELECT product_id FROM cart WHERE customer_id = $1 AND product_id = $2 AND deleted_at IS NULL", customerID, v.ProductID)
				if err != nil {
					c.JSON(http.StatusBadRequest, gin.H{
						"status":  false,
						"message": err.Error(),
					})
					return
				}
				defer func() {
					if err := rowCart.Close(); err != nil {
						c.JSON(http.StatusBadRequest, gin.H{
							"status":  false,
							"message": err.Error(),
						})
						return
					}
				}()

				var productId string

				for rowCart.Next() {
					if err := rowCart.Scan(&productId); err != nil {
						c.JSON(http.StatusBadRequest, gin.H{
							"status":  false,
							"message": err.Error(),
						})
						return
					}
				}

				if productId == "" {

					resultCartInsert, err := db.Query("INSERT INTO cart (customer_id,product_id,quantity_of_product) VALUES ($1,$2,$3)", customerID, v.ProductID, v.QuantityOfProduct)
					if err != nil {
						c.JSON(http.StatusBadRequest, gin.H{
							"status":  false,
							"message": err.Error(),
						})
						return
					}
					defer func() {
						if err := resultCartInsert.Close(); err != nil {
							c.JSON(http.StatusBadRequest, gin.H{
								"status":  false,
								"message": err.Error(),
							})
							return
						}
					}()

				} else {

					reultCartUpdate, err := db.Query("UPDATE cart SET quantity_of_product = $1 WHERE customer_id = $2 AND product_id = $3 AND deleted_at IS NULL", v.QuantityOfProduct, customerID, v.ProductID)
					if err != nil {
						c.JSON(http.StatusBadRequest, gin.H{
							"status":  false,
							"message": err.Error(),
						})
						return
					}
					defer func() {
						if err := reultCartUpdate.Close(); err != nil {
							c.JSON(http.StatusBadRequest, gin.H{
								"status":  false,
								"message": err.Error(),
							})
							return
						}
					}()

				}

			}

		}

		products, err := GetCartProducts(langShortName, customerID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status":   true,
			"products": products,
		})

	} else {

		products, err := GetCartProducts(langShortName, customerID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}

		if len(products) != 0 {
			c.JSON(http.StatusOK, gin.H{
				"status":   true,
				"products": products,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status":  true,
				"message": "cart empty",
			})
		}

	}

}

func GetCartProducts(langShortName, customerID string) ([]ProductOfCart, error) {

	db, err := config.ConnDB()
	if err != nil {
		return []ProductOfCart{}, err
	}
	defer db.Close()

	// GET DATA FROM ROUTE PARAMETER
	// langShortName := c.Param("lang")

	// GET language id
	langID, err := backController.GetLangID(langShortName)
	if err != nil {
		return []ProductOfCart{}, err
	}

	rowsProduct, err := db.Query("SELECT p.id,p.brend_id,p.price,p.old_price,p.amount,p.limit_amount,p.is_new,c.quantity_of_product FROM products p LEFT JOIN cart c ON c.product_id = p.id WHERE c.customer_id = $1 AND c.deleted_at IS NULL AND p.deleted_at IS NULL", customerID)
	if err != nil {
		return []ProductOfCart{}, err
	}
	defer rowsProduct.Close()

	var products []ProductOfCart

	for rowsProduct.Next() {
		var product ProductOfCart

		if err := rowsProduct.Scan(&product.ID, &product.BrendID, &product.Price, &product.OldPrice, &product.Amount, &product.LimitAmount, &product.IsNew, &product.QuantityOfProduct); err != nil {
			return []ProductOfCart{}, err
		}

		if product.OldPrice != 0 {
			product.Percentage = -math.Round(((product.OldPrice - product.Price) * 100) / product.OldPrice)
		} else {
			product.Percentage = 0
		}

		rowMainImage, err := db.Query("SELECT small,medium,large FROM main_image WHERE product_id = $1 AND deleted_at IS NULL", product.ID)
		if err != nil {
			return []ProductOfCart{}, err
		}
		defer rowMainImage.Close()

		var mainImage models.MainImage

		for rowMainImage.Next() {
			if err := rowMainImage.Scan(&mainImage.Small, &mainImage.Medium, &mainImage.Large); err != nil {
				return []ProductOfCart{}, err
			}
		}

		product.MainImage = mainImage

		rowsImages, err := db.Query("SELECT small,large FROM images WHERE product_id = $1 AND deleted_at IS NULL", product.ID)
		if err != nil {
			return []ProductOfCart{}, err
		}
		defer rowsImages.Close()

		var images []models.Images

		for rowsImages.Next() {
			var image models.Images

			if err := rowsImages.Scan(&image.Small, &image.Large); err != nil {
				return []ProductOfCart{}, err
			}

			images = append(images, image)
		}

		product.Images = images

		rowTrProduct, err := db.Query("SELECT name,description FROM translation_product WHERE lang_id = $1 AND product_id = $2", langID, product.ID)
		if err != nil {
			return []ProductOfCart{}, err
		}
		defer rowTrProduct.Close()

		var trProduct models.TranslationProduct

		for rowTrProduct.Next() {
			if err := rowTrProduct.Scan(&trProduct.Name, &trProduct.Description); err != nil {
				return []ProductOfCart{}, err
			}
		}

		product.TranslationProduct = trProduct

		products = append(products, product)
	}

	return products, nil

}

func GetCustomerCartProducts(c *gin.Context) {

	custID, hasCustomer := c.Get("customer_id")
	if !hasCustomer {
		c.JSON(http.StatusBadRequest, "customer_id is required")
		return
	}
	customerID, ok := custID.(string)
	if !ok {
		c.JSON(http.StatusBadRequest, "customer_id must be string")
	}

	langShortName := c.Param("lang")

	products, err := GetCartProducts(langShortName, customerID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":   true,
		"products": products,
	})

}

func RemoveCart(c *gin.Context) {

	custID, hasCustomer := c.Get("customer_id")
	if !hasCustomer {
		c.JSON(http.StatusBadRequest, "customer_id is required")
		return
	}
	customerID, ok := custID.(string)
	if !ok {
		c.JSON(http.StatusBadRequest, "customer_id must be string")
	}

	// productID := c.PostForm("product_id")
	productID := ""

	if err := DeleteCart(customerID, productID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "product successfull deleted from cart",
	})

}

func DeleteCart(customerID, productID string) error {

	db, err := config.ConnDB()
	if err != nil {
		return err
	}
	defer func() error {
		if err := db.Close(); err != nil {
			return err
		}
		return nil
	}()

	if productID != "" {

		rowCart, err := db.Query("SELECT product_id FROM cart WHERE customer_id = $1 AND product_id = $2 AND deleted_at IS NULL", customerID, productID)
		if err != nil {
			return err
		}
		defer func() error {
			if err := rowCart.Close(); err != nil {
				return err
			}
			return nil
		}()

		var product_id string

		for rowCart.Next() {
			if err := rowCart.Scan(&product_id); err != nil {
				return err
			}
		}

		if product_id == "" {
			return errors.New("this product not found in this customer")
		}

		resultCart, err := db.Query("DELETE FROM cart WHERE customer_id = $1 AND product_id = $2 AND deleted_at IS NULL", customerID, productID)
		if err != nil {
			return err
		}
		defer func() error {
			if err := resultCart.Close(); err != nil {
				return err
			}
			return nil
		}()

	} else {

		resultCart, err := db.Query("DELETE FROM cart WHERE customer_id = $1 AND deleted_at IS NULL", customerID)
		if err != nil {
			return err
		}
		defer func() error {
			if err := resultCart.Close(); err != nil {
				return err
			}
			return nil
		}()

	}
	return nil

}
