package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"rest-and-go/api"
)

/**
 *
 * :=  created by:  Shuza
 * :=  create date:  23-Apr-18
 * :=  (C) CopyRight Shuza
 * :=  www.shuza.ninja
 * :=  shuza.sa@gmail.com
 * :=  Fun  :  Coffee  :  Code
 *
 **/

var routes = mux.NewRouter()

func main() {
	routes.HandleFunc("/", api.IndexHandler).Methods("GET", "POST", "PATCH", "OPTIONS", "PUT", "DELETE")

	v1 := routes.NewRoute().PathPrefix("/api/v1").Subrouter()
	v1.HandleFunc("/user/create", api.CreateUser).Methods("POST")
	v1.HandleFunc("/user/auth/login", api.LoginUser).Methods("POST")

	v1.HandleFunc("/friend/create", api.Authorization(api.AddFriend)).Methods("POST")
	v1.HandleFunc("/friend/list", api.Authorization(api.GetFriendList)).Methods("GET")

	fmt.Println("Application running On : 8000....")
	http.ListenAndServe(":8000", routes)
}
