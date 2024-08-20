package controllers

import (
	"html/template"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/todologico/golang-crud/models"
)

// Product Handler
func ProductHandler(w http.ResponseWriter, r *http.Request) {

	db, err := models.OpenDB()

	if err != nil {
		http.Error(w, "Error connecting to the database: "+err.Error(), http.StatusInternalServerError)
		return
	}

	defer db.Close()

	products, err := models.GetProducts(db)
	if err != nil {
		http.Error(w, "Error executing the query: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Load HTML template
	tmplPath := filepath.Join("views", "products.html")
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		http.Error(w, "Error loading template: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Render template with products data
	w.Header().Set("Content-Type", "text/html")
	if err := tmpl.Execute(w, products); err != nil {
		http.Error(w, "Error rendering template: "+err.Error(), http.StatusInternalServerError)
		return
	}

}

//----------------------------------------------------------

// Delete Handler
func DeleteProductHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// get parameters
	idStr := r.FormValue("id")
	prodToken := r.FormValue("prod_token")

	// Cconvert the string id to int
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	// db connection
	db, err := models.OpenDB()
	if err != nil {
		http.Error(w, "Error connecting to the database: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// delete product model
	err = models.DeleteProduct(db, id, prodToken)
	if err != nil {
		http.Error(w, "Failed to delete product: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// redirect to product list
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
