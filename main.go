package main

import (
	"fmt"
	"github.com/shuza/rest-and-go/storage"
)

func main() {
	initSqlite()
	fmt.Println("Server is under construction")
}

func initSqlite() {
	storage.DbClient = &storage.SqliteConnectin{}
	if err := storage.DbClient.Init(); err != nil {
		panic(err)
	}
}
