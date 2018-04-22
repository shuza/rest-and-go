package data

import (
	"../db"
	"fmt"
)

type Friend struct {
	FriendId           int    `json:"friend_id"`
	UserId             int    `json:"user_id"`
	FriendName         string `json:"friend_name"`
	FriendMobileNumber string `json:"friend_mobile_number"`
}

func (f *Friend) Save() bool {
	query := "CREATE TABLE IF NOT EXISTS friends(friend_id INTEGER PRIMARY KEY AUTOINCREMENT," +
		"user_id TEXT, friend_name TEXT, friend_mobile_number TEXT);"
	if db.ExecuteQuery(query) {
		query = "INSERT INTO friends (user_id, friend_name, friend_mobile_number) VALUES ('%d', '%s', '%s');"
		query = fmt.Sprintf(query, f.UserId, f.FriendName, f.FriendMobileNumber)
		return db.ExecuteQuery(query)
	}
	return false
}

func GetFriendByUserId(userId int) []Friend {
	query := "SELECT * FROM friends WHERE user_id='%d'"
	query = fmt.Sprintf(query, userId)
	rows, err := db.GetRews(query)
	if err != nil {
		return []Friend{}
	}
	defer rows.Close()
	friends := []Friend{}
	for rows.Next() {
		friend := Friend{}
		rows.Scan(&friend.FriendId, &friend.UserId, &friend.FriendName, &friend.FriendMobileNumber)
		friends = append(friends, friend)
	}
	return friends
}
