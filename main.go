package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

//--------------------------------------
// product list with database access (mariadb)
//--------------------------------------
func home(w http.ResponseWriter, r *http.Request) {
    
    //secrets - only for dev
    dsn := "golang:00000000@tcp(golang-db:3310)/golang"

    // database connection
    db, err := sql.Open("mysql", dsn)

    if err != nil {
        http.Error(w, "Error connecting to the database: "+err.Error(), http.StatusInternalServerError)
        return
    }
    defer db.Close()

    // Confirm the connection
    if err := db.Ping(); err != nil {
        http.Error(w, "Error verifying the connection: "+err.Error(), http.StatusInternalServerError)
        return
    }

    // Query
    products, err := db.Query("SELECT id, name, quantity FROM products")
    if err != nil {
        http.Error(w, "Error executing the query: "+err.Error(), http.StatusInternalServerError)
        return
    }
    defer products.Close()

    // HTML BUFFER
    var htmlResponse string

    htmlResponse += "<html><body><h1>Product List with Golang</h1><table border='1'><tr><th>ID</th><th>Name</th><th>Quantity</th></tr>"

    // HTML Table to display results
    for products.Next() {
        var id int
        var name string
        var quantity int
        
        err := products.Scan(&id, &name, &quantity)
        if err != nil {
            http.Error(w, "Error reading row: "+err.Error(), http.StatusInternalServerError)
            return
        }
        
        htmlResponse += fmt.Sprintf("<tr><td>%d</td><td>%s</td><td>%d</td></tr>", id, name, quantity)
    }

    // Verify for errors while iterating over products
    if err := products.Err(); err != nil {
        http.Error(w, "Error during row iteration: "+err.Error(), http.StatusInternalServerError)
        return
    }

    htmlResponse += "</table></body></html>"

    // Response
    w.Header().Set("Content-Type", "text/html")
    w.Write([]byte(htmlResponse))
}

//--------------------------------------
// services
//--------------------------------------
func services(w http.ResponseWriter, r *http.Request) {

    fmt.Fprintf(w, "services url")
}


//--------------------------------------
// products
//--------------------------------------
func products(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "products url!\n")        

    numeros := make([]int, 5)

    numeros = append(numeros, 6)
    
    fmt.Println(numeros)   

}

//--------------------------------------
// main method routing
//--------------------------------------
func main() {

    http.HandleFunc("/", home)

    http.HandleFunc("/services", services)

    http.HandleFunc("/products", products)

    log.Fatal(http.ListenAndServe(":8080", nil))
}