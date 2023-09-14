package controllers

import (
	"context"
	"github/abbgo/isleg/isleg-backend/config"
	"github/abbgo/isleg/isleg-backend/helpers"
	"github/abbgo/isleg/isleg-backend/models"
	"github/abbgo/isleg/isleg-backend/pkg"
	"path/filepath"
	"strconv"
	"strings"

	"math"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gosimple/slug"
	"github.com/lib/pq"
	"github.com/xuri/excelize/v2"
	"gopkg.in/guregu/null.v4"
)

type ProductForFront struct {
	ID           string                                 `json:"id,omitempty"`
	Price        float64                                `json:"price,omitempty"`
	OldPrice     float64                                `json:"old_price,omitempty"`
	Percentage   float64                                `json:"percentage,omitempty"`
	MainImage    string                                 `json:"main_image,omitempty"`
	Brend        Brend                                  `json:"brend,omitempty"`
	LimitAmount  int                                    `json:"limit_amount,omitempty"`
	Amount       int                                    `json:"amount,omitempty"`
	IsNew        bool                                   `json:"is_new,omitempty"`
	Benefit      null.Float                             `json:"-"`
	Images       []string                               `json:"images,omitempty"`
	Translations []map[string]models.TranslationProduct `json:"translations"`
}

type DeleteImage struct {
	Image string `json:"image"`
}

type ProductForAdmin struct {
	ID                 string                      `json:"id,omitempty"`
	BrendID            null.String                 `json:"brend_id,omitempty"`
	ShopID             null.String                 `json:"shop_id,omitempty"`
	Price              float64                     `json:"price,omitempty" binding:"required"`
	OldPrice           float64                     `json:"old_price"`
	Benefit            null.Float                  `json:"benefit"`
	Amount             uint                        `json:"amount,omitempty" binding:"required"`
	LimitAmount        uint                        `json:"limit_amount,omitempty" binding:"required"`
	IsNew              bool                        `json:"is_new,omitempty"`
	MainImage          string                      `json:"main_image,omitempty"`
	Images             []string                    `json:"images,omitempty"`                                 // one to many
	TranslationProduct []models.TranslationProduct `json:"translation_product,omitempty" binding:"required"` // one to many
	Categories         []string                    `json:"categories,omitempty" binding:"required"`
}

func DeleteProductImages(c *gin.Context) {
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	var image DeleteImage
	if err := c.Bind(&image); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	if image.Image == "" {
		helpers.HandleError(c, 400, "path of image is required")
		return
	}

	var helperImageID string
	if err := db.QueryRow(context.Background(), "SELECT id FROM helper_images WHERE image = $1 AND deleted_at IS NULL", image.Image).Scan(&helperImageID); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	if helperImageID != "" {
		_, err := db.Exec(context.Background(), "DELETE FROM helper_images WHERE id = $1", helperImageID)
		if err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}
	}

	if err := os.Remove(pkg.ServerPath + image.Image); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "image successfully deleted",
	})
}

func CreateProductImage(c *gin.Context) {
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	var path, file_name string
	imageType := c.Query("image")

	oldImage := c.PostForm("old_path")
	if oldImage != "" {
		if err := os.Remove(pkg.ServerPath + oldImage); err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}

		_, err := db.Exec(context.Background(), "DELETE FROM helper_images WHERE image = $1", oldImage)
		if err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}
	}

	switch imageType {
	case "product":
		fileName := c.Query("type")
		if fileName != "main_image" && fileName != "image" {
			helpers.HandleError(c, 400, "invalid file name")
			return
		}
		path = "product/" + fileName
		file_name = fileName
	case "category":
		path = "category"
		file_name = "image"
	case "brend":
		path = "brend"
		file_name = "image"
	case "language":
		path = "language"
		file_name = "image"
	case "banner":
		path = "banner"
		file_name = "image"
	case "afisa":
		path = "afisa"
		file_name = "image"
	default:
		helpers.HandleError(c, 400, "invalid image")
		return
	}

	image, err := pkg.FileUpload(file_name, path, "image", c)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	_, err = db.Exec(context.Background(), "INSERT INTO helper_images (image) VALUES ($1)", image)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"image":  image,
	})
}

