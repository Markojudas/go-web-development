package main

import (
	"database/sql"
	"fmt"
	"io"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	db, err = sql.Open("mysql", "DBMASTERUSERNAME:DBMASTERPASSWD@tcp(godb.c5tkerwyxnm3.us-east-1.rds.amazonaws.com:3306)/test?charset=utf8")
	check(err)
	defer db.Close()

	err = db.Ping()
	check(err)

	http.HandleFunc("/", index)
	http.HandleFunc("/amigos", amigos)
	http.HandleFunc("/create", create)
	http.HandleFunc("/insert", insert)
	http.HandleFunc("/read", read)
	http.HandleFunc("/update", update)
	http.HandleFunc("/delete", del)
	http.HandleFunc("/drop", drop)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	err := http.ListenAndServe(":8080", nil)
	check(err)

}

func index(w http.ResponseWriter, req *http.Request) {
	_, err = io.WriteString(w, "at index")
	check(err)
}

func amigos(w http.ResponseWriter, req *http.Request) {

	//accessing the table

	rows, err := db.Query(`SELECT aName FROM amigos;`)
	check(err)
	defer rows.Close()

	//data to be used in query
	var str, name string
	str = "RETRIVED RECORDS:\n"

	for rows.Next() {
		err = rows.Scan(&name)
		check(err)
		str += name + "\n"
	}
	fmt.Fprintln(w, str)
}

func create(w http.ResponseWriter, req *http.Request) {

	stmt, err := db.Prepare(`CREATE TABLE customer (name VARCHAR(20));`)
	check(err)
	defer stmt.Close()

	r, err := stmt.Exec()
	check(err)

	n, err := r.RowsAffected()
	check(err)

	fmt.Fprintln(w, "CREATED TABLE customer", n)
}

func insert(w http.ResponseWriter, req *http.Request) {

	stmt, err := db.Prepare(`INSERT INTO customer VALUES ("James");`)
	check(err)
	defer stmt.Close()

	r, err := stmt.Exec()
	check(err)

	n, err := r.RowsAffected()
	check(err)

	fmt.Fprintln(w, "INSERTED RECORD", n)
}

func read(w http.ResponseWriter, req *http.Request) {

	rows, err := db.Query(`SELECT * FROM customer;`)
	check(err)
	defer rows.Close()

	var name string

	for rows.Next() {
		err = rows.Scan(&name)
		check(err)
		fmt.Fprintln(w, "RETRIVED RECORD:", name)
	}
}

func update(w http.ResponseWriter, req *http.Request) {

	stmt, err := db.Prepare(`UPDATE customer SET name="Jimmy" WHERE name="James";`)
	check(err)
	defer stmt.Close()

	r, err := stmt.Exec()
	check(err)

	n, err := r.RowsAffected()
	check(err)

	fmt.Fprintln(w, "UPDATED RECORD:", n)
}

func del(w http.ResponseWriter, req *http.Request) {
	stmt, err := db.Prepare(`DELETE FROM customer WHERE name="Jimmy";`)
	check(err)

	r, err := stmt.Exec()
	check(err)

	n, err := r.RowsAffected()
	check(err)

	fmt.Fprintln(w, "DELETED RECORD:", n)
}

func drop(w http.ResponseWriter, req *http.Request) {
	stmt, err := db.Prepare(`DROP TABLE customer;`)
	check(err)
	defer stmt.Close()

	_, err = stmt.Exec()
	check(err)

	fmt.Fprintln(w, "DROPPED TABLE customer")
}
