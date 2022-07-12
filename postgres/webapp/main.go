package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/lib/pq"
)

//create a DB object
var db *sql.DB

//init
func init() {
	var err error

	db, err = sql.Open("postgres", "postgres://bond:password@localhost/bookstore?sslmode=disable")
	checkErr(err)

	if err = db.Ping(); err != nil {
		panic(err)
	}

	fmt.Println("Successfully Connected to Database")
}

//Creating a struct to hold the values
type Book struct {
	isbn   string
	title  string
	author string
	price  float32
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func booksIndex(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	rows, err := db.Query("SELECT * FROM books")
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	defer rows.Close()

	bks := make([]Book, 0)

	for rows.Next() {
		bk := Book{}

		err := rows.Scan(&bk.isbn, &bk.title, &bk.author, &bk.price)
		if err != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}

		bks = append(bks, bk)
	}
	if err = rows.Err(); err != nil {
		http.Error(w, http.StatusText(500), 500)
	}

	for _, bk := range bks {
		fmt.Fprintf(w, "%s, %s, %s, $%.2f\n", bk.isbn, bk.title, bk.author, bk.price)
	}
}

func main() {
	http.HandleFunc("/books", booksIndex)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}
