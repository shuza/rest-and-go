package api

import (
	"net/http"
	"encoding/json"
	"../data"
)

func IndexHandler(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json; charset=UFT-8")
	json.NewEncoder(w).Encode(data.Response{
		Code:http.StatusOK,
		Message:"RESTFull API is running......",
	})
}
