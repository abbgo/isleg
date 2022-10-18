package controllers

import (
	"errors"
	"github/abbgo/isleg/isleg-backend/config"
	backController "github/abbgo/isleg/isleg-backend/controllers/back"
	"github/abbgo/isleg/isleg-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Cart struct {
	CustomerID string        `json:"customer_id"`
	Products   []CartProduct `json:"products"`
}

type CartProduct struct {
	ProductID         string `json:"product_id" binding:"required"`
	QuantityOfProduct int    `json:"quantity_of_product" binding:"required"`
}

type ProductOfCart struct {
	ID                 uuid.UUID                 `json:"id"`
	BrendID            uuid.UUID                 `json:"brend_id"`
	Price              float64                   `json:"price"`
	OldPrice           float64                   `json:"old_price"`
	Amount             uint                      `json:"amount"`
	ProductCode        string                    `json:"product_code"`
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

	// langShortName := c.Param("lang")
	var cart Cart

	if err := c.BindJSON(&cart); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	rowCustomer, err := db.Query("SELECT id FROM customers WHERE id = $1 AND deleted_at IS NULL", cart.CustomerID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer func() {
		if err := rowCustomer.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

	var customer_id string

	for rowCustomer.Next() {
		if err := rowCustomer.Scan(&customer_id); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}

	if customer_id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "Customer not found",
		})
		return
	}

	for k, v := range cart.Products {

		if v.QuantityOfProduct < 1 {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": "quantity of product cannot be less than 1",
			})
			return
		}

		for _, x := range cart.Products[(k + 1):] {
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

		if product_id == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": "product not found",
			})
			return
		}

		rowCart, err := db.Query("SELECT product_id FROM cart WHERE customer_id = $1 AND product_id = $2 AND deleted_at IS NULL", cart.CustomerID, v.ProductID)
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

			resultCartInsert, err := db.Query("INSERT INTO cart (customer_id,product_id,quantity_of_product) VALUES ($1,$2,$3)", cart.CustomerID, v.ProductID, v.QuantityOfProduct)
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

			reultCartUpdate, err := db.Query("UPDATE cart SET quantity_of_product = $1 WHERE customer_id = $2 AND product_id = $3 AND deleted_at IS NULL", v.QuantityOfProduct, cart.CustomerID, v.ProductID)
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

	// products, err := GetCartProducts(langShortName, cart.CustomerID)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"status":  false,
	// 		"message": err.Error(),
	// 	})
	// 	return
	// }

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "products have been successfully added to cart",
	})

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

	// customerID := c.Param("customer_id")

	rowCustomer, err := db.Query("SELECT id FROM customers WHERE id = $1 AND deleted_at IS NULL", customerID)
	if err != nil {
		return []ProductOfCart{}, err
	}
	defer rowCustomer.Close()

	var customer_id string

	for rowCustomer.Next() {
		if err := rowCustomer.Scan(&customer_id); err != nil {
			return []ProductOfCart{}, err
		}
	}

	if customer_id == "" {
		return []ProductOfCart{}, errors.New("customer not found")
	}

	rowsProduct, err := db.Query("SELECT p.id,p.brend_id,p.price,p.old_price,p.amount,p.product_code,p.limit_amount,p.is_new,c.quantity_of_product FROM products p LEFT JOIN cart c ON c.product_id = p.id WHERE c.customer_id = $1 AND c.deleted_at IS NULL AND p.deleted_at IS NULL", customerID)
	if err != nil {
		return []ProductOfCart{}, err
	}
	defer rowsProduct.Close()

	var products []ProductOfCart

	for rowsProduct.Next() {
		var product ProductOfCart

		if err := rowsProduct.Scan(&product.ID, &product.BrendID, &product.Price, &product.OldPrice, &product.Amount, &product.ProductCode, &product.LimitAmount, &product.IsNew, &product.QuantityOfProduct); err != nil {
			return []ProductOfCart{}, err
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

func RemoveCart(c *gin.Context) {

	db, err := config.ConnDB()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer db.Close()

	customerID := c.Query("customer_id")
	productID := c.Query("product_id")

	rowCustomer, err := db.Query("SELECT id FROM customers WHERE id = $1 AND deleted_at IS NULL", customerID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer rowCustomer.Close()

	var customer_id string

	for rowCustomer.Next() {
		if err := rowCustomer.Scan(&customer_id); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}

	if customer_id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "customer not found",
		})
		return
	}

	if productID != "" {

		rowCart, err := db.Query("SELECT product_id FROM cart WHERE customer_id = $1 AND product_id = $2 AND deleted_at IS NULL", customerID, productID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
		defer rowCart.Close()

		var product_id string

		for rowCart.Next() {
			if err := rowCart.Scan(&product_id); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  false,
					"message": err.Error(),
				})
				return
			}
		}

		if product_id == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": "this product not found in this customer",
			})
			return
		}

		resultCart, err := db.Query("DELETE FROM cart WHERE customer_id = $1 AND product_id = $2 AND deleted_at IS NULL", customerID, productID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
		defer resultCart.Close()

	} else {

		resultCart, err := db.Query("DELETE FROM cart WHERE customer_id = $1 AND deleted_at IS NULL", customerID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
		defer resultCart.Close()

	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "product successfull deleted from cart",
	})

}
