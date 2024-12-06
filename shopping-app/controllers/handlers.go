package controllers

import (
	"database/sql"
	"fmt"
	"net/http"
	"shopping-app/models"
	"html/template"
)

// LoginHandler handles the login route.
func LoginHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Sample logic: Render login page
		tmpl, err := template.ParseFiles("views/templates/login.html")
		if err != nil {
			http.Error(w, "Could not load template", http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, nil)
	}
}

// ProductsHandler handles the products page.
func ProductsHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Sample logic: Render product page
		products := []models.Product{
			{ID: 1, Name: "Laptop", Price: 1000},
			{ID: 2, Name: "Smartphone", Price: 500},
			{ID: 3, Name: "Headphones", Price: 200},
			// Add more products here
		}
		tmpl, err := template.ParseFiles("views/templates/products.html")
		if err != nil {
			http.Error(w, "Could not load template", http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, products)
	}
}

// AddToCartHandler handles adding a product to the cart.
func AddToCartHandler(w http.ResponseWriter, r *http.Request) {
	// Sample logic: Add product to cart and show confirmation
	// Assuming you will handle cart logic in session or database
	fmt.Fprintln(w, "Product added to cart")
}

// CartHandler displays the contents of the cart.
func CartHandler(w http.ResponseWriter, r *http.Request) {
	// Sample logic: Render cart page (with added products)
	cart := []models.Product{
		{ID: 1, Name: "Laptop", Price: 1000},
		{ID: 2, Name: "Smartphone", Price: 500},
	}
	tmpl, err := template.ParseFiles("views/templates/cart.html")
	if err != nil {
		http.Error(w, "Could not load template", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, cart)
}

