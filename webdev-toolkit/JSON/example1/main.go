package main

// MARSHAL AND ENCODE example

import (
	"encoding/json"
	"log"
	"net/http"
)

type person struct {
	Fname string
	Lname string
	Items []string
}

func main() {

	http.HandleFunc("/", foo)
	http.HandleFunc("/mshl", mshl)
	http.HandleFunc("/encd", encd)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)

}

func foo(w http.ResponseWriter, req *http.Request) {

	str := `<!DOCTYPE html>
	<html lang="en">
	<head>
	<meta charset="UTF-8">
	<title>FOO</title>
	</head>
	<body>
	You are at foo
	</body>
	</html>`

	w.Write([]byte(str))
}

func mshl(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	p1 := person{
		"James",
		"Bond",
		[]string{"suit", "gun", "wry sense of humor"},
	}

	json, err := json.Marshal(p1)
	if err != nil {
		log.Println(err)
	}

	w.Write(json)
}

func encd(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	p1 := person{
		"James",
		"Bond",
		[]string{"suit", "gun", "wry sense of humor"},
	}

	err := json.NewEncoder(w).Encode(p1)
	if err != nil {
		log.Println(err)
	}
}
