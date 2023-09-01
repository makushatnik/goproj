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

	res := ApiResponse{["ok"], 200, ""}
	jsonData, _ = json.Marshal(res)
	_, _ := w.Write(jsonData)
}