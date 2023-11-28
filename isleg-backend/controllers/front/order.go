package controllers

import (
	"context"
	"fmt"
	"github/abbgo/isleg/isleg-backend/config"
	backController "github/abbgo/isleg/isleg-backend/controllers/back"
	"github/abbgo/isleg/isleg-backend/helpers"
	"github/abbgo/isleg/isleg-backend/models"
	"github/abbgo/isleg/isleg-backend/pkg"
	"math"
	"net/http"
	"os"
	"strconv"

	_ "image/jpeg"
	_ "image/png"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/lib/pq"
	"github.com/xuri/excelize/v2"
	"gopkg.in/guregu/null.v4"
)

type OrderForAdmin struct {
	ID            string            `json:"id"`
	CustomerID    string            `json:"-"`
	FullName      string            `json:"full_name"`
	PhoneNumber   string            `json:"phone_number"`
	Address       string            `json:"address"`
	CustomerMark  string            `json:"customer_mark"`
	OrderTime     string            `json:"order_time"`
	PaymentType   string            `json:"payment_type"`
	TotalPrice    float64           `json:"total_price"`
	ShippingPrice float64           `json:"shipping_price"`
	CreatedAt     string            `json:"created_at"`
	Excel         null.String       `json:"excel,omitempty"`
	OrderNumber   int               `json:"order_number,omitempty"`
	Products      []ProductForAdmin `json:"products"`
}

type ProductForAdmin struct {
	ID                uuid.UUID                 `json:"id"`
	BrendID           uuid.NullUUID             `json:"brend_id,omitempty"`
	Price             float64                   `json:"price"`
	OldPrice          float64                   `json:"old_price"`
	Percentage        float64                   `json:"percentage"`
	Amount            uint                      `json:"amount"`
	LimitAmount       uint                      `json:"limit_amount"`
	IsNew             bool                      `json:"is_new"`
	Benefit           null.Float                `json:"-"`
	QuantityOfProduct int                       `json:"quantity_of_product"`
	MainImage         string                    `json:"main_image"`
	Translations      models.TranslationProduct `json:"translations"`
	Code              null.String               `json:"code,omitempty"`
}

type Order struct {
	FullName      string        `json:"full_name" binding:"required,min=3"`
	PhoneNumber   string        `json:"phone_number" binding:"required,e164,len=12"`
	Address       string        `json:"address" binding:"required,min=3"`
	CustomerMark  string        `json:"customer_mark"`
	OrderTime     string        `json:"order_time" binding:"required"`
	PaymentType   uint8         `json:"payment_type" binding:"required"`
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
	Name   string      `json:"name"`
	Price  float64     `json:"price"`
	Amount uint        `json:"amount"`
	Code   null.String `json:"code"`
}

