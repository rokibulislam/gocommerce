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

var packs []domain.Package

func GetPackages(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		packs := []domain.Package{}
		rows, err := db.Query("SELECT * FROM packages")
		log.Println(err)
		defer rows.Close()

		for rows.Next() {
			var pack domain.Package
			err := rows.Scan(&pack.Id, &pack.Name, &pack.Description, &pack.Price, &pack.Quantity)
			if err != nil {
				panic(err.Error())
			}
			packs = append(packs, pack)
		}

		if r.Header.Get("content-type") == "application/xml" {
			w.Header().Add("content-type", "application/xml")
			xml.NewEncoder(w).Encode(packs)
		} else {
			w.Header().Add("content-type", "application/json")
			json.NewEncoder(w).Encode(packs)
		}
	}
}

func GetPackage(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id, _ := strconv.Atoi(params["id"])

		var pack domain.Package
		rows := db.QueryRow("SELECT * FROM packages where id=?", id)
		err := rows.Scan(&pack.Id, &pack.Name, &pack.Description, &pack.Price, &pack.Quantity)
		log.Println(err)

		if r.Header.Get("content-type") == "application/xml" {
			w.Header().Add("content-type", "application/xml")
			xml.NewEncoder(w).Encode(pack)
		} else {
			w.Header().Add("content-type", "application/json")
			json.NewEncoder(w).Encode(pack)
		}
	}
}

func CreatePackage(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var pack domain.Package
		var packID int
		json.NewDecoder(r.Body).Decode(&pack)
		err := db.QueryRow("INSERT INTO reviews (name,description,price,qty)  VALUES ($1,$2,) RETURNING id;",
			pack.Name, pack.Description, pack.Price, pack.Quantity).Scan(&packID)
		log.Println(err)
		packs = append(packs, pack)
		json.NewEncoder(w).Encode(packs)
	}
}

func UpdatePackage(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var pack domain.Package
		json.NewDecoder(r.Body).Decode(&pack)
		result, _ := db.Exec("UPDATE packages set name=$1 description=$2 price=$3 qty=$4 where id=$4 RETURNING id;",
			pack.Name, pack.Description, pack.Price,
			pack.Quantity, pack.Id)
		// log.Println(er)
		rowupdated, _ := result.RowsAffected()
		json.NewEncoder(w).Encode(rowupdated)
	}
}

func DeletePackage(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id, _ := strconv.Atoi(params["id"])
		result, _ := db.Exec("DELETE packages where id=$1 RETURNING id;", id)
		rowdeleted, _ := result.RowsAffected()
		json.NewEncoder(w).Encode(rowdeleted)
	}
}
