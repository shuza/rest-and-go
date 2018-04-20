package main

import (
	"github.com/gorilla/mux"
	"./api"
	"fmt"
	"net/http"
)

var routes = mux.NewRouter()

func main() {
	routes.HandleFunc("/", api.IndexHandler).Methods("GET", "POST", "PATCH", "OPTIONS", "PUT", "DELETE")

	v1 := routes.NewRoute().PathPrefix("/api/v1").Subrouter()
	v1.HandleFunc("/user/create", api.CreateUser).Methods("POST")
	v1.HandleFunc("/user/auth/login", api.LoginUser).Methods("POST")

	fmt.Println("Application running On : 8000....")
	http.ListenAndServe(":8000", routes)
}
