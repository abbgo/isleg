package controllers

import (
	"context"
	"math"
	"net/http"

	"github/abbgo/isleg/isleg-backend/config"
	backController "github/abbgo/isleg/isleg-backend/controllers/back"
	"github/abbgo/isleg/isleg-backend/helpers"
	"github/abbgo/isleg/isleg-backend/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gopkg.in/guregu/null.v4"
)

type HomePageCategory struct {
	ID       string    `json:"id,omitempty"`
	Name     string    `json:"name,omitempty"`
	Products []Product `json:"products,omitempty"`
}

type Product struct {
	ID          string  `json:"id,omitempty"`
	Price       float64 `json:"price,omitempty"`
	OldPrice    float64 `json:"old_price,omitempty"`
	MainImage   string  `json:"main_image,omitempty"`
	Brend       Brend   `json:"brend,omitempty"`
	LimitAmount int     `json:"limit_amount,omitempty"`
	IsNew       bool    `json:"is_new,omitempty"`
	Amount      int     `json:"amount,omitempty"`
	// Translation models.TranslationProduct `json:"translation,omitempty"`
	Translations []map[string]models.TranslationProduct `json:"translations,omitempty"`
	Benefit      null.Float                             `json:"-"`
	Percentage   float64                                `json:"percentage,omitempty"`
}

type Brend struct {
	ID   uuid.NullUUID `json:"id,omitempty"`
	Name null.String   `json:"name,omitempty"`
}

// ahli brendlerin suratlaryny we id - lerini getiryar
func GetBrends(c *gin.Context) {
	// get all brend from brend controller
	brends, err := backController.GetAllBrendForHomePage()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"brends": brends,
	})
}

func GetHomePageCategories(c *gin.Context) {
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	langID, err := backController.CheckLanguage(c)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	// get all homepage category where translation category id equal langID
	categoryRows, err := db.Query(context.Background(), "SELECT c.id,t.name FROM categories c LEFT JOIN translation_category t ON c.id=t.category_id WHERE t.lang_id = $1 AND c.is_home_category = true AND t.deleted_at IS NULL AND c.deleted_at IS NULL", langID)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer categoryRows.Close()

	var homePageCategories []HomePageCategory
	for categoryRows.Next() {
		var homePageCategory HomePageCategory
		if err := categoryRows.Scan(&homePageCategory.ID, &homePageCategory.Name); err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}

		// get all product where category id equal homePageCategory.ID and lang_id equal langID
		productRows, err := db.Query(context.Background(), "SELECT p.id,p.price,p.old_price,p.limit_amount,p.is_new,p.amount,p.main_image,p.benefit FROM products p LEFT JOIN category_product c ON p.id=c.product_id WHERE c.category_id = $1 AND p.deleted_at IS NULL AND c.deleted_at IS NULL AND p.amount > 0 AND p.limit_amount > 0 ORDER BY p.created_at DESC LIMIT 4", homePageCategory.ID)
		if err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}
		defer productRows.Close()

		var products []Product
		for productRows.Next() {
			var product Product
			if err := productRows.Scan(&product.ID, &product.Price, &product.OldPrice, &product.LimitAmount, &product.IsNew, &product.Amount, &product.MainImage, &product.Benefit); err != nil {
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
				db.QueryRow(context.Background(), "SELECT name,description FROM translation_product WHERE lang_id = $1 AND product_id = $2 AND deleted_at IS NULL", v.ID, product.ID).Scan(&trProduct.Name, &trProduct.Description)
				translation[v.NameShort] = trProduct
				product.Translations = append(product.Translations, translation)
			}

			// get brend where id of product brend_id
			var brend Brend
			db.QueryRow(context.Background(), "SELECT b.id,b.name FROM products p LEFT JOIN brends b ON p.brend_id=b.id WHERE p.id = $1 AND p.deleted_at IS NULL AND b.deleted_at IS NULL", product.ID).Scan(&brend.ID, &brend.Name)
			product.Brend = brend
			products = append(products, product)
		}
		homePageCategory.Products = products
		homePageCategories = append(homePageCategories, homePageCategory)
	}

	var home_page_categories []HomePageCategory
	for _, v := range homePageCategories {
		if len(v.Products) != 0 {
			home_page_categories = append(home_page_categories, v)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"status":              true,
		"homepage_categories": home_page_categories,
	})
}
