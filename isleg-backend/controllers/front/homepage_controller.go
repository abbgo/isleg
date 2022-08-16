package controllers

import (
	"net/http"

	"github/abbgo/isleg/isleg-backend/config"
	backController "github/abbgo/isleg/isleg-backend/controllers/back"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

type HomePageCategory struct {
	ID       string    `json:"id"`
	Name     string    `json:"name"`
	Products []Product `json:"products"`
}

type Product struct {
	ID          string         `json:"id"`
	Name        string         `json:"name"`
	Price       float64        `json:"price"`
	OldPrice    float64        `json:"old_price"`
	MainImage   string         `json:"main_image"`
	ProductCode string         `json:"product_code"`
	Images      pq.StringArray `json:"images"`
	Brend       Brend          `json:"brend"`
}

type Brend struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

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

	langID, err := backController.CheckLanguage(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	// get all homepage category where translation category id equal langID
	categoryRows, err := config.ConnDB().Query("SELECT c.id,t.name FROM categories c LEFT JOIN translation_category t ON c.id=t.category_id WHERE t.lang_id = $1 AND c.is_home_category = true AND t.deleted_at IS NULL AND c.deleted_at IS NULL", langID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer categoryRows.Close()

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
		productRows, err := config.ConnDB().Query("SELECT p.id,t.name,p.price,p.old_price,p.main_image,p.product_code,p.images FROM products p LEFT JOIN category_product c ON p.id=c.product_id LEFT JOIN translation_product t ON p.id=t.product_id WHERE t.lang_id = $1 AND c.category_id = $2 AND p.deleted_at IS NULL AND c.deleted_at IS NULL AND t.deleted_at IS NULL ORDER BY p.created_at DESC LIMIT 4", langID, homePageCategory.ID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
		defer productRows.Close()

		var products []Product
		for productRows.Next() {
			var product Product
			if err := productRows.Scan(&product.ID, &product.Name, &product.Price, &product.OldPrice, &product.MainImage, &product.ProductCode, &product.Images); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  false,
					"message": err.Error(),
				})
				return
			}

			// get brend where id of product brend_id
			brendRows, err := config.ConnDB().Query("SELECT b.id,b.name FROM products p LEFT JOIN brends b ON p.brend_id=b.id WHERE p.id = $1 AND p.deleted_at IS NULL AND b.deleted_at IS NULL", product.ID)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  false,
					"message": err.Error(),
				})
				return
			}
			defer brendRows.Close()

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
		homePageCategory.Products = products
		homePageCategories = append(homePageCategories, homePageCategory)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":              true,
		"homepage_categories": homePageCategories,
	})

}