func CreateProduct(c *gin.Context) {
	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	var product ProductForAdmin
	if err := c.BindJSON(&product); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	if product.MainImage == "" {
		helpers.HandleError(c, 400, "main_image is required")
		return
	}

	benefit, _, price, oldPrice, amount, limitAmount, isNew, err := models.ValidateProductModel("", product.Benefit.Float64, "", product.Price, product.OldPrice, product.Amount, product.LimitAmount, product.IsNew, product.Categories)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	var shopID, brendID interface{}
	if product.BrendID.String == "" {
		brendID = nil
	} else {
		brendID = product.BrendID.String
	}
	if product.ShopID.String == "" {
		shopID = nil
	} else {
		shopID = product.ShopID.String
	}

	// create product
	var productID string
	if err := db.QueryRow(context.Background(), "INSERT INTO products (brend_id,price,old_price,amount,limit_amount,is_new,shop_id,main_image,benefit) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9) RETURNING id", brendID, price, oldPrice, amount, limitAmount, isNew, shopID, product.MainImage, benefit).Scan(&productID); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	if len(product.Images) != 0 {
		// create images of product
		_, err := db.Exec(context.Background(), "INSERT INTO images (product_id,image) VALUES ($1,unnest($2::varchar[]))", productID, pq.Array(product.Images))
		if err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}

		_, err = db.Exec(context.Background(), "DELETE FROM helper_images WHERE image = ANY($1)", pq.Array(product.Images))
		if err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}
	}

	for _, v := range product.TranslationProduct {
		_, err := db.Exec(context.Background(), "INSERT INTO translation_product (lang_id,product_id,name,description,slug) VALUES ($1,$2,$3,$4,$5)", v.LangID, productID, v.Name, v.Description, slug.MakeLang(v.Name, "en"))
		if err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}
	}

	// create category product
	_, err = db.Exec(context.Background(), "INSERT INTO category_product (category_id,product_id) VALUES (unnest($1::uuid[]),$2)", pq.Array(product.Categories), productID)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	_, err = db.Exec(context.Background(), "DELETE FROM helper_images WHERE image = $1", product.MainImage)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":     true,
		"message":    "data successfully added",
		"product_id": productID,
	})
}

