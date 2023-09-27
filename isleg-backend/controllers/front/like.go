package controllers

import (
	"context"
	"github/abbgo/isleg/isleg-backend/config"
	"github/abbgo/isleg/isleg-backend/helpers"
	"github/abbgo/isleg/isleg-backend/models"
	"github/abbgo/isleg/isleg-backend/pkg"
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/lib/pq"
	"gopkg.in/guregu/null.v4"
)

type LikeProduct struct {
	ID           uuid.UUID                              `json:"id"`
	BrendID      uuid.NullUUID                          `json:"brend_id,omitempty"`
	Price        float64                                `json:"price"`
	OldPrice     float64                                `json:"old_price"`
	Percentage   float64                                `json:"percentage"`
	Amount       uint                                   `json:"amount"`
	LimitAmount  uint                                   `json:"limit_amount"`
	IsNew        bool                                   `json:"is_new"`
	Benefit      null.Float                             `json:"-"`
	MainImage    string                                 `json:"main_image"`
	Translations []map[string]models.TranslationProduct `json:"translations"`
	Code         null.String                            `json:"code"`
}

type ProductID struct {
	IDS []string `json:"product_ids"`
}

func AddOrRemoveLike(c *gin.Context) {
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	// bu yerde api query - den status ady bilen status alynyar
	// eger status = true bolsa onda halanlarym sahypa haryt gosulyar
	// eger status = false bolsa onda halanlarym sahypadan haryt ayrylyar
	statusStr := c.Query("status")
	status, err := strconv.ParseBool(statusStr)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	customerID, err := pkg.ValidateMiddlewareData(c, "customer_id")
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	// front - dan gelen maglumatlar bind edilyar
	var productIds ProductID
	if err := c.BindJSON(&productIds); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	if status { // eger status = true bolsa halanlarym sahyp haryt gosulyar
		if len(productIds.IDS) != 0 { // eger front - dan gelen id bar bolsa onda halanlarym sahypa haryt gosulyar
			for _, v := range productIds.IDS {
				// front - dan gelen id - lere den gelyan bazada haryt barmy yokmy sol barlanyar
				var product_id string
				db.QueryRow(context.Background(), "SELECT id FROM products WHERE id = $1 AND deleted_at IS NULL", v).Scan(&product_id)

				if product_id != "" { // eger haryt products tablida bar bolsa onda sol haryt on gelen musderinin
					// halanlarynyn arasynda barmy ya-da yokmy sol barlanyar
					var product string
					db.QueryRow(context.Background(), "SELECT product_id FROM likes WHERE customer_id = $1 AND product_id = $2 AND deleted_at IS NULL", customerID, v).Scan(&product)

					if product == "" { // eger haryt musderinin halanlarym harytlarynyn arasynda yok bolsa
						// onda haryt sol musderinin halanlarym tablisasyna gosulyar
						_, err := db.Exec(context.Background(), "INSERT INTO likes (customer_id,product_id) VALUES ($1,$2)", customerID, v)
						if err != nil {
							helpers.HandleError(c, 400, err.Error())
							return
						}
					}
				}
			}

			// front - dan gelen harytlar halanlarym sahypa gosulandan son
			// yzyna sol harytlar ddoly maglumatlary bilen berilyar
			products, err := GetLikes(customerID)
			if err != nil {
				helpers.HandleError(c, 400, err.Error())
				return
			}

			c.JSON(http.StatusOK, gin.H{
				"status":   true,
				"products": products,
			})
		} else { // eger front hic hile id gelmese onda musderinin onki bazadaky halan harytlaryny fronta bermeli
			products, err := GetLikes(customerID)
			if err != nil {
				helpers.HandleError(c, 400, err.Error())
				return
			}

			if len(products) == 0 {
				c.JSON(http.StatusOK, gin.H{
					"status":  true,
					"message": "like empty",
				})
			} else {
				c.JSON(http.StatusOK, gin.H{
					"status":   true,
					"products": products,
				})
			}
		}

	} else { // eger status = false gelse onda front - dan gele id - li harydy sol musderinin halanlarym harytlaryndan pozmaly
		if len(productIds.IDS) != 0 { // front - dan maglumat gelyarmi ya-da gelenokmy sony barlayas
			// front - dan gelen id - ler halanlarym tablisada barmy ya-da yokmy sony barlayas
			var product_id string
			db.QueryRow(context.Background(), "SELECT product_id FROM likes WHERE customer_id = $1 AND product_id = $2 AND deleted_at IS NULL", customerID, productIds.IDS[0]).Scan(&product_id)

			// eger haryt halanlarym tablisada yok bolsa
			// yzyna yalnyslyk goyberyas
			if product_id == "" {
				helpers.HandleError(c, 404, "this product not found in this customer")
				return
			}

			// haryt halanlarym tablisada bar bolsa onda ony pozyas
			_, err = db.Exec(context.Background(), "DELETE FROM likes WHERE customer_id = $1 AND product_id = $2 AND deleted_at IS NULL", customerID, productIds.IDS[0])
			if err != nil {
				helpers.HandleError(c, 400, err.Error())
				return
			}

			c.JSON(http.StatusOK, gin.H{
				"status":  true,
				"message": "like successfull deleted",
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": "product id is required",
			})
		}
	}
}

