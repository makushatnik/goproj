package api

import (
	"encoding/json"
	"net/http"
)

type ApiResponse struct {
	Result []any
	Code int
	Err string
}

func handler(w http.ResponseWriter, r *http.Request) {

	res := ApiResponse{Result: ["ok"], Code: 200, Err: ""}
	jsonData, _ = json.Marshal(res)
	_, _ := w.Write(jsonData)
}