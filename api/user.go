package api

import (
	"encoding/json"
	"github.com/shuza/rest-and-go/data"
	"github.com/shuza/rest-and-go/utils"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {

	user := data.User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err == nil {
		if user.UserName != "" && user.UserPassword != "" {
			if user.Save() {
				json.NewEncoder(w).Encode(data.Response{
					Code:    http.StatusOK,
					Message: "User created successfully",
				})
			} else {
				json.NewEncoder(w).Encode(data.Response{
					Code:    http.StatusInternalServerError,
					Message: "Something went wrong!!!",
				})
			}
		}
	}
	json.NewEncoder(w).Encode(data.Response{
		Code:    http.StatusBadRequest,
		Message: "Malformed data received!!",
	})
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UFT-8")
	user := data.User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err == nil {
		if user.UserName != "" && user.UserPassword != "" {
			if user.FindUser() {
				session := data.Session{}
				session.AccessToken = utils.GetUUI()
				session.RefreshToken = utils.GetUUI()
				session.UserId = data.GetUser(user.UserName).UserId
				if session.Save() {
					json.NewEncoder(w).Encode(data.Response{
						Code:    http.StatusOK,
						Message: "Successfully Logged in",
						Data:    session,
					})
					return
				}
			} else {
				json.NewEncoder(w).Encode(data.Response{
					Code:    http.StatusNonAuthoritativeInfo,
					Message: "Username and password not match",
				})
				return
			}
		}
	}
	json.NewEncoder(w).Encode(data.Response{
		Code:    http.StatusBadRequest,
		Message: "Malformed data received!!",
	})
}
