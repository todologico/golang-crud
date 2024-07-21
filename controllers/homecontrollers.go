package controllers

import (
	"html/template"
	"net/http"
	"path/filepath"

	"go-crud/models"
)

// HomeHandler
func HomeHandler(w http.ResponseWriter, r *http.Request) {

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
