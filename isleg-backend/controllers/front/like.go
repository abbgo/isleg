package controllers

import (
	"github/abbgo/isleg/isleg-backend/config"
	controllers "github/abbgo/isleg/isleg-backend/controllers/back"
	"github/abbgo/isleg/isleg-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ForLikeCustomer struct {
	ID          string                 `json:"id"`
	Brend       OneBrend               `json:"brend"`
	Price       float64                `json:"price"`
	OldPrice    float64                `json:"old_price"`
	MainImage   string                 `json:"main_image"`
	Translation TranslationLikeProduct `json:"translations"`
	LimitAmount uint                   `json:"limit_amount"`
}

type TranslationLikeProduct struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type OneBrend struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func AddLike(c *gin.Context) {

	db, err := config.ConnDB()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer db.Close()

	customerID := c.PostForm("customer_id")
	productIds, _ := c.GetPostFormArray("products")

	err = models.ValidateCustomerLike(customerID, productIds)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	for _, productID := range productIds {
		resultLike, err := db.Query("INSERT INTO likes (product_id,customer_id) VALUES ($1,$2)", productID, customerID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
		defer resultLike.Close()
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "like successfully added",
	})

}

func GetCustomerLikes(c *gin.Context) {

	db, err := config.ConnDB()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer db.Close()

	customerID := c.Param("customer_id")
	lang := c.Param("lang")

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

	langID, err := controllers.GetLangID(lang)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	rowsLikes, err := db.Query("SELECT p.id,p.price,p.old_price,p.main_image,p.limit_amount,t.name,t.description FROM products p LEFT JOIN likes l ON l.product_id = p.id LEFT JOIN translation_product t ON t.product_id = p.id WHERE l.customer_id = $1 AND t.lang_id = $2 AND p.deleted_at IS NULL AND l.deleted_at IS NULL AND t.deleted_at IS NULL", customerID, langID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer rowsLikes.Close()

	var likes []ForLikeCustomer

	for rowsLikes.Next() {
		var like ForLikeCustomer
		if err := rowsLikes.Scan(&like.ID, &like.Price, &like.OldPrice, &like.MainImage, &like.LimitAmount, &like.Translation.Name, &like.Translation.Description); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}

		// get brend where id equal brend_id of product
		brendRows, err := db.Query("SELECT b.id,b.name FROM products p LEFT JOIN brends b ON p.brend_id=b.id WHERE p.id = $1 AND p.deleted_at IS NULL AND b.deleted_at IS NULL", like.ID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
		defer brendRows.Close()

		var brend OneBrend

		for brendRows.Next() {
			if err := brendRows.Scan(&brend.ID, &brend.Name); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  false,
					"message": err.Error(),
				})
				return
			}
		}

		like.Brend = brend

		likes = append(likes, like)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":         true,
		"customer_likes": likes,
	})

}

func RemoveLike(c *gin.Context) {

	db, err := config.ConnDB()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer db.Close()

	customerID := c.Param("customer_id")
	productID := c.Param("product_id")

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

	rowProduct, err := db.Query("SELECT id FROM products WHERE id = $1 AND deleted_at IS NULL", productID)
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

	resultLike, err := db.Query("DELETE FROM likes WHERE product_id = $1 AND customer_id = $2", productID, customerID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer resultLike.Close()

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "like successfully removed",
	})

}
