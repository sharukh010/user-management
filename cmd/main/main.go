package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sharukh010/user-managment/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterUserRoutes(r)
	http.Handle("/",r)
	fmt.Println("Server running on port 8001")
	err := http.ListenAndServe("localhost:8001",r)
	if err != nil {
		log.Fatalf("Error occured:%v",err.Error())
	}
}