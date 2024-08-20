package models

import (
	"database/sql"
)

// Product represents a record in the database
type Product struct {
	Id            int
	Prod_name     string
	Prod_quantity int
	Prod_token    string
}

// GetProducts retrieves products from the database
func GetProducts(db *sql.DB) ([]Product, error) {

	rows, err := db.Query("SELECT id, prod_name, prod_quantity, prod_token FROM products")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	//struct slice
	var products []Product

	//struct
	var single Product

	//sending rows to the struct product
	for rows.Next() {

		if err := rows.Scan(&single.Id, &single.Prod_name, &single.Prod_quantity, &single.Prod_token); err != nil {
			return nil, err
		}

		//slice
		products = append(products, single)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}

// ---------------------------------------------------

// Delete a product from the database by id and prod_token
func DeleteProduct(db *sql.DB, id int, prodToken string) error {
	_, err := db.Query("DELETE FROM products WHERE id = ? AND prod_token = ?", id, prodToken)
	return err
}
