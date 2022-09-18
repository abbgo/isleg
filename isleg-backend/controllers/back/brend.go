package controllers

import (
	"github/abbgo/isleg/isleg-backend/config"
	"github/abbgo/isleg/isleg-backend/models"
	"github/abbgo/isleg/isleg-backend/pkg"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type BrendForHomePage struct {
	ID    uuid.UUID `json:"id"`
	Image string    `json:"image"`
}

type OneBrend struct {
	Name  string `json:"name"`
	Image string `json:"image"`
}

// type ProductImages struct {
// 	MainImage string         `json:"main_image"`
// 	Images    pq.StringArray `json:"images"`
// }

func CreateBrend(c *gin.Context) {

	db, err := config.ConnDB()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer db.Close()

	// GET DATA FROM REQUEST
	name := c.PostForm("name")

	// VALIDATE DATA
	if name == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "brend name is required",
		})
		return
	}

	// FILE UPLOAD
	newFileName, err := pkg.FileUpload("image", "brend", c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	// CREATE BREND
	result, err := db.Query("INSERT INTO brends (name,image) VALUES ($1,$2)", name, "uploads/brend/"+newFileName)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer result.Close()

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "brend successfully added",
	})

}

func UpdateBrendByID(c *gin.Context) {

	db, err := config.ConnDB()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer db.Close()

	ID := c.Param("id")
	name := c.PostForm("name")
	var fileName string

	rowBrend, err := db.Query("SELECT image FROM brends WHERE id = $1 AND deleted_at IS NULL", ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer rowBrend.Close()

	var image string

	for rowBrend.Next() {
		if err := rowBrend.Scan(&image); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}

	if image == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "record not found",
		})
		return
	}

	if name == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "name of brend is required",
		})
		return
	}

	file, err := c.FormFile("image")
	if err != nil {
		fileName = image
	} else {
		extensionFile := filepath.Ext(file.Filename)

		if extensionFile != ".jpg" && extensionFile != ".jpeg" && extensionFile != ".png" && extensionFile != ".gif" {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": "the file must be an image",
			})
			return
		}

		newFileName := uuid.New().String() + extensionFile
		c.SaveUploadedFile(file, "./uploads/brend/"+newFileName)

		if err := os.Remove("./" + image); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}

		fileName = "uploads/brend/" + newFileName
	}

	currentTime := time.Now()

	resultBrend, err := db.Query("UPDATE brends SET name = $1 , image = $2 , updated_at = $4 WHERE id = $3", name, fileName, ID, currentTime)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer resultBrend.Close()

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "brend successfully updated",
	})

}

func GetBrendByID(c *gin.Context) {

	db, err := config.ConnDB()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer db.Close()

	ID := c.Param("id")

	rowBrend, err := db.Query("SELECT name,image FROM brends WHERE id = $1 AND deleted_at IS NULL", ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer rowBrend.Close()

	var brend OneBrend

	for rowBrend.Next() {
		if err := rowBrend.Scan(&brend.Name, &brend.Image); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}

	if brend.Name == "" || brend.Image == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "record not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"brend":  brend,
	})

}

func GetBrends(c *gin.Context) {

	db, err := config.ConnDB()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer db.Close()

	rowBrends, err := db.Query("SELECT name,image FROM brends WHERE deleted_at IS NULL")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer rowBrends.Close()

	var brends []OneBrend

	for rowBrends.Next() {
		var brend OneBrend

		if err := rowBrends.Scan(&brend.Name, &brend.Image); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}

		brends = append(brends, brend)
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"brends": brends,
	})

}

