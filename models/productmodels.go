package models

import (
	"database/sql"
	"errors"
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

//---------------------------------------------------

// GetProduct retrieves a product from the database based on id and prod_token
func GetProduct(db *sql.DB, id int, prod_token string) (*Product, error) {

	var product Product

	query := "SELECT id, prod_name, prod_quantity, prod_token FROM products WHERE id = ? AND prod_token = ?"

	err := db.QueryRow(query, id, prod_token).Scan(&product.Id, &product.Prod_name, &product.Prod_quantity, &product.Prod_token)

	if err != nil {
		return nil, err
	}

	return &product, nil
}

// ---------------------------------------------------

// Delete a product by id and prod_token
func DeleteProduct(db *sql.DB, id int, prod_token string) error {

	_, err := db.Query("DELETE FROM products WHERE id = ? AND prod_token = ?", id, prod_token)

	return err
}

// ---------------------------------------------------

// Update a product by id and prod_token
func UpdateProduct(db *sql.DB, id int, prod_token string, prodName string, prodQuantity int) error {

	query := "UPDATE products SET prod_name = ?, prod_quantity = ? WHERE id = ? AND prod_token = ?"

	result, err := db.Exec(query, prodName, prodQuantity, id, prod_token)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("no rows affected")
	}

	return nil
}
