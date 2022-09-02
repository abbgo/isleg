package models

import (
	"errors"
	"github/abbgo/isleg/isleg-backend/config"

	"github.com/google/uuid"
)

type Like struct {
	ID         uuid.UUID `json:"id"`
	ProductID  uuid.UUID `json:"product_id"`
	CustomerID uuid.UUID `json:"customer_id"`
	CreatedAt  string    `json:"-"`
	UpdatedAt  string    `json:"-"`
	DeletedAt  string    `json:"-"`
}

func ValidateCustomerLike(customerID string, productIDs []string) error {

	db, err := config.ConnDB()
	if err != nil {
		return nil
	}
	defer db.Close()

	if customerID == "" {
		return errors.New("customer_id is required")
	}

	rowCustomer, err := db.Query("SELECT id FROM customers WHERE id = $1 AND deleted_at IS NULL", customerID)
	if err != nil {
		return err
	}

	var customer_id string
	for rowCustomer.Next() {
		if err := rowCustomer.Scan(&customer_id); err != nil {
			return err
		}
	}

	if customer_id == "" {
		return errors.New("customer does not exist")
	}

	if len(productIDs) == 0 {
		return errors.New("product_id is required")
	}

	for _, productID := range productIDs {
		rowProduct, err := db.Query("SELECT id FROM products WHERE id = $1 AND deleted_at IS NULL", productID)
		if err != nil {
			return err
		}

		var product_id string
		for rowProduct.Next() {
			if err := rowProduct.Scan(&product_id); err != nil {
				return err
			}
		}

		if product_id == "" {
			return errors.New("product does not exist")
		}
		rows, err := db.Query("SELECT product_id FROM likes WHERE customer_id = $1 AND deleted_at IS NULL", customerID)
		if err != nil {
			return err
		}

		var product_ids []string

		for rows.Next() {
			if err := rows.Scan(&product_id); err != nil {
				return err
			}
			product_ids = append(product_ids, product_id)
		}

		for _, v := range product_ids {
			if productID == v {
				return errors.New("this product already exists in this customer")
			}
		}

	}

	return nil

}
