package controllers

import (
	"fmt"
	"github/abbgo/isleg/isleg-backend/config"
	backController "github/abbgo/isleg/isleg-backend/controllers/back"
	"github/abbgo/isleg/isleg-backend/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
)

type Order struct {
	FullName     string        `json:"full_name" binding:"required,min=3"`
	PhoneNumber  string        `json:"phone_number" binding:"required,e164,len=12"`
	Address      string        `json:"address" binding:"required,min=3"`
	CustomerMark string        `json:"customer_mark"`
	OrderTime    string        `json:"order_time" binding:"required"`
	PaymentType  string        `json:"payment_type" binding:"required"`
	TotalPrice   float64       `json:"total_price" binding:"required"`
	Products     []CartProduct `json:"products" binding:"required"`
}

type GetOrder struct {
	ID         string          `json:"id"`
	Date       string          `json:"date"`
	TotalPrice float64         `json:"total_price"`
	Products   []ProductOfCart `json:"products"`
}

type OrderedProduct struct {
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	ProductCode string  `json:"product_code"`
	Amount      uint    `json:"amount"`
}

func ToOrder(c *gin.Context) {

	db, err := config.ConnDB()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer db.Close()

	// GET DATA FROM ROUTE PARAMETER
	langShortName := c.Param("lang")

	// GET language id
	langID, err := backController.GetLangID(langShortName)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	var order Order

	if err := c.BindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	rowCustomer, err := db.Query("SELECT id,phone_number FROM customers WHERE phone_number = $1 AND deleted_at IS NULL", order.PhoneNumber)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer rowCustomer.Close()

	var customerPhoneNumber string
	var customerID string

	for rowCustomer.Next() {
		if err := rowCustomer.Scan(&customerID, &customerPhoneNumber); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}

	if customerPhoneNumber != "" {

		rowsCustomerAddress, err := db.Query("SELECT address FROM customer_address WHERE customer_id = $1 AND address = $2 AND deleted_at IS NULL", customerID, order.Address)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
		defer rowsCustomerAddress.Close()

		var customerAddress string

		for rowsCustomerAddress.Next() {
			if err := rowsCustomerAddress.Scan(&customerAddress); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  false,
					"message": err.Error(),
				})
				return
			}
		}

		if customerAddress == "" {

			resultCustomerAddres, err := db.Query("INSERT INTO customer_address (customer_id,address,is_active) VALUES ($1,$2,$3)", customerID, order.Address, false)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  false,
					"message": err.Error(),
				})
				return
			}
			defer resultCustomerAddres.Close()

		}

	} else {

		resultCustomer, err := db.Query("INSERT INTO customers (full_name,phone_number,is_register) VALUES ($1,$2,$3)", order.FullName, order.PhoneNumber, false)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
		defer resultCustomer.Close()

		lastCustomerID, err := db.Query("SELECT id FROM customers ORDER BY created_at DESC LIMIT 1")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
		defer lastCustomerID.Close()

		var customer_id string

		for lastCustomerID.Next() {
			if err := lastCustomerID.Scan(&customer_id); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  false,
					"message": err.Error(),
				})
				return
			}
		}

		resultCustomerAddress, err := db.Query("INSERT INTO customer_address (customer_id,address,is_active) VALUES ($1,$2,$3)", customer_id, order.Address, false)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
		defer resultCustomerAddress.Close()

		customerID = customer_id

	}

	resultOrders, err := db.Query("INSERT INTO orders (customer_id,customer_mark,order_time,payment_type,total_price) VALUES ($1,$2,$3,$4,$5)", customerID, order.CustomerMark, order.OrderTime, order.PaymentType, order.TotalPrice)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer resultOrders.Close()

	lastOrderID, err := db.Query("SELECT id FROM orders ORDER BY created_at DESC LIMIT 1")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer lastOrderID.Close()

	var order_id string

	for lastOrderID.Next() {
		if err := lastOrderID.Scan(&order_id); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}

	for _, v := range order.Products {

		if v.QuantityOfProduct < 1 {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": "quantity of product cannot be less than 1",
			})
			return
		}

		rowProduct, err := db.Query("SELECT id FROM products WHERE id = $1 AND deleted_at IS NULL", v.ProductID)
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

		resultOrderedProduct, err := db.Query("INSERT INTO ordered_products (product_id,quantity_of_product,order_id) VALUES ($1,$2,$3)", v.ProductID, v.QuantityOfProduct, order_id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
		defer resultOrderedProduct.Close()

	}

	rowCompanyPhone, err := db.Query("SELECT phone FROM company_phone ORDER BY created_at DESC LIMIT 1")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer rowCompanyPhone.Close()

	var companyPhone string

	for rowCompanyPhone.Next() {
		if err := rowCompanyPhone.Scan(&companyPhone); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}

	rowCompanySetting, err := db.Query("SELECT email,instagram FROM company_setting ORDER BY created_at DESC LIMIT 1")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer rowCompanySetting.Close()

	var email, instagram string

	for rowCompanySetting.Next() {
		if err := rowCompanySetting.Scan(&email, &instagram); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}

	rowOrder, err := db.Query("SELECT order_number,TO_CHAR(created_at,'DD.MM.YYYY HH24:MI'),order_time,customer_mark,total_price,payment_type FROM orders WHERE id = $1 AND deleted_at IS NULL", order_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer rowOrder.Close()

	var sargyt models.Orders

	for rowOrder.Next() {
		if err := rowOrder.Scan(&sargyt.OrderNumber, &sargyt.CreatedAt, &sargyt.OrderTime, &sargyt.CustomerMark, &sargyt.TotalPrice, &sargyt.PaymentType); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}

	rowsCustomer, err := db.Query("SELECT full_name,phone_number FROM customers WHERE id = $1 AND deleted_at IS NULL", customerID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer rowsCustomer.Close()

	var customerName, customerPhone string

	for rowsCustomer.Next() {
		if err := rowsCustomer.Scan(&customerName, &customerPhone); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}

	rowsOrderedProducts, err := db.Query("SELECT product_id,quantity_of_product FROM ordered_products WHERE order_id = $1 AND deleted_at IS NULL", order_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer rowsOrderedProducts.Close()

	var orderedProducts []models.OrderedProducts

	for rowsOrderedProducts.Next() {
		var orderedProduct models.OrderedProducts

		if err := rowsOrderedProducts.Scan(&orderedProduct.ProductID, &orderedProduct.QuantityOfProduct); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
		orderedProducts = append(orderedProducts, orderedProduct)

	}

	var products []OrderedProduct

	for _, v := range orderedProducts {

		var product OrderedProduct

		product.Amount = v.QuantityOfProduct

		row, err := db.Query("SELECT price,product_code FROM products WHERE id= $1 AND deleted_at IS NULL", v.ProductID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
		defer row.Close()

		for row.Next() {
			if err := row.Scan(&product.Price, &product.ProductCode); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  false,
					"message": err.Error(),
				})
				return
			}
		}

		rowTr, err := db.Query("SELECT name FROM translation_product WHERE product_id = $1 AND lang_id = $2 AND deleted_at IS NULL", v.ProductID, langID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
		defer rowTr.Close()

		for rowTr.Next() {
			if err := rowTr.Scan(&product.Name); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  false,
					"message": err.Error(),
				})
				return
			}
		}

		products = append(products, product)
	}

	f, err := excelize.OpenFile("./uploads/orders/order.xlsx")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	defer func() {
		if err := f.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

	f.SetCellValue("Лист1", "c1", "Telefon: "+companyPhone)
	f.SetCellValue("Лист1", "c2", "IMO: "+companyPhone)
	f.SetCellValue("Лист1", "c3", "Instagram: "+instagram)
	f.SetCellValue("Лист1", "c4", "Mail: "+email)
	f.SetCellValue("Лист1", "a6", "Sargyt No: "+strconv.Itoa(sargyt.OrderNumber))
	f.SetCellValue("Лист1", "a9", "Ady: "+customerName)
	f.SetCellValue("Лист1", "a10", "Telefon nomer: "+customerPhone)
	f.SetCellValue("Лист1", "a11", "Salgy: "+order.Address)
	f.SetCellValue("Лист1", "a12", "Bellik: "+sargyt.CustomerMark)
	f.SetCellValue("Лист1", "B9", "Sargyt edilen senesi: "+sargyt.CreatedAt)
	f.SetCellValue("Лист1", "b10", "Eltip bermeli wagty: "+sargyt.OrderTime)
	f.SetCellValue("Лист1", "b11", "Toleg sekili: "+sargyt.PaymentType)
	f.SetCellValue("Лист1", "b12", "Jemi: "+strconv.FormatFloat(sargyt.TotalPrice, 'f', 6, 64))

	for i := 0; i < len(products); i++ {

		if err = f.InsertRow("Лист1", 16); err != nil {
			log.Fatal(err)
		}

		style, err := f.NewStyle(&excelize.Style{
			Border: []excelize.Border{
				{Type: "left", Color: "#000000", Style: 1},
				{Type: "top", Color: "#000000", Style: 1},
				{Type: "bottom", Color: "#000000", Style: 1},
				{Type: "right", Color: "#000000", Style: 1},
			},
			Font: &excelize.Font{
				Bold:   false,
				Italic: false,
				Family: "Calibri",
				Size:   9,
				Color:  "#000000",
			},
			Alignment: &excelize.Alignment{
				Horizontal: "center",
			},
		})
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}

		style1, err := f.NewStyle(&excelize.Style{
			Border: []excelize.Border{
				{Type: "left", Color: "#000000", Style: 1},
			},
		})
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}

		if err = f.SetCellStyle("Лист1", "a16", "a16", style); err != nil {
			log.Fatal(err)
		}

		if err = f.SetCellStyle("Лист1", "b16", "b16", style); err != nil {
			log.Fatal(err)
		}

		if err = f.SetCellStyle("Лист1", "c16", "c16", style); err != nil {
			log.Fatal(err)
		}

		if err = f.SetCellStyle("Лист1", "d16", "d16", style); err != nil {
			log.Fatal(err)
		}

		if err = f.SetCellStyle("Лист1", "e16", "e16", style); err != nil {
			log.Fatal(err)
		}

		if err = f.SetCellStyle("Лист1", "f16", "f16", style1); err != nil {
			log.Fatal(err)
		}

	}

	var totalPrice float64 = 0

	for k, v2 := range products {

		f.SetCellValue("Лист1", "a"+strconv.Itoa(16+k), v2.Name)
		f.SetCellValue("Лист1", "b"+strconv.Itoa(16+k), v2.Amount)
		f.SetCellValue("Лист1", "c"+strconv.Itoa(16+k), v2.ProductCode)
		f.SetCellValue("Лист1", "d"+strconv.Itoa(16+k), v2.Price)
		f.SetCellValue("Лист1", "e"+strconv.Itoa(16+k), float64(v2.Amount)*v2.Price)

		totalPrice = totalPrice + float64(v2.Amount)*v2.Price

	}

	f.SetCellValue("Лист1", "d20", totalPrice)

	if err := f.SaveAs("./uploads/orders/" + strconv.Itoa(int(sargyt.OrderNumber)) + ".xlsx"); err != nil {
		fmt.Println(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":    true,
		"message":   "success",
		"file_path": "uploads/orders/" + strconv.Itoa(int(sargyt.OrderNumber)) + ".xlsx",
	})

}

