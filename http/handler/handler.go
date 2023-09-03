package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

type UserDto struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Age     int64  `json:"age"`
	IsAdmin bool   `json:"is_admin"`
}

type HttpResponse struct {
	Message string `json:"message"`
}

func ParseGetOrder(w http.ResponseWriter, r *http.Request) {
	orderId := r.URL.Query().Get("id")
	fmt.Printf("Order ID: %s\n", orderId)
}

func ParseSendMessage(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	message := r.FormValue("message")

	file, _, err := r.FormFile("file")
	if err != nil {
		fmt.Fprintln(w, err.Error())
	}
	defer file.Close()

	buf := bytes.NewBuffer(nil)
	io.Copy(buf, file)
	fmt.Fprintf(w, "email: %s, message: %s, file_data: %s",
		email, message, string(buf.Bytes()))
}

func ParsePostUser(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintln(w, err.Error())
	}
	r.Body.Close()

	var obj UserDto
	if err := json.Unmarshal(data, &obj); err != nil {
		fmt.Fprintln(w, err.Error())
	}

	fmt.Fprintf(w, "%v", obj)
}

func BasicAuth(w http.ResponseWriter, r *http.Request) {
	auth := r.Header.Get("Authorization")
	fmt.Fprintln(w, auth)
}

func JsonResponse(w http.ResponseWriter, r *http.Request) {
	resp := HttpResponse{"JsonResponse has sent!"}
	jsonData, err := json.Marshal(&resp)
	if err != nil {
		http.Error(w, "invalid json", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func GetOrderById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprintf(w, "Looking for Order by ID: %s", vars["id"])
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Webapp on Go. It works!")
}
