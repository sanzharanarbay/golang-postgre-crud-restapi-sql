package controllers

import (
	"RestApi/models" // models package where User schema is defined
	"RestApi/services"
	"encoding/json" // package to encode and decode the json into struct and vice versa
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http" // used to access the request and response object of the api
	"strconv"

	//"strconv"  // package used to covert string into int type

	//"github.com/gorilla/mux" // used to get the params from the route
)

// response format
type Response struct {
	Status      int64  `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
}

func  CreateUser(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var user models.User
	err:= json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Fatalf("Unable to decode the request body.  %v", err)
	}

	// call insert user function and pass the user
	insertID := services.InsertUser(user)

	// format a response object
	res := Response{
		Status:      insertID,
		Message: "User created successfully",
	}

	// send the response
	json.NewEncoder(w).Encode(res)
}

// GetUser will return a single user by its id
func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// get the userid from the request params, key is "id"
	params := mux.Vars(r)
	// convert the id type from string to int
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}
	// call the getUser function with user id to retrieve a single user
	user, err := services.GetUser(int64(id))
	if err != nil {
		log.Fatalf("Unable to get user. %v", err)
	}
	// send the response
	json.NewEncoder(w).Encode(user)
}

// GetAllUser will return all the users
func GetAllUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// get all the users in the db
	users, err := services.GetAllUsers()
	if err != nil {
		log.Fatalf("Unable to get all user. %v", err)
	}
	// send all the users as response
	json.NewEncoder(w).Encode(users)
}

// UpdateUser update user's detail in the postgres db
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	// get the userid from the request params, key is "id"
	params := mux.Vars(r)
	// convert the id type from string to int
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}
	// create an empty user of type models.User
	var user models.User
	// decode the json request to user
	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Fatalf("Unable to decode the request body.  %v", err)
	}
	// call update user to update the user
	updatedRows := services.UpdateUser(int64(id), user)
	// format the message string
	msg := fmt.Sprintf("User updated successfully. Total rows/record affected %v", updatedRows)

	// format the response message
	res := Response{
		Status:      int64(id),
		Message: msg,
	}

	// send the response
	json.NewEncoder(w).Encode(res)
}

// DeleteUser delete user's detail in the postgres db
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	// get the userid from the request params, key is "id"
	params := mux.Vars(r)
	// convert the id in string to int
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}
	// call the deleteUser, convert the int to int64
	deletedRows := services.DeleteUser(int64(id))
	// format the message string
	msg := fmt.Sprintf("User updated successfully. Total rows/record affected %v", deletedRows)
	// format the reponse message
	res := Response{
		Status:      int64(id),
		Message: msg,
	}
	// send the response
	json.NewEncoder(w).Encode(res)
}
