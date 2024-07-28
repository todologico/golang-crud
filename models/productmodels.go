package models

import (
	"database/sql"
)

// Product represents a product record in the database
type Product struct {
    ID       int
    Name     string
    Quantity int
}

// GetProducts retrieves products from the database
func GetProducts(db *sql.DB) ([]Product, error) {

    rows, err := db.Query("SELECT id, name, quantity FROM products")

    if err != nil {
        return nil, err
    }

    defer rows.Close()

    var products []Product

    //paso las rows al Product struct
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