func DeleteBrendByID(c *gin.Context) {

	db, err := config.ConnDB()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer db.Close()

	ID := c.Param("id")

	rowBrend, err := db.Query("SELECT image FROM brends WHERE id = $1 AND deleted_at IS NULL", ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer rowBrend.Close()

	var image string

	for rowBrend.Next() {
		if err := rowBrend.Scan(&image); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}

	if image == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "record not found",
		})
		return
	}

	currentTime := time.Now()

	resultBrends, err := db.Query("UPDATE brends SET deleted_at = $1 WHERE id = $2", currentTime, ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer resultBrends.Close()

	resultProducts, err := db.Query("UPDATE products SET deleted_at = $1 WHERE brend_id = $2", currentTime, ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer resultProducts.Close()

	resultTRProduct, err := db.Query("UPDATE translation_product SET deleted_at = $1 FROM products WHERE translation_product.product_id=products.id AND products.brend_id = $2", currentTime, ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer resultTRProduct.Close()

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "brend successfully deleted",
	})

}

func RestoreBrendByID(c *gin.Context) {

	db, err := config.ConnDB()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer db.Close()

	ID := c.Param("id")

	rowBrend, err := db.Query("SELECT image FROM brends WHERE id = $1 AND deleted_at IS NOT NULL", ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer rowBrend.Close()

	var image string

	for rowBrend.Next() {
		if err := rowBrend.Scan(&image); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}

	if image == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "record not found",
		})
		return
	}

	resultBrends, err := db.Query("UPDATE brends SET deleted_at = NULL WHERE id = $1", ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer resultBrends.Close()

	resultProducts, err := db.Query("UPDATE products SET deleted_at = NULL WHERE brend_id = $1", ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer resultProducts.Close()

	resultTRProduct, err := db.Query("UPDATE translation_product SET deleted_at = NULL FROM products WHERE translation_product.product_id=products.id AND products.brend_id = $1", ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer resultTRProduct.Close()

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "brend successfully restored",
	})

}

func DeletePermanentlyBrendByID(c *gin.Context) {

	db, err := config.ConnDB()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer db.Close()

	ID := c.Param("id")

	rowBrend, err := db.Query("SELECT image FROM brends WHERE id = $1 AND deleted_at IS NOT NULL", ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer rowBrend.Close()

	var image string

	for rowBrend.Next() {
		if err := rowBrend.Scan(&image); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}

	if image == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "record not found",
		})
		return
	}

	if err := os.Remove("./" + image); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	rowsMainImage, err := db.Query("SELECT m.small,m.medium,m.large FROM main_image m INNER JOIN products p ON p.id = m.product_id WHERE p.brend_id = $1", ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer rowsMainImage.Close()

	var mainImages []models.MainImage

	for rowsMainImage.Next() {
		var mainImage models.MainImage

		if err := rowsMainImage.Scan(&mainImage.Small, &mainImage.Medium, &mainImage.Large); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}

		mainImages = append(mainImages, mainImage)
	}

	for _, v := range mainImages {
		if err := os.Remove("./" + v.Small); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}

		if err := os.Remove("./" + v.Medium); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}

		if err := os.Remove("./" + v.Large); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}

	rowsImages, err := db.Query("SELECT i.small,i.large FROM images i INNER JOIN products p ON p.id = i.product_id WHERE p.brend_id = $1", ID)
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

		if err := rowsImages.Scan(&image.Small, &image.Large); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}

		images = append(images, image)
	}

	for _, v := range images {
		if err := os.Remove("./" + v.Small); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}

		if err := os.Remove("./" + v.Large); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}

	resultBrends, err := db.Query("DELETE FROM brends WHERE id = $1", ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer resultBrends.Close()

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "brend successfully deleted",
	})

}

func GetAllBrendForHomePage() ([]BrendForHomePage, error) {

	db, err := config.ConnDB()
	if err != nil {
		return []BrendForHomePage{}, nil
	}
	defer db.Close()

	var brends []BrendForHomePage

	// get all brends
	rows, err := db.Query("SELECT id,image FROM brends WHERE deleted_at IS NULL")
	if err != nil {
		return []BrendForHomePage{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var brend BrendForHomePage
		if err := rows.Scan(&brend.ID, &brend.Image); err != nil {
			return []BrendForHomePage{}, err
		}

		brends = append(brends, brend)
	}

	return brends, nil

}
