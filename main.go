package main

import (
//	"encoding/json"
    "log"
    "net/http"
    "github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/login", LoginUser).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", router))
}

type User struct {
	Email string `json:"email"`
	Name string `json:"name"`
}

type Credential struct {
	Email string `json:"email"`
	Password string `json: "password"`
}

func LoginUser(w http.ResponseWriter, r *http.Request) {

}