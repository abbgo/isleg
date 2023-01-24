package controllers

import (
	"database/sql"
	"github/abbgo/isleg/isleg-backend/config"
	backController "github/abbgo/isleg/isleg-backend/controllers/back"
	"github/abbgo/isleg/isleg-backend/models"
	"github/abbgo/isleg/isleg-backend/pkg"
	"log"
	"math"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"github.com/xuri/excelize/v2"
)

type OrderForAdmin struct {
	ID            string          `json:"id"`
	CustomerID    string          `json:"-"`
	FullName      string          `json:"full_name"`
	PhoneNumber   string          `json:"phone_number"`
	Address       string          `json:"address"`
	CustomerMark  string          `json:"customer_mark"`
	OrderTime     string          `json:"order_time"`
	PaymentType   string          `json:"payment_type"`
	TotalPrice    float64         `json:"total_price"`
	ShippingPrice float64         `json:"shipping_price"`
	CreatedAt     string          `json:"created_at"`
	Excel         string          `json:"excel"`
	Products      []ProductOfCart `json:"products"`
}

type Order struct {
	FullName      string        `json:"full_name" binding:"required,min=3"`
	PhoneNumber   string        `json:"phone_number" binding:"required,e164,len=12"`
	Address       string        `json:"address" binding:"required,min=3"`
	CustomerMark  string        `json:"customer_mark"`
	OrderTime     string        `json:"order_time" binding:"required"`
	PaymentType   string        `json:"payment_type" binding:"required"`
	TotalPrice    float64       `json:"total_price" binding:"required"`
	ShippingPrice float64       `json:"shipping_price,omitempty"`
	Products      []CartProduct `json:"products" binding:"required"`
}

type GetOrder struct {
	ID         string          `json:"id"`
	Date       string          `json:"date"`
	TotalPrice float64         `json:"total_price"`
	Products   []ProductOfCart `json:"products"`
}