func CreateProductsByExcelFile(c *gin.Context) {
	countOfRows := 0
	count_of_err := 0

	columns := []string{"c", "d", "e", "f", "g"}
	var sheetName string

	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	// OPEN EXCEL FILE
	f, err := excelize.OpenFile(pkg.ServerPath + "uploads/product/products.xlsx")
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer f.Close()

	sheetName = f.GetSheetName(f.GetActiveSheetIndex())

	// CHECK IS EXISTS "uploads/product/main_image DIRECT" , IF DIRECT NOT FOUND CREATE THIS DIRECT
	_, err = os.Stat(pkg.ServerPath + "uploads/product/main_image")
	if err != nil {
		if err := os.MkdirAll(pkg.ServerPath+"uploads/product/main_image", os.ModePerm); err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}
	}

	// CHECK IS EXISTS "uploads/product/image" DIRECT , IF DIRECT NOT FOUND CREATE THIS DIRECT
	_, err = os.Stat(pkg.ServerPath + "uploads/product/image")
	if err != nil {
		if err := os.MkdirAll(pkg.ServerPath+"uploads/product/image", os.ModePerm); err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}
	}

	//////////////////////  GET COUNT OF ROWS FROM EXCEL FILE ----------------------------------------
	for {
		value, err := f.GetCellValue(sheetName, "c"+strconv.Itoa(countOfRows+3))
		if err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}
		if value != "" {
			countOfRows++
		} else {
			break
		}
	}

	for i := 3; i < countOfRows+3; i++ {
		countOfErr := 0
		errString := ""
		var product ProductForAdmin

		mainImageFileID, err := f.GetCellValue(sheetName, columns[0]+strconv.Itoa(i))
		if err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}
		if mainImageFileID == "" {
			countOfErr++
			errString = errString + "| esasy surat hokman gerekli | "
		}

		if strings.Contains(mainImageFileID, " ") || strings.Contains(mainImageFileID, "_") || strings.Contains(mainImageFileID, "-") {
			mainImageFileID = strings.ReplaceAll(mainImageFileID, " ", "")
			mainImageFileID = strings.ReplaceAll(mainImageFileID, "_", "")
			mainImageFileID = strings.ReplaceAll(mainImageFileID, "-", "")
		}
		mainImageFileID = slug.MakeLang(strings.Trim(mainImageFileID, filepath.Ext(mainImageFileID)), "en") + filepath.Ext(mainImageFileID)

		newFileName, err := pkg.CopyFile("uploads/images/", "uploads/product/main_image/", mainImageFileID)
		if err != nil {
			countOfErr++
			errString = errString + err.Error() + " | "
		}

		_, err = db.Exec(context.Background(), "INSERT INTO helper_images (image) VALUES ($1)", "uploads/product/main_image/"+newFileName)
		if err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}

		product.MainImage = "uploads/product/main_image/" + newFileName

		for _, vColumn := range columns[1:] {
			imageFileID, err := f.GetCellValue(sheetName, vColumn+strconv.Itoa(i))
			if err != nil {
				helpers.HandleError(c, 400, err.Error())
				return
			}

			if imageFileID != "" {
				if strings.Contains(imageFileID, " ") || strings.Contains(imageFileID, "_") || strings.Contains(imageFileID, "-") {
					imageFileID = strings.ReplaceAll(imageFileID, " ", "")
					imageFileID = strings.ReplaceAll(imageFileID, "_", "")
					imageFileID = strings.ReplaceAll(imageFileID, "-", "")
				}
				imageFileID = slug.MakeLang(strings.Trim(imageFileID, filepath.Ext(imageFileID)), "en") + filepath.Ext(imageFileID)

				newFileName, err := pkg.CopyFile("uploads/images/", "uploads/product/image/", imageFileID)
				if err != nil {
					countOfErr++
					errString = errString + err.Error() + " | "

				}

				_, err = db.Exec(context.Background(), "INSERT INTO helper_images (image) VALUES ($1)", "uploads/product/image/"+newFileName)
				if err != nil {
					helpers.HandleError(c, 400, err.Error())
					return
				}
				product.Images = append(product.Images, "uploads/product/image/"+newFileName)
			}
		}

		// //////////////////////      GET CATEGORIES FROM EXCEL FILE ----------------------------------------
		namesOfCategories, err := f.GetCellValue(sheetName, "h"+strconv.Itoa(i))
		if err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}
		if namesOfCategories == "" {
			countOfErr++
			errString = errString + "harydyn kategoriyalary hokman gerek | "
		}

		names_of_categories := strings.Split(namesOfCategories, " | ")
		for _, v := range names_of_categories {
			if strings.Contains(v, " ") {
				strs := strings.Split(v, " ")
				for _, str := range strs {
					if str != "" {
						v = str
					}
				}
			}

			var categoryID string
			if err := db.QueryRow(context.Background(), "SELECT c.id FROM categories c INNER JOIN translation_category tc ON tc.category_id=c.id INNER JOIN languages l ON l.id=tc.lang_id WHERE l.name_short = $1 AND tc.name = $2 AND c.deleted_at IS NULL AND l.deleted_at IS NULL AND tc.deleted_at IS NULL", "tm", v).Scan(&categoryID); err != nil {
				helpers.HandleError(c, 400, err.Error())
				return
			}
			product.Categories = append(product.Categories, categoryID)
		}

		// //////////////////////      GET BREND FROM EXCEL FILE ----------------------------------------
		nameOfBrend, err := f.GetCellValue(sheetName, "n"+strconv.Itoa(i))
		if err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}

		var shopID, brendID interface{}
		if nameOfBrend != "" {
			if err := db.QueryRow(context.Background(), "SELECT id FROM brends WHERE name = $1 AND deleted_at IS NULL", nameOfBrend).Scan(&product.BrendID); err != nil {
				helpers.HandleError(c, 400, err.Error())
				return
			}

			if product.BrendID.String == "" {
				countOfErr++
				errString = errString + nameOfBrend + " atly brend yok | "

			}
		}

		if product.BrendID.String == "" {
			brendID = nil
		} else {
			brendID = product.BrendID.String
		}

		// //////////////////////      GET SHOP FROM EXCEL FILE ----------------------------------------
		shopID = nil // default value
		// shopPhoneNumber, err := f.GetCellValue(sheetName, "g"+strconv.Itoa(vRow))
		// if err != nil {
		// 	c.JSON(http.StatusBadRequest, gin.H{
		// 		"status":  false,
		// 		"message": err.Error(),
		// 	})
		// 	return
		// }

		// rowShop, err := db.Query(context.Background(),"SELECT id FROM shops WHERE phone_number = $1 AND deleted_at IS NULL", shopPhoneNumber)
		// if err != nil {
		// 	c.JSON(http.StatusBadRequest, gin.H{
		// 		"status":  false,
		// 		"message": err.Error(),
		// 	})
		// 	return
		// }
		// defer func() {
		// 	if err := rowShop.Close(); err != nil {
		// 		c.JSON(http.StatusBadRequest, gin.H{
		// 			"status":  false,
		// 			"message": err.Error(),
		// 		})
		// 		return
		// 	}
		// }()
		// for rowShop.Next() {
		// 	if err := rowShop.Scan(&product.ShopID); err != nil {
		// 		c.JSON(http.StatusBadRequest, gin.H{
		// 			"status":  false,
		// 			"message": err.Error(),
		// 		})
		// 		return
		// 	}
		// }
		// if product.ShopID.String == "" {
		// 	shopID = nil
		// } else {
		// 	shopID = product.ShopID.String
		// }

		// //////////////////////      GET PRICE FROM EXCEL FILE ----------------------------------------
		priceStr, err := f.GetCellValue(sheetName, "i"+strconv.Itoa(i))
		if err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}
		product.Price, err = strconv.ParseFloat(priceStr, 64)
		if err != nil {
			countOfErr++
			errString = errString + "harydyn alynan bahasy hokman gerek ya-da ol onluk san , bitin bolmaly | "

		}

		// //////////////////////      GET OLD PRICE FROM EXCEL FILE ----------------------------------------
		product.OldPrice = 0 // default value
		// oldPriceStr, err := f.GetCellValue(sheetName, "i"+strconv.Itoa(vRow))
		// if err != nil {
		// 	c.JSON(http.StatusBadRequest, gin.H{
		// 		"status":  false,
		// 		"message": err.Error(),
		// 	})
		// 	return
		// }
		// product.OldPrice, err = strconv.ParseFloat(oldPriceStr, 64)
		// if err != nil {
		// 	c.JSON(http.StatusBadRequest, gin.H{
		// 		"status":  false,
		// 		"message": err.Error(),
		// 	})
		// 	return
		// }

		// //////////////////////      GET BENEFIT FROM EXCEL FILE ----------------------------------------
		benefitStr, err := f.GetCellValue(sheetName, "j"+strconv.Itoa(i))
		if err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}
		if strings.Contains(benefitStr, "%") {
			benefitStr = strings.Trim(benefitStr, "%")
		}

		product.Benefit.Float64, err = strconv.ParseFloat(benefitStr, 64)
		if err != nil {
			countOfErr++
			errString = errString + "haryda goyulyan goterim hokman gerek ya-da ol onluk san , bitin bolmaly | "
		}

		// //////////////////////      GET AMOUNT FROM EXCEL FILE ----------------------------------------
		amountStr, err := f.GetCellValue(sheetName, "l"+strconv.Itoa(i))
		if err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}
		amount64, err := strconv.ParseUint(amountStr, 10, 32)
		if err != nil {
			countOfErr++
			errString = errString + "harydyn ammardaky mukdary hokman gerek ya-da ol onluk san , bitin bolmaly | "

		}
		product.Amount = uint(amount64)

		// //////////////////////      GET LIMIT AMOUNT FROM EXCEL FILE ----------------------------------------
		limitAmountStr, err := f.GetCellValue(sheetName, "m"+strconv.Itoa(i))
		if err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}
		limitAmount64, err := strconv.ParseUint(limitAmountStr, 10, 32)
		if err != nil {
			countOfErr++
			errString = errString + "haryda goyulyan limit hokman gerek ya-da ol onluk san , bitin bolmaly | "

		}
		product.LimitAmount = uint(limitAmount64)

		// //////////////////////      GET IS NEW FROM EXCEL FILE ----------------------------------------
		isNewStr, err := f.GetCellValue(sheetName, "o"+strconv.Itoa(i))
		if err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}
		product.IsNew, err = strconv.ParseBool(isNewStr)
		if err != nil {
			countOfErr++
			errString = errString + "harydyn tazeligine hokman true ya-da false yazylan bolmaly | "

		}

		benefit, _, price, oldPrice, amount, limitAmount, isNew, err := models.ValidateProductModel("", product.Benefit.Float64, "", product.Price, product.OldPrice, product.Amount, product.LimitAmount, product.IsNew, product.Categories)
		if err != nil {
			countOfErr++
			errString = errString + err.Error() + " | "
		}

		// //////////////////////      GET TRANSLATIONS OF PRODUCT FROM EXCEL FILE ----------------------------------------
		trTitleTM, err := f.GetCellValue(sheetName, "p"+strconv.Itoa(i))
		if err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}

		if trTitleTM == "" {
			countOfErr++
			errString = errString + "harydyn turkmence ady hokman gerek | "
		}

		trDescTM, err := f.GetCellValue(sheetName, "q"+strconv.Itoa(i))
		if err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}

		var tr models.TranslationProduct
		tr.Name = trTitleTM
		tr.Description = trDescTM

		rowLang, err := db.Query(context.Background(), "SELECT id FROM languages WHERE name_short = $1 AND deleted_at IS NULL", "tm")
		if err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}

		for rowLang.Next() {
			if err := rowLang.Scan(&tr.LangID); err != nil {
				helpers.HandleError(c, 400, err.Error())
				return
			}
		}

		product.TranslationProduct = append(product.TranslationProduct, tr)

		trTitleRU, err := f.GetCellValue(sheetName, "r"+strconv.Itoa(i))
		if err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}

		if trTitleRU == "" {
			countOfErr++
			errString = errString + "harydyn osrca ady hokman gerek | "
		}

		trDescRU, err := f.GetCellValue(sheetName, "s"+strconv.Itoa(i))
		if err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}

		tr.Name = trTitleRU
		tr.Description = trDescRU
		rowLang, err = db.Query(context.Background(), "SELECT id FROM languages WHERE name_short = $1 AND deleted_at IS NULL", "ru")
		if err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}

		for rowLang.Next() {
			if err := rowLang.Scan(&tr.LangID); err != nil {
				helpers.HandleError(c, 400, err.Error())
				return
			}
		}
		product.TranslationProduct = append(product.TranslationProduct, tr)

		if countOfErr == 0 {
			// create product
			resultProducts, err := db.Query(context.Background(), "INSERT INTO products (brend_id,price,old_price,amount,limit_amount,is_new,shop_id,main_image,benefit) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9) RETURNING id", brendID, price, oldPrice, amount, limitAmount, isNew, shopID, product.MainImage, benefit)
			if err != nil {
				helpers.HandleError(c, 400, err.Error())
				return
			}

			var productID string
			for resultProducts.Next() {
				if err := resultProducts.Scan(&productID); err != nil {
					helpers.HandleError(c, 400, err.Error())
					return
				}
			}

			if len(product.Images) != 0 {
				// create images of product
				_, err := db.Exec(context.Background(), "INSERT INTO images (product_id,image) VALUES ($1,unnest($2::varchar[]))", productID, pq.Array(product.Images))
				if err != nil {
					helpers.HandleError(c, 400, err.Error())
					return
				}

				_, err = db.Exec(context.Background(), "DELETE FROM helper_images WHERE image = ANY($1)", pq.Array(product.Images))
				if err != nil {
					helpers.HandleError(c, 400, err.Error())
					return
				}
			}

			for _, v := range product.TranslationProduct {
				_, err := db.Exec(context.Background(), "INSERT INTO translation_product (lang_id,product_id,name,description,slug) VALUES ($1,$2,$3,$4,$5)", v.LangID, productID, v.Name, v.Description, slug.MakeLang(v.Name, "en"))
				if err != nil {
					helpers.HandleError(c, 400, err.Error())
					return
				}
			}

			// create category product
			_, err = db.Exec(context.Background(), "INSERT INTO category_product (category_id,product_id) VALUES (unnest($1::uuid[]),$2)", pq.Array(product.Categories), productID)
			if err != nil {
				helpers.HandleError(c, 400, err.Error())
				return
			}

			_, err = db.Exec(context.Background(), "DELETE FROM helper_images WHERE image = $1", product.MainImage)
			if err != nil {
				helpers.HandleError(c, 400, err.Error())
				return
			}

			if err := f.RemoveRow(sheetName, i); err != nil {
				helpers.HandleError(c, 400, err.Error())
				return
			}
			countOfRows--
			i--
		} else {
			count_of_err++
			if err := f.SetCellStr(sheetName, "t"+strconv.Itoa(i), errString); err != nil {
				helpers.HandleError(c, 400, err.Error())
				return
			}
		}

	}

	if err := f.Save(); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	if count_of_err == 0 {
		c.JSON(http.StatusOK, gin.H{
			"status":  true,
			"message": "data successfully added",
		})
		return
	}

	c.JSON(201, gin.H{
		"status":  false,
		"message": "There are " + strconv.Itoa(count_of_err) + " errors in excel file",
	})
}

