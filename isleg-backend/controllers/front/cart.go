package controllers

import (
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
	ProductID         string `json:"product_id"`
	QuantityOfProduct int    `json:"quantity_of_product"`
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
	defer db.Close()

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
		defer rowProduct.Close()

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
		defer rowCart.Close()

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
			defer resultCartInsert.Close()

		} else {

			reultCartUpdate, err := db.Query("UPDATE cart SET quantity_of_product = $1 WHERE customer_id = $2 AND product_id = $3 AND deleted_at IS NULL", v.QuantityOfProduct, cart.CustomerID, v.ProductID)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  false,
					"message": err.Error(),
				})
				return
			}
			defer reultCartUpdate.Close()

		}

	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "product added successfull to cart",
	})

}

func GetCartProducts(c *gin.Context) {

	db, err := config.ConnDB()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer db.Close()

	// GET DATA FROM ROUTE PARAMETER
	langShortName := c.Param("lang")

	// GET language id
	langID, err := backController.GetLangID(langShortName)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	customerID := c.Param("customer_id")

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

	rowsProduct, err := db.Query("SELECT p.id,p.brend_id,p.price,p.old_price,p.amount,p.product_code,p.limit_amount,p.is_new,c.quantity_of_product FROM products p LEFT JOIN cart c ON c.product_id = p.id WHERE c.customer_id = $1 AND c.deleted_at IS NULL AND p.deleted_at IS NULL", customerID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer rowsProduct.Close()

	var products []ProductOfCart

	for rowsProduct.Next() {
		var product ProductOfCart

		if err := rowsProduct.Scan(&product.ID, &product.BrendID, &product.Price, &product.OldPrice, &product.Amount, &product.ProductCode, &product.LimitAmount, &product.IsNew, &product.QuantityOfProduct); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}

		rowMainImage, err := db.Query("SELECT small,medium,large FROM main_image WHERE product_id = $1 AND deleted_at IS NULL", product.ID)
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

		rowsImages, err := db.Query("SELECT small,medium,large FROM images WHERE product_id = $1 AND deleted_at IS NULL", product.ID)
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

			if err := rowsImages.Scan(&image.Small, &image.Medium, &image.Large); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  false,
					"message": err.Error(),
				})
				return
			}

			images = append(images, image)
		}

		product.Images = images

		rowTrProduct, err := db.Query("SELECT name,description FROM translation_product WHERE lang_id = $1 AND product_id = $2", langID, product.ID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
		defer rowTrProduct.Close()

		var trProduct models.TranslationProduct

		for rowTrProduct.Next() {
			if err := rowTrProduct.Scan(&trProduct.Name, &trProduct.Description); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  false,
					"message": err.Error(),
				})
				return
			}
		}

		product.TranslationProduct = trProduct

		products = append(products, product)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":   true,
		"products": products,
	})

}

// func AddProductToBasket(c *gin.Context) {

// 	db, err := config.ConnDB()
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"status":  false,
// 			"message": err.Error(),
// 		})
// 		return
// 	}
// 	defer db.Close()

// 	customerID := c.PostForm("customer_id")
// 	productID := c.PostForm("product_id")
// 	quantityOfProduct := c.PostForm("quantity_of_product")

// 	quantityInt, err := models.ValidateCustomerBasket(customerID, productID, quantityOfProduct)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"status":  false,
// 			"message": err.Error(),
// 		})
// 		return
// 	}

// 	rowBasket, err := db.Query("SELECT COUNT(*) FROM basket WHERE product_id = $1 AND customer_id = $2 AND deleted_at IS NULL", productID, customerID)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"status":  false,
// 			"message": err.Error(),
// 		})
// 		return
// 	}

// 	var quantity int

// 	for rowBasket.Next() {
// 		if err := rowBasket.Scan(&quantity); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{
// 				"status":  false,
// 				"message": err.Error(),
// 			})
// 			return
// 		}
// 	}

// 	if quantity == 0 {
// 		resultBasket, err := db.Query("INSERT INTO basket (product_id,customer_id,quantity_of_product) VALUES ($1,$2,$3)", productID, customerID, quantityInt)
// 		if err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{
// 				"status":  false,
// 				"message": err.Error(),
// 			})
// 			return
// 		}
// 		defer resultBasket.Close()
// 	} else {
// 		resultBasket, err := db.Query("UPDATE basket SET quantity_of_product = $1 WHERE product_id = $2 AND customer_id = $3", quantityInt, productID, customerID)
// 		if err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{
// 				"status":  false,
// 				"message": err.Error(),
// 			})
// 			return
// 		}
// 		defer resultBasket.Close()
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"status":  true,
// 		"message": "product successfully added to card",
// 	})

// }