type OrderedProduct struct {
	Name   string  `json:"name"`
	Price  float64 `json:"price"`
	Amount uint    `json:"amount"`
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
	defer func() {
		if err := db.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

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

	// haryt sargyt etyan musderi on bazada barmy ya-da yokmy sony bilmek ucin order.PhoneNumber telefon belgi boyunca sol musderini
	// bazadan gozletyas
	rowCustomer, err := db.Query("SELECT id,phone_number FROM customers WHERE phone_number = $1 AND deleted_at IS NULL", order.PhoneNumber)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer func() {
		if err := rowCustomer.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

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

	if customerPhoneNumber != "" { // eger musderi on bazada bar bolsa onda , yene-de frontdan gelen order.Address musderinin
		// adresi on bazada barmy ya-da yokmy sony barlayas

		rowsCustomerAddress, err := db.Query("SELECT address FROM customer_address WHERE customer_id = $1 AND address = $2 AND deleted_at IS NULL", customerID, order.Address)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
		defer func() {
			if err := rowsCustomerAddress.Close(); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  false,
					"message": err.Error(),
				})
				return
			}
		}()

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

		if customerAddress == "" { // eger musderinin adresi bazada yok bolsa onda gelen order.Address adresi sol musdera
			// taze adres hokmunde baza yazdyryas

			resultCustomerAddres, err := db.Query("INSERT INTO customer_address (customer_id,address,is_active) VALUES ($1,$2,$3)", customerID, order.Address, false)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  false,
					"message": err.Error(),
				})
				return
			}
			defer func() {
				if err := resultCustomerAddres.Close(); err != nil {
					c.JSON(http.StatusBadRequest, gin.H{
						"status":  false,
						"message": err.Error(),
					})
					return
				}
			}()

		}

	} else { // bu yerde bolsa eger musderi bazada yok bolsa , onda musderini baza gosyas

		resultCustomer, err := db.Query("INSERT INTO customers (full_name,phone_number,is_register) VALUES ($1,$2,$3) RETURNING id", order.FullName, order.PhoneNumber, false)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
		defer func() {
			if err := resultCustomer.Close(); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  false,
					"message": err.Error(),
				})
				return
			}
		}()

		var customer_id string

		for resultCustomer.Next() {
			if err := resultCustomer.Scan(&customer_id); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  false,
					"message": err.Error(),
				})
				return
			}
		}

		// musderi baza gosulandan son bolsa sol musderinin adresini baza gosyas
		resultCustomerAddress, err := db.Query("INSERT INTO customer_address (customer_id,address,is_active) VALUES ($1,$2,$3)", customer_id, order.Address, false)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
		defer func() {
			if err := resultCustomerAddress.Close(); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  false,
					"message": err.Error(),
				})
				return
			}
		}()

		customerID = customer_id

	}

	// musderinin baza bilen barlag isleri gutarandan son musderinin sargydyny baza gosyas we
	// gosulan sargydyn id - sini alyarys, ordered_prodcuts tablisa ucin
	resultOrders, err := db.Query("INSERT INTO orders (customer_id,customer_mark,order_time,payment_type,total_price,shipping_price,address) VALUES ($1,$2,$3,$4,$5,$6,$7) RETURNING id", customerID, order.CustomerMark, order.OrderTime, order.PaymentType, order.TotalPrice, order.ShippingPrice, order.Address)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer func() {
		if err := resultOrders.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

	var order_id string

	for resultOrders.Next() {
		if err := resultOrders.Scan(&order_id); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}

	// edilen sargyt baza gosulandan son sol sargyda degisli harytlary baza gosyas
	for _, v := range order.Products {

		// eger gelyan harydyn mukdary 1 - den kici bolsa
		// ondan yzyna error ugratyas. Sebabi 0 mukdarda haryt sargyt edip bolmayar
		if v.QuantityOfProduct < 1 {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": "quantity of product cannot be less than 1",
			})
			return
		}

		// gelen product_id boyunca sol haryt bazadaky harytlaryn arasynda barmy ya-da yokmy sony barlayas
		rowProduct, err := db.Query("SELECT id FROM products WHERE id = $1 AND deleted_at IS NULL", v.ProductID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
		defer func() {
			if err := rowProduct.Close(); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  false,
					"message": err.Error(),
				})
				return
			}
		}()

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

		// eger sargyt edilen haryt bazada yok bolsa , onda yzyna error iberyas/
		// sebabi bazada yok bolan harydy sargyt edip bolmayar
		if product_id == "" {
			c.JSON(http.StatusNotFound, gin.H{
				"status":  false,
				"message": "product not found",
			})
			return
		}

		// harydyn barlaglary gutarandan son bolsa sargyt edilen harytlary ordered_products tablisa gosyas
		resultOrderedProduct, err := db.Query("INSERT INTO ordered_products (product_id,quantity_of_product,order_id) VALUES ($1,$2,$3)", v.ProductID, v.QuantityOfProduct, order_id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
		defer func() {
			if err := resultOrderedProduct.Close(); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  false,
					"message": err.Error(),
				})
				return
			}
		}()

		// haryt ordered_products tablisa gosulandan son products tablisadan sargyt edilen
		// harytlaryn mukdaryny azaltyas
		resultProduct, err := db.Query("UPDATE products SET amount = amount - $1 WHERE id = $2", v.QuantityOfProduct, v.ProductID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
		defer func() {
			if err := resultProduct.Close(); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  false,
					"message": err.Error(),
				})
				return
			}
		}()

	}

	// harytlar sargyt edilenden son sargyt eden musderinin sebedindaki harytlary ayyryas
	resultCart, err := db.Query("DELETE FROM cart WHERE customer_id = $1", customerID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer func() {
		if err := resultCart.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

	// excel fayly doldurmak ucin bazadan firmanyn telefon belgisini almaly
	rowCompanyPhone, err := db.Query("SELECT phone FROM company_phone ORDER BY created_at DESC LIMIT 1")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer func() {
		if err := rowCompanyPhone.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

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

	// excel fayly doldurmak ucin firmanyn email adresini we instagram sahypasyny bazadan almaly
	rowCompanySetting, err := db.Query("SELECT email,instagram FROM company_setting ORDER BY created_at DESC LIMIT 1")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer func() {
		if err := rowCompanySetting.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

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

	// excel fayly doldurmak ucin edilen sargydyn maglumatlaryny bazadan almaly
	rowOrder, err := db.Query("SELECT order_number,TO_CHAR(created_at,'DD.MM.YYYY HH24:MI'),order_time,customer_mark,total_price,shipping_price,payment_type FROM orders WHERE id = $1 AND deleted_at IS NULL", order_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer func() {
		if err := rowOrder.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

	var sargyt models.Orders

	for rowOrder.Next() {
		if err := rowOrder.Scan(&sargyt.OrderNumber, &sargyt.CreatedAt, &sargyt.OrderTime, &sargyt.CustomerMark, &sargyt.TotalPrice, &sargyt.ShippingPrice, &sargyt.PaymentType); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}

	// excel fayly doldurmak ucin sargyt eden musderinin maglumatlaryny bazadan almaly
	rowsCustomer, err := db.Query("SELECT full_name,phone_number FROM customers WHERE id = $1 AND deleted_at IS NULL", customerID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer func() {
		if err := rowsCustomer.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

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

	// excel fayly doldurmak ucin sargyt edilen harytlary bazadan almaly
	rowsOrderedProducts, err := db.Query("SELECT product_id,quantity_of_product FROM ordered_products WHERE order_id = $1 AND deleted_at IS NULL", order_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer func() {
		if err := rowsOrderedProducts.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

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

		row, err := db.Query("SELECT price FROM products WHERE id= $1 AND deleted_at IS NULL", v.ProductID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
		defer func() {
			if err := row.Close(); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  false,
					"message": err.Error(),
				})
				return
			}
		}()

		for row.Next() {
			if err := row.Scan(&product.Price); err != nil {
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
		defer func() {
			if err := rowTr.Close(); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  false,
					"message": err.Error(),
				})
				return
			}
		}()

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

	// dolduryljak excel fayly acmaly
	f, err := excelize.OpenFile(pkg.ServerPath + "uploads/orders/order.xlsx")
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

	// excel fayly maglumatlar bilen doldurmaly
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
	f.SetCellValue("Лист1", "b12", "Eltip bermek hyzmaty: "+strconv.FormatFloat(pkg.RoundFloat(sargyt.ShippingPrice, 2), 'f', -1, 64))

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

	style2, err := f.NewStyle(&excelize.Style{
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
			Horizontal: "left",
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

	// sargyt edilen harytlar excel fayla yazdyrylyar
	for i := 0; i < len(products); i++ {

		if err = f.InsertRow("Лист1", 16); err != nil {
			log.Fatal(err)
		}

		if err := f.MergeCell("Лист1", "a16", "b16"); err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}

		if err = f.SetCellStyle("Лист1", "a16", "a16", style2); err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}

		if err = f.SetCellStyle("Лист1", "c16", "c16", style); err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}

		if err = f.SetCellStyle("Лист1", "d16", "d16", style); err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}

		if err = f.SetCellStyle("Лист1", "e16", "e16", style); err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}

		if err = f.SetCellStyle("Лист1", "f16", "f16", style1); err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}

	}

	var totalPrice float64 = 0
	var counter int

	for k, v2 := range products {

		f.SetCellValue("Лист1", "a"+strconv.Itoa(16+k), v2.Name)
		f.SetCellValue("Лист1", "c"+strconv.Itoa(16+k), v2.Amount)
		f.SetCellValue("Лист1", "d"+strconv.Itoa(16+k), v2.Price)
		f.SetCellValue("Лист1", "e"+strconv.Itoa(16+k), float64(v2.Amount)*v2.Price)

		totalPrice = totalPrice + float64(v2.Amount)*v2.Price

		counter++

	}

	// sargyt edilen harytlaryn jemi bahasy we sargydyn jemi bahasy excel fayla yazdyrylyar
	f.SetCellValue("Лист1", "d"+strconv.Itoa(17+counter), totalPrice)
	f.SetCellValue("Лист1", "b13", "Jemi: "+strconv.FormatFloat(pkg.RoundFloat(sargyt.TotalPrice, 2), 'f', -1, 64))

	// tayyar bolan excel fayl uploads papka yazdyrylyar
	if err := f.SaveAs(pkg.ServerPath + "uploads/orders/" + strconv.Itoa(int(sargyt.OrderNumber)) + ".xlsx"); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	// excel fayl tayyar bolanson sargydyn fayly hokmunde baza yazdyrylyar
	resultOrderUpdate, err := db.Query("UPDATE orders SET excel = $1 WHERE id = $2", "uploads/orders/"+strconv.Itoa(int(sargyt.OrderNumber))+".xlsx", order_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer func() {
		if err := resultOrderUpdate.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "success",
	})

}

func GetOrders(c *gin.Context) {

	db, err := config.ConnDB()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer func() {
		if err := db.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

	// GET language id
	langID, err := backController.GetLangID("tm")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	// get limit from param
	limitStr := c.Param("limit")
	if limitStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "limit is required",
		})
		return
	}
	limit, err := strconv.ParseUint(limitStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	// get page from param
	pageStr := c.Param("page")
	if pageStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "page is required",
		})
		return
	}
	page, err := strconv.ParseUint(pageStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	offset := limit * (page - 1)
	countOfOrders := 0

	statusStr := c.Query("status")
	status, err := strconv.ParseBool(statusStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	var countAllCustomer, rowsCustomerID, rowsOrder *sql.Rows
	if status {
		rows, err := db.Query("SELECT customer_id FROM orders WHERE deleted_at IS NOT NULL")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
		countAllCustomer = rows
	} else {
		rows, err := db.Query("SELECT customer_id FROM orders WHERE deleted_at IS NULL")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
		countAllCustomer = rows
	}
	defer func() {
		if err := countAllCustomer.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

	for countAllCustomer.Next() {
		countOfOrders++
	}

	if status {
		rows, err := db.Query("SELECT customer_id FROM orders WHERE deleted_at IS NOT NULL GROUP BY customer_id")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
		rowsCustomerID = rows
	} else {
		rows, err := db.Query("SELECT customer_id FROM orders WHERE deleted_at IS NULL GROUP BY customer_id")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
		rowsCustomerID = rows
	}
	defer func() {
		if err := rowsCustomerID.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

	var customerIDs []string

	for rowsCustomerID.Next() {
		var customerID string

		if err := rowsCustomerID.Scan(&customerID); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}

		customerIDs = append(customerIDs, customerID)
	}

	var orders []OrderForAdmin

	if status {
		rows, err := db.Query("SELECT customer_id,id,customer_mark,order_time,payment_type,total_price,shipping_price,excel,address,TO_CHAR(created_at, 'DD.MM.YYYY') FROM orders WHERE customer_id = ANY($1) AND deleted_at IS NOT NULL LIMIT $2 OFFSET $3", pq.Array(customerIDs), limit, offset)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
		rowsOrder = rows
	} else {
		rows, err := db.Query("SELECT customer_id,id,customer_mark,order_time,payment_type,total_price,shipping_price,excel,address,TO_CHAR(created_at, 'DD.MM.YYYY') FROM orders WHERE customer_id = ANY($1) AND deleted_at IS NULL LIMIT $2 OFFSET $3", pq.Array(customerIDs), limit, offset)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
		rowsOrder = rows
	}
	defer func() {
		if err := rowsOrder.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

	for rowsOrder.Next() {
		var order OrderForAdmin
		if err := rowsOrder.Scan(&order.CustomerID, &order.ID, &order.CustomerMark, &order.OrderTime, &order.PaymentType, &order.TotalPrice, &order.ShippingPrice, &order.Excel, &order.Address, &order.CreatedAt); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}

		rowCustomer, err := db.Query("SELECT full_name,phone_number FROM customers WHERE deleted_at IS NULL AND id = $1", order.CustomerID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
		defer func() {
			if err := rowCustomer.Close(); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  false,
					"message": err.Error(),
				})
				return
			}
		}()

		for rowCustomer.Next() {
			if err := rowCustomer.Scan(&order.FullName, &order.PhoneNumber); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  false,
					"message": err.Error(),
				})
				return
			}
		}

		rowsOrderedProducts, err := db.Query("SELECT product_id,quantity_of_product FROM ordered_products WHERE order_id = $1 AND deleted_at IS NULL", order.ID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
		defer func() {
			if err := rowsOrderedProducts.Close(); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  false,
					"message": err.Error(),
				})
				return
			}
		}()

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

			rowProduct, err := db.Query("SELECT brend_id,price,old_price,amount,limit_amount,is_new,main_image FROM products WHERE id = $1 AND deleted_at IS NULL", product.ID)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  false,
					"message": err.Error(),
				})
				return
			}
			defer func() {
				if err := rowProduct.Close(); err != nil {
					c.JSON(http.StatusBadRequest, gin.H{
						"status":  false,
						"message": err.Error(),
					})
					return
				}
			}()

			for rowProduct.Next() {
				if err := rowProduct.Scan(&product.BrendID, &product.Price, &product.OldPrice, &product.Amount, &product.LimitAmount, &product.IsNew, &product.MainImage); err != nil {
					c.JSON(http.StatusBadRequest, gin.H{
						"status":  false,
						"message": err.Error(),
					})
					return
				}
			}

			if product.OldPrice != 0 {
				product.Percentage = -math.Round(((product.OldPrice - product.Price) * 100) / product.OldPrice)
			} else {
				product.Percentage = 0
			}

			rowTrProduct, err := db.Query("SELECT name,description FROM translation_product WHERE product_id = $1 AND lang_id = $2 AND deleted_at IS NULL", product.ID, langID)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  false,
					"message": err.Error(),
				})
				return
			}
			defer func() {
				if err := rowTrProduct.Close(); err != nil {
					c.JSON(http.StatusBadRequest, gin.H{
						"status":  false,
						"message": err.Error(),
					})
					return
				}
			}()

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

			products = append(products, product)

		}

		order.Products = products

		orders = append(orders, order)

	}

	c.JSON(http.StatusOK, gin.H{
		"status":          true,
		"orders":          orders,
		"count_of_orders": countOfOrders,
	})

}

