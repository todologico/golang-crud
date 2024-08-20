package main

//-------------------------------------------------------------------------------
// MVC basic structure with golang, for learning purposes only - Arturo - 2024
//-------------------------------------------------------------------------------

import (
	"fmt"
	"log"
	"net/http"

	"github.com/todologico/golang-crud/controllers"
)

// --------------------------------------
// main method routing
// --------------------------------------
func main() {

	http.HandleFunc("/", controllers.ProductHandler)

	http.HandleFunc("/delete", controllers.DeleteProductHandler)

	http.HandleFunc("/services", services)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

// --------------------------------------
// services
// --------------------------------------
func services(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "services url")
}
