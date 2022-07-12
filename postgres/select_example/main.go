package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

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

func main() {

	//Connect to the DB; get a pointer to the DB
	db, err := sql.Open("postgres", "postgres://bond:password@localhost/bookstore?sslmode=disable")
	checkErr(err)
	//defer close connection
	defer db.Close()

	//test the connection with a ping
	if err = db.Ping(); err != nil {
		panic(err)
	}
	//if successful
	fmt.Println("Successfully Connected to Database")

	//QUERY
	rows, err := db.Query("SELECT * FROM books;")
	checkErr(err)

	//defer close query
	defer rows.Close()

	//making a list books on the DB
	bks := make([]Book, 0)

	//LOOP through the records found
	for rows.Next() {
		bk := Book{} //empty composite literal

		//create a book with the data returned from query
		err := rows.Scan(&bk.isbn, &bk.title, &bk.author, &bk.price) //order matters!!
		checkErr(err)

		//append to the list
		bks = append(bks, bk)
	}
	//check for error
	if err = rows.Err(); err != nil {
		panic(err)
	}

	//print the results
	for _, bk := range bks {
		fmt.Printf("%s, %s, %s, $%.2f\n", bk.isbn, bk.title, bk.author, bk.price)
	}

}
