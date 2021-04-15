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

var attributes []domain.Attribute

func GetAttributes(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		attributes := []domain.Attribute{}
		rows, err := db.Query("SELECT * FROM attributes")
		log.Println(err)
		defer rows.Close()

		for rows.Next() {
			var attribute domain.Attribute
			err := rows.Scan(&attribute.Id, &attribute.Name)
			if err != nil {
				panic(err.Error())
			}
			attributes = append(attributes, attribute)
		}

		if r.Header.Get("content-type") == "application/xml" {
			w.Header().Add("content-type", "application/xml")
			xml.NewEncoder(w).Encode(attributes)
		} else {
			w.Header().Add("content-type", "application/json")
			json.NewEncoder(w).Encode(attributes)
		}
	}
}

func GetAttribute(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id, _ := strconv.Atoi(params["id"])
		var attribute domain.Attribute
		rows := db.QueryRow("SELECT * FROM attributes where id=?", id)
		err := rows.Scan(&attribute.Id, &attribute.Name)
		log.Println(err)

		if r.Header.Get("content-type") == "application/xml" {
			w.Header().Add("content-type", "application/xml")
			xml.NewEncoder(w).Encode(attribute)
		} else {
			w.Header().Add("content-type", "application/json")
			json.NewEncoder(w).Encode(attribute)
		}
	}
}

func CreateAttribute(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var attribute domain.Attribute
		var attributeID int
		json.NewDecoder(r.Body).Decode(&attribute)
		err := db.QueryRow(
			`INSERT INTO attributes (name) VALUES($1) RETURNING id`, attribute.Name).Scan(&attributeID)
		log.Println(err)
		attributes = append(attributes, attribute)
		json.NewEncoder(w).Encode(attributes)
	}
}

func UpdateAttribute(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var attribute domain.Attribute
		json.NewDecoder(r.Body).Decode(&attribute)
		result, _ := db.Exec("UPDATE attributes set name=$1 where id=$2 RETURNING id;",
			attribute.Name, attribute.Id)
		// log.Println(er)
		rowupdated, _ := result.RowsAffected()
		json.NewEncoder(w).Encode(rowupdated)
	}
}

func DeleteAttribute(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id, _ := strconv.Atoi(params["id"])
		log.Println(id)
		result, err := db.Exec("DELETE FROM attributes WHERE id=$1 RETURNING id;", id)
		log.Println(err)
		rowdeleted, _ := result.RowsAffected()
		json.NewEncoder(w).Encode(rowdeleted)
	}
}
