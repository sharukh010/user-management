package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sharukh010/user-managment/pkg/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	userId := mux.Vars(r)["userId"]

	oid, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	user, err := models.GetUserById(oid)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	userJSON, err := json.Marshal(user)
	if err != nil {
		http.Error(w, "Failed to serialize user", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(userJSON)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if user.Name == "" || user.Gender == "" {
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}

	err := user.CreateUser()
	if err != nil {
		http.Error(w, "Failed to create user", http.StatusBadRequest)
		return
	}

	userJSON, err := json.Marshal(user)
	if err != nil {
		http.Error(w, "Failed to serialize user", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(userJSON)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	userId := mux.Vars(r)["userId"]

	oid, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	err = models.DeleteUser(oid)
	if err != nil {
		http.Error(w, "Failed to delete user", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Deleted user with ID: %s\n", userId)
}
