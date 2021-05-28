package routes

import (
	"RestApi/controllers"
	"github.com/gorilla/mux"
)

// Router is exported and used in main.go
func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/users/user/{id}", controllers.GetUser).Methods("GET")
	router.HandleFunc("/api/users/all", controllers.GetAllUser).Methods("GET")
	router.HandleFunc("/api/users/create", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/api/users/update/{id}", controllers.UpdateUser).Methods("PUT")
	router.HandleFunc("/api/users/delete/{id}", controllers.DeleteUser).Methods("DELETE")
	return router
}