// func RemoveLike(c *gin.Context) {

// 	db, err := config.ConnDB()
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"status":  false,
// 			"message": err.Error(),
// 		})
// 		return
// 	}
// 	defer func() {
// 		if err := db.Close(); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{
// 				"status":  false,
// 				"message": err.Error(),
// 			})
// 			return
// 		}
// 	}()

// 	custID, hasCustomer := c.Get("customer_id")
// 	if !hasCustomer {
// 		c.JSON(http.StatusBadRequest, "customer_id is required")
// 		return
// 	}
// 	customerID, ok := custID.(string)
// 	if !ok {
// 		c.JSON(http.StatusBadRequest, "customer_id must be string")
// 	}

// 	productID := c.PostForm("product_id") // Su parametri nireden almaly ( postForm , query , parameter - mi )

// 	rowCustomer, err := db.Query(context.Background(),"SELECT id FROM customers WHERE id = $1 AND deleted_at IS NULL", customerID)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"status":  false,
// 			"message": err.Error(),
// 		})
// 		return
// 	}
// 	defer func() {
// 		if err := rowCustomer.Close(); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{
// 				"status":  false,
// 				"message": err.Error(),
// 			})
// 			return
// 		}
// 	}()

// 	var customer_id string

// 	for rowCustomer.Next() {
// 		if err := rowCustomer.Scan(&customer_id); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{
// 				"status":  false,
// 				"message": err.Error(),
// 			})
// 			return
// 		}
// 	}

// 	if customer_id == "" {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"status":  false,
// 			"message": "customer not found",
// 		})
// 		return
// 	}

// 	rowLike, err := db.Query(context.Background(),"SELECT product_id FROM likes WHERE customer_id = $1 AND product_id = $2 AND deleted_at IS NULL", customerID, productID)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"status":  false,
// 			"message": err.Error(),
// 		})
// 		return
// 	}
// 	defer func() {
// 		if err := rowLike.Close(); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{
// 				"status":  false,
// 				"message": err.Error(),
// 			})
// 			return
// 		}
// 	}()

// 	var product_id string

// 	for rowLike.Next() {
// 		if err := rowLike.Scan(&product_id); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{
// 				"status":  false,
// 				"message": err.Error(),
// 			})
// 			return
// 		}
// 	}

// 	if product_id == "" {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"status":  false,
// 			"message": "this product not found in this customer",
// 		})
// 		return
// 	}

// 	resultLike, err := db.Query(context.Background(),"DELETE FROM likes WHERE customer_id = $1 AND product_id = $2 AND deleted_at IS NULL", customerID, productID)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"status":  false,
// 			"message": err.Error(),
// 		})
// 		return
// 	}
// 	defer func() {
// 		if err := resultLike.Close(); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{
// 				"status":  false,
// 				"message": err.Error(),
// 			})
// 			return
// 		}
// 	}()

