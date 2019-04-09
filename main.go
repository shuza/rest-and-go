package main

import (
	"fmt"
	"github.com/shuza/rest-and-go/api"
	"net/http"
)

/**
 *
 * :=  created by:  Shuza
 * :=  create date:  23-Apr-18
 * :=  (C) CopyRight Shuza
 * :=  www.shuza.ninja
 * :=  shuza.sa@gmail.com
 * :=  Fun  :  Coffee  :  Code
 *
 **/

func main() {
	routes := api.GetRoutes()
	fmt.Println("Application running On : 8000....")
	http.ListenAndServe(":8000", routes)
}
