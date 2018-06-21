package main

import (
    "encoding/json"
    "log"
    "net/http"
    "github.com/gorilla/mux"
    //"github.com/dgrijalva/jwt-go"
)

func main() {
    insecureRouter := mux.NewRouter()
    secureRouter := mux.NewRouter()
    amw := authMiddleWare{}
    secureRouter.Use(amw.ValidateJWT)
    insecureRouter.HandleFunc("/login", LoginUser).Methods("POST")
    secureRouter.HandleFunc("/me", FetchLoggedInUser).Methods("GET")
    log.Fatal(http.ListenAndServe(":8000", insecureRouter, secureRouter))
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
    jsonDecoder := json.NewDecoder(r.Body)
    var cred Credential
    err := jsonDecoder.Decode(&cred)
    if err != nil {
        panic(err)
    }
    //These values should come from database
    const validEmail = "ak.sskg@gmail.com"
    const validPassword = "password" //password should have been hashed
    if cred.Email == validEmail && cred.Password == validPassword {
        w.Write([]byte("head.pay.sign"))
    }
}

func FetchLoggedInUser(w http.ResponseWriter, r *http.Request) {
    
}


type authMiddleWare struct {

}

func (amw *authMiddleWare) ValidateJWT(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        //token := r.Header.Get("Authorization")

        //if valid user {
            next.ServeHTTP(w, r)
        //} else {
            
            //http.Error(w, "Forbidden", http.StatusForbidden)
        //}
    })
}