// 	c.JSON(http.StatusOK, gin.H{
// 		"status":  true,
// 		"message": "like successfull deleted",
// 	})

// }

func GetLikes(customerID string) ([]LikeProduct, error) {
	db, err := config.ConnDB()
	if err != nil {
		return []LikeProduct{}, err
	}
	defer db.Close()

	rowsProduct, err := db.Query(context.Background(), "SELECT p.id,p.brend_id,p.price,p.old_price,p.amount,p.limit_amount,p.is_new,p.main_image,p.benefit FROM products p LEFT JOIN likes l ON l.product_id = p.id WHERE l.customer_id = $1 AND p.amount > 0 AND p.limit_amount > 0 AND l.deleted_at IS NULL AND p.deleted_at IS NULL", customerID)
	if err != nil {
		return []LikeProduct{}, err
	}
	defer rowsProduct.Close()

	var products []LikeProduct
	for rowsProduct.Next() {
		var product LikeProduct
		if err := rowsProduct.Scan(&product.ID, &product.BrendID, &product.Price, &product.OldPrice, &product.Amount, &product.LimitAmount, &product.IsNew, &product.MainImage, &product.Benefit); err != nil {
			return []LikeProduct{}, err
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

		rowsLang, err := db.Query(context.Background(), "SELECT id,name_short FROM languages WHERE deleted_at IS NULL")
		if err != nil {
			return []LikeProduct{}, err
		}
		defer rowsLang.Close()

		var languages []models.Language
		for rowsLang.Next() {
			var language models.Language
			if err := rowsLang.Scan(&language.ID, &language.NameShort); err != nil {
				return []LikeProduct{}, err
			}
			languages = append(languages, language)
		}

		for _, v := range languages {
			var trProduct models.TranslationProduct
			translation := make(map[string]models.TranslationProduct)
			db.QueryRow(context.Background(), "SELECT name,description FROM translation_product WHERE lang_id = $1 AND product_id = $2", v.ID, product.ID).Scan(&trProduct.Name, &trProduct.Description)

			translation[v.NameShort] = trProduct
			product.Translations = append(product.Translations, translation)
		}
		products = append(products, product)
	}

	return products, nil
}

func GetCustomerLikes(c *gin.Context) {
	customerID, err := pkg.ValidateMiddlewareData(c, "customer_id")
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	products, err := GetLikes(customerID)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":   true,
		"products": products,
	})
}

// func GetLikedProductsWithoutCustomer(c *gin.Context) {

// 	// databaza konnektion acylyar
// 	db, err := config.ConnDB()
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"status":  false,
// 			"message": err.Error(),
// 		})
// 		return
// 	}

// 	// in sonunda databaza konnektion yapylyar
// 	defer func() {
// 		if err := db.Close(); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{
// 				"status":  false,
// 				"message": err.Error(),
// 			})
// 			return
// 		}
// 	}()

// 	// front - dan gelen maglumatlar bind edilyar
// 	var productIds ProductID
// 	if err := c.BindJSON(&productIds); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"status":  false,
// 			"message": err.Error(),
// 		})
// 		return
// 	}

// 	// front - dan gelen id - ler prodcuts tablisada barmy ya-da yokmy sol barlanyar
// 	for _, v := range productIds.IDS {
// 		rowProduct, err := db.Query(context.Background(),"SELECT id FROM products WHERE id = $1 AND deleted_at IS NULL", v)
// 		if err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{
// 				"status":  false,
// 				"message": err.Error(),
// 			})
// 			return
// 		}
// 		defer func() {
// 			if err := rowProduct.Close(); err != nil {
// 				c.JSON(http.StatusBadRequest, gin.H{
// 					"status":  false,
// 					"message": err.Error(),
// 				})
// 				return
// 			}
// 		}()

// 		var product_id string

