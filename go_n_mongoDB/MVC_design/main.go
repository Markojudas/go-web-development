package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/markojudas/web-development/go_n_mongoDB/MVC_design/controllers"
)

func main() {

	router := httprouter.New()

	userControl := controllers.NewUserController()

	router.GET("/user/:id", userControl.GetUser)

	router.POST("/user", userControl.CreateUser)

	router.DELETE("/user/:id", userControl.DeleteUser)

	http.ListenAndServe("localhost:8080", router)

}
