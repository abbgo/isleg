package models

type Cart struct {
	ID                string `json:"id,omitempty"`
	ProductID         string `json:"product_id,omitempty"`
	CustomerID        string `json:"customer_id,omitempty"`
	QuantityOfProduct int    `json:"quantity_of_product,omitempty"`
	CreatedAt         string `json:"-"`
	UpdatedAt         string `json:"-"`
	DeletedAt         string `json:"-"`
}

// func ValidateCustomerBasket(customerID, productID, quantityOfProduct string) (int, error) {

// 	db, err := config.ConnDB()
// 	if err != nil {
// 		return 0, nil
// 	}
// 	defer db.Close()

// 	if customerID == "" {
// 		return 0, errors.New("customer_id is required")
// 	}

// 	rowCustomer, err := db.Query("SELECT id FROM customers WHERE id = $1 AND deleted_at IS NULL", customerID)
// 	if err != nil {
// 		return 0, err
// 	}

// 	var customer_id string
// 	for rowCustomer.Next() {
// 		if err := rowCustomer.Scan(&customer_id); err != nil {
// 			return 0, err
// 		}
// 	}

// 	if customer_id == "" {
// 		return 0, errors.New("customer does not exist")
// 	}

// 	if productID == "" {
// 		return 0, errors.New("product_id is required")
// 	}

// 	rowProduct, err := db.Query("SELECT id FROM products WHERE id = $1 AND deleted_at IS NULL", productID)
// 	if err != nil {
// 		return 0, err
// 	}

// 	var product_id string
// 	for rowProduct.Next() {
// 		if err := rowProduct.Scan(&product_id); err != nil {
// 			return 0, err
// 		}
// 	}

// 	if product_id == "" {
// 		return 0, errors.New("product does not exist")
// 	}

// 	if quantityOfProduct == "" {
// 		return 0, errors.New("quantity of product required")
// 	}

// 	quantityInt, err := strconv.Atoi(quantityOfProduct)
// 	if err != nil {
// 		return 0, err
// 	}

// 	if quantityInt < 1 {
// 		return 0, errors.New("quantity of product cannot be less than 1")
// 	}

// 	return quantityInt, nil

// }
