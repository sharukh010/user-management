package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sharukh010/user-managment/pkg/models"
	"gopkg.in/mgo.v2/bson"
)

func GetUser(w http.ResponseWriter,r *http.Request){
	userId := mux.Vars(r)["userId"]

	if !bson.IsObjectIdHex(userId){
		w.WriteHeader(http.StatusNotFound)
		return
	}

	oid := bson.ObjectIdHex(userId)
	user,err := models.GetUserById(oid)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	userJSON,err := json.Marshal(user)

	if err != nil {
		fmt.Printf("error : %v\n",err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type","Application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(userJSON)
}

func CreateUser(w http.ResponseWriter,r *http.Request){
	user := models.User{}

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		fmt.Println("Failed to decode JSON:", err.Error())
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
	}

	if user.Name == "" || user.Gender == "" {
		fmt.Println("Age : ",user.Age)
        http.Error(w, "Missing required fields", http.StatusBadRequest)
        return
    }

	err := user.CreateUser()
	
	if err != nil {
		fmt.Println("Error occured :",err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	userJSON,err := json.Marshal(user)
	if err != nil {
		fmt.Printf("error : %v\n",err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type","Application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(userJSON)

}