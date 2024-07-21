package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

//var db *sql.DB

//--------------------------------------
// home
//--------------------------------------
func home(w http.ResponseWriter, r *http.Request) {  
    
    

    dsn := "golang:00000000@tcp(golang-db:3310)/golang"



    // Abrir una conexión a la base de datos
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        fmt.Println("Error al conectar a la base de datos:", err)
        return
    }
    defer db.Close()

    // Verificar la conexión
    if err := db.Ping(); err != nil {
        fmt.Println("Error al verificar la conexión:", err)
        return
    }

    // Ejecutar una consulta
    rows, err := db.Query("SELECT id,name,quantity FROM products")
    if err != nil {
        fmt.Println("Error al ejecutar la consulta:", err)
        return
    }
    defer rows.Close()

    // Leer los resultados
    for rows.Next() {
        var id int
        var name string
        var quantity int
        
        err := rows.Scan(&id, &name, &quantity)
        if err != nil {
            fmt.Println("Error al leer una fila:", err)
            return
        }
        
        fmt.Printf("ID: %d, Nombre: %s, Edad: %d\n", id, name, quantity)
    }

    // Verificar si ocurrieron errores durante la iteración de las filas
    if err := rows.Err(); err != nil {
        fmt.Println("Error durante la iteración de filas:", err)
    }


    fmt.Fprintf(w, "home url 2")

    
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