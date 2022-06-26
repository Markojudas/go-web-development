package main

import (
	"crypto/sha1"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	uuid "github.com/satori/go.uuid"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {

	cookie := getCookie(w, req)

	// process form submission
	if req.Method == http.MethodPost {
		//gettting a multi-part file and fileheader
		multiFile, fileHead, err := req.FormFile("nf")
		if err != nil {
			fmt.Println(err)
		}
		defer multiFile.Close()

		//encrypt (sha) file name
		ext := strings.Split(fileHead.Filename, ".")[1] // getting the extension

		//using cyrpto/sha1 package
		//hashing
		h := sha1.New()

		//
		io.Copy(h, multiFile)

		//add the ext to the hashed filename as hex
		fname := fmt.Sprintf("%x", h.Sum(nil)) + "." + ext

		// create new file
		wd, err := os.Getwd() // get the working directory
		if err != nil {
			fmt.Println(err)
		}

		path := filepath.Join(wd, "resources", "pics", fname)

		newFile, err := os.Create(path)
		if err != nil {
			fmt.Println(err)
		}
		defer newFile.Close()

		//reset the heads to the begging of the file
		multiFile.Seek(0, 0)

		// add/copy the multiFile to the newFile
		io.Copy(newFile, multiFile)

		//add filename to the user's cookie
		cookie = appendValue(w, cookie, fname)
	}

	xs := strings.Split(cookie.Value, "|")

	tpl.ExecuteTemplate(w, "index.gohtml", xs)
}

func getCookie(w http.ResponseWriter, req *http.Request) *http.Cookie {

	//get the cookie
	cookie, err := req.Cookie("session")

	//if the cookie is not found
	if err != nil {
		sessionID := uuid.NewV4()

		cookie = &http.Cookie{
			Name:  "session",
			Value: sessionID.String(),
		}

		http.SetCookie(w, cookie)
	}

	return cookie
}

func appendValue(w http.ResponseWriter, cookie *http.Cookie, fname string) *http.Cookie {
	str := cookie.Value

	if !strings.Contains(str, fname) {
		str += "|" + fname
	}

	cookie.Value = str
	http.SetCookie(w, cookie)

	return cookie
}