func UploadExcelFile(c *gin.Context) {
	excel, err := pkg.FileUpload("products", "product", "excel", c)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	_, err = pkg.MultipartFileUpload("images", "images", c)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": excel + " file successfully uploaded",
	})
}

func RemoveExcelFile(c *gin.Context) {
	err := os.Remove(pkg.ServerPath + "uploads/product/products.xlsx") //remove the file
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "excel file successfully deleted",
	})
}

func DownloadErrExcelFile(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "uploads/product/products.xlsx",
	})
}

func UpdateProductByID(c *gin.Context) {
	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	var product ProductForAdmin
	if err := c.BindJSON(&product); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	// get id from request parameter
	ID := c.Param("id")

	// validate data
	benefit, mainImage, price, oldPrice, amount, limitAmount, isNew, err := models.ValidateProductModel(product.MainImage, product.Benefit.Float64, ID, product.Price, product.OldPrice, product.Amount, product.LimitAmount, product.IsNew, product.Categories)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	var shopID, brendID interface{}
	if product.BrendID.String == "" {
		brendID = nil
	} else {
		brendID = product.BrendID.String
	}
	if product.ShopID.String == "" {
		shopID = nil
	} else {
		shopID = product.ShopID.String
	}

	_, err = db.Exec(context.Background(), "UPDATE products SET brend_id = $1 , price = $2 , old_price = $3, amount = $4, limit_amount = $5 , is_new = $6, shop_id = $8 , main_image = $9 , benefit = $10 WHERE id = $7", brendID, price, oldPrice, amount, limitAmount, isNew, ID, shopID, mainImage, benefit)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	// update translation product
	for _, v := range product.TranslationProduct {
		_, err := db.Exec(context.Background(), "UPDATE translation_product SET name = $1, description = $2, slug = $3 WHERE product_id = $4 AND lang_id = $5", v.Name, v.Description, slug.MakeLang(v.Name, "en"), ID, v.LangID)
		if err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}
	}

	_, err = db.Exec(context.Background(), "DELETE FROM images WHERE product_id = $1", ID)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	if len(product.Images) != 0 {
		_, err := db.Exec(context.Background(), "INSERT INTO images (image,product_id) VALUES (unnest($1::varchar[]),$2)", pq.Array(product.Images), ID)
		if err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}

		_, err = db.Exec(context.Background(), "DELETE FROM helper_images WHERE image = ANY($1)", pq.Array(product.Images))
		if err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}
	}

	// update category product
	_, err = db.Exec(context.Background(), "DELETE FROM category_product WHERE product_id = $1", ID)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	_, err = db.Exec(context.Background(), "INSERT INTO category_product (category_id,product_id) VALUES (unnest($1::uuid[]),$2)", pq.Array(product.Categories), ID)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":     true,
		"message":    "data successfully updated",
		"product_id": ID,
	})
}

