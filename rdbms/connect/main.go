package main

import (
	"database/sql"
	"io"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func main() {
	db, err = sql.Open("mysql", "markojudas:Anewfire=123123!!@tcp(database-1.c5tkerwyxnm3.us-east-1.rds.amazonaws.com:3306)")
}

func index(w http.ResponseWriter, req *http.Request) {
	_, err = io.WriteString(w, "Successfully completed")
	check(err)
}