func OrderConfirmation(c *gin.Context) {

	db, err := config.ConnDB()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer func() {
		if err := db.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

	// get id of language from request parameter
	orderID := c.Param("id")

	// check orderID
	rowOrder, err := db.Query("SELECT id FROM orders WHERE id = $1 AND deleted_at IS NULL", orderID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer func() {
		if err := rowOrder.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

	var order_id string

	for rowOrder.Next() {
		if err := rowOrder.Scan(&order_id); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}

	if order_id == "" {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  false,
			"message": "order not found",
		})
		return
	}

	resultOrder, err := db.Query("UPDATE orders SET deleted_at = now() WHERE id = $1", orderID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer func() {
		if err := resultOrder.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "order confirmed",
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
	defer func() {
		if err := db.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

	custID, hasCustomer := c.Get("customer_id")
	if !hasCustomer {
		c.JSON(http.StatusBadRequest, "customer_id is required")
		return
	}
	customerID, ok := custID.(string)
	if !ok {
		c.JSON(http.StatusBadRequest, "customer_id must be string")
	}

	// get limit from param
	limitStr := c.Param("limit")
	if limitStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "limit is required",
		})
		return
	}
	limit, err := strconv.ParseUint(limitStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	// get page from param
	pageStr := c.Param("page")
	if pageStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "page is required",
		})
		return
	}
	page, err := strconv.ParseUint(pageStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	offset := limit * (page - 1)

	var countOfOrders uint

	// programmany ulanyp otyran musderinin sargytlarynyn sanyny alyas frontda pagination ucin
	countOrders, err := db.Query("SELECT COUNT(id) FROM orders WHERE customer_id = $1", customerID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer func() {
		if err := countOrders.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

	for countOrders.Next() {
		if err := countOrders.Scan(&countOfOrders); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}

	// musderinin sargytlaryny alyas bazadan
	rowsOrders, err := db.Query("SELECT id,TO_CHAR(created_at, 'DD.MM.YYYY'),total_price FROM orders WHERE customer_id = $1 ORDER BY created_at ASC LIMIT $2 OFFSET $3", customerID, limit, offset)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer func() {
		if err := rowsOrders.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

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
		defer func() {
			if err := rowsOrderedProducts.Close(); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  false,
					"message": err.Error(),
				})
				return
			}
		}()

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

			rowProduct, err := db.Query("SELECT brend_id,price,old_price,amount,limit_amount,is_new,main_image FROM products WHERE id = $1 AND deleted_at IS NULL", product.ID)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  false,
					"message": err.Error(),
				})
				return
			}
			defer func() {
				if err := rowProduct.Close(); err != nil {
					c.JSON(http.StatusBadRequest, gin.H{
						"status":  false,
						"message": err.Error(),
					})
					return
				}
			}()

			for rowProduct.Next() {
				if err := rowProduct.Scan(&product.BrendID, &product.Price, &product.OldPrice, &product.Amount, &product.LimitAmount, &product.IsNew, &product.MainImage); err != nil {
					c.JSON(http.StatusBadRequest, gin.H{
						"status":  false,
						"message": err.Error(),
					})
					return
				}
			}

			if product.OldPrice != 0 {
				product.Percentage = -math.Round(((product.OldPrice - product.Price) * 100) / product.OldPrice)
			} else {
				product.Percentage = 0
			}

			rowsLang, err := db.Query("SELECT id,name_short FROM languages WHERE deleted_at IS NULL")
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  false,
					"message": err.Error(),
				})
				return
			}
			defer func() {
				if err := rowsLang.Close(); err != nil {
					c.JSON(http.StatusBadRequest, gin.H{
						"status":  false,
						"message": err.Error(),
					})
					return
				}
			}()

			var languages []models.Language

			for rowsLang.Next() {
				var language models.Language

				if err := rowsLang.Scan(&language.ID, &language.NameShort); err != nil {
					c.JSON(http.StatusBadRequest, gin.H{
						"status":  false,
						"message": err.Error(),
					})
					return
				}

				languages = append(languages, language)
			}

			for _, v := range languages {

				rowTrProduct, err := db.Query("SELECT name,description FROM translation_product WHERE product_id = $1 AND lang_id = $2 AND deleted_at IS NULL", product.ID, v.ID)
				if err != nil {
					c.JSON(http.StatusBadRequest, gin.H{
						"status":  false,
						"message": err.Error(),
					})
					return
				}
				defer func() {
					if err := rowTrProduct.Close(); err != nil {
						c.JSON(http.StatusBadRequest, gin.H{
							"status":  false,
							"message": err.Error(),
						})
						return
					}
				}()

				var trProduct models.TranslationProduct

				translation := make(map[string]models.TranslationProduct)

				for rowTrProduct.Next() {
					if err := rowTrProduct.Scan(&trProduct.Name, &trProduct.Description); err != nil {
						c.JSON(http.StatusBadRequest, gin.H{
							"status":  false,
							"message": err.Error(),
						})
						return
					}
				}

				translation[v.NameShort] = trProduct

				product.Translations = append(product.Translations, translation)

			}

			products = append(products, product)

		}

		order.Products = products

		orders = append(orders, order)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":          true,
		"orders":          orders,
		"count_of_orders": countOfOrders,
	})

}

