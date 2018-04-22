package data

import (
	"../db"
	"fmt"
)

type Session struct {
	SessionId    int    `json:"session_id",omitempty`
	UserId       int    `json:"user_id"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func (s *Session) Save() bool {
	query := "CREATE TABLE IF NOT EXISTS sessions(session_id INTEGER PRIMARY KEY AUTOINCREMENT," +
		"user_id INTEGER, access_token TEXT, refresh_token TEXT);"
	result := db.ExecuteQuery(query)
	if result {
		query = "INSERT INTO sessions (user_id, access_token, refresh_token) VALUES ('%d', '%s', '%s');"
		query = fmt.Sprintf(query, s.UserId, s.AccessToken, s.RefreshToken)
		return db.ExecuteQuery(query)
	}
	return false
}

func GetSession(accessToken string) Session {
	query := "SELECT * FROM sessions WHERE access_token='%s'"
	query = fmt.Sprintf(query, accessToken)
	rows, err := db.GetRews(query)
	session := Session{}
	if err == nil {
		if rows.Next() {
			rows.Scan(&session.SessionId, &session.UserId, &session.AccessToken, &session.RefreshToken)
			rows.Close()
		}
	}
	return session
}
