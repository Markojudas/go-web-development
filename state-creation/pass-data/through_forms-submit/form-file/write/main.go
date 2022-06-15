package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("templates/index.gohtml"))
}

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

		//write/store
		//ceates a file under the path ./user/
		// uses the original filename as the name for the new file
		dst, err := os.Create(filepath.Join("./user/", header.Filename))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer dst.Close()

		//writes the content of read file to the new created file
		_, err = dst.Write(bs)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	//RESPONSE HEADER
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	tpl.ExecuteTemplate(w, "index.gohtml", str)
}
