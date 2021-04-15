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

var orders []domain.Order

func GetOrders(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		orders := []domain.Order{}
		rows, err := db.Query("SELECT * FROM orders")
		log.Println(err)
		defer rows.Close()

		for rows.Next() {
			var order domain.Order
			err := rows.Scan(&order.Id, &order.Name, &order.City)
			if err != nil {
				panic(err.Error())
			}
			orders = append(orders, order)
		}

		if r.Header.Get("content-type") == "application/xml" {
			w.Header().Add("content-type", "application/xml")
			xml.NewEncoder(w).Encode(orders)
		} else {
			w.Header().Add("content-type", "application/json")
			json.NewEncoder(w).Encode(orders)
		}
	}
}

func GetOrder(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id, _ := strconv.Atoi(params["id"])
		var order domain.Order
		rows := db.QueryRow("SELECT * from orders where id=?", id)
		err := rows.Scan(&order.Id, &order.Name, &order.City)
		log.Println(err)

		if r.Header.Get("content-type") == "application/xml" {
			w.Header().Add("content-type", "application/xml")
			xml.NewEncoder(w).Encode(order)
		} else {
			w.Header().Add("content-type", "application/json")
			json.NewEncoder(w).Encode(order)
		}
	}
}

func CreateOrder(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var order domain.Order
		var orderID int
		json.NewDecoder(r.Body).Decode(&order)
		err := db.QueryRow("INSERT INTO orders (name) VALUES ($1) RETURNING id;",
			order.Name).Scan(&orderID)
		log.Println(err)
		orders = append(orders, order)
		json.NewEncoder(w).Encode(orders)
	}
}

func UpdateOrder(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var order domain.Order
		json.NewDecoder(r.Body).Decode(&order)
		result, _ := db.Exec("UPDATE orders set name=$1 where id=$2 RETURNING id;",
			order.Name, order.Id)
		// log.Println(er)
		rowupdated, _ := result.RowsAffected()
		json.NewEncoder(w).Encode(rowupdated)
	}
}

func DeleteOrder(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id, _ := strconv.Atoi(params["id"])
		result, _ := db.Exec("DELETE orders where id=$1 RETURNING id;", id)
		rowdeleted, _ := result.RowsAffected()
		json.NewEncoder(w).Encode(rowdeleted)
	}
}
