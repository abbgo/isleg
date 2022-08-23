package controllers

import (
	"github/abbgo/isleg/isleg-backend/config"
	controllers "github/abbgo/isleg/isleg-backend/controllers/back"
	"github/abbgo/isleg/isleg-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ForLikeCustomer struct {
	ID          string                 `json:"id"`
	BrendID     uuid.UUID              `json:"brend_id"`
	Price       float64                `json:"price"`
	OldPrice    float64                `json:"old_price"`
	MainImage   string                 `json:"main_image"`
	Translation TranslationLikeProduct `json:"translations"`
}

type TranslationLikeProduct struct {
	Name        string `json:"name"`
	Description string `json:"description"`
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
	productID := c.PostForm("product_id")

	err = models.ValidateCustomerLike(customerID, productID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	resultLike, err := db.Query("INSERT INTO likes (product_id,customer_id) VALUES ($1,$2)", productID, customerID)
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
		"message": "like successfully added",
	})

}

func GetLikes(c *gin.Context) {

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

	rowsLikes, err := db.Query("SELECT p.id,p.brend_id,p.price,p.old_price,p.main_image,t.name,t.description FROM products p LEFT JOIN likes l ON l.product_id = p.id LEFT JOIN translation_product t ON t.product_id = p.id WHERE l.customer_id = $1 AND t.lang_id = $2 AND p.deleted_at IS NULL AND l.deleted_at IS NULL AND t.deleted_at IS NULL", customerID, langID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	var likes []ForLikeCustomer

	for rowsLikes.Next() {
		var like ForLikeCustomer
		if err := rowsLikes.Scan(&like.ID, &like.BrendID, &like.Price, &like.OldPrice, &like.MainImage, &like.Translation.Name, &like.Translation.Description); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
		likes = append(likes, like)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":         true,
		"customer_likes": likes,
	})

}
