package api

import (
	"encoding/json"
	"net/http"
	"rest-and-go/data"
	"strconv"
)

func Authorization(h http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		userId, err := strconv.Atoi(request.Header.Get("user_id"))
		if err == nil {
			accessToken := request.Header.Get("access_token")
			if userId > 0 && accessToken != "" {
				session := data.GetSession(accessToken)
				if session.AccessToken == accessToken && session.UserId == userId {
					h.ServeHTTP(writer, request)
					return
				}
			}
		}
		writer.Header().Set("Content-Type", "application/json; charset=UFT-8")
		json.NewEncoder(writer).Encode(data.Response{
			Code:    http.StatusForbidden,
			Message: "Unauthorized user..",
		})
		return
	}
}