// func GetOrderedProductsWithoutCustomer(c *gin.Context) {

// 	db, err := config.ConnDB()
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"status":  false,
// 			"message": err.Error(),
// 		})
// 		return
// 	}
// 	defer func() {
// 		if err := db.Close(); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{
// 				"status":  false,
// 				"message": err.Error(),
// 			})
// 			return
// 		}
// 	}()

// 	// front - dan maglumatlar bind edilyar
// 	var productIds ProductID
// 	if err := c.BindJSON(&productIds); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"status":  false,
// 			"message": err.Error(),
// 		})
// 		return
// 	}

// 	// front - dan gelen id - li harytlar bazada barmy ya-da yokmy sol barlanyar
// 	for _, v := range productIds.IDS {

// 		rowProduct, err := db.Query("SELECT id FROM products WHERE id = $1 AND deleted_at IS NULL", v)
// 		if err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{
// 				"status":  false,
// 				"message": err.Error(),
// 			})
// 			return
// 		}
// 		defer func() {
// 			if err := rowProduct.Close(); err != nil {
// 				c.JSON(http.StatusBadRequest, gin.H{
// 					"status":  false,
// 					"message": err.Error(),
// 				})
// 				return
// 			}
// 		}()

// 		var product_id string

// 		for rowProduct.Next() {
// 			if err := rowProduct.Scan(&product_id); err != nil {
// 				c.JSON(http.StatusBadRequest, gin.H{
// 					"status":  false,
// 					"message": err.Error(),
// 				})
// 				return
// 			}
// 		}

