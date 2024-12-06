package routes

import (
	"net/http"
	"shopping-app/controllers"
)

func SetupRoutes() {
	// Static Files
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("views/static"))))

	// Templates
	http.HandleFunc("/", controllers.IndexPage)
	http.HandleFunc("/register", controllers.RegisterPage)
	http.HandleFunc("/login", controllers.LoginPage)
	http.HandleFunc("/products", controllers.ProductsPage)
	http.HandleFunc("/cart", controllers.CartPage)
	http.HandleFunc("/checkout", controllers.CheckoutPage)
	http.HandleFunc("/add-to-cart", controllers.AddToCart)
}

