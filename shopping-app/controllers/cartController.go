package controllers

import (
	"html/template"
	"net/http"
)

func CartPage(w http.ResponseWriter, r *http.Request) {
	cartItems := []string{"Laptop", "Smartphone"} // Simulate cart
	tmpl := template.Must(template.ParseFiles("views/templates/cart.html"))
	tmpl.Execute(w, cartItems)
}

func CheckoutPage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("views/templates/checkout.html"))
	tmpl.Execute(w, nil)
}