func GetProductByID(c *gin.Context) {
	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	// get id from request parameter
	ID := c.Param("id")

	var product ProductForAdmin
	if err := db.QueryRow(context.Background(), "SELECT id,brend_id,price,old_price,amount,limit_amount,is_new,main_image,benefit FROM products WHERE id = $1 AND deleted_at IS NULL", ID).Scan(&product.ID, &product.BrendID, &product.Price, &product.OldPrice, &product.Amount, &product.LimitAmount, &product.IsNew, &product.MainImage, &product.Benefit); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	if product.ID == "" {
		helpers.HandleError(c, 404, "record not found")
		return
	}

	rowsImages, err := db.Query(context.Background(), "SELECT image FROM images WHERE deleted_at IS NULL AND product_id = $1", ID)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	var images []string
	for rowsImages.Next() {
		var image string
		if err := rowsImages.Scan(&image); err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}
		images = append(images, image)
	}

	product.Images = images
	rowsCategoryProduct, err := db.Query(context.Background(), "SELECT category_id FROM category_product WHERE product_id = $1 AND deleted_at IS NULL", ID)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	var categories []string
	for rowsCategoryProduct.Next() {
		var category string
		if err := rowsCategoryProduct.Scan(&category); err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}
		categories = append(categories, category)
	}

	if len(categories) == 0 {
		helpers.HandleError(c, 404, "record not found")
		return
	}

	product.Categories = categories

	rowTranslationProduct, err := db.Query(context.Background(), "SELECT lang_id,name,description FROM translation_product WHERE product_id = $1 AND deleted_at IS NULL", ID)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	var translations []models.TranslationProduct
	for rowTranslationProduct.Next() {
		var translation models.TranslationProduct
		if err := rowTranslationProduct.Scan(&translation.LangID, &translation.Name, &translation.Description); err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}
		if translation.Name == "" {
			helpers.HandleError(c, 404, "record not found")
			return
		}
		translations = append(translations, translation)
	}

	product.TranslationProduct = translations

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"product": product,
	})
}