// 		// eger front - dan gelen id li haryt bazada yok bolsa onda yzyna yalnyslyk goyberilyar
// 		if product_id == "" {
// 			c.JSON(http.StatusBadRequest, gin.H{
// 				"status":  false,
// 				"message": "product not found",
// 			})
// 			return
// 		}

// 	}

// 	// front - dan gelen id - lere id - si den bolan harytlar yzyna ugradylyar
// 	rowOrders, err := db.Query("SELECT id,brend_id,price,old_price,amount,limit_amount,is_new FROM products WHERE id = ANY($1) AND deleted_at IS NULL", pq.Array(productIds.IDS))
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"status":  false,
// 			"message": err.Error(),
// 		})
// 		return
// 	}
// 	defer func() {
// 		if err := rowOrders.Close(); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{
// 				"status":  false,
// 				"message": err.Error(),
// 			})
// 			return
// 		}
// 	}()

// 	var products []LikeProduct

// 	for rowOrders.Next() {
// 		var product LikeProduct

// 		if err := rowOrders.Scan(&product.ID, &product.BrendID, &product.Price, &product.OldPrice, &product.Amount, &product.LimitAmount, &product.IsNew); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{
// 				"status":  false,
// 				"message": err.Error(),
// 			})
// 			return
// 		}

// 		if product.OldPrice != 0 {
// 			product.Percentage = -math.Round(((product.OldPrice - product.Price) * 100) / product.OldPrice)
// 		} else {
// 			product.Percentage = 0
// 		}

