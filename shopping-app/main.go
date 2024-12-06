package main

import (
	"fmt"
	"html/template"
	"log"
	"net"
	"net/http"
	"os/exec"
	"strings"
)

// Detects the public IP of the server
func GetPublicIP() string {
	cmd := exec.Command("curl", "http://checkip.amazonaws.com")
	out, err := cmd.Output()
	if err != nil {
		log.Println("Error fetching public IP:", err)
		return "localhost" // Default fallback
	}
	return strings.TrimSpace(string(out))
}

// Detects the local IP of the server
func GetLocalIP() string {
	interfaces, err := net.Interfaces()
	if err != nil {
		log.Fatal(err)
	}
	for _, iface := range interfaces {
		addrs, err := iface.Addrs()
		if err != nil {
			log.Fatal(err)
		}
		for _, addr := range addrs {
			if ip, ok := addr.(*net.IPNet); ok && !ip.IP.IsLoopback() && ip.IP.To4() != nil {
				return ip.IP.String()
			}
		}
	}
	return "localhost" // Default fallback
}

func renderTemplate(w http.ResponseWriter, page string) {
	tmplPath := fmt.Sprintf("views/templates/%s", page)
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		http.Error(w, "Template not found!", http.StatusNotFound)
		return
	}
	tmpl.Execute(w, nil)
}

func main() {
	// Dynamically bind server IP
	publicIP := GetPublicIP()
	localIP := GetLocalIP()
	port := "3000"

	mux := http.NewServeMux()

	// Routes
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "index.html")
	})
	mux.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "register.html")
	})
	mux.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "login.html")
	})
	mux.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "products.html")
	})
	mux.HandleFunc("/cart", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "cart.html")
	})

	// Static files (CSS, JS)
	fileServer := http.FileServer(http.Dir("views/static"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	log.Printf("Server running on http://%s:%s (Public IP: http://%s:%s)", localIP, port, publicIP, port)
	log.Fatal(http.ListenAndServe("0.0.0.0:"+port, mux))
}

