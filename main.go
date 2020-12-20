package main

import (
	"fmt"
	"github.com/shuza/rest-and-go/api"
	"net/http"
)

func main() {
	routes := api.GetRoutes()
	fmt.Println("Application running On : 8000....")
	http.ListenAndServe(":8000", routes)
}
