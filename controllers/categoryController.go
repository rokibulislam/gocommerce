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

var categories []domain.Category

func GetCategories(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		categories := []domain.Category{}
		rows, err := db.Query("SELECT * FROM category")
		log.Println(err)
		defer rows.Close()

		for rows.Next() {
			var category domain.Category
			err := rows.Scan(&category.Id, &category.Name, &category.Slug, &category.Description)
			if err != nil {
				panic(err.Error())
			}
			categories = append(categories, category)
		}

		if r.Header.Get("content-type") == "application/xml" {
			w.Header().Add("content-type", "application/xml")
			xml.NewEncoder(w).Encode(categories)
		} else {
			w.Header().Add("content-type", "application/json")
			json.NewEncoder(w).Encode(categories)
		}
	}
}

func GetCategory(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id, _ := strconv.Atoi(params["id"])
		var category domain.Category
		rows := db.QueryRow("SELECT * FROM category where id=?", id)
		err := rows.Scan(&category.Id, &category.Name, &category.Slug, &category.Description)
		log.Println(err)

		if r.Header.Get("content-type") == "application/xml" {
			w.Header().Add("content-type", "application/xml")
			xml.NewEncoder(w).Encode(category)
		} else {
			w.Header().Add("content-type", "application/json")
			json.NewEncoder(w).Encode(category)
		}
	}
}

func CreateCategory(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var category domain.Category
		var categoryID int
		json.NewDecoder(r.Body).Decode(category)
		err := db.QueryRow("INSERT INTO category (name,slug,description) VALUES ($1,$2,$3) RETURNING id;",
			category.Name, category.Slug, category.Description).Scan(&categoryID)
		log.Println(err)
		categories = append(categories, category)
		json.NewEncoder(w).Encode(categories)
	}
}

func UpdateCategory(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var category domain.Category
		json.NewDecoder(r.Body).Decode(&category)
		result, _ := db.Exec("UPDATE category set name=$1 slug=$2 description=$3 where id=$4 RETURNING id;",
			category.Name, category.Slug, category.Description, category.Id)
		// log.Println(er)
		rowupdated, _ := result.RowsAffected()
		json.NewEncoder(w).Encode(rowupdated)
	}
}

func DeleteCategory(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id, _ := strconv.Atoi(params["id"])
		result, _ := db.Exec("DELETE category where id=$1 RETURNING id;", id)
		rowdeleted, _ := result.RowsAffected()
		json.NewEncoder(w).Encode(rowdeleted)
	}
}
