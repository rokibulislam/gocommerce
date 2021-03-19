package controllers

import (
	"encoding/json"
	"encoding/xml"
	"net/http"

	"github.com/rokibulislam/gocommerce/domain"
)

func GetManufactures(w http.ResponseWriter, r *http.Request) {
	manufactures := []domain.Manufacture{
		{Name: "Rokib", City: "Dhaka"},
		{Name: "Kamrul", City: "Dhaka"},
	}

	if r.Header.Get("content-type") == "application/xml" {
		w.Header().Add("content-type", "application/xml")
		xml.NewEncoder(w).Encode(manufactures)
	} else {
		w.Header().Add("content-type", "application/json")
		json.NewEncoder(w).Encode(manufactures)
	}
}

func GetManufacture(w http.ResponseWriter, r *http.Request) {

}

func CreateManufacture(w http.ResponseWriter, r *http.Request) {

}

func DeleteManufacture(w http.ResponseWriter, r *http.Request) {

}
