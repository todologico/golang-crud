package controllers

import (
	"html/template"
	"net/http"
	"path/filepath"
	"strconv"
	"github.com/todologico/golang-crud/utilities"
	"github.com/todologico/golang-crud/models"
)

//------------------------------------------------------------

// Product Handler
func ListProduct(w http.ResponseWriter, r *http.Request) {

	products, err := models.GetProducts()
	if err != nil {
		http.Error(w, "Error retrieving products: "+err.Error(), http.StatusInternalServerError)
		return
	}

	tmplPath := filepath.Join("views", "products.html")
	tmpl, err := template.ParseFiles(tmplPath)

	//sending a header to the browser
	w.Header().Set("Content-Type", "text/html")
	

	if err := tmpl.Execute(w, products); err != nil {
		http.Error(w, "Error rendering template: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

//------------------------------------------------------------

// InsertProduct
func InsertProduct(w http.ResponseWriter, r *http.Request) {

	tmplPath := filepath.Join("views", "insert.html")
	tmpl, err := template.ParseFiles(tmplPath)

	if err != nil {
		http.Error(w, "Error loading template: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")

	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, "Error rendering template: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

//------------------------------------------------------------

// InsertProcessProduct Handler
func InsertProcessProduct(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	prodToken, err := utils.GenerateRandomToken(50)

	if err != nil {
		http.Error(w, "Error generating token: "+err.Error(), http.StatusInternalServerError)
		return
	}

	prodName := r.FormValue("prod_name")
	prodQuantityStr := r.FormValue("prod_quantity")
	prodQuantity, err := strconv.Atoi(prodQuantityStr)

	if err != nil {
		http.Error(w, "Invalid product quantity", http.StatusBadRequest)
		return
	}

	err = models.InsertProduct(prodName, prodQuantity, prodToken)

	if err != nil {
		http.Error(w, "Failed to insert product: "+err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

//------------------------------------------------------------

// EditProduct
func EditProduct(w http.ResponseWriter, r *http.Request) {

	idStr := r.URL.Query().Get("id")
	prodToken := r.URL.Query().Get("prod_token")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	product, err := models.GetProduct(id, prodToken)

	if err != nil {
		http.Error(w, "Error retrieving product: "+err.Error(), http.StatusInternalServerError)
		return
	}

	tmplPath := filepath.Join("views", "edit.html")
	tmpl, err := template.ParseFiles(tmplPath)

	if err != nil {
		http.Error(w, "Error loading template: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")

	if err := tmpl.Execute(w, product); err != nil {
		http.Error(w, "Error rendering template: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

//------------------------------------------------------------

// EditProcessProduct
func EditProcessProduct(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.FormValue("id")
	prodToken := r.FormValue("prod_token")
	prodName := r.FormValue("prod_name")
	prodQuantityStr := r.FormValue("prod_quantity")
	prodQuantity, err := strconv.Atoi(prodQuantityStr)

	if err != nil {
		http.Error(w, "Invalid product quantity", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	err = models.UpdateProduct(id, prodToken, prodName, prodQuantity)

	if err != nil {
		http.Error(w, "Failed to update product: "+err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

//------------------------------------------------------------

// DeleteProduct
func DeleteProduct(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.FormValue("id")
	prodToken := r.FormValue("prod_token")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	err = models.DeleteProduct(id, prodToken)

	if err != nil {
		http.Error(w, "Failed to delete product: "+err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