// 		rowMainImage, err := db.Query("SELECT small,medium,large FROM main_image WHERE product_id = $1 AND deleted_at IS NULL", product.ID)
// 		if err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{
// 				"status":  false,
// 				"message": err.Error(),
// 			})
// 			return
// 		}
// 		defer func() {
// 			if err := rowMainImage.Close(); err != nil {
// 				c.JSON(http.StatusBadRequest, gin.H{
// 					"status":  false,
// 					"message": err.Error(),
// 				})
// 				return
// 			}
// 		}()

// 		var mainImage models.MainImage

// 		for rowMainImage.Next() {
// 			if err := rowMainImage.Scan(&mainImage.Small, &mainImage.Medium, &mainImage.Large); err != nil {
// 				c.JSON(http.StatusBadRequest, gin.H{
// 					"status":  false,
// 					"message": err.Error(),
// 				})
// 				return
// 			}
// 		}

// 		product.MainImage = mainImage

// 		rowsImages, err := db.Query("SELECT small,large FROM images WHERE product_id = $1 AND deleted_at IS NULL", product.ID)
// 		if err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{
// 				"status":  false,
// 				"message": err.Error(),
// 			})
// 			return
// 		}
// 		defer func() {
// 			if err := rowsImages.Close(); err != nil {
// 				c.JSON(http.StatusBadRequest, gin.H{
// 					"status":  false,
// 					"message": err.Error(),
// 				})
// 				return
// 			}
// 		}()