// 		for rowProduct.Next() {
// 			if err := rowProduct.Scan(&product_id); err != nil {
// 				c.JSON(http.StatusBadRequest, gin.H{
// 					"status":  false,
// 					"message": err.Error(),
// 				})
// 				return
// 			}
// 		}

// 		// eger id products tablisada yok bolsa , onda yzyna yalnyslyk ugradylyar
// 		if product_id == "" {
// 			c.JSON(http.StatusBadRequest, gin.H{
// 				"status":  false,
// 				"message": "product not found",
// 			})
// 			return
// 		}
// 	}

// 	// front - dan gelen id - ler boyunca id - si gelen id den bolan harytlar yzyna ugradylyar
// 	rowLikes, err := db.Query(context.Background(),"SELECT id,brend_id,price,old_price,amount,limit_amount,is_new FROM products WHERE id = ANY($1) AND deleted_at IS NULL", pq.Array(productIds.IDS))
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"status":  false,
// 			"message": err.Error(),
// 		})
// 		return
// 	}
// 	defer func() {
// 		if err := rowLikes.Close(); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{
// 				"status":  false,
// 				"message": err.Error(),
// 			})
// 			return
// 		}
// 	}()

// 	var products []LikeProduct

// 	for rowLikes.Next() {
// 		var product LikeProduct

// 		if err := rowLikes.Scan(&product.ID, &product.BrendID, &product.Price, &product.OldPrice, &product.Amount, &product.LimitAmount, &product.IsNew); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{
// 				"status":  false,
// 				"message": err.Error(),
// 			})
// 			return
// 		}

// 		if product.OldPrice != 0 {
// 			product.Percentage = -math.Round(((product.OldPrice - product.Price) * 100) / product.OldPrice)
// 		} else {
// 			product.Percentage = 0
// 		}

// 		rowMainImage, err := db.Query(context.Background(),"SELECT small,medium,large FROM main_image WHERE product_id = $1 AND deleted_at IS NULL", product.ID)
// 		if err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{
// 				"status":  false,
// 				"message": err.Error(),
// 			})
// 			return
// 		}
// 		defer func() {
// 			if err := rowMainImage.Close(); err != nil {
// 				c.JSON(http.StatusBadRequest, gin.H{
// 					"status":  false,
// 					"message": err.Error(),
// 				})
// 				return
// 			}
// 		}()

// 		var mainImage models.MainImage

// 		for rowMainImage.Next() {
// 			if err := rowMainImage.Scan(&mainImage.Small, &mainImage.Medium, &mainImage.Large); err != nil {
// 				c.JSON(http.StatusBadRequest, gin.H{
// 					"status":  false,
// 					"message": err.Error(),
// 				})
// 				return
// 			}
// 		}

// 		product.MainImage = mainImage

// 		rowsImages, err := db.Query(context.Background(),"SELECT small,large FROM images WHERE product_id = $1 AND deleted_at IS NULL", product.ID)
// 		if err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{
// 				"status":  false,
// 				"message": err.Error(),
// 			})
// 			return
// 		}
// 		defer func() {
// 			if err := rowsImages.Close(); err != nil {
// 				c.JSON(http.StatusBadRequest, gin.H{
// 					"status":  false,
// 					"message": err.Error(),
// 				})
// 				return
// 			}
// 		}()

// 		var images []models.Images

// 		for rowsImages.Next() {
// 			var image models.Images

// 			if err := rowsImages.Scan(&image.Small, &image.Large); err != nil {
// 				c.JSON(http.StatusBadRequest, gin.H{
// 					"status":  false,
// 					"message": err.Error(),
// 				})
// 				return
// 			}

// 			images = append(images, image)
// 		}

// 		product.Images = images

// 		rowsLang, err := db.Query(context.Background(),"SELECT id,name_short FROM languages WHERE deleted_at IS NULL")
// 		if err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{
// 				"status":  false,
// 				"message": err.Error(),
// 			})
// 			return
// 		}
// 		defer func() {
// 			if err := rowsLang.Close(); err != nil {
// 				c.JSON(http.StatusBadRequest, gin.H{
// 					"status":  false,
// 					"message": err.Error(),
// 				})
// 				return
// 			}
// 		}()

