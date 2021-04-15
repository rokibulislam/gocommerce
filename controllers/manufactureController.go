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

var manufactures []domain.Manufacture

func GetManufactures(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		manufactures := []domain.Manufacture{}
		rows, err := db.Query("SELECT * FROM manufactures")
		log.Println(err)
		defer rows.Close()

		for rows.Next() {
			var manufacture domain.Manufacture
			err := rows.Scan(&manufacture.Id, &manufacture.Name, &manufacture.Slug, &manufacture.Description)
			if err != nil {
				panic(err.Error())
			}
			manufactures = append(manufactures, manufacture)
		}

		if r.Header.Get("content-type") == "application/xml" {
			w.Header().Add("content-type", "application/xml")
			xml.NewEncoder(w).Encode(manufactures)
		} else {
			w.Header().Add("content-type", "application/json")
			json.NewEncoder(w).Encode(manufactures)
		}
	}
}

func GetManufacture(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id, _ := strconv.Atoi(params["id"])

		var manufacture domain.Manufacture
		rows := db.QueryRow("SELECT * FROM manufactures where id=?", id)
		err := rows.Scan(&manufacture.Id, &manufacture.Name, &manufacture.Slug, &manufacture.Description)
		log.Println(err)

		if r.Header.Get("content-type") == "application/xml" {
			w.Header().Add("content-type", "application/xml")
			xml.NewEncoder(w).Encode(manufacture)
		} else {
			w.Header().Add("content-type", "application/json")
			json.NewEncoder(w).Encode(manufacture)
		}
	}
}

func CreateManufacture(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var manufacture domain.Manufacture
		var manufactureID int
		json.NewDecoder(r.Body).Decode(&manufacture)
		err := db.QueryRow("INSERT INTO tags (name,slug,description) VALUES ($1,$2,$3) RETURNING id;",
			manufacture.Name, manufacture.Slug, manufacture.Description).Scan(&manufactureID)
		log.Println(err)
		manufactures = append(manufactures, manufacture)
		json.NewEncoder(w).Encode(manufactures)
	}
}

func UpdateManufacture(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var manufacture domain.Manufacture
		json.NewDecoder(r.Body).Decode(&manufacture)
		result, _ := db.Exec("UPDATE manufactures set name=$1 slug=$2 description=$3 where id=$4 RETURNING id;",
			manufacture.Name, manufacture.Slug, manufacture.Description, manufacture.Id)
		// log.Println(er)
		rowupdated, _ := result.RowsAffected()
		json.NewEncoder(w).Encode(rowupdated)
	}
}

func DeleteManufacture(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id, _ := strconv.Atoi(params["id"])
		result, _ := db.Exec("DELETE manufactures where id=$1 RETURNING id;", id)
		rowdeleted, _ := result.RowsAffected()
		json.NewEncoder(w).Encode(rowdeleted)
	}
}
