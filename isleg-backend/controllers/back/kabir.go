package controllers

import (
	"github/abbgo/isleg/isleg-backend/config"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

type OneProduct struct {
	BrendID     uuid.UUID          `json:"brend_id"`
	Price       float64            `json:"price"`
	OldPrice    float64            `json:"old_price"`
	Amount      uint               `json:"amount"`
	ProductCode string             `json:"product_code"`
	MainImage   string             `json:"main_image"`
	Images      pq.StringArray     `json:"images"`
	Categories  []string           `json:"categories"`
	Translation TranslationProduct `json:"translation"`
}

type TranslationProduct struct {
	LanguageID  string `json:"lang_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func GetProductByID(c *gin.Context) {

	ID := c.Param("id")

	rowProduct, err := config.ConnDB().Query("SELECT brend_id,price,old_price,amount,product_code,main_image,images FROM products WHERE id = $1 AND deleted_at IS NULL", ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	var product OneProduct

	for rowProduct.Next() {
		if err := rowProduct.Scan(&product.BrendID, &product.Price, &product.OldPrice, &product.Amount, &product.ProductCode, &product.MainImage, &product.Images); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}

	if product.MainImage == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "record not found",
		})
		return
	}

	rowsCategoryProduct, err := config.ConnDB().Query("SELECT category_id FROM category_product WHERE product_id = $1 AND deleted_at IS NULL", ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

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

	rowTranslationProduct, err := config.ConnDB().Query("SELECT lang_id,name,description FROM translation_product WHERE product_id = $1 AND deleted_at IS NULL", ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	var translation TranslationProduct

	for rowTranslationProduct.Next() {
		if err := rowTranslationProduct.Scan(&translation.LanguageID, &translation.Name, &translation.Description); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}

	if translation.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "record not found",
		})
		return
	}

	product.Translation = translation

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"product": product,
	})

}
