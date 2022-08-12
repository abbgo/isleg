package controllers

import (
	"github/abbgo/isleg/isleg-backend/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetProducts(c *gin.Context) {

	rowsProduct, err := config.ConnDB().Query("SELECT id,brend_id,price,old_price,amount,product_code,main_image,images FROM products WHERE deleted_at IS NULL")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	var products []OneProduct
	var ids []string
	var categories []string
	var translation TranslationProduct

	for rowsProduct.Next() {
		var product OneProduct
		var id string

		if err := rowsProduct.Scan(&id, &product.BrendID, &product.Price, &product.OldPrice, &product.Amount, &product.ProductCode, &product.MainImage, &product.Images); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}

		ids = append(ids, id)

		for _, ID := range ids {
			rowsCategoryProduct, err := config.ConnDB().Query("SELECT category_id FROM category_product WHERE product_id = $1 AND deleted_at IS NULL", ID)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  false,
					"message": err.Error(),
				})
				return
			}

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
		}

		product.Categories = categories

		for _, ID := range ids {
			rowTranslationProduct, err := config.ConnDB().Query("SELECT lang_id,name,description FROM translation_product WHERE product_id = $1 AND deleted_at IS NULL", ID)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  false,
					"message": err.Error(),
				})
				return
			}

			for rowTranslationProduct.Next() {
				if err := rowTranslationProduct.Scan(&translation.LanguageID, &translation.Name, &translation.Description); err != nil {
					c.JSON(http.StatusBadRequest, gin.H{
						"status":  false,
						"message": err.Error(),
					})
					return
				}
			}
		}

		product.Translation = translation

		products = append(products, product)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":   true,
		"products": products,
	})

}
