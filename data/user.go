package data

type User struct {
	UserId       int    `json:"user_id"`
	UserName     string `json:"user_name"`
	UserPassword string `json:"user_password"`
}

func (u *User) Save() bool {
	query := "CREATE TABLE IF NOT EXISTS users(user_id INTEGER PRIMARY KEY AUTOINCREMENT, user_name TEXT, user_password TEXT)"

}