// 		var images []models.Images

// 		for rowsImages.Next() {
// 			var image models.Images

// 			if err := rowsImages.Scan(&image.Small, &image.Large); err != nil {
// 				c.JSON(http.StatusBadRequest, gin.H{
// 					"status":  false,
// 					"message": err.Error(),
// 				})
// 				return
// 			}

// 			images = append(images, image)
// 		}

// 		product.Images = images

// 		rowsLang, err := db.Query("SELECT id,name_short FROM languages WHERE deleted_at IS NULL")
// 		if err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{
// 				"status":  false,
// 				"message": err.Error(),
// 			})
// 			return
// 		}
// 		defer func() {
// 			if err := rowsLang.Close(); err != nil {
// 				c.JSON(http.StatusBadRequest, gin.H{
// 					"status":  false,
// 					"message": err.Error(),
// 				})
// 				return
// 			}
// 		}()

// 		var languages []models.Language

// 		for rowsLang.Next() {
// 			var language models.Language

// 			if err := rowsLang.Scan(&language.ID, &language.NameShort); err != nil {
// 				c.JSON(http.StatusBadRequest, gin.H{
// 					"status":  false,
// 					"message": err.Error(),
// 				})
// 				return
// 			}

// 			languages = append(languages, language)
// 		}

// 		for _, v := range languages {

