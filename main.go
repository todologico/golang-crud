package main

//-------------------------------------------------------------------------------
// MVC basic structure with golang, for learning purposes only - Arturo - 2024
//-------------------------------------------------------------------------------

import (
	"fmt"
	"go-crud/controllers"
	"log"
	"net/http"
)

//--------------------------------------
// main method routing
//--------------------------------------
func main() {

    http.HandleFunc("/", controllers.ProductHandler)

    http.HandleFunc("/services", services)

    log.Fatal(http.ListenAndServe(":8080", nil))
}

//--------------------------------------
// services
//--------------------------------------
func services(w http.ResponseWriter, r *http.Request) {

    fmt.Fprintf(w, "services url")
}

