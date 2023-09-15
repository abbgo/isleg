package controllers

import (
	"context"
	"errors"
	"github/abbgo/isleg/isleg-backend/config"
	"github/abbgo/isleg/isleg-backend/helpers"
	"github/abbgo/isleg/isleg-backend/models"
	"math"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gopkg.in/guregu/null.v4"
)

// type DataForAddCart struct {
// 	Products []CartProduct `json:"products"`
// }

type CartProduct struct {
	ProductID         string `json:"product_id" binding:"required"`
	QuantityOfProduct uint   `json:"quantity_of_product"`
}

type ProductOfCart struct {
	ID                uuid.UUID                              `json:"id"`
	BrendID           uuid.NullUUID                          `json:"brend_id,omitempty"`
	Price             float64                                `json:"price"`
	OldPrice          float64                                `json:"old_price"`
	Percentage        float64                                `json:"percentage"`
	Amount            uint                                   `json:"amount"`
	LimitAmount       uint                                   `json:"limit_amount"`
	IsNew             bool                                   `json:"is_new"`
	Benefit           null.Float                             `json:"-"`
	QuantityOfProduct int                                    `json:"quantity_of_product"`
	MainImage         string                                 `json:"main_image"`
	Translations      []map[string]models.TranslationProduct `json:"translations"`
}

