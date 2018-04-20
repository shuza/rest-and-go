package db

import (
	_ "github.com/mattn/go-sqlite3"
	"database/sql"
	"fmt"
)

func NewConnection() *sql.DB {
	db, err := sql.Open("sqlite3", "rest-and-go.db")
	if err != nil {
		panic(err)
	}
	return db
}

func GetRews(query string) (*sql.Rows, error) {
	return NewConnection().Query(query)
}

func ExecuteQuery(query string) bool {
	result, err := NewConnection().Exec(query)
	fmt.Println(result)
	fmt.Println(err)
	return err == nil
}
