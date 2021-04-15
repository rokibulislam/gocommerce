package app

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rokibulislam/gocommerce/controllers"
	"github.com/rokibulislam/gocommerce/database"
)

func Start() {
	db := database.ConnectDB()

	// mux := http.NewServeMux()
	router := mux.NewRouter()
	//Auth
	router.HandleFunc("/login", controllers.Login(db)).Methods("POST")
	router.HandleFunc("/register", controllers.Register(db)).Methods("POST")
	//User
	router.HandleFunc("/users", controllers.GetUsers(db)).Methods("GET")
	router.HandleFunc("/user/{id}", controllers.GetUser(db)).Methods("GET")
	router.HandleFunc("/user", controllers.CreateUser(db)).Methods("POST")
	router.HandleFunc("/user/{id}", controllers.UpdateUser(db)).Methods("PUT")
	router.HandleFunc("/user/{id}", controllers.DeleteUser(db)).Methods("DELETE")

	//product
	router.HandleFunc("/products", controllers.GetProducts(db)).Methods("GET")
	router.HandleFunc("/product/{id}", controllers.GetProduct(db)).Methods("GET")
	router.HandleFunc("/product", controllers.CreateProduct(db)).Methods("POST")
	router.HandleFunc("/product/{id}", controllers.UpdateProduct(db)).Methods("PUT")
	router.HandleFunc("/product/{id}", controllers.DeleteProduct(db)).Methods("DELETE")

	//categories
	router.HandleFunc("/categories", controllers.GetCategories(db)).Methods("GET")
	router.HandleFunc("/categories/{id}", controllers.GetCategory(db)).Methods("GET")
	router.HandleFunc("/categories", controllers.CreateCategory(db)).Methods("POST")
	router.HandleFunc("/categories/{id}", controllers.UpdateCategory(db)).Methods("PUT")
	router.HandleFunc("/categories/{id}", controllers.DeleteCategory(db)).Methods("DELETE")

	//attribute
	router.HandleFunc("/attributes", controllers.GetAttributes(db)).Methods("GET")
	router.HandleFunc("/attributes/{id}", controllers.GetAttribute(db)).Methods("GET")
	router.HandleFunc("/attributes", controllers.CreateAttribute(db)).Methods("POST")
	router.HandleFunc("/attributes/{id}", controllers.UpdateAttribute(db)).Methods("PUT")
	router.HandleFunc("/attributes/{id}", controllers.DeleteAttribute(db)).Methods("DELETE")
	// Tag
	router.HandleFunc("/tags", controllers.GetTags(db)).Methods("GET")
	router.HandleFunc("/tag/{id}", controllers.GetTag(db)).Methods("GET")
	router.HandleFunc("/tags", controllers.CreateTag(db)).Methods("POST")
	router.HandleFunc("/tags/{id}", controllers.UpdateTag(db)).Methods("PUT")
	router.HandleFunc("/tags/{id}", controllers.DeleteTag(db)).Methods("DELETE")
	// Manufacture
	router.HandleFunc("/manufactures", controllers.GetManufactures(db)).Methods("GET")
	router.HandleFunc("/manufacture/{id}", controllers.GetManufacture(db)).Methods("GET")
	router.HandleFunc("/manufactures", controllers.CreateManufacture(db)).Methods("POST")
	router.HandleFunc("/manufactures/{id}", controllers.UpdateManufacture(db)).Methods("PUT")
	router.HandleFunc("/manufactures/{id}", controllers.DeleteManufacture(db)).Methods("DELETE")

	// Coupon
	router.HandleFunc("/coupons", controllers.GetCoupons(db)).Methods("GET")
	router.HandleFunc("/coupon/{id}", controllers.GetCoupon(db)).Methods("GET")
	router.HandleFunc("/coupons", controllers.CreateCoupon(db)).Methods("POST")
	router.HandleFunc("/coupons/{id}", controllers.UpdateCoupon(db)).Methods("PUT")
	router.HandleFunc("/coupons/{id}", controllers.DeleteCoupon(db)).Methods("DELETE")

	// Reviews
	router.HandleFunc("/reviews", controllers.GetReviews(db)).Methods("GET")
	router.HandleFunc("/review/{id}", controllers.GetReview(db)).Methods("GET")
	router.HandleFunc("/reviews", controllers.CreateReview(db)).Methods("POST")
	router.HandleFunc("/reviews/{id}", controllers.UpdateReview(db)).Methods("PUT")
	router.HandleFunc("/reviews/{id}", controllers.DeleteReview(db)).Methods("DELETE")

	//Orders
	router.HandleFunc("/orders", controllers.GetOrders(db)).Methods("GET")
	router.HandleFunc("/order/{id}", controllers.GetOrder(db)).Methods("GET")
	router.HandleFunc("/orders", controllers.CreateOrder(db)).Methods("POST")
	router.HandleFunc("/orders/{id}", controllers.UpdateOrder(db)).Methods("PUT")
	router.HandleFunc("/orders/{id}", controllers.DeleteOrder(db)).Methods("DELETE")

	//Packages
	router.HandleFunc("/packages", controllers.GetPackages(db)).Methods("GET")
	router.HandleFunc("/package/{id}", controllers.GetPackage(db)).Methods("GET")
	router.HandleFunc("/packages", controllers.CreatePackage(db)).Methods("POST")
	router.HandleFunc("/packages/{id}", controllers.UpdatePackage(db)).Methods("PUT")
	router.HandleFunc("/packages/{id}", controllers.DeletePackage(db)).Methods("DELETE")

	// router.HandleFunc("/payments", controllers.GetPayments).Methods("GET")
	// router.HandleFunc("/payments/{id}", controllers.GetPayment).Methods("GET")

	log.Println("Listen on port 8005...")
	log.Fatal(http.ListenAndServe(":8005", router))
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
