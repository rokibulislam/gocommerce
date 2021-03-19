package controllers

import (
	"encoding/json"
	"encoding/xml"
	"net/http"

	"github.com/rokibulislam/gocommerce/domain"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	products := []domain.Product{
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

func GetProduct(w http.ResponseWriter, r *http.Request) {

}

func CreateProduct(w http.ResponseWriter, r *http.Request) {

}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {

}
