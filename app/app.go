package app

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/rokibulislam/gocommerce/controllers"
)

func Start() {
	mux := http.NewServeMux()
	mux.HandleFunc("/login", controllers.Login)
	mux.HandleFunc("/register", controllers.Register)
	mux.HandleFunc("/customers", controllers.GetUsers)
	mux.HandleFunc("/products", controllers.GetProducts)
	mux.HandleFunc("/categories", controllers.GetCategories)
	mux.HandleFunc("/tags", controllers.GetCategories)
	mux.HandleFunc("/manufactures", controllers.GetManufactures)
	mux.HandleFunc("/coupons", controllers.GetCoupons)
	mux.HandleFunc("/reviews", controllers.GetReviews)
	mux.HandleFunc("/orders", controllers.GetPayments)
	mux.HandleFunc("/payments", controllers.GetPayments)

	log.Println("Listen on port 8003...")
	log.Fatal(http.ListenAndServe(":8003", mux))
}

func greet(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]bool{"greet": true})
}

func signup(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Signup")
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Login")
}

func home(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]bool{"ok": true})
}
