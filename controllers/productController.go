package controllers

import (
	"database/sql"
	"encoding/json"
	"encoding/xml"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rokibulislam/gocommerce/domain"
)

var products []domain.Product

func GetProducts(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		products := []domain.Product{}
		rows, err := db.Query("SELECT * FROM products")
		log.Println(err)
		defer rows.Close()

		for rows.Next() {
			var product domain.Product
			err := rows.Scan(&product.Id, &product.Title, &product.Description, &product.Price)
			if err != nil {
				panic(err.Error())
			}
			products = append(products, product)
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

func GetProduct(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id, _ := strconv.Atoi(params["id"])

		var product domain.Product
		rows := db.QueryRow("SELECT * FROM products where id=?", id)
		err := rows.Scan(&product.Id, &product.Title, &product.Description, &product.Price)
		log.Println(err)

		if r.Header.Get("content-type") == "application/xml" {
			w.Header().Add("content-type", "application/xml")
			xml.NewEncoder(w).Encode(product)
		} else {
			w.Header().Add("content-type", "application/json")
			json.NewEncoder(w).Encode(product)
		}
	}
}

func CreateProduct(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var product domain.Product
		var productID int
		json.NewDecoder(r.Body).Decode(&product)
		err := db.QueryRow("INSERT INTO products (name,description,price)  VALUES ($1,$2,$3) RETURNING id;",
			product.Title, product.Description, product.Price).Scan(&productID)
		log.Println(err)
		products = append(products, product)
		json.NewEncoder(w).Encode(products)
	}
}

func UpdateProduct(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var product domain.Product
		json.NewDecoder(r.Body).Decode(&product)
		result, _ := db.Exec("UPDATE products set title=$1 description=$2 price=$3 where id=$4 RETURNING id;",
			product.Title, product.Description, product.Price, product.Id)
		// log.Println(er)
		rowupdated, _ := result.RowsAffected()
		json.NewEncoder(w).Encode(rowupdated)
	}
}

func DeleteProduct(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id, _ := strconv.Atoi(params["id"])
		result, _ := db.Exec("DELETE products where id=$1 RETURNING id;", id)
		rowdeleted, _ := result.RowsAffected()
		json.NewEncoder(w).Encode(rowdeleted)
	}
}
