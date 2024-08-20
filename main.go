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

	http.HandleFunc("/", controllers.ListProduct)

	http.HandleFunc("/edit", controllers.EditProduct)

	http.HandleFunc("/delete", controllers.DeleteProduct)

	http.HandleFunc("/services", services)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

// --------------------------------------
// services
// --------------------------------------
func services(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "services url another package")
}
