package controllers

import (
	"encoding/json"
	"encoding/xml"
	"net/http"

	"github.com/rokibulislam/gocommerce/domain"
)

func GetReviews(w http.ResponseWriter, r *http.Request) {
	products := []domain.Review{
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

func GetReview(w http.ResponseWriter, r *http.Request) {

}

func CreateReview(w http.ResponseWriter, r *http.Request) {

}

func DeleteReview(w http.ResponseWriter, r *http.Request) {

}
