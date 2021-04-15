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

var coupons []domain.Coupon

func GetCoupons(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		coupons := []domain.Coupon{}
		rows, err := db.Query("SELECT * FROM coupons")
		log.Println(err)
		defer rows.Close()

		for rows.Next() {
			var coupon domain.Coupon
			err := rows.Scan(&coupon.Id, &coupon.Code, &coupon.Amount)
			if err != nil {
				panic(err.Error())
			}
			coupons = append(coupons, coupon)
		}

		if r.Header.Get("content-type") == "application/xml" {
			w.Header().Add("content-type", "application/xml")
			xml.NewEncoder(w).Encode(coupons)
		} else {
			w.Header().Add("content-type", "application/json")
			json.NewEncoder(w).Encode(coupons)
		}
	}
}

func GetCoupon(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id, _ := strconv.Atoi(params["id"])

		var coupon domain.Coupon
		rows := db.QueryRow("SELECT * FROM coupons where id=?", id)
		err := rows.Scan(&coupon.Id, &coupon.Code, &coupon.Amount)
		log.Println(err)

		if r.Header.Get("content-type") == "application/xml" {
			w.Header().Add("content-type", "application/xml")
			xml.NewEncoder(w).Encode(coupon)
		} else {
			w.Header().Add("content-type", "application/json")
			json.NewEncoder(w).Encode(coupon)
		}
	}
}

func CreateCoupon(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var coupon domain.Coupon
		var couponID int
		json.NewDecoder(r.Body).Decode(coupon)
		err := db.QueryRow("INSERT INTO coupons (code,amount) VALUES ($1,$2,) RETURNING id;",
			coupon.Code, coupon.Amount).Scan(&couponID)
		log.Println(err)
		coupons = append(coupons, coupon)
		json.NewEncoder(w).Encode(coupons)
	}
}

func UpdateCoupon(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var coupon domain.Coupon
		json.NewDecoder(r.Body).Decode(&coupon)
		result, _ := db.Exec("UPDATE coupons set code=$1 amount=$2 where id=$4 RETURNING id;",
			coupon.Code, coupon.Amount)
		// log.Println(er)
		rowupdated, _ := result.RowsAffected()
		json.NewEncoder(w).Encode(rowupdated)
	}
}

func DeleteCoupon(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id, _ := strconv.Atoi(params["id"])
		result, _ := db.Exec("DELETE coupons where id=$1 RETURNING id;", id)
		rowdeleted, _ := result.RowsAffected()
		json.NewEncoder(w).Encode(rowdeleted)
	}
}
