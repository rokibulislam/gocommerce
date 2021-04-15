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

var users []domain.User

func GetUsers(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		println("testing users")
		users := []domain.User{}
		rows, err := db.Query("SELECT * FROM users")
		log.Println(err)
		defer rows.Close()

		for rows.Next() {
			var user domain.User
			println(rows)
			err := rows.Scan(&user.Id, &user.Name, &user.City, &user.Zipcode, &user.DateofBirth, &user.Email, &user.Status, &user.Password)
			if err != nil {
				panic(err.Error())
			}
			users = append(users, user)
		}

		if r.Header.Get("content-type") == "application/xml" {
			w.Header().Add("content-type", "application/xml")
			xml.NewEncoder(w).Encode(users)
		} else {
			w.Header().Add("content-type", "application/json")
			json.NewEncoder(w).Encode(users)
		}
	}
}

func GetUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id, _ := strconv.Atoi(params["id"])
		var user domain.User
		rows := db.QueryRow("SELECT * FROM users where id=?", id)
		err := rows.Scan(&user.Id, &user.Name, &user.City, &user.Zipcode, &user.DateofBirth, &user.Email, &user.Status, &user.Password)
		log.Println(err)

		if r.Header.Get("content-type") == "application/xml" {
			w.Header().Add("content-type", "application/xml")
			xml.NewEncoder(w).Encode(user)
		} else {
			w.Header().Add("content-type", "application/json")
			json.NewEncoder(w).Encode(user)
		}
	}
}

func CreateUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user domain.User
		var userID int
		json.NewDecoder(r.Body).Decode(&user)
		err := db.QueryRow("INSERT INTO users (name,email,password,city,zipcode,dateofbirth,status) VALUES ($1,$2,$3,$4,$5,$6,$7) RETURNING id;",
			user.Name, user.Email, user.Password, user.City, user.Zipcode, user.DateofBirth, user.Status).Scan(&userID)
		log.Println(err)
		users = append(users, user)
		json.NewEncoder(w).Encode(users)
	}
}

func UpdateUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user domain.User

		json.NewDecoder(r.Body).Decode(&user)
		result, _ := db.Exec("UPDATE users set name=$1 email=$2 password=$3 city=$4  zipcode=$5 dateofbirth=$6 status=$7 where id=$8 RETURNING id;",
			user.Name, user.Email, user.Password, user.City, user.Zipcode, user.DateofBirth, user.Status)
		// log.Println(er)
		rowupdated, _ := result.RowsAffected()
		json.NewEncoder(w).Encode(rowupdated)
	}
}

func DeleteUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id, _ := strconv.Atoi(params["id"])
		result, _ := db.Exec("DELETE users where id=$1 RETURNING id;", id)
		rowdeleted, _ := result.RowsAffected()
		json.NewEncoder(w).Encode(rowdeleted)
	}
}
