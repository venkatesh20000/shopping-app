package controllers

import (
	"html/template"
	"net/http"
)

func ProductsPage(w http.ResponseWriter, r *http.Request) {
	products := []string{
		"Laptop", "Smartphone", "Headphones", "Smartwatch",
		"Tablet", "Camera", "Printer", "Monitor", "Keyboard", "Mouse",
	}
	tmpl := template.Must(template.ParseFiles("views/templates/products.html"))
	tmpl.Execute(w, products)
}

func AddToCart(w http.ResponseWriter, r *http.Request) {
	// Simulate adding to cart
	http.Redirect(w, r, "/cart", http.StatusSeeOther)
}