// 			rowTrProduct, err := db.Query("SELECT name,description FROM translation_product WHERE product_id = $1 AND lang_id = $2 AND deleted_at IS NULL", product.ID, v.ID)
// 			if err != nil {
// 				c.JSON(http.StatusBadRequest, gin.H{
// 					"status":  false,
// 					"message": err.Error(),
// 				})
// 				return
// 			}
// 			defer func() {
// 				if err := rowTrProduct.Close(); err != nil {
// 					c.JSON(http.StatusBadRequest, gin.H{
// 						"status":  false,
// 						"message": err.Error(),
// 					})
// 					return
// 				}
// 			}()

// 			var trProduct models.TranslationProduct

// 			translation := make(map[string]models.TranslationProduct)

// 			for rowTrProduct.Next() {
// 				if err := rowTrProduct.Scan(&trProduct.Name, &trProduct.Description); err != nil {
// 					c.JSON(http.StatusBadRequest, gin.H{
// 						"status":  false,
// 						"message": err.Error(),
// 					})
// 					return
// 				}
// 			}

// 			translation[v.NameShort] = trProduct

// 			product.Translations = append(product.Translations, translation)

// 		}

// 		products = append(products, product)
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"status":   true,
// 		"products": products,
// 	})

// }

func ReturnOrder(c *gin.Context) {

	db, err := config.ConnDB()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer func() {
		if err := db.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

	orderID := c.Param("id")

	rowOrder, err := db.Query("SELECT id,excel FROM orders WHERE id = $1 AND deleted_at IS NULL", orderID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer func() {
		if err := rowOrder.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

	var order_id, excel string

	for rowOrder.Next() {
		if err := rowOrder.Scan(&order_id, &excel); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}

	if order_id == "" {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  false,
			"message": "order not found",
		})
		return
	}

	if err := os.Remove(pkg.ServerPath + excel); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	rowsOrderedProduct, err := db.Query("SELECT product_id,quantity_of_product FROM ordered_products WHERE order_id = $1 AND deleted_at IS NULL", order_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer func() {
		if err := rowsOrderedProduct.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

	var orderedProducts []models.OrderedProducts

	for rowsOrderedProduct.Next() {
		var orderedProduct models.OrderedProducts

		if err := rowsOrderedProduct.Scan(&orderedProduct.ProductID, &orderedProduct.QuantityOfProduct); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}

		orderedProducts = append(orderedProducts, orderedProduct)
	}

	for _, v := range orderedProducts {

		resultProduct, err := db.Query("UPDATE products SET amount = amount + $1 WHERE id = $2", v.QuantityOfProduct, v.ProductID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
		defer func() {
			if err := resultProduct.Close(); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  false,
					"message": err.Error(),
				})
				return
			}
		}()

	}

	resultOrder, err := db.Query("DELETE FROM orders WHERE id = $1", order_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer func() {
		if err := resultOrder.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "order returned success",
	})

}
