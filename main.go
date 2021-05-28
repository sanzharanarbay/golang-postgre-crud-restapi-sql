package main

import (
	"RestApi/routes"
	"fmt"
	"log"
	"net/http"
)


func main() {
	r := routes.Router()
	fmt.Println("Starting server on the port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
