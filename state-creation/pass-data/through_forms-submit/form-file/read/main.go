package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {

	//variable to hold the data read
	var str string
	fmt.Println(req.Method) //just to see it
	//http.MethodPost is a constant == "POST"
	if req.Method == http.MethodPost {

		//open fp .. FormFile catches a file submitted by user
		fp, header, err := req.FormFile("q")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer fp.Close()

		//for your information
		fmt.Println("\nfile:", fp, "\nheader:", header, "\nerr", err)

		//read
		bs, err := ioutil.ReadAll(fp)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		str = string(bs)
	}

	//RESPONSE HEADER
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	//WRITING A FORM
	io.WriteString(w, `
	<form method="POST" enctype="multipart/form-data">
	<input type="file" name="q">
	<input type="submit">
	</form>
	<br>`+str)
}
