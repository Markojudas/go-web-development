package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/markojudas/web-development/go_n_mongoDB/MVC_design/models"
)

//creating a userController
type UserController struct{}

func NewUserController() *UserController {
	return &UserController{}
}

func (userControl UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	_user := models.User{
		Name:   "James Bond",
		Gender: "male",
		Age:    37,
		Id:     p.ByName("id"),
	}

	_userjson, err := json.Marshal(_user)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", _userjson)
}

func (userControl UserController) CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	_user := models.User{}

	json.NewDecoder(r.Body).Decode(&_user)

	_user.Id = "007"

	_userjson, err := json.Marshal(_user)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", _userjson)
}

func (userControl UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	//TODO: only wite code to delete user
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "write code to delete user.\n")
}
