package api

import (
	"net/http"
	"strconv"
	"../data"
	"encoding/json"
)

func AddFriend(w http.ResponseWriter, r *http.Request) {
	userId, err := strconv.Atoi(r.Header.Get("user_id"))
	if err == nil {
		friend := data.Friend{}
		err := json.NewDecoder(r.Body).Decode(&friend)
		if err == nil {
			friend.UserId = userId
			if friend.Save() {
				w.Header().Set("Content-Type", "application/json; charset=UFT-8")
				json.NewEncoder(w).Encode(data.Response{
					Code:    http.StatusOK,
					Message: "Friend Added Successfully",
				})
				return
			}
		}
	}
	json.NewEncoder(w).Encode(data.Response{
		Code: http.StatusBadRequest,
	})
}

func GetFriendList(w http.ResponseWriter, r *http.Request) {
	userId, err := strconv.Atoi(r.Header.Get("user_id"))
	if err == nil {
		friends := data.GetFriendByUserId(userId)
		w.Header().Set("Content-Type", "application/json; charset=UFT-8")
		json.NewEncoder(w).Encode(data.Response{
			Code: http.StatusOK,
			Data: friends,
		})
		return
	}
	json.NewEncoder(w).Encode(data.Response{
		Code: http.StatusBadRequest,
	})
}