func GetCustomerOrders(c *gin.Context) {

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

	// GET DATA FROM ROUTE PARAMETER
	langShortName := c.Param("lang")

	// GET language id
	langID, err := backController.GetLangID(langShortName)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	rowCustomer, err := db.Query("SELECT id FROM customers WHERE id = $1 AND is_register = true AND deleted_at IS NULL", customerID)
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
		if err := rowCustomer.Scan(&customer_id); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": "customer not found",
			})
			return
		}
	}

	rowsOrders, err := db.Query("SELECT id,TO_CHAR(created_at, 'DD.MM.YYYY'),total_price FROM orders WHERE customer_id = $1 AND deleted_at IS NULL", customerID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer rowsOrders.Close()

	var orders []GetOrder

	for rowsOrders.Next() {
		var order GetOrder

		if err := rowsOrders.Scan(&order.ID, &order.Date, &order.TotalPrice); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}

		rowsOrderedProducts, err := db.Query("SELECT product_id,quantity_of_product FROM ordered_products WHERE order_id = $1 AND deleted_at IS NULL", order.ID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
		defer rowsOrderedProducts.Close()

		var products []ProductOfCart

		for rowsOrderedProducts.Next() {
			var product ProductOfCart

			if err := rowsOrderedProducts.Scan(&product.ID, &product.QuantityOfProduct); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  false,
					"message": err.Error(),
				})
				return
			}

			rowProduct, err := db.Query("SELECT brend_id,price,old_price,amount,product_code,limit_amount,is_new FROM products WHERE id = $1 AND deleted_at IS NULL", product.ID)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  false,
					"message": err.Error(),
				})
				return
			}
			defer rowProduct.Close()

			for rowProduct.Next() {
				if err := rowProduct.Scan(&product.BrendID, &product.Price, &product.OldPrice, &product.Amount, &product.ProductCode, &product.LimitAmount, &product.IsNew); err != nil {
					c.JSON(http.StatusBadRequest, gin.H{
						"status":  false,
						"message": err.Error(),
					})
					return
				}
			}

			rowMainImage, err := db.Query("SELECT small,medium,large FROM main_image WHERE product_id = $1 AND deleted_at IS NULL", product.ID)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  false,
					"message": err.Error(),
				})
				return
			}
			defer rowMainImage.Close()

			var mainImage models.MainImage

			for rowMainImage.Next() {
				if err := rowMainImage.Scan(&mainImage.Small, &mainImage.Medium, &mainImage.Large); err != nil {
					c.JSON(http.StatusBadRequest, gin.H{
						"status":  false,
						"message": err.Error(),
					})
					return
				}
			}

			product.MainImage = mainImage

			rowsImages, err := db.Query("SELECT small,large FROM images WHERE product_id = $1 AND deleted_at IS NULL", product.ID)
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

			product.Images = images

			rowTrProduct, err := db.Query("SELECT name,description FROM translation_product WHERE product_id = $1 AND lang_id = $2 AND deleted_at IS NULL", product.ID, langID)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  false,
					"message": err.Error(),
				})
				return
			}
			defer rowTrProduct.Close()

			var trProduct models.TranslationProduct

			for rowTrProduct.Next() {
				if err := rowTrProduct.Scan(&trProduct.Name, &trProduct.Description); err != nil {
					c.JSON(http.StatusBadRequest, gin.H{
						"status":  false,
						"message": err.Error(),
					})
					return
				}
			}
			product.TranslationProduct = trProduct

			products = append(products, product)

		}

		order.Products = products

		orders = append(orders, order)
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"orders": orders,
	})

}
