package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("templates/index.gohtml"))
}

/* type person struct {
	FirstName  string
	LastName   string
	Subscribed bool
} */

func foo(w http.ResponseWriter, req *http.Request) {

	//BODY
	byteslice := make([]byte, req.ContentLength) //the since of the slice is the content length

	req.Body.Read(byteslice) //reads INTO byteslice
	body := string(byteslice)

	err := tpl.ExecuteTemplate(w, "index.gohtml", body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalln(err)
	}
}

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

/*
	output:

BODY: -----------------------------234089674320769898433723809626 Content-Disposition: form-data; name="first" fdsf -----------------------------234089674320769898433723809626 Content-Disposition: form-data; name="last" fdfsf -----------------------------234089674320769898433723809626 Content-Disposition: form-data; name="subscribe" on -----------------------------234089674320769898433723809626 Content-Disposition: form-data; name="poem"; filename="blindguardian.txt" Content-Type: text/plain Even though the storm's calmed down the bitter end is just a matter of time shall we dare the dragon? merciless is poisoning our hearts our hearts -----------------------------234089674320769898433723809626--
*/