func GetProducts(c *gin.Context) {
	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	rowsProduct, err := db.Query(context.Background(), "SELECT id,brend_id,price,old_price,amount,limit_amount,is_new,main_image,benefit FROM products WHERE deleted_at IS NULL")
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	var products []models.Product
	// var ids []string
	for rowsProduct.Next() {
		var product models.Product
		if err := rowsProduct.Scan(&product.ID, &product.BrendID, &product.Price, &product.OldPrice, &product.Amount, &product.LimitAmount, &product.IsNew, &product.MainImage, &product.Benefit); err != nil {
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

		rowsImages, err := db.Query(context.Background(), "SELECT image FROM images WHERE deleted_at IS NULL AND product_id = $1", product.ID)
		if err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}

		var images []string
		for rowsImages.Next() {
			var image string
			if err := rowsImages.Scan(&image); err != nil {
				helpers.HandleError(c, 400, err.Error())
				return
			}
			images = append(images, image)
		}

		product.Images = images

		rowsCategoryProduct, err := db.Query(context.Background(), "SELECT category_id FROM category_product WHERE product_id = $1 AND deleted_at IS NULL", product.ID)
		if err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}

		var categories []string
		for rowsCategoryProduct.Next() {
			var category string
			if err := rowsCategoryProduct.Scan(&category); err != nil {
				helpers.HandleError(c, 400, err.Error())
				return
			}
			categories = append(categories, category)
		}

		product.Categories = categories

		rowTranslationProduct, err := db.Query(context.Background(), "SELECT lang_id,name,description FROM translation_product WHERE product_id = $1 AND deleted_at IS NULL", product.ID)
		if err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}

		var translations []models.TranslationProduct
		for rowTranslationProduct.Next() {
			var translation models.TranslationProduct
			if err := rowTranslationProduct.Scan(&translation.LangID, &translation.Name, &translation.Description); err != nil {
				helpers.HandleError(c, 400, err.Error())
				return
			}
			translations = append(translations, translation)
		}

		product.TranslationProduct = translations
		products = append(products, product)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":   true,
		"products": products,
	})
}