func ToOrder(c *gin.Context) {

	langID, err := pkg.ValidateMiddlewareData(c, "lang_id")
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	var order Order
	if err := c.BindJSON(&order); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	var paymentType string
	db.QueryRow(context.Background(), "SELECT name FROM payment_types WHERE deleted_at IS NULL AND lang_id = $1 AND value = $2", langID, order.PaymentType).Scan(&paymentType)
	if paymentType == "" {
		helpers.HandleError(c, 404, "payment type not found")
		return
	}

	// haryt sargyt etyan musderi on bazada barmy ya-da yokmy sony bilmek ucin order.PhoneNumber telefon belgi boyunca sol musderini
	// bazadan gozletyas
	var customerPhoneNumber, customerID string
	db.QueryRow(context.Background(), "SELECT id,phone_number FROM customers WHERE phone_number = $1 AND deleted_at IS NULL", order.PhoneNumber).Scan(&customerID, &customerPhoneNumber)

	if customerPhoneNumber != "" { // eger musderi on bazada bar bolsa onda , yene-de frontdan gelen order.Address musderinin
		// adresi on bazada barmy ya-da yokmy sony barlayas
		var customerAddress string
		db.QueryRow(context.Background(), "SELECT address FROM customer_address WHERE customer_id = $1 AND address = $2 AND deleted_at IS NULL", customerID, order.Address).Scan(&customerAddress)

		if customerAddress == "" { // eger musderinin adresi bazada yok bolsa onda gelen order.Address adresi sol musdera
			// taze adres hokmunde baza yazdyryas
			_, err := db.Exec(context.Background(), "INSERT INTO customer_address (customer_id,address,is_active) VALUES ($1,$2,$3)", customerID, order.Address, false)
			if err != nil {
				helpers.HandleError(c, 400, err.Error())
				return
			}
		}
	} else { // bu yerde bolsa eger musderi bazada yok bolsa , onda musderini baza gosyas
		var customer_id string
		db.QueryRow(context.Background(), "INSERT INTO customers (full_name,phone_number,is_register) VALUES ($1,$2,$3) RETURNING id", order.FullName, order.PhoneNumber, false).Scan(&customer_id)

		// musderi baza gosulandan son bolsa sol musderinin adresini baza gosyas
		_, err = db.Exec(context.Background(), "INSERT INTO customer_address (customer_id,address,is_active) VALUES ($1,$2,$3)", customer_id, order.Address, false)
		if err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}
		customerID = customer_id
	}

	// musderinin baza bilen barlag isleri gutarandan son musderinin sargydyny baza gosyas we
	// gosulan sargydyn id - sini alyarys, ordered_prodcuts tablisa ucin
	var order_id, createdAt string
	var orderNumber uint
	db.QueryRow(context.Background(), "INSERT INTO orders (customer_id,customer_mark,order_time,payment_type,total_price,shipping_price,address) VALUES ($1,$2,$3,$4,$5,$6,$7) RETURNING id,TO_CHAR(created_at,'DD.MM.YYYY HH24:MI'),order_number", customerID, order.CustomerMark, order.OrderTime, order.PaymentType, order.TotalPrice, order.ShippingPrice, order.Address).Scan(&order_id, &createdAt, &orderNumber)

	// edilen sargyt baza gosulandan son sol sargyda degisli harytlary baza gosyas
	for _, v := range order.Products {
		// eger gelyan harydyn mukdary 1 - den kici bolsa
		// ondan yzyna error ugratyas. Sebabi 0 mukdarda haryt sargyt edip bolmayar
		if v.QuantityOfProduct < 1 {
			helpers.HandleError(c, 400, "quantity of product cannot be less than 1")
			return
		}

		// harydyn barlaglary gutarandan son bolsa sargyt edilen harytlary ordered_products tablisa gosyas
		_, err := db.Exec(context.Background(), "INSERT INTO ordered_products (product_id,quantity_of_product,order_id) VALUES ($1,$2,$3)", v.ProductID, v.QuantityOfProduct, order_id)
		if err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}

		// haryt ordered_products tablisa gosulandan son products tablisadan sargyt edilen
		// harytlaryn mukdaryny azaltyas
		_, err = db.Exec(context.Background(), "UPDATE products SET amount = amount - $1 WHERE id = $2", v.QuantityOfProduct, v.ProductID)
		if err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}
	}

	// harytlar sargyt edilenden son sargyt eden musderinin sebedindaki harytlary ayyryas
	_, err = db.Exec(context.Background(), "DELETE FROM cart WHERE customer_id = $1", customerID)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	// excel fayly doldurmak ucin bazadan firmanyn telefon belgisini almaly
	var companyPhone string
	db.QueryRow(context.Background(), "SELECT phone FROM company_phone ORDER BY created_at DESC LIMIT 1").Scan(&companyPhone)

	// excel fayly doldurmak ucin firmanyn email adresini we instagram sahypasyny bazadan almaly
	var email, instagram string
	db.QueryRow(context.Background(), "SELECT email,instagram FROM company_setting ORDER BY created_at DESC LIMIT 1").Scan(&email, &instagram)

	var products []OrderedProduct
	for _, v := range order.Products {
		var product OrderedProduct
		product.Amount = v.QuantityOfProduct
		var benefit null.Float
		db.QueryRow(context.Background(), "SELECT price,code,benefit FROM products WHERE id= $1 AND deleted_at IS NULL", v.ProductID).Scan(&product.Price, &product.Code, &benefit)
		if benefit.Float64 != 0 {
			product.Price = product.Price + (product.Price*benefit.Float64)/100
		}

		db.QueryRow(context.Background(), "SELECT name FROM translation_product WHERE product_id = $1 AND lang_id = $2 AND deleted_at IS NULL", v.ProductID, langID).Scan(&product.Name)
		products = append(products, product)
	}

	// dolduryljak excel fayly acmaly
	f, err := excelize.OpenFile(pkg.ServerPath + "uploads/orders/order.xlsx")
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer func() {
		if err := f.Close(); err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}
	}()

	// excel fayly maglumatlar bilen doldurmaly
	f.SetCellValue("Sheet1", "f1", "Telefon: "+companyPhone)
	f.SetCellValue("Sheet1", "f2", "IMO: "+companyPhone)
	f.SetCellValue("Sheet1", "f3", "Instagram: "+instagram)
	f.SetCellValue("Sheet1", "f4", "Mail: "+email)
	f.SetCellValue("Sheet1", "a6", "Sargyt No: "+strconv.Itoa(int(orderNumber)))
	f.SetCellValue("Sheet1", "a9", "Ady: "+order.FullName)
	f.SetCellValue("Sheet1", "a10", "Telefon nomer: "+order.PhoneNumber)
	f.SetCellValue("Sheet1", "a11", "Salgy: "+order.Address)
	f.SetCellValue("Sheet1", "a12", "Bellik: "+order.CustomerMark)
	f.SetCellValue("Sheet1", "e9", "Sargyt edilen senesi: "+createdAt)
	f.SetCellValue("Sheet1", "e10", "Eltip bermeli wagty: "+order.OrderTime)
	f.SetCellValue("Sheet1", "e11", "Toleg sekili: "+paymentType)
	f.SetCellValue("Sheet1", "e12", "Eltip bermek hyzmaty: "+strconv.FormatFloat(pkg.RoundFloat(order.ShippingPrice, 2), 'f', -1, 64)+" TMT")
	f.SetCellValue("Sheet1", "e13", "Jemi: "+strconv.FormatFloat(pkg.RoundFloat(order.TotalPrice, 2), 'f', -1, 64)+" TMT")

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
			Size:   10,
			Color:  "#000000",
		},
		Alignment: &excelize.Alignment{
			Horizontal: "center",
		},
	})
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
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
			Size:   10,
			Color:  "#000000",
		},
		Alignment: &excelize.Alignment{
			Horizontal: "left",
		},
	})
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	// sargyt edilen harytlar excel fayla yazdyrylyar
	lenProductsArr := len(products)
	for i := 0; i < lenProductsArr; i++ {
		f.InsertRows("Sheet1", 16, 1)
		f.MergeCell("Sheet1", "a16", "e16")
		f.SetCellStyle("Sheet1", "a16", "a16", style2)
		f.SetCellStyle("Sheet1", "b16", "b16", style2)
		f.SetCellStyle("Sheet1", "c16", "c16", style2)
		f.SetCellStyle("Sheet1", "d16", "d16", style2)
		f.SetCellStyle("Sheet1", "e16", "e16", style2)
		f.SetCellStyle("Sheet1", "f16", "f16", style)
		f.SetCellStyle("Sheet1", "g16", "g16", style)
		f.SetCellStyle("Sheet1", "h16", "h16", style)
		f.SetCellStyle("Sheet1", "i16", "h16", style)
	}

	var totalPrice float64 = 0

	for k, v2 := range products {
		f.SetCellValue("Sheet1", "a"+strconv.Itoa(16+k), v2.Name)
		f.SetCellValue("Sheet1", "f"+strconv.Itoa(16+k), v2.Amount)
		f.SetCellValue("Sheet1", "g"+strconv.Itoa(16+k), v2.Code.String)
		f.SetCellValue("Sheet1", "h"+strconv.Itoa(16+k), v2.Price)
		f.SetCellValue("Sheet1", "i"+strconv.Itoa(16+k), float64(v2.Amount)*v2.Price)

		totalPrice = totalPrice + float64(v2.Amount)*v2.Price
	}

	// sargyt edilen harytlaryn jemi bahasy we sargydyn jemi bahasy excel fayla yazdyrylyar
	f.SetCellValue("Sheet1", "g"+strconv.Itoa(17+lenProductsArr), strconv.FormatFloat(totalPrice, 'f', -1, 64)+" TMT")

	// if 2*(20+lenProductsArr) > 54 {
	if 2*(20+lenProductsArr) > 60 {
		// count := 54 - (20 + lenProductsArr)
		count := 60 - (20 + lenProductsArr)
		for i := 0; i < count; i++ {
			f.InsertRows("Sheet1", 20+lenProductsArr, 1)
		}

		if err := f.AddPicture("Sheet1", "b"+strconv.Itoa(lenProductsArr+count+21), pkg.ServerPath+"uploads/orders/isleg.png", &excelize.GraphicOptions{
			ScaleX:  0.5,
			ScaleY:  0.65,
			OffsetX: 50,
			OffsetY: 7,
		}); err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}

		f.SetCellValue("Sheet1", "f"+strconv.Itoa(lenProductsArr+count+21), "Telefon: "+companyPhone)
		f.SetCellValue("Sheet1", "f"+strconv.Itoa(lenProductsArr+count+22), "IMO: "+companyPhone)
		f.SetCellValue("Sheet1", "f"+strconv.Itoa(lenProductsArr+count+23), "Instagram: "+instagram)
		f.SetCellValue("Sheet1", "f"+strconv.Itoa(lenProductsArr+count+24), "Mail: "+email)
		f.SetCellValue("Sheet1", "a"+strconv.Itoa(lenProductsArr+count+26), "Sargyt No: "+strconv.Itoa(int(orderNumber)))
		f.SetCellValue("Sheet1", "a"+strconv.Itoa(lenProductsArr+count+29), "Ady: "+order.FullName)
		f.SetCellValue("Sheet1", "a"+strconv.Itoa(lenProductsArr+count+30), "Telefon nomer: "+order.PhoneNumber)
		f.SetCellValue("Sheet1", "a"+strconv.Itoa(lenProductsArr+count+31), "Salgy: "+order.Address)
		f.SetCellValue("Sheet1", "a"+strconv.Itoa(lenProductsArr+count+32), "Bellik: "+order.CustomerMark)
		f.SetCellValue("Sheet1", "e"+strconv.Itoa(lenProductsArr+count+29), "Sargyt edilen senesi: "+createdAt)
		f.SetCellValue("Sheet1", "e"+strconv.Itoa(lenProductsArr+count+30), "Eltip bermeli wagty: "+order.OrderTime)
		f.SetCellValue("Sheet1", "e"+strconv.Itoa(lenProductsArr+count+31), "Toleg sekili: "+paymentType)
		f.SetCellValue("Sheet1", "e"+strconv.Itoa(lenProductsArr+count+32), "Eltip bermek hyzmaty: "+strconv.FormatFloat(pkg.RoundFloat(order.ShippingPrice, 2), 'f', -1, 64)+" TMT")
		f.SetCellValue("Sheet1", "e"+strconv.Itoa(lenProductsArr+count+33), "Jemi: "+strconv.FormatFloat(pkg.RoundFloat(order.TotalPrice, 2), 'f', -1, 64)+" TMT")

		for i := 0; i < lenProductsArr; i++ {
			f.InsertRows("Sheet1", 37+lenProductsArr+count, 1)
			f.MergeCell("Sheet1", "a"+strconv.Itoa(lenProductsArr+count+37), "e"+strconv.Itoa(lenProductsArr+count+37))
			f.SetCellStyle("Sheet1", "a"+strconv.Itoa(lenProductsArr+count+37), "a"+strconv.Itoa(lenProductsArr+count+37), style2)
			f.SetCellStyle("Sheet1", "b"+strconv.Itoa(lenProductsArr+count+37), "b"+strconv.Itoa(lenProductsArr+count+37), style2)
			f.SetCellStyle("Sheet1", "c"+strconv.Itoa(lenProductsArr+count+37), "c"+strconv.Itoa(lenProductsArr+count+37), style2)
			f.SetCellStyle("Sheet1", "d"+strconv.Itoa(lenProductsArr+count+37), "d"+strconv.Itoa(lenProductsArr+count+37), style2)
			f.SetCellStyle("Sheet1", "e"+strconv.Itoa(lenProductsArr+count+37), "e"+strconv.Itoa(lenProductsArr+count+37), style2)
			f.SetCellStyle("Sheet1", "f"+strconv.Itoa(lenProductsArr+count+37), "f"+strconv.Itoa(lenProductsArr+count+37), style)
			f.SetCellStyle("Sheet1", "g"+strconv.Itoa(lenProductsArr+count+37), "g"+strconv.Itoa(lenProductsArr+count+37), style)
			f.SetCellStyle("Sheet1", "h"+strconv.Itoa(lenProductsArr+count+37), "h"+strconv.Itoa(lenProductsArr+count+37), style)
			f.SetCellStyle("Sheet1", "i"+strconv.Itoa(lenProductsArr+count+37), "h"+strconv.Itoa(lenProductsArr+count+37), style)
		}

		for k, v2 := range products {
			f.SetCellValue("Sheet1", "a"+strconv.Itoa(37+k+lenProductsArr+count), v2.Name)
			f.SetCellValue("Sheet1", "f"+strconv.Itoa(37+k+lenProductsArr+count), v2.Amount)
			f.SetCellValue("Sheet1", "g"+strconv.Itoa(37+k+lenProductsArr+count), v2.Code.String)
			f.SetCellValue("Sheet1", "h"+strconv.Itoa(37+k+lenProductsArr+count), v2.Price)
			f.SetCellValue("Sheet1", "i"+strconv.Itoa(37+k+lenProductsArr+count), float64(v2.Amount)*v2.Price)
		}

		f.SetCellValue("Sheet1", "g"+strconv.Itoa(38+2*lenProductsArr+count), strconv.FormatFloat(totalPrice, 'f', -1, 64)+" TMT")
	} else {
		if err := f.AddPicture("Sheet1", "b"+strconv.Itoa(lenProductsArr+21), pkg.ServerPath+"uploads/orders/isleg.png", &excelize.GraphicOptions{
			ScaleX:  0.7,
			ScaleY:  0.6,
			OffsetX: 50,
			OffsetY: 7,
		}); err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}

		f.SetCellValue("Sheet1", "f"+strconv.Itoa(lenProductsArr+21), "Telefon: "+companyPhone)
		f.SetCellValue("Sheet1", "f"+strconv.Itoa(lenProductsArr+22), "IMO: "+companyPhone)
		f.SetCellValue("Sheet1", "f"+strconv.Itoa(lenProductsArr+23), "Instagram: "+instagram)
		f.SetCellValue("Sheet1", "f"+strconv.Itoa(lenProductsArr+24), "Mail: "+email)
		f.SetCellValue("Sheet1", "a"+strconv.Itoa(lenProductsArr+26), "Sargyt No: "+strconv.Itoa(int(orderNumber)))
		f.SetCellValue("Sheet1", "a"+strconv.Itoa(lenProductsArr+29), "Ady: "+order.FullName)
		f.SetCellValue("Sheet1", "a"+strconv.Itoa(lenProductsArr+30), "Telefon nomer: "+order.PhoneNumber)
		f.SetCellValue("Sheet1", "a"+strconv.Itoa(lenProductsArr+31), "Salgy: "+order.Address)
		f.SetCellValue("Sheet1", "a"+strconv.Itoa(lenProductsArr+32), "Bellik: "+order.CustomerMark)
		f.SetCellValue("Sheet1", "e"+strconv.Itoa(lenProductsArr+29), "Sargyt edilen senesi: "+createdAt)
		f.SetCellValue("Sheet1", "e"+strconv.Itoa(lenProductsArr+30), "Eltip bermeli wagty: "+order.OrderTime)
		f.SetCellValue("Sheet1", "e"+strconv.Itoa(lenProductsArr+31), "Toleg sekili: "+paymentType)
		f.SetCellValue("Sheet1", "e"+strconv.Itoa(lenProductsArr+32), "Eltip bermek hyzmaty: "+strconv.FormatFloat(pkg.RoundFloat(order.ShippingPrice, 2), 'f', -1, 64)+" TMT")
		f.SetCellValue("Sheet1", "e"+strconv.Itoa(lenProductsArr+33), "Jemi: "+strconv.FormatFloat(pkg.RoundFloat(order.TotalPrice, 2), 'f', -1, 64)+" TMT")

		for i := 0; i < lenProductsArr; i++ {
			f.InsertRows("Sheet1", 37+lenProductsArr, 1)
			f.MergeCell("Sheet1", "a"+strconv.Itoa(lenProductsArr+37), "e"+strconv.Itoa(lenProductsArr+37))
			f.SetCellStyle("Sheet1", "a"+strconv.Itoa(lenProductsArr+37), "a"+strconv.Itoa(lenProductsArr+37), style2)
			f.SetCellStyle("Sheet1", "b"+strconv.Itoa(lenProductsArr+37), "b"+strconv.Itoa(lenProductsArr+37), style2)
			f.SetCellStyle("Sheet1", "c"+strconv.Itoa(lenProductsArr+37), "c"+strconv.Itoa(lenProductsArr+37), style2)
			f.SetCellStyle("Sheet1", "d"+strconv.Itoa(lenProductsArr+37), "d"+strconv.Itoa(lenProductsArr+37), style2)
			f.SetCellStyle("Sheet1", "e"+strconv.Itoa(lenProductsArr+37), "e"+strconv.Itoa(lenProductsArr+37), style2)
			f.SetCellStyle("Sheet1", "f"+strconv.Itoa(lenProductsArr+37), "f"+strconv.Itoa(lenProductsArr+37), style)
			f.SetCellStyle("Sheet1", "g"+strconv.Itoa(lenProductsArr+37), "g"+strconv.Itoa(lenProductsArr+37), style)
			f.SetCellStyle("Sheet1", "h"+strconv.Itoa(lenProductsArr+37), "h"+strconv.Itoa(lenProductsArr+37), style)
			f.SetCellStyle("Sheet1", "i"+strconv.Itoa(lenProductsArr+37), "h"+strconv.Itoa(lenProductsArr+37), style)
		}

		for k, v2 := range products {
			f.SetCellValue("Sheet1", "a"+strconv.Itoa(37+k+lenProductsArr), v2.Name)
			f.SetCellValue("Sheet1", "f"+strconv.Itoa(37+k+lenProductsArr), v2.Amount)
			f.SetCellValue("Sheet1", "g"+strconv.Itoa(37+k+lenProductsArr), v2.Code.String)
			f.SetCellValue("Sheet1", "h"+strconv.Itoa(37+k+lenProductsArr), v2.Price)
			f.SetCellValue("Sheet1", "i"+strconv.Itoa(37+k+lenProductsArr), float64(v2.Amount)*v2.Price)
		}

		f.SetCellValue("Sheet1", "g"+strconv.Itoa(38+2*lenProductsArr), strconv.FormatFloat(totalPrice, 'f', -1, 64)+" TMT")
	}

	// tayyar bolan excel fayl uploads papka yazdyrylyar
	if err := f.SaveAs(pkg.ServerPath + "uploads/orders/" + strconv.Itoa(int(orderNumber)) + ".xlsx"); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	// excel fayl tayyar bolanson sargydyn fayly hokmunde baza yazdyrylyar
	_, err = db.Exec(context.Background(), "UPDATE orders SET excel = $1 WHERE id = $2", "uploads/orders/"+strconv.Itoa(int(orderNumber))+".xlsx", order_id)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "success",
	})

}

