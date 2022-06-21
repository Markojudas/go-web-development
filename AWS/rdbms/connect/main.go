package main

import (
	"database/sql"
	"io"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func check(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	db, err = sql.Open("mysql", "marko:Anewfire=123123!!@tcp(godb.c5tkerwyxnm3.us-east-1.rds.amazonaws.com:3306)/test?charset=utf8")
	check(err)
	defer db.Close()

	err = db.Ping()
	check(err)

	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	err := http.ListenAndServe(":8080", nil)
	check(err)

}

func index(w http.ResponseWriter, req *http.Request) {
	_, err = io.WriteString(w, "Successfully completed")
	check(err)
}