func DeleteProductByID(c *gin.Context) {
	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	// get id from request parameter
	ID := c.Param("id")

	// check id
	var productID string
	if err := db.QueryRow(context.Background(), "SELECT id FROM products WHERE id = $1 AND deleted_at IS NULL", ID).Scan(&productID); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	if productID == "" {
		helpers.HandleError(c, 404, "record not found")
		return
	}

	_, err = db.Exec(context.Background(), "CALL delete_product($1)", ID)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "data successfully deleted",
	})
}

func RestoreProductByID(c *gin.Context) {
	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	// get id from request parameter
	ID := c.Param("id")

	// check id
	var productID string
	if err := db.QueryRow(context.Background(), "SELECT id FROM products WHERE id = $1 AND deleted_at IS NOT NULL", ID).Scan(&productID); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	if productID == "" {
		helpers.HandleError(c, 404, "record not found")
		return
	}

	_, err = db.Exec(context.Background(), "CALL restore_product($1)", ID)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "data successfully restored",
	})
}

func DeletePermanentlyProductByID(c *gin.Context) {
	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	// get id from request parameter
	ID := c.Param("id")

	// check id
	var productID, mainImage string
	if err := db.QueryRow(context.Background(), "SELECT id,main_image FROM products WHERE id = $1 AND deleted_at IS NOT NULL", ID).Scan(&productID, &mainImage); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	if productID == "" {
		helpers.HandleError(c, 404, "product not found")
		return
	}

	// remove main image of product
	if err := os.Remove(pkg.ServerPath + mainImage); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	// get images of product
	rowsImages, err := db.Query(context.Background(), "SELECT image FROM images WHERE product_id = $1", productID)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	var images []models.Images
	for rowsImages.Next() {
		var image models.Images
		if err := rowsImages.Scan(&image.Image); err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}
		images = append(images, image)
	}

	// remove images of product
	if len(images) != 0 {
		for _, v := range images {
			if err := os.Remove(pkg.ServerPath + v.Image); err != nil {
				helpers.HandleError(c, 400, err.Error())
				return
			}
		}
	}

	_, err = db.Exec(context.Background(), "DELETE FROM cart WHERE product_id = $1", ID)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	_, err = db.Exec(context.Background(), "DELETE FROM likes WHERE product_id = $1", ID)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	_, err = db.Exec(context.Background(), "DELETE FROM ordered_products WHERE product_id = $1", ID)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	_, err = db.Exec(context.Background(), "DELETE FROM products WHERE id = $1", ID)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "data successfully deleted",
	})
}