func GetOrders(c *gin.Context) {
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	// GET language id
	langID, err := backController.GetLangID("tm")
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	// get limit from param
	limitStr := c.Param("limit")
	if limitStr == "" {
		helpers.HandleError(c, 400, "limit is required")
		return
	}
	limit, err := strconv.ParseUint(limitStr, 10, 64)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	// get page from param
	pageStr := c.Param("page")
	if pageStr == "" {
		helpers.HandleError(c, 400, "page is required")
		return
	}
	page, err := strconv.ParseUint(pageStr, 10, 64)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	offset := limit * (page - 1)
	countOfOrders := 0

	statusStr := c.Query("status")
	status, err := strconv.ParseBool(statusStr)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	searchQuery := c.Query("search")
	search := fmt.Sprintf("%%%s%%", searchQuery)
	var countAllCustomer, rowsCustomerID pgx.Rows
	var rowsOrderQuery string
	if status {
		if searchQuery == "" {
			countAllCustomer, err = db.Query(context.Background(), "SELECT customer_id FROM orders WHERE deleted_at IS NOT NULL")
			if err != nil {
				helpers.HandleError(c, 400, err.Error())
				return
			}
		} else {
			countAllCustomer, err = db.Query(context.Background(), "SELECT c.id FROM orders o INNER JOIN customers c ON o.customer_id=c.id WHERE c.deleted_at IS NOT NULL AND o.deleted_at IS NOT NULL AND c.phone_number LIKE $1", search)
			if err != nil {
				helpers.HandleError(c, 400, err.Error())
				return
			}
		}
	} else {
		if searchQuery == "" {
			countAllCustomer, err = db.Query(context.Background(), "SELECT customer_id FROM orders WHERE deleted_at IS NULL")
			if err != nil {
				helpers.HandleError(c, 400, err.Error())
				return
			}
		} else {
			countAllCustomer, err = db.Query(context.Background(), "SELECT c.id FROM orders o INNER JOIN customers c ON o.customer_id=c.id WHERE c.deleted_at IS NULL AND o.deleted_at IS NULL AND c.phone_number LIKE $1", search)
			if err != nil {
				helpers.HandleError(c, 400, err.Error())
				return
			}
		}
	}
	defer countAllCustomer.Close()

	for countAllCustomer.Next() {
		countOfOrders++
	}

	if status {
		if searchQuery == "" {
			rowsCustomerID, err = db.Query(context.Background(), "SELECT customer_id FROM orders WHERE deleted_at IS NOT NULL GROUP BY customer_id")
			if err != nil {
				helpers.HandleError(c, 400, err.Error())
				return
			}
		} else {
			rowsCustomerID, err = db.Query(context.Background(), "SELECT c.id FROM orders o INNER JOIN customers c ON c.id=o.customer_id WHERE c.deleted_at IS NOT NULL AND o.deleted_at IS NOT NULL AND c.phone_number LIKE $1 GROUP BY c.id", search)
			if err != nil {
				helpers.HandleError(c, 400, err.Error())
				return
			}
		}
	} else {
		if searchQuery == "" {
			rowsCustomerID, err = db.Query(context.Background(), "SELECT customer_id FROM orders WHERE deleted_at IS NULL GROUP BY customer_id")
			if err != nil {
				helpers.HandleError(c, 400, err.Error())
				return
			}
		} else {
			rowsCustomerID, err = db.Query(context.Background(), "SELECT c.id FROM orders o INNER JOIN customers c ON c.id=o.customer_id WHERE c.deleted_at IS NULL AND o.deleted_at IS NULL AND c.phone_number LIKE $1 GROUP BY c.id", search)
			if err != nil {
				helpers.HandleError(c, 400, err.Error())
				return
			}
		}
	}
	defer rowsCustomerID.Close()

	var customerIDs []string
	for rowsCustomerID.Next() {
		var customerID string
		if err := rowsCustomerID.Scan(&customerID); err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}
		customerIDs = append(customerIDs, customerID)
	}

	var orders []OrderForAdmin
	if status {
		rowsOrderQuery = `SELECT customer_id,id,customer_mark,order_time,payment_type,total_price,shipping_price,excel,address,TO_CHAR(created_at, 'DD.MM.YYYY'),order_number FROM orders WHERE customer_id = ANY($1) AND deleted_at IS NOT NULL ORDER BY updated_at DESC LIMIT $2 OFFSET $3`
	} else {
		rowsOrderQuery = `SELECT customer_id,id,customer_mark,order_time,payment_type,total_price,shipping_price,excel,address,TO_CHAR(created_at, 'DD.MM.YYYY'),order_number FROM orders WHERE customer_id = ANY($1) AND deleted_at IS NULL ORDER BY updated_at DESC LIMIT $2 OFFSET $3`
	}

	rowsOrder, err := db.Query(context.Background(), rowsOrderQuery, pq.Array(customerIDs), limit, offset)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer rowsOrder.Close()

	for rowsOrder.Next() {
		var order OrderForAdmin
		var paymentType uint8
		if err := rowsOrder.Scan(&order.CustomerID, &order.ID, &order.CustomerMark, &order.OrderTime, &paymentType, &order.TotalPrice, &order.ShippingPrice, &order.Excel, &order.Address, &order.CreatedAt, &order.OrderNumber); err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}

		db.QueryRow(context.Background(), "SELECT name FROM payment_types pt INNER JOIN languages l ON l.id=pt.lang_id WHERE pt.value = $1 AND l.name_short = 'tm' AND pt.deleted_at IS NULL", paymentType).Scan(&order.PaymentType)
		if order.PaymentType == "" {
			helpers.HandleError(c, 404, "payment type not found")
			return
		}

		rowCustomer, err := db.Query(context.Background(), "SELECT full_name,phone_number FROM customers WHERE deleted_at IS NULL AND id = $1", order.CustomerID)
		if err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}
		defer rowCustomer.Close()

		for rowCustomer.Next() {
			var customerFullName null.String
			if err := rowCustomer.Scan(&customerFullName, &order.PhoneNumber); err != nil {
				helpers.HandleError(c, 400, err.Error())
				return
			}

			order.FullName = customerFullName.String
		}

		rowsOrderedProducts, err := db.Query(context.Background(), "SELECT product_id,quantity_of_product FROM ordered_products WHERE order_id = $1 AND deleted_at IS NULL", order.ID)
		if err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}
		defer rowsOrderedProducts.Close()

		var products []ProductForAdmin
		for rowsOrderedProducts.Next() {
			var product ProductForAdmin
			if err := rowsOrderedProducts.Scan(&product.ID, &product.QuantityOfProduct); err != nil {
				helpers.HandleError(c, 400, err.Error())
				return
			}

			db.QueryRow(context.Background(), "SELECT brend_id,price,old_price,amount,limit_amount,is_new,main_image,benefit,code FROM products WHERE id = $1 AND deleted_at IS NULL", product.ID).Scan(&product.BrendID, &product.Price, &product.OldPrice, &product.Amount, &product.LimitAmount, &product.IsNew, &product.MainImage, &product.Benefit, &product.Code)

			if product.Benefit.Float64 != 0 {
				product.Price = product.Price + (product.Price*product.Benefit.Float64)/100
				product.OldPrice = product.OldPrice + (product.OldPrice*product.Benefit.Float64)/100
			}

			if product.OldPrice != 0 {
				product.Percentage = -math.Round(((product.OldPrice - product.Price) * 100) / product.OldPrice)
			} else {
				product.Percentage = 0
			}

			var trProduct models.TranslationProduct
			db.QueryRow(context.Background(), "SELECT name,description FROM translation_product WHERE product_id = $1 AND lang_id = $2 AND deleted_at IS NULL", product.ID, langID).Scan(&trProduct.Name, &trProduct.Description)
			product.Translations = trProduct
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
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	// get id of language from request parameter
	orderID := c.Param("id")

	statusStr := c.Query("status")
	status, err := strconv.ParseBool(statusStr)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	// check orderID
	var order_id string
	if status {
		db.QueryRow(context.Background(), "SELECT id FROM orders WHERE id = $1 AND deleted_at IS NOT NULL", orderID).Scan(&order_id)
	} else {
		db.QueryRow(context.Background(), "SELECT id FROM orders WHERE id = $1 AND deleted_at IS NULL", orderID).Scan(&order_id)
	}

	if order_id == "" {
		helpers.HandleError(c, 404, "order not found")
		return
	}

	if status {
		_, err := db.Exec(context.Background(), "UPDATE orders SET deleted_at = NULL WHERE id = $1", orderID)
		if err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}
	} else {
		_, err := db.Exec(context.Background(), "UPDATE orders SET deleted_at = now() WHERE id = $1", orderID)
		if err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}
	}

	if status {
		c.JSON(http.StatusOK, gin.H{
			"status":  true,
			"message": "the order has been moved to the unconfirmed table",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "order confirmed",
	})
}

