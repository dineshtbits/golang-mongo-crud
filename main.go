package main

import (
	"encoding/json"
	"github.com/dineshtbits/golang-mongo-crud/model"
	"github.com/dineshtbits/golang-mongo-crud/mongo"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type UserRequest struct {
	FirstName string   `json:"first_name"`
	LastName  string   `json:"last_name"`
	Roles     []string `json:"roles"`
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/users", createUserHandler).Methods("POST")
	r.HandleFunc("/users", getUsersHandler).Methods("GET")
	r.HandleFunc("/users/{id}", deleteUsersHandler).Methods("DELETE")
	r.HandleFunc("/users/{id}", updateUserHandler).Methods("PATCH")

	log.Fatal(http.ListenAndServe(":8000", r))
}

func getUsersHandler(w http.ResponseWriter, r *http.Request) {
	result := mongo.FindAllUsers()
	json.NewEncoder(w).Encode(result)
}

func deleteUsersHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	result := mongo.DeleteUser(vars["id"])
	json.NewEncoder(w).Encode(result)
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	var u UserRequest
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user := model.User{
		ID:        uuid.New().String(),
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Roles:     u.Roles,
	}

	result := mongo.CreateUser(user)
	json.NewEncoder(w).Encode(result)
}

func updateUserHandler(w http.ResponseWriter, r *http.Request) {
	var u UserRequest
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user := model.User{
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Roles:     u.Roles,
	}

	vars := mux.Vars(r)
	result := mongo.UpdateUser(vars["id"], user)
	json.NewEncoder(w).Encode(result)
}
