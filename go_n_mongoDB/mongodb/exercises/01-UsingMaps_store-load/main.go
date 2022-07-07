package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/markojudas/web-development/go_n_mongoDB/mongodb/exercises/01-UsingMaps_store-load/controllers"
	"github.com/markojudas/web-development/go_n_mongoDB/mongodb/exercises/01-UsingMaps_store-load/models"
)

func main() {
	r := httprouter.New()

	uc := controllers.NewUserController(getSession())

	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	http.ListenAndServe("localhost:8080", r)

}

func getSession() map[string]models.User {
	return models.LoadUsers()
}