func GetProductByIDForFront(c *gin.Context) {
	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	// get id from request parameter
	ID := c.Param("id")

	var product ProductForFront
	if err := db.QueryRow(context.Background(), "SELECT id,price,old_price,amount,limit_amount,is_new,main_image,benefit FROM products WHERE id = $1 AND deleted_at IS NULL", ID).Scan(&product.ID, &product.Price, &product.OldPrice, &product.Amount, &product.LimitAmount, &product.IsNew, &product.MainImage, &product.Benefit); err != nil {
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

	if product.ID == "" {
		helpers.HandleError(c, 404, "record not found")
		return
	}

	rowsImages, err := db.Query(context.Background(), "SELECT image FROM images WHERE deleted_at IS NULL AND product_id = $1", ID)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	var images []string
	for rowsImages.Next() {
		var image string
		if err := rowsImages.Scan(&image); err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}
		images = append(images, image)
	}

	product.Images = images

	rowsLang, err := db.Query(context.Background(), "SELECT id,name_short FROM languages WHERE deleted_at IS NULL")
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

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
		rowTrProduct, err := db.Query(context.Background(), "SELECT name,description FROM translation_product WHERE lang_id = $1 AND product_id = $2 AND deleted_at IS NULL", v.ID, product.ID)
		if err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}

		var trProduct models.TranslationProduct
		translation := make(map[string]models.TranslationProduct)
		for rowTrProduct.Next() {
			if err := rowTrProduct.Scan(&trProduct.Name, &trProduct.Description); err != nil {
				helpers.HandleError(c, 400, err.Error())
				return
			}
		}

		translation[v.NameShort] = trProduct
		product.Translations = append(product.Translations, translation)
	}

	// get brend where id equal brend_id of product
	var brend Brend
	if err := db.QueryRow(context.Background(), "SELECT b.id,b.name FROM products p LEFT JOIN brends b ON p.brend_id=b.id WHERE p.id = $1 AND p.deleted_at IS NULL AND b.deleted_at IS NULL", product.ID).Scan(&brend.ID, &brend.Name); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	product.Brend = brend

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"product": product,
	})
}
