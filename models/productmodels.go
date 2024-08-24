package models

import (
	"database/sql"
	"errors"
)

// My product represents a record in the database
type Product struct {
	Id            int
	Prod_name     string
	Prod_quantity int
	Prod_token    string
}

//------------------------------------------------------------

// GetProducts retrieves products from the database
func GetProducts() ([]Product, error) {
	db, err := OpenDB()
	if err != nil {
		return nil, err
	}
	defer CloseDB(db)

	rows, err := db.Query("SELECT id, prod_name, prod_quantity, prod_token FROM products")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []Product

	for rows.Next() {

		var single Product
		
		if err := rows.Scan(&single.Id, &single.Prod_name, &single.Prod_quantity, &single.Prod_token); err != nil {
			return nil, err
		}

		products = append(products, single)
	}

	if err := rows.Err(); err != nil {

		return nil, err
	}

	return products, nil
}

//------------------------------------------------------------

// GetProduct retrieves a product from the database based on id and prod_token
func GetProduct(id int, prod_token string) (*Product, error) {

	db, err := OpenDB()
	if err != nil {
		return nil, err
	}
	defer CloseDB(db)

	var product Product

	query := "SELECT id, prod_name, prod_quantity, prod_token FROM products WHERE id = ? AND prod_token = ?"
	err = db.QueryRow(query, id, prod_token).Scan(&product.Id, &product.Prod_name, &product.Prod_quantity, &product.Prod_token)

	if err != nil {
		return nil, err
	}

	return &product, nil
}

//------------------------------------------------------------

// DeleteProduct deletes a product by id and prod_token
func DeleteProduct(id int, prod_token string) error {

	db, err := OpenDB()
	if err != nil {
		return err
	}
	defer CloseDB(db)

	_, err = db.Exec("DELETE FROM products WHERE id = ? AND prod_token = ?", id, prod_token)

	return err
}

//------------------------------------------------------------

// UpdateProduct updates a product by id and prod_token
func UpdateProduct(id int, prod_token string, prodName string, prodQuantity int) error {

	db, err := OpenDB()
	if err != nil {
		return err
	}
	defer CloseDB(db)

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

// InsertProduct inserts a new product into the database
func InsertProduct(prodName string, prodQuantity int, prodToken string) error {

	db, err := OpenDB()
	if err != nil {
		return err
	}
	defer CloseDB(db)

	query := "INSERT INTO products (prod_name, prod_quantity, prod_token) VALUES (?, ?, ?)"
	result, err := db.Exec(query, prodName, prodQuantity, prodToken)

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
