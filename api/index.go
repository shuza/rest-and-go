package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/shuza/rest-and-go/data"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UFT-8")
	json.NewEncoder(w).Encode(data.Response{
		Code:    http.StatusOK,
		Message: "RESTFull API is running......",
	})
}

func GetRoutes() *mux.Router {
	routes := mux.NewRouter()
	routes.HandleFunc("/", IndexHandler).Methods("GET", "POST", "PATCH", "OPTIONS", "PUT", "DELETE")

	v1 := routes.NewRoute().PathPrefix("/api/v1").Subrouter()
	v1.HandleFunc("/user/create", CreateUser).Methods("POST")
	v1.HandleFunc("/user/auth/login", LoginUser).Methods("POST")

	v1.HandleFunc("/friend/create", Authorization(AddFriend)).Methods("POST")
	v1.HandleFunc("/friend/list", Authorization(GetFriendList)).Methods("GET")

	return routes
}