func AddCart(c *gin.Context) {
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	custID, hasCustomer := c.Get("customer_id")
	if !hasCustomer {
		helpers.HandleError(c, 400, "customer_id is required")
		return
	}
	customerID, ok := custID.(string)
	if !ok {
		helpers.HandleError(c, 400, "customer_id must be string")
		return
	}

	// frontdan maglumaty bind etyar
	var cart []CartProduct
	if err := c.BindJSON(&cart); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	// frontdan maglumat gelipdirmi gelmandirmi sony barlayas
	if len(cart) != 0 { // eger frontdan maglumatdan gelyan bolsa gelen harytlary sebede gosyas
		for k, v := range cart {
			// eger frontdan gelen harydyn mukdary 1 - den kici bolsa
			// sol musderinin sol harydyny sebetden ayyryas , yagny mukdary nol bolan haryt sebetde durup bilmez
			// sonun ucin eger musderi harydyn sanyny nol etse ony sebetdebn ayyryas
			if v.QuantityOfProduct < 1 {
				if err := DeleteCart(customerID, v.ProductID); err != nil {
					helpers.HandleError(c, 400, err.Error())
					return
				}
				break
			}

			// bu yerde frontdan 1 haryt 2 gezek gaytalanyp gelipdirmi ya-da gelmandirmi
			// sony barlayas. Eger 1 haryt 2 gezek gelen bolsa yzyna osibka yazyp ugartyas
			for _, x := range cart[(k + 1):] {
				if v.ProductID == x.ProductID {
					helpers.HandleError(c, 400, "the same product cannot be repeated twice")
					return
				}
			}

			// bu yerde frontdan gelen haryt bazada barmy ya-da yokmy sol barlanyar
			var product_id string
			db.QueryRow(context.Background(), "SELECT id FROM products WHERE id = $1 AND deleted_at IS NULL", v.ProductID).Scan(&product_id)

			if product_id != "" {
				// eger haryt bazada bar bolsa onda sol haryt programmany ulanyp otyran musderinin sebedinde barmy ya-da yok
				// sony barlayas
				var productId string
				db.QueryRow(context.Background(), "SELECT product_id FROM cart WHERE customer_id = $1 AND product_id = $2 AND deleted_at IS NULL", customerID, v.ProductID).Scan(&productId)

				if productId == "" {
					// eger sol haryt programmany ulanyp otyran musderinin sebedinde yok bolsa
					// sol harydy sol musderinin sebedine gosyas
					_, err := db.Exec(context.Background(), "INSERT INTO cart (customer_id,product_id,quantity_of_product) VALUES ($1,$2,$3)", customerID, v.ProductID, v.QuantityOfProduct)
					if err != nil {
						helpers.HandleError(c, 400, err.Error())
						return
					}
				} else {
					// eger sol haryt programmany ulanyp otyran musderinin sebedinde bar bolsa
					// onda sol harydyn mukdaryny update etyas
					_, err := db.Exec(context.Background(), "UPDATE cart SET quantity_of_product = $1 WHERE customer_id = $2 AND product_id = $3 AND deleted_at IS NULL", v.QuantityOfProduct, customerID, v.ProductID)
					if err != nil {
						helpers.HandleError(c, 400, err.Error())
						return
					}
				}
			}
		}

		// haryt sebede gosulandan sonra programmany ulanyp otyran musdera degisli
		// harytlary yzyna gaytaryp beryas
		products, err := GetCartProducts(customerID)
		if err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status":   true,
			"products": products,
		})
	} else {
		// eger frontdan maglumat gelmese programmany ulanyp otyran musderinin
		// sebedindaki harytlary yzyna ugratyas
		products, err := GetCartProducts(customerID)
		if err != nil {
			helpers.HandleError(c, 400, err.Error())
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

// GetCartProducts funksiya bazadan belli bir musdera degisli sebetdaki harytlary alyp beryar
func GetCartProducts(customerID string) ([]ProductOfCart, error) {
	db, err := config.ConnDB()
	if err != nil {
		return []ProductOfCart{}, err
	}
	defer db.Close()

	rowsProduct, err := db.Query(context.Background(), "SELECT p.id,p.brend_id,p.price,p.old_price,p.amount,p.limit_amount,p.is_new,c.quantity_of_product,p.main_image,p.benefit FROM products p LEFT JOIN cart c ON c.product_id = p.id WHERE c.customer_id = $1 AND c.deleted_at IS NULL AND p.deleted_at IS NULL", customerID)
	if err != nil {
		return []ProductOfCart{}, err
	}
	defer rowsProduct.Close()

	var products []ProductOfCart
	for rowsProduct.Next() {
		var product ProductOfCart
		if err := rowsProduct.Scan(&product.ID, &product.BrendID, &product.Price, &product.OldPrice, &product.Amount, &product.LimitAmount, &product.IsNew, &product.QuantityOfProduct, &product.MainImage, &product.Benefit); err != nil {
			return []ProductOfCart{}, err
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
			return []ProductOfCart{}, err
		}
		defer rowsLang.Close()

		var languages []models.Language
		for rowsLang.Next() {
			var language models.Language
			if err := rowsLang.Scan(&language.ID, &language.NameShort); err != nil {
				return []ProductOfCart{}, err
			}
			languages = append(languages, language)
		}

		for _, v := range languages {
			var trProduct models.TranslationProduct
			translation := make(map[string]models.TranslationProduct)
			db.QueryRow(context.Background(), "SELECT name,description FROM translation_product WHERE lang_id = $1 AND product_id = $2 AND deleted_at IS NULL", v.ID, product.ID).Scan(&trProduct.Name, &trProduct.Description)
			if err != nil {
				return []ProductOfCart{}, err
			}

			translation[v.NameShort] = trProduct
			product.Translations = append(product.Translations, translation)
		}
		products = append(products, product)
	}

	return products, nil
}

func GetCustomerCartProducts(c *gin.Context) {
	custID, hasCustomer := c.Get("customer_id")
	if !hasCustomer {
		helpers.HandleError(c, 400, "customer_id is required")
		return
	}
	customerID, ok := custID.(string)
	if !ok {
		helpers.HandleError(c, 400, "customer_id must be string")
		return
	}

	products, err := GetCartProducts(customerID)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
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
		helpers.HandleError(c, 400, "customer_id is required")
		return
	}
	customerID, ok := custID.(string)
	if !ok {
		helpers.HandleError(c, 400, "customer_id must be string")
		return
	}

	// productID := c.PostForm("product_id")

	if err := DeleteCart(customerID, ""); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "product successfull deleted from cart",
	})
}

// DeleteCart funksiya muderinin sebedinden haryt pozmak ucin ulanylyar
func DeleteCart(customerID, productID string) error {
	db, err := config.ConnDB()
	if err != nil {
		return err
	}
	defer db.Close()

	if productID != "" {
		var product_id string
		db.QueryRow(context.Background(), "SELECT product_id FROM cart WHERE customer_id = $1 AND product_id = $2 AND deleted_at IS NULL", customerID, productID).Scan(&product_id)

		if product_id == "" {
			return errors.New("this product not found in this customer")
		}

		_, err = db.Exec(context.Background(), "DELETE FROM cart WHERE customer_id = $1 AND product_id = $2 AND deleted_at IS NULL", customerID, productID)
		if err != nil {
			return err
		}
	} else {
		_, err := db.Exec(context.Background(), "DELETE FROM cart WHERE customer_id = $1 AND deleted_at IS NULL", customerID)
		if err != nil {
			return err
		}
	}
	return nil
}
