package routes

import (
	"github.com/gorilla/mux"
	"github.com/sharukh010/user-managment/pkg/controllers"
)

func RegisterUserRoutes(router *mux.Router){
	router.HandleFunc("/user/:id",controllers.GetUser).Methods("GET")
	router.HandleFunc("/user",controllers.CreateUser).Methods("POST")
	router.HandleFunc("/user/:id",controllers.DeleteUser).Methods("DELETE")
}