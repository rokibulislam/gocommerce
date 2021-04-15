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

var reviews []domain.Review

func GetReviews(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		reviews := []domain.Review{}
		rows, err := db.Query("SELECT * FROM reviews")
		log.Println(err)
		defer rows.Close()

		for rows.Next() {
			var review domain.Review
			err := rows.Scan(&review.Id, &review.Comment, &review.Rating)
			if err != nil {
				panic(err.Error())
			}
			reviews = append(reviews, review)
		}

		if r.Header.Get("content-type") == "application/xml" {
			w.Header().Add("content-type", "application/xml")
			xml.NewEncoder(w).Encode(reviews)
		} else {
			w.Header().Add("content-type", "application/json")
			json.NewEncoder(w).Encode(reviews)
		}
	}
}

func GetReview(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id, _ := strconv.Atoi(params["id"])

		var review domain.Review
		rows := db.QueryRow("SELECT * FROM reviews where id=?", id)
		err := rows.Scan(&review.Id, &review.Comment, &review.Rating)
		log.Println(err)

		if r.Header.Get("content-type") == "application/xml" {
			w.Header().Add("content-type", "application/xml")
			xml.NewEncoder(w).Encode(review)
		} else {
			w.Header().Add("content-type", "application/json")
			json.NewEncoder(w).Encode(review)
		}
	}
}

func CreateReview(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var review domain.Review
		var reviewID int
		json.NewDecoder(r.Body).Decode(review)
		err := db.QueryRow("INSERT INTO reviews (comment,rating)  VALUES ($1,$2,) RETURNING id;",
			review.Comment, review.Rating).Scan(&reviewID)
		log.Println(err)
		reviews = append(reviews, review)
		json.NewEncoder(w).Encode(reviews)
	}
}

func UpdateReview(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var review domain.Review
		json.NewDecoder(r.Body).Decode(&review)
		result, _ := db.Exec("UPDATE reviews set comment=$1 rating=$2 where id=$4 RETURNING id;",
			review.Comment, review.Rating, review.Id)
		// log.Println(er)
		rowupdated, _ := result.RowsAffected()
		json.NewEncoder(w).Encode(rowupdated)
	}
}

func DeleteReview(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id, _ := strconv.Atoi(params["id"])
		result, _ := db.Exec("DELETE reviews where id=$1 RETURNING id;", id)
		rowdeleted, _ := result.RowsAffected()
		json.NewEncoder(w).Encode(rowdeleted)
	}
}