func GetCustomerOrders(c *gin.Context) {
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	customerID, err := pkg.ValidateMiddlewareData(c, "customer_id")
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	// get limit from param
	limitStr := c.Param("limit")
	if limitStr == "" {
		helpers.HandleError(c, 400, "limit is required")
		return
	}
	limit, err := strconv.ParseUint(limitStr, 10, 64)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	// get page from param
	pageStr := c.Param("page")
	if pageStr == "" {
		helpers.HandleError(c, 400, "page is required")
		return
	}
	page, err := strconv.ParseUint(pageStr, 10, 64)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	offset := limit * (page - 1)

	// programmany ulanyp otyran musderinin sargytlarynyn sanyny alyas frontda pagination ucin
	var countOfOrders uint
	db.QueryRow(context.Background(), "SELECT COUNT(id) FROM orders WHERE customer_id = $1", customerID).Scan(&countOfOrders)

	// musderinin sargytlaryny alyas bazadan
	rowsOrders, err := db.Query(context.Background(), "SELECT id,TO_CHAR(created_at, 'DD.MM.YYYY'),total_price FROM orders WHERE customer_id = $1 ORDER BY created_at ASC LIMIT $2 OFFSET $3", customerID, limit, offset)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer rowsOrders.Close()

	var orders []GetOrder
	for rowsOrders.Next() {
		var order GetOrder
		if err := rowsOrders.Scan(&order.ID, &order.Date, &order.TotalPrice); err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}

		rowsOrderedProducts, err := db.Query(context.Background(), "SELECT product_id,quantity_of_product FROM ordered_products WHERE order_id = $1 AND deleted_at IS NULL", order.ID)
		if err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}
		defer rowsOrderedProducts.Close()

		var products []ProductOfCart
		for rowsOrderedProducts.Next() {
			var product ProductOfCart
			if err := rowsOrderedProducts.Scan(&product.ID, &product.QuantityOfProduct); err != nil {
				helpers.HandleError(c, 400, err.Error())
				return
			}

			rowProduct, err := db.Query(context.Background(), "SELECT brend_id,price,old_price,amount,limit_amount,is_new,main_image,benefit,code FROM products WHERE id = $1 AND deleted_at IS NULL", product.ID)
			if err != nil {
				helpers.HandleError(c, 400, err.Error())
				return
			}
			defer rowProduct.Close()

			for rowProduct.Next() {
				if err := rowProduct.Scan(&product.BrendID, &product.Price, &product.OldPrice, &product.Amount, &product.LimitAmount, &product.IsNew, &product.MainImage, &product.Benefit, &product.Code); err != nil {
					helpers.HandleError(c, 400, err.Error())
					return
				}
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
				db.QueryRow(context.Background(), "SELECT name,description FROM translation_product WHERE product_id = $1 AND lang_id = $2 AND deleted_at IS NULL", product.ID, v.ID).Scan(&trProduct.Name, &trProduct.Description)
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

// 		rowProduct, err := db.Query(context.Background(),"SELECT id FROM products WHERE id = $1 AND deleted_at IS NULL", v)
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
// 	rowOrders, err := db.Query(context.Background(),"SELECT id,brend_id,price,old_price,amount,limit_amount,is_new FROM products WHERE id = ANY($1) AND deleted_at IS NULL", pq.Array(productIds.IDS))
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

// 		rowMainImage, err := db.Query(context.Background(),"SELECT small,medium,large FROM main_image WHERE product_id = $1 AND deleted_at IS NULL", product.ID)
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

// 		rowsImages, err := db.Query(context.Background(),"SELECT small,large FROM images WHERE product_id = $1 AND deleted_at IS NULL", product.ID)
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

// 		rowsLang, err := db.Query(context.Background(),"SELECT id,name_short FROM languages WHERE deleted_at IS NULL")
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

// 			rowTrProduct, err := db.Query(context.Background(),"SELECT name,description FROM translation_product WHERE product_id = $1 AND lang_id = $2 AND deleted_at IS NULL", product.ID, v.ID)
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
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	orderID := c.Param("id")

	var order_id, excel string
	db.QueryRow(context.Background(), "SELECT id,excel FROM orders WHERE id = $1", orderID).Scan(&order_id, &excel)

	if order_id == "" {
		helpers.HandleError(c, 404, "order not found")
		return
	}

	if err := os.Remove(pkg.ServerPath + excel); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	rowsOrderedProduct, err := db.Query(context.Background(), "SELECT product_id,quantity_of_product FROM ordered_products WHERE order_id = $1", order_id)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer rowsOrderedProduct.Close()

	var orderedProducts []models.OrderedProducts
	for rowsOrderedProduct.Next() {
		var orderedProduct models.OrderedProducts
		if err := rowsOrderedProduct.Scan(&orderedProduct.ProductID, &orderedProduct.QuantityOfProduct); err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}
		orderedProducts = append(orderedProducts, orderedProduct)
	}

	for _, v := range orderedProducts {
		_, err := db.Exec(context.Background(), "UPDATE products SET amount = amount + $1 WHERE id = $2", v.QuantityOfProduct, v.ProductID)
		if err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}
	}

	_, err = db.Exec(context.Background(), "DELETE FROM orders WHERE id = $1", order_id)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "order returned success",
	})
}
