package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var MysqlCN = ConnectDB()

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func ConnectDB() *sql.DB {
	client, err := sql.Open("mysql", "admin:admin@tcp(127.0.0.1:3306)/library")

	if err != nil {
		panic(err)
	}

	err = client.Ping()

	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println("Connect successfully")

	return client
}

/*ConnectionOK : Check the connection and return true or false */
func ConnectionOK() bool {
	err := MysqlCN.Ping()
	if err != nil {
		log.Fatal(err.Error())
		return false
	}
	return true
}
