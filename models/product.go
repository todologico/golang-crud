// models/product.go
package models

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// Database connection details
const dsn = "golang:00000000@tcp(golang-db:3310)/golang"

// OpenDB opens a connection to the database
func OpenDB() (*sql.DB, error) {
    
    db, err := sql.Open("mysql", dsn)
   
    if err != nil {
        return nil, err
    }

    if err := db.Ping(); err != nil {
        return nil, err
    }

    return db, nil
}

// products from the database
func GetProducts(db *sql.DB) ([]Product, error) {
    rows, err := db.Query("SELECT id, name, quantity FROM products")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var products []Product
    for rows.Next() {
        var product Product
        if err := rows.Scan(&product.ID, &product.Name, &product.Quantity); err != nil {
            return nil, err
        }
        products = append(products, product)
    }

    if err := rows.Err(); err != nil {
        return nil, err
    }

    return products, nil
}

// Product represents a product record in the database
type Product struct {
    ID       int
    Name     string
    Quantity int
}