// 		var languages []models.Language

// 		for rowsLang.Next() {
// 			var language models.Language

// 			if err := rowsLang.Scan(&language.ID, &language.NameShort); err != nil {
// 				c.JSON(http.StatusBadRequest, gin.H{
// 					"status":  false,
// 					"message": err.Error(),
// 				})
// 				return
// 			}

// 			languages = append(languages, language)
// 		}

// 		for _, v := range languages {

// 			rowTrProduct, err := db.Query(context.Background(),"SELECT name,description FROM translation_product WHERE product_id = $1 AND lang_id = $2 AND deleted_at IS NULL", product.ID, v.ID)
// 			if err != nil {
// 				c.JSON(http.StatusBadRequest, gin.H{
// 					"status":  false,
// 					"message": err.Error(),
// 				})
// 				return
// 			}
// 			defer func() {
// 				if err := rowTrProduct.Close(); err != nil {
// 					c.JSON(http.StatusBadRequest, gin.H{
// 						"status":  false,
// 						"message": err.Error(),
// 					})
// 					return
// 				}
// 			}()

// 			var trProduct models.TranslationProduct

// 			translation := make(map[string]models.TranslationProduct)

// 			for rowTrProduct.Next() {
// 				if err := rowTrProduct.Scan(&trProduct.Name, &trProduct.Description); err != nil {
// 					c.JSON(http.StatusBadRequest, gin.H{
// 						"status":  false,
// 						"message": err.Error(),
// 					})
// 					return
// 				}
// 			}

// 			translation[v.NameShort] = trProduct

// 			product.Translations = append(product.Translations, translation)

// 		}

// 		products = append(products, product)
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"status":   true,
// 		"products": products,
// 	})

// }

func GetLikedOrOrderedProductsWithoutCustomer(c *gin.Context) {
	// databaza konnektion acylyar
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	// in sonunda databaza konnektion yapylyar
	defer db.Close()

	// front - dan gelen maglumatlar bind edilyar
	var productIds ProductID
	if err := c.BindJSON(&productIds); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	// front - dan gelen id - ler boyunca id - si gelen id den bolan harytlar yzyna ugradylyar
	rowLikes, err := db.Query(context.Background(), "SELECT id,brend_id,price,old_price,amount,limit_amount,is_new,main_image,benefit,code FROM products WHERE id = ANY($1) AND amount > 0 AND limit_amount > 0 AND deleted_at IS NULL", pq.Array(productIds.IDS))
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer rowLikes.Close()

	var products []LikeProduct
	for rowLikes.Next() {
		var product LikeProduct
		if err := rowLikes.Scan(&product.ID, &product.BrendID, &product.Price, &product.OldPrice, &product.Amount, &product.LimitAmount, &product.IsNew, &product.MainImage, &product.Benefit, &product.Code); err != nil {
			helpers.HandleError(c, 400, err.Error())
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

		rowsLang, err := db.Query(context.Background(), "SELECT id,name_short FROM languages WHERE deleted_at IS NULL")
		if err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}
		defer rowsLang.Close()

		var languages []models.Language
		for rowsLang.Next() {
			var language models.Language
			if err := rowsLang.Scan(&language.ID, &language.NameShort); err != nil {
				helpers.HandleError(c, 400, err.Error())
				return
			}
			languages = append(languages, language)
		}

		for _, v := range languages {
			var trProduct models.TranslationProduct
			translation := make(map[string]models.TranslationProduct)
			db.QueryRow(context.Background(), "SELECT name,description FROM translation_product WHERE product_id = $1 AND lang_id = $2 AND deleted_at IS NULL", product.ID, v.ID).Scan(&trProduct.Name, &trProduct.Description)
			translation[v.NameShort] = trProduct
			product.Translations = append(product.Translations, translation)
		}
		products = append(products, product)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":   true,
		"products": products,
	})
}
