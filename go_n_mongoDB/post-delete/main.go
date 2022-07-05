package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/markojudas/web-development/go_n_mongoDB/json/models"
)

func main() {

	router := httprouter.New()

	router.GET("/", index)

	//add another route
	router.GET("/user/:id", getUser)

	router.POST("/user", createUser)

	router.DELETE("/user/:id", deleteUser)

	http.ListenAndServe("localhost:8080", router)

}

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	s := `<!DOCTYPE html>
	<html lang="en">
	<head>
	<meta charset="UTF-8">
	<title>Index</title>
	</head>
	<body>
	<a href="/user/9872309847">GO TO: http://localhost:8080/user/9872309847</a>
	</body>
	</html>
		`
	s += "\r" //carriage return

	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(s))
}

func getUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	_user := models.User{
		Name:   "James Bond",
		Gender: "male",
		Age:    37,
		Id:     p.ByName("id"),
	}

	//marshal into JSON
	_userJson, err := json.Marshal(_user)
	if err != nil {
		fmt.Println(err)
	}

	//write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", _userJson)

}

func createUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	//composite literal
	_user := models.User{}

	/// encode/decode for sending/receiving JSON to/from a stream
	json.NewDecoder(r.Body).Decode(&_user)

	//change ID
	_user.Id = "007"

	// marshal/unmarshal for having JSON assigned to a variable
	_userjson, _ := json.Marshal(_user)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", _userjson)
}

func deleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "write code to delete user\n")
}
