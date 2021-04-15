package controllers

import (
	"database/sql"
	"encoding/json"
	"encoding/xml"
	"net/http"

	"github.com/rokibulislam/gocommerce/domain"
)

func GetPayments(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		products := []domain.Payment{
			{Name: "Rokib", City: "Dhaka"},
			{Name: "Kamrul", City: "Dhaka"},
		}

		if r.Header.Get("content-type") == "application/xml" {
			w.Header().Add("content-type", "application/xml")
			xml.NewEncoder(w).Encode(products)
		} else {
			w.Header().Add("content-type", "application/json")
			json.NewEncoder(w).Encode(products)
		}
	}
}

func GetPayment(w http.ResponseWriter, r *http.Request) {
	// params := mux.Vars(r)
	// id, _ := strconv.Atoi(params["id"])
}

func CreatePayment(w http.ResponseWriter, r *http.Request) {

}

func UpdatePayment(w http.ResponseWriter, r *http.Request) {

}

func DeletePayment(w http.ResponseWriter, r *http.Request) {
	// params := mux.Vars(r)
	// id, _ := strconv.Atoi(params["id"])
}
