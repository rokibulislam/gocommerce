package controllers

import (
	"encoding/json"
	"encoding/xml"
	"net/http"

	"github.com/rokibulislam/gocommerce/domain"
)

func GetCoupons(w http.ResponseWriter, r *http.Request) {
	products := []domain.Coupon{
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

func GetCoupon(w http.ResponseWriter, r *http.Request) {

}

func CreateCoupon(w http.ResponseWriter, r *http.Request) {

}

func DeleteCoupon(w http.ResponseWriter, r *http.Request) {

}
