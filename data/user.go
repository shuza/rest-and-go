package data

import (
	"../db"
	"fmt"
)

type User struct {
	UserId       int    `json:"user_id"`
	UserName     string `json:"user_name"`
	UserPassword string `json:"user_password"`
}

func (u *User) Save() bool {
	query := "CREATE TABLE IF NOT EXISTS users(user_id INTEGER PRIMARY KEY AUTOINCREMENT, user_name TEXT, user_password TEXT);"
	result := db.ExecuteQuery(query)
	if result {
		query = "INSERT INTO users(user_name, user_password) VALUES('%s', '%s');"
		query = fmt.Sprintf(query, u.UserName, u.UserPassword)
		result = db.ExecuteQuery(query)
		return result
	}
	return false
}

func (u *User) FindUser() bool {
	query := "SELECT * FROM users WHERE user_name='%s' AND user_password='%s';"
	query = fmt.Sprintf(query, u.UserName, u.UserPassword)
	rows, err := db.GetRews(query)
	if err != nil {
		return false
	}
	result := rows.Next()
	rows.Close()
	return result
}

func GetUser(userName string) User {
	query := "SELECT * FROM users WHERE user_name='%s';"
	query = fmt.Sprintf(query, userName)
	user := User{}
	rows, err := db.GetRews(query)
	if err == nil {
		if rows.Next() {
			rows.Scan(&user.UserId, &user.UserName, &user.UserPassword)
			rows.Close()
		}
	}
	return user
}
