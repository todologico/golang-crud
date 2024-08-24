package controllers

import (
	"html/template"
	"net/http"
	"path/filepath"
	"strconv"
	"github.com/todologico/golang-crud/utilities"
	"github.com/todologico/golang-crud/models"
)

// Product Handler
func ListProduct(w http.ResponseWriter, r *http.Request) {

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

//---------------------------------------------------------

func InsertProduct(w http.ResponseWriter, r *http.Request) {
	// HTML
	tmplPath := filepath.Join("views", "insert.html")
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		http.Error(w, "Error loading template: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Render
	w.Header().Set("Content-Type", "text/html")
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, "Error rendering template: "+err.Error(), http.StatusInternalServerError)
		return
	}
}


//---------------------------------------------------

// InsertProcessProduct Handler
func InsertProcessProduct(w http.ResponseWriter, r *http.Request) {

	//token generator utils
	prodToken, err := utils.GenerateRandomToken(50)

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Obtener los parámetros del formulario
	prodName := r.FormValue("prod_name")
	prodQuantityStr := r.FormValue("prod_quantity")

	// Convertir la cantidad del producto de cadena a entero
	prodQuantity, err := strconv.Atoi(prodQuantityStr)
	if err != nil {
		http.Error(w, "Invalid product quantity", http.StatusBadRequest)
		return
	}

	// Token fijo para pruebas maximo 50 caracteres
	//prodToken := "kfgjlñsdkfgjsdlsdkfgjlñsdkgj"

	// Conexión a la base de datos
	db, err := models.OpenDB()
	if err != nil {
		http.Error(w, "Error connecting to the database: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Insertar producto
	err = models.InsertProduct(db, prodName, prodQuantity, prodToken)
	if err != nil {
		http.Error(w, "Failed to insert product: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Redirigir a la lista de productos
	http.Redirect(w, r, "/", http.StatusSeeOther)
}


//---------------------------------------------------------

// EditProduct with the product edit page
func EditProduct(w http.ResponseWriter, r *http.Request) {

	// parameters
	idStr := r.URL.Query().Get("id")
	prod_token := r.URL.Query().Get("prod_token")

	// string to integer
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

	// get the product !!
	product, err := models.GetProduct(db, id, prod_token)
	if err != nil {
		http.Error(w, "Error retrieving product: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// HTML
	tmplPath := filepath.Join("views", "edit.html")
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		http.Error(w, "Error loading template: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Render
	w.Header().Set("Content-Type", "text/html")
	if err := tmpl.Execute(w, product); err != nil {
		http.Error(w, "Error rendering template: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

//----------------------------------------------------------

// Delete Handler
func DeleteProduct(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// post parameters
	idStr := r.FormValue("id")
	prod_token := r.FormValue("prod_token")

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
	err = models.DeleteProduct(db, id, prod_token)
	if err != nil {
		http.Error(w, "Failed to delete product: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// redirect to product list
	http.Redirect(w, r, "/", http.StatusSeeOther)

}


//----------------------------------------------------------

// EditProcessProduct Handler
func EditProcessProduct(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// post parameters
	idStr := r.FormValue("id")
	prod_token := r.FormValue("prod_token")
	prod_name := r.FormValue("prod_name")
	prod_quantityStr := r.FormValue("prod_quantity")

	// Convert the string id to int
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	// Convert the string prod_quantity to int
	prod_quantity, err := strconv.Atoi(prod_quantityStr)
	if err != nil {
		http.Error(w, "Invalid product quantity", http.StatusBadRequest)
		return
	}

	// db connection
	db, err := models.OpenDB()
	if err != nil {
		http.Error(w, "Error connecting to the database: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// update product model
	err = models.UpdateProduct(db, id, prod_token, prod_name, prod_quantity)
	if err != nil {
		http.Error(w, "Failed to update product: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// redirect to product list
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

