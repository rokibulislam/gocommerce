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

var tags []domain.Tag

func GetTags(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		tags := []domain.Tag{}
		rows, err := db.Query("SELECT * FROM tags")
		log.Println(err)
		defer rows.Close()

		for rows.Next() {
			var tag domain.Tag
			err := rows.Scan(&tag.Id, &tag.Name, &tag.Slug, &tag.Description)
			if err != nil {
				panic(err.Error())
			}
			tags = append(tags, tag)
		}

		if r.Header.Get("content-type") == "application/xml" {
			w.Header().Add("content-type", "application/xml")
			xml.NewEncoder(w).Encode(tags)
		} else {
			w.Header().Add("content-type", "application/json")
			json.NewEncoder(w).Encode(tags)
		}
	}
}

func GetTag(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id, _ := strconv.Atoi(params["id"])
		var tag domain.Tag
		rows := db.QueryRow("SELECT * FROM tags where id=?", id)
		err := rows.Scan(&tag.Id, &tag.Name, &tag.Slug, &tag.Description)
		log.Println(err)

		if r.Header.Get("content-type") == "application/xml" {
			w.Header().Add("content-type", "application/xml")
			xml.NewEncoder(w).Encode(tag)
		} else {
			w.Header().Add("content-type", "application/json")
			json.NewEncoder(w).Encode(tag)
		}
	}
}

func CreateTag(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var tag domain.Tag
		var tagID int
		json.NewDecoder(r.Body).Decode(&tag)
		err := db.QueryRow("INSERT INTO tags (name,slug,description) VALUES ($1,$2,$3) RETURNING id;",
			tag.Name, tag.Slug, tag.Description).Scan(&tagID)
		log.Println(err)
		tags = append(tags, tag)
		json.NewEncoder(w).Encode(tags)
	}
}

func UpdateTag(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var tag domain.Tag
		json.NewDecoder(r.Body).Decode(&tag)
		result, _ := db.Exec("UPDATE tags set name=$1 slug=$2 description=$3 where id=$4 RETURNING id;",
			tag.Name, tag.Slug, tag.Description, tag.Id)
		// log.Println(er)
		rowupdated, _ := result.RowsAffected()
		json.NewEncoder(w).Encode(rowupdated)
	}
}

func DeleteTag(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id, _ := strconv.Atoi(params["id"])
		result, _ := db.Exec("DELETE tags where id=$1 RETURNING id;", id)
		rowdeleted, _ := result.RowsAffected()
		json.NewEncoder(w).Encode(rowdeleted)
	}
}